package main

import (
	"fmt"
	"github.com/giskook/charging_pile_das"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	charging_pile_das.ReadConfig(charging_pile_das.CONF_FILE_PATH)
	log.Println(charging_pile_das.GetConf())

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
