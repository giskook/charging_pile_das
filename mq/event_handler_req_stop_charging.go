package mq

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/protocol"
	"log"
)

func event_handler_req_stop_charging(tid uint64, serial uint32, param []*Report.Param) {
	log.Println(tid)
	pkg := protocol.ParseNsqStopCharging(tid, serial, param)
	connection := conn.NewConns().GetConn(tid)
	if connection != nil {
		connection.SendToTerm(pkg)
	}
}
