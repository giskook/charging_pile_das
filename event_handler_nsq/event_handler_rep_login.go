package event_handler_nsq

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
)

func event_handler_rep_login(tid uint64, serial uint32, param []*Report.Param) {
	conn.NewConns().GetConn(tid)
}
