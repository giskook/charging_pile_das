package mq

import (
	"log"
	//"sync"

	"github.com/bitly/go-nsq"
	"github.com/giskook/charging_pile_das/conf"
)

type NsqProducer struct {
	//waitGroup *sync.WaitGroup
	conf     *conf.ProducerConf
	producer *nsq.Producer
}

func NewNsqProducer(config *conf.ProducerConf) *NsqProducer {
	return &NsqProducer{
		conf: config,
		//waitGroup: &sync.WaitGroup{},
	}
}

func (s *NsqProducer) Send(topic string, value []byte) error {
	err := s.producer.PublishAsync(topic, value, nil, nil)
	log.Printf("<OUT_NSQ> topic %s %s\n", topic, value)
	if err != nil {
		log.Println("error occur")
		log.Println(err.Error())
	}

	return err
}

func (s *NsqProducer) Start() {
	//s.waitGroup.Add(1)
	defer func() {
		err := recover()
		//s.waitGroup.Done()
		if err != nil {
			log.Println("err found")
		}

	}()
	config := nsq.NewConfig()

	var errmsg error
	s.producer, errmsg = nsq.NewProducer(s.conf.Addr, config)

	if errmsg != nil {
		//	log.Printf("create producer error" + errmsg.Error())
		panic("create producer error " + errmsg.Error())
	} else {
		log.Println("producer start ok")
	}
}

func (s *NsqProducer) Stop() {
	//	s.waitGroup.Done()
	//	s.waitGroup.Wait()

	s.producer.Stop()
}
