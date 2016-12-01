package main

import (
	"fmt"
	"github.com/giskook/charging_pile_das"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/event_handler"
	"github.com/giskook/charging_pile_das/mq"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// read configuration
	configuration, err := conf.ReadConfig("./conf.json")

	checkError(err)
	// creates a tcp listener
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+configuration.Server.BindPort)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// creates a tcp server
	config := &gotcp.Config{
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}
	srv := gotcp.NewServer(config, &event_handler.Callback{}, &charging_pile_das.Charging_Pile_Protocol{})

	// create a mq socket
	mq_socket := mq.NewNsqSocket(configuration.Nsq)
	// create charging_pile_das server
	server_conf := &server.ServerConfig{
		Listener:      listener,
		AcceptTimeout: time.Duration(configuration.Server.ConnTimeout) * time.Second,
		NsqConf:       configuration.Nsq,
	}
	cpd_server := server.NewServer(srv, mq_socket, server_conf)
	server.SetServer(cpd_server)
	// starts service
	fmt.Println("listening:", listener.Addr())
	cpd_server.Start()

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)

	// stops service
	cpd_server.Stop()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
