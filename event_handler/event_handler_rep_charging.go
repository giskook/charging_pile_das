package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
	"log"
)

func event_handler_rep_charging(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicWeiXin, p.Serialize())

	charging_pkg := p.Packet.(*protocol.ChargingPacket)
	log.Println(charging_pkg)
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, charging_pkg.SerializeTss())
}
