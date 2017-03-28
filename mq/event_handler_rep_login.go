package mq

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/protocol"
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
		connection.Charging_Pile.PhaseMode = uint8(param[3].Npara)
		connection.Charging_Pile.AuthMode = uint8(param[4].Npara)
		connection.Charging_Pile.LockMode = uint8(param[5].Npara)

		if connection.Charging_Pile.Status == 1 {
			p := &protocol.ChargingStoppedPacket{
				Uuid:            conf.GetConf().Uuid,
				Tid:             connection.ID,
				EndMeterReading: connection.Charging_Pile.EndMeterReading,
				UserID:          connection.Charging_Pile.UserID,
				StopTime:        connection.Charging_Pile.StopTime,
				TransactionID:   connection.Charging_Pile.TransactionID,
				DBID:            connection.Charging_Pile.DB_ID,
				StationID:       connection.Charging_Pile.Station_ID,
				Timestamp:       connection.Charging_Pile.Timestamp,
			}
			log.Println("send offline")
			log.Println(p)
			socket.Send(conf.GetConf().Nsq.Producer.TopicStatus, p.SerializeTss())
		}
	}

}
