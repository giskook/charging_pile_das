package mq

import (
	"github.com/bitly/go-nsq"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/event_handler_nsq"
	"log"
)

type NsqSocket struct {
	conf      *conf.NsqConfiguration
	Producers []*NsqProducer
	Consumers []*NsqConsumer
}

func NewNsqSocket(config *conf.NsqConfiguration) *NsqSocket {
	nsq_producer_count := config.Producer.Count
	return &NsqSocket{
		conf:      config,
		Producers: make([]*NsqProducer, nsq_producer_count),
		Consumers: make([]*NsqConsumer, 0),
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
	for _, ch := range socket.conf.Consumer.Channels {
		consumer_conf := &NsqConsumerConf{
			Addr:    socket.conf.Consumer.Addr,
			Topic:   socket.conf.Consumer.Topic,
			Channel: ch,
			Handler: nsq.HandlerFunc(func(message *nsq.Message) error {
				data := message.Body
				event_handler_nsq.ProcessNsq(data)

				return nil
			}),
		}

		consumer := NewNsqConsumer(consumer_conf)
		consumer.Start()
		socket.Consumers = append(socket.Consumers, consumer)
	}

	for _, ch := range socket.conf.ConsumerNotify.Channels {
		consumer_conf := &NsqConsumerConf{
			Addr:    socket.conf.ConsumerNotify.Addr,
			Topic:   socket.conf.ConsumerNotify.Topic,
			Channel: ch,
			Handler: nsq.HandlerFunc(func(message *nsq.Message) error {
				data := message.Body
				event_handler_nsq.ProcessNsqNotify(data)

				return nil
			}),
		}

		consumer := NewNsqConsumer(consumer_conf)
		consumer.Start()
		socket.Consumers = append(socket.Consumers, consumer)
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
