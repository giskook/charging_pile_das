package main

import (
	"fmt"
	"github.com/bitly/go-nsq"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	config := nsq.NewConfig()

	consumer, errmsg := nsq.NewConsumer("tms", "ch2", config)
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		data := message.Body
		log.Printf("<IN_NSQ>  %s\n", data)

		return nil
	}))

	if errmsg != nil {
		log.Println("create producer error" + errmsg.Error())
	}
	consumer.ConnectToNSQD("192.168.2.67:4150")

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
