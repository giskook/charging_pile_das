package mq

import (
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

func ProcessNsq(socket *NsqSocket, message []byte) {
	command := &Report.Command{}
	err := proto.Unmarshal(message, command)
	if err != nil {
		log.Println("unmarshal error")
	} else {
		log.Printf("<IN NSQ> %s %d %x\n", command.Uuid, command.Tid, command.Type)
		switch command.Type {
		case Report.Command_CMT_REP_LOGIN:
			event_handler_rep_login(socket, command.Tid, command.SerialNumber, command.Paras)
			break
		case Report.Command_CMT_REP_SETTING:
			event_handler_rep_setting(command.Tid, command.SerialNumber, command.Paras)
			break
		case Report.Command_CMT_REP_PRICE:
			event_handler_rep_price(command.Tid, command.SerialNumber, command.Paras)
		case Report.Command_CMT_REQ_GET_GUN_STATUS:
			event_handler_req_gun_status(command.Tid)
		case Report.Command_CMT_REQ_CHARGING:
			event_handler_req_charging(command.Tid, command.SerialNumber, command.Paras)
		case Report.Command_CMT_REQ_STOP_CHARGING:
			event_handler_req_stop_charging(command.Tid)
		case Report.Command_CMT_REQ_PIN:
			event_handler_req_pin(command.Tid)

		}
	}
}
