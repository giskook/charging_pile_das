package mq

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/protocol"
	"log"
	"strconv"
	"strings"
)

func event_handler_req_notify_set_price(tid uint64, serial uint32, param []*Report.Param) {
	cpids_str := param[0].Strpara
	cpids := strings.Split(cpids_str, ",")
	for _, cpid_str := range cpids {
		log.Println(cpid_str)
		cpid, _ := strconv.ParseUint(cpid_str, 10, 64)
		log.Println(cpid)
		pkg := protocol.ParseNsqNotifySetPrice(cpid, serial, param)
		connection := conn.NewConns().GetConn(cpid)
		if connection != nil {
			connection.SendToTerm(pkg)
		}
	}
}
