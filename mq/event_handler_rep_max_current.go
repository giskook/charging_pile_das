package mq

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/protocol"
)

func event_handler_rep_max_current(tid uint64, serial uint32, param []*Report.Param) {
	pkg := protocol.ParseNsqMaxCurrent(tid, param)
	connection := conn.NewConns().GetConn(tid)
	if connection != nil {
		connection.SendToTerm(pkg)
	}
}
