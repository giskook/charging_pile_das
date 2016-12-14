package mq

import (
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

func ProcessNsqNotify(message []byte) {
	command := &Report.Command{}
	err := proto.Unmarshal(message, command)
	if err != nil {
		log.Println("unmarshal error")
	} else {
		log.Printf("<IN NSQ> %s %x \n", command.Uuid, command.Tid)
		switch command.Type {
		case Report.Command_CMT_REQ_NOTIFY_SET_PRICE:
			event_handler_req_notify_set_price(command.Tid, command.SerialNumber, command.Paras)
			break
		}
	}
}
