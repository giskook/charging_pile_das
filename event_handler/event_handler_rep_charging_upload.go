package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
	"log"
)

func event_handler_rep_charging_upload(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	charging_upload := p.Packet.(*protocol.ChargingUploadPacket)
	log.Println(charging_upload)
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, p.Serialize())
}
