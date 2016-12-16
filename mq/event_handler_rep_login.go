package mq

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/golang/protobuf/proto"
	"log"
)

func event_handler_rep_login(socket *NsqSocket, tid uint64, serial uint32, param []*Report.Param) {
	pkg := protocol.ParseNsqLogin(tid, param)
	connection := conn.NewConns().GetConn(tid)
	log.Println(param)
	if connection != nil && param[0].Npara == 0 {
		connection.SendToTerm(pkg)
		connection.Status = conn.ConnSuccess
		connection.Charging_Pile.DB_ID = uint32(param[1].Npara)
		connection.Charging_Pile.Station_ID = uint32(param[2].Npara)
		status := &Report.ChargingPileStatus{
			DasUuid:   conf.GetConf().Uuid,
			Cpid:      tid,
			Id:        uint32(param[1].Npara),
			StationId: uint32(param[2].Npara),
			Status:    Report.ChargingPileStatus_ChargingPileStatusType(param[3].Npara),
			Timestamp: param[4].Npara,
		}
		data, _ := proto.Marshal(status)
		socket.Send(conf.GetConf().Nsq.Producer.TopicStatus, data)
		log.Println(param)
	}

}
