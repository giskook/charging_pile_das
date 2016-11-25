package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
)

func event_handler_heart(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		connection.SendToTerm(p)
	}
	heart_pkg := p.Packet.(*protocol.HeartPacket)
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, heart_pkg.SerializeTss())
}
