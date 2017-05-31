package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
	"time"
)

func event_handler_rep_charging_stopped(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {

	rep_charging_stopped_pkg := p.Packet.(*protocol.ChargingStoppedPacket)
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		if rep_charging_stopped_pkg.Timestamp-uint64(rep_charging_stopped_pkg.StartTime) > uint64(conf.GetConf().Limit.SendChargeStoppedThreshold) {
			server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, rep_charging_stopped_pkg.SerializeTss())
			connection.Charging_Pile.StopSendTime = uint32(rep_charging_stopped_pkg.Timestamp) // in case of after stopped. terminal send heart immidiate.

		} else {
			payload := rep_charging_stopped_pkg.SerializeTss()
			go send_delay(payload)
			connection.Charging_Pile.StopSendTime = uint32(rep_charging_stopped_pkg.Timestamp) + uint32(conf.GetConf().Limit.SendChargeStoppedDelay) // in case of after stopped. terminal send heart immidiate.

		}
	}
}

func send_delay(payload []byte) {
	for {
		select {
		case <-time.After(time.Duration(conf.GetConf().Limit.SendChargeStoppedDelay) * time.Second):
			server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, payload)

			return
		}
	}
}
