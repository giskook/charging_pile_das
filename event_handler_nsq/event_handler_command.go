package event_handler_nsq

import (
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

func ProcessNsq(message []byte) {
	command := &Report.Command{}
	err := proto.Unmarshal(message, command)
	if err != nil {
		log.Println("unmarshal error")
	} else {
		log.Printf("<IN NSQ> %s %d \n", command.Uuid, command.Tid)
		switch command.Type {
		case Report.Command_CMT_REP_LOGIN:
			event_handler_rep_login(command.Tid, command.SerialNumber, command.Paras)
			break
		case Report.Command_CMT_REP_SETTING:
			event_handler_rep_setting(command.Tid, command.SerialNumber, command.Paras)
			break
		case Report.Command_CMT_REP_PRICE:
			event_handler_rep_price(command.Tid, command.SerialNumber, command.Paras)
		case Report.Command_CMT_REP_MODE:
			event_handler_rep_mode(command.Tid, command.SerialNumber, command.Paras)
		case Report.Command_CMT_REP_MAX_CURRENT:
			event_handler_rep_max_current(command.Tid, command.SerialNumber, command.Paras)
		case Report.Command_CMT_REQ_CHARGING_PREPARE:
			event_handler_req_charging_prepare(command.Tid, command.SerialNumber, command.Paras)
		case Report.Command_CMT_REQ_CHARGING:
			event_handler_req_charging(command.Tid, command.SerialNumber, command.Paras)
		case Report.Command_CMT_REQ_STOP_CHARGING:
			event_handler_req_stop_charging(command.Tid, command.SerialNumber, command.Paras)

		}
	}
}
