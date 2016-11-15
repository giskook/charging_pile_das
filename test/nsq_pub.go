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
	pub, errmsg := nsq.NewProducer("192.168.2.67:4150", config)
	if errmsg != nil {
		log.Printf("create producer error" + errmsg.Error())
	}

	err := pub.PublishAsync("test_nsq", []byte("I am puber"), nil, nil)
	if err != nil {
		log.Printf("create producer error" + errmsg.Error())
	}

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
