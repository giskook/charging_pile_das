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

func event_handler_rep_charging(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicWeChat, p.Serialize())

	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		rep_charging_pkg := p.Packet.(*protocol.RepChargingPacket)
		status := &Report.ChargingPileStatus{
			DasUuid:   conf.GetConf().Uuid,
			Cpid:      connection.ID,
			Id:        connection.Charging_Pile.DB_ID,
			StationId: connection.Charging_Pile.Station_ID,
			Status:    0,
			Timestamp: rep_charging_pkg.Timestamp,
		}

		data, _ := proto.Marshal(status)
		server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, data)
	}
}
