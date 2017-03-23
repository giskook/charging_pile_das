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

	rep_charging_started_pkg := p.Packet.(*protocol.ChargingStartedPacket)
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, rep_charging_started_pkg.SerializeTss())
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicWeChat, rep_charging_started_pkg.SerializeWeChat())
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		//connection.SendToTerm(p)
		connection.Charging_Pile.StartTime = uint32(rep_charging_started_pkg.Timestamp)
		connection.Charging_Pile.StartMeterReading = rep_charging_started_pkg.StartMeterReading
		connection.Charging_Pile.TransactionID = rep_charging_started_pkg.TransactionID
	}
}
