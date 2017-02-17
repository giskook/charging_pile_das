package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
	"log"
)

func event_handler_rep_charging_stopped(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	log.Println("event_handler_rep_charging_stopped")
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, p.Serialize())
}
