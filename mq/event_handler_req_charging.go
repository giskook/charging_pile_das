package mq

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/protocol"
	"log"
)

func event_handler_req_charging(tid uint64, serial uint32, param []*Report.Param) {
	log.Println("event_handler_req_charging")
	_pkg := protocol.ParseNsqCharging(tid, serial, param)
	log.Println(conn.NewConns())
	connection := conn.NewConns().GetConn(tid)
	if connection != nil {
		connection.SendToTerm(_pkg)
		connection.Charging_Pile.TransactionID = _pkg.TransactionID
	}
}
