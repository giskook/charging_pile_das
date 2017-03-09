package mq

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/protocol"
	"log"
)

func event_handler_rep_charging_cost(tid uint64, param []*Report.Param) {
	log.Println("event_handler_rep_charging_cost")
	pkg := protocol.ParseNsqRepCost(tid, param)
	connection := conn.NewConns().GetConn(tid)
	if connection != nil {
		connection.SendToTerm(pkg)
	}

}
