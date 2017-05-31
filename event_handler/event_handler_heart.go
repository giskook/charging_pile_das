package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
	"github.com/golang/protobuf/proto"
)

func event_handler_heart(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		connection.SendToTerm(p)
	}
	heart_pkg := p.Packet.(*protocol.HeartPacket)
	if connection.Charging_Pile.Status == 1 { // 1 means re on line
		connection.Charging_Pile.Status = 0
		if heart_pkg.Timestamp-connection.Charging_Pile.Timestamp > uint64(conf.GetConf().Server.SendHeartAfterLogin) {
			status := &Report.ChargingPileStatus{
				DasUuid:   conf.GetConf().Uuid,
				Cpid:      connection.ID,
				Id:        connection.Charging_Pile.DB_ID,
				StationId: connection.Charging_Pile.Station_ID,
				Status:    uint32(heart_pkg.Status),
				Timestamp: heart_pkg.Timestamp,
			}
			data, _ := proto.Marshal(status)
			server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, data)
		}

	} else {
		if uint64(connection.Charging_Pile.StopSendTime)+uint64(conf.GetConf().Limit.SendHeartThreshold) < heart_pkg.Timestamp { // in case of a heart pkg follow the stopped pkg
			status := &Report.ChargingPileStatus{
				DasUuid:   conf.GetConf().Uuid,
				Cpid:      connection.ID,
				Id:        connection.Charging_Pile.DB_ID,
				StationId: connection.Charging_Pile.Station_ID,
				Status:    uint32(heart_pkg.Status),
				Timestamp: heart_pkg.Timestamp,
			}
			data, _ := proto.Marshal(status)
			server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, data)
		}
	}
}
