package mq

import (
	"github.com/giskook/charging_pile_das/conf"
	"log"
)

type NsqSocket struct {
	conf      *conf.NsqConfiguration
	producers []*NsqProducer
	consumers []*NsqConsumer
}

func NewNsqSocket(config *conf.NsqConfiguration) *NsqSocket {
	nsq_producer_count := config.Producer.Count
	nsq_consumer_count := len(config.Consumer.Channels)
	return &NsqSocket{
		conf:      config,
		producers: make([]*NsqProducer, nsq_producer_count),
		consumers: make([]*NsqConsumer, nsq_consumer_count),
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
	socket.producers = append(socket.producers, producer)
	socket.ConsumerStart()
}

func (socket *NsqSocket) ConsumerStart() {
	for _, ch := range socket.conf.Consumer.Channels {
		consumer := NewNsqConsumer(socket.conf.Consumer, ch)
		consumer.Start()
		socket.consumers = append(socket.consumers, consumer)
	}
}

func (socket *NsqSocket) Stop() {
	for _, producer := range socket.producers {
		producer.Stop()
	}

	for _, consumer := range socket.consumers {
		consumer.Stop()
	}
}

func (socket *NsqSocket) Send(topic string, value []byte) error {
	return socket.producers[0].Send(topic, value)
}
