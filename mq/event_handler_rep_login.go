package mq

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

func event_handler_rep_login(socket *NsqSocket, tid uint64, serial uint32, param []*Report.Param) {
	pkg := protocol.ParseNsqLogin(tid, param)
	connection := conn.NewConns().GetConn(tid)
	if connection != nil {
		connection.SendToTerm(pkg)
		connection.Status = conn.ConnSuccess
		log.Println(socket)
		status := &Report.ChargingPileStatus{
			DasUuid:   conf.GetConf().Uuid,
			Cpid:      tid,
			Status:    Report.ChargingPileStatus_ChargingPileStatusType(param[1].Npara),
			Timestamp: uint64(time.Now().Unix()),
		}
		data, _ := proto.Marshal(status)
		socket.Send(conf.GetConf().Nsq.Producer.TopicStatus, data)
	}

}
