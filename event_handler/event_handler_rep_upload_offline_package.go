package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
)

func event_handler_rep_upload_offline_package(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	upload_offline_data := p.Packet.(*protocol.UploadOfflineDataPacket)
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, upload_offline_data.SerializeTss())
	//server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicWeChat, rep_charging_stopped_pkg.SerializeWeChat())
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		connection.SendToTerm(p)
	}
}
