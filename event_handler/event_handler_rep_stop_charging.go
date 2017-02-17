package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
	"github.com/golang/protobuf/proto"
	"log"
)

func event_handler_rep_stop_charging(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	stop_charging_pkg := p.Packet.(*protocol.StopChargingPacket)
	log.Println(stop_charging_pkg)
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicWeiXin, p.Serialize())
	connection := c.GetExtraData().(*conn.Conn)
	//server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, heart_pkg.SerializeTss())
	status := &Report.ChargingPileStatus{
		DasUuid:   conf.GetConf().Uuid,
		Cpid:      connection.ID,
		Id:        connection.Charging_Pile.DB_ID,
		StationId: connection.Charging_Pile.Station_ID,
		Status:    uint32(protocol.PROTOCOL_CHARGING_PILE_IDLE),
		Timestamp: stop_charging_pkg.Timestamp,
	}
	data, _ := proto.Marshal(status)
	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicStatus, data)
}
