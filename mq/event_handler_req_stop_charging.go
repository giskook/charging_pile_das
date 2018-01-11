package mq

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/protocol"
	"log"
)

func event_handler_req_stop_charging(tid uint64) {
	log.Println(tid)
	pkg := protocol.ParseNsqStopCharging(tid)
	connection := conn.NewConns().GetConn(tid)
	if connection != nil && connection.Charging_Pile.StatusEx == base.STATUS_CHARGING {
		connection.SendToTerm(pkg)
	}
}
