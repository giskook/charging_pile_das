package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
)

func event_handler_login(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	connection := c.GetExtraData().(*conn.Conn)
	login_pkg := p.Packet.(*protocol.LoginPacket)
	connection.ID = login_pkg.Tid
	conn.NewConns().SetID(login_pkg.Tid, connection)
	if login_pkg.Status == 1 {
		connection.Charging_Pile.StartTime = login_pkg.StartTime
		connection.Charging_Pile.StartMeterReading = login_pkg.StartMeterReading
	}
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicAuth, p.Serialize())
}
