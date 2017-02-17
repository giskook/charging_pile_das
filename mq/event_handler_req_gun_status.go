package mq

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/protocol"
	"log"
)

func event_handler_req_gun_status(tid uint64) {
	log.Println("event_handler_req_gun_status")
	pkg := protocol.ParseNsqGunStatus(tid)
	connection := conn.NewConns().GetConn(tid)
	if connection != nil {
		connection.SendToTerm(pkg)
	}
}
