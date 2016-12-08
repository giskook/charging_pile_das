package mq

import (
	"log"
	"sync"

	"github.com/bitly/go-nsq"
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
	s.consumer.AddHandler(s.config.Handler)

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
