package mq

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/protocol"
	"log"
)

func event_handler_rep_setting(tid uint64, serial uint32, param []*Report.Param) {
	log.Println("rep setting")
	pkg := protocol.ParseNsqSetting(tid, param)
	connection := conn.NewConns().GetConn(tid)
	if connection != nil {
		connection.SendToTerm(pkg)
	}

}
