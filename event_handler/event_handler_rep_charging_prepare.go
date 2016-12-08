package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
	"log"
)

func event_handler_rep_charging_prepare(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	log.Println("rep_charging_prepare")
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicWeiXin, p.Serialize())
}
