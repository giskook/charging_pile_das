package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
)

func event_handler_rep_charging_started(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, p.Serialize())

	rep_charging_started_pkg := p.Packet.(*protocol.ChargingStartedPacket)
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicWeChat, rep_charging_started_pkg.SerializeWeChat())
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		connection.Charging_Pile.StartTime = uint32(rep_charging_started_pkg.Timestamp)
		connection.Charging_Pile.StartMeterReading = rep_charging_started_pkg.StartMeterReading
	}
}
