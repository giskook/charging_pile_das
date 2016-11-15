package mq

import (
	"github.com/giskook/charging_pile_das/conf"
	"log"
)

type NsqSocket struct {
	conf      *conf.NsqConfiguration
	Producers []*NsqProducer
	Consumers []*NsqConsumer
}

func NewNsqSocket(config *conf.NsqConfiguration) *NsqSocket {
	nsq_producer_count := config.Producer.Count
	nsq_consumer_count := len(config.Consumer.Channels)
	return &NsqSocket{
		conf:      config,
		Producers: make([]*NsqProducer, nsq_producer_count),
		Consumers: make([]*NsqConsumer, nsq_consumer_count),
	}
}

func (socket *NsqSocket) Start() {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()

	producer := NewNsqProducer(socket.conf.Producer)
	producer.Start()
	socket.Producers[0] = producer
	socket.ConsumerStart()
}

func (socket *NsqSocket) ConsumerStart() {
	var i int = 0
	for _, ch := range socket.conf.Consumer.Channels {
		consumer := NewNsqConsumer(socket.conf.Consumer, ch)
		consumer.Start()
		socket.Consumers[i] = consumer
		i++
	}
}

func (socket *NsqSocket) Stop() {
	for _, producer := range socket.Producers {
		producer.Stop()
	}

	for _, consumer := range socket.Consumers {
		consumer.Stop()
	}
}

func (socket *NsqSocket) Send(topic string, value []byte) error {
	err := socket.Producers[0].Send(topic, value)

	return err
}
