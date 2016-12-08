package main

import (
	"fmt"
	"github.com/bitly/go-nsq"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

type NsqConsumerConf struct {
	Addr    string
	Topic   string
	Channel string

	Handler nsq.HandlerFunc
}

type NsqConsumer struct {
	config *NsqConsumerConf

	waitGroup *sync.WaitGroup
	consumer  *nsq.Consumer
}

func NewNsqConsumer(conf *NsqConsumerConf) *NsqConsumer {
	return &NsqConsumer{
		config:    conf,
		waitGroup: &sync.WaitGroup{},
	}
}

func (s *NsqConsumer) Start() {
	s.waitGroup.Add(1)
	defer func() {
		s.waitGroup.Done()
		errmsg := recover()
		if errmsg != nil {
			log.Println(errmsg)
			s.Stop()
		}

	}()

	config := nsq.NewConfig()

	var errmsg error
	s.consumer, errmsg = nsq.NewConsumer(s.config.Topic, s.config.Channel, config)

	if errmsg != nil {
		//	panic("create consumer error -> " + errmsg.Error())
		log.Println("create consumer error -> " + errmsg.Error())
	}
	log.Println("aaa")
	s.consumer.AddHandler(s.config.Handler)
	log.Println("bbb")

	err := s.consumer.ConnectToNSQD(s.config.Addr)
	if err != nil {
		panic("Counld not connect to nsq -> " + err.Error())
	}
}

func (s *NsqConsumer) Stop() {
	s.waitGroup.Wait()

	errmsg := s.consumer.DisconnectFromNSQD(s.config.Addr)

	if errmsg != nil {
		log.Printf("stop consumer error ", errmsg.Error())
	}

	s.consumer.Stop()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	config := &NsqConsumerConf{
		Addr:    "192.168.2.67:4150",
		Topic:   "test",
		Channel: "test",
		Handler: nsq.HandlerFunc(func(message *nsq.Message) error {
			data := message.Body
			log.Printf("<IN_NSQ>  %s\n", data)

			return nil
		}),
	}

	consumer := NewNsqConsumer(config)
	consumer.Start()
	//	config := nsq.NewConfig()
	//
	//	consumer, errmsg := nsq.NewConsumer("tms", "ch2", config)
	//	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
	//		data := message.Body
	//		log.Printf("<IN_NSQ>  %s\n", data)
	//
	//		return nil
	//	}))
	//
	//	if errmsg != nil {
	//		log.Println("create producer error" + errmsg.Error())
	//	}
	//	consumer.ConnectToNSQD("192.168.2.67:4150")

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
