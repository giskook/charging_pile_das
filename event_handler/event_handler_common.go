package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/gotcp"
	"log"
)

type Callback struct{}

func (this *Callback) OnConnect(c *gotcp.Conn) bool {
	checkinterval := conf.GetConf().Server.ConnCheckInterval
	readlimit := conf.GetConf().Server.ReadLimit
	writelimit := conf.GetConf().Server.WriteLimit
	config := &conn.ConnConfig{
		ConnCheckInterval: uint16(checkinterval),
		ReadLimit:         uint16(readlimit),
		WriteLimit:        uint16(writelimit),
	}
	connection := conn.NewConn(c, config)

	c.PutExtraData(connection)

	connection.Do()
	conn.NewConns().Add(connection)

	return true
}

func (this *Callback) OnClose(c *gotcp.Conn) {
	connection := c.GetExtraData().(*conn.Conn)
	log.Printf("close %d\n", connection.ID)
	connection.Close()
	conn.NewConns().Remove(connection)
	log.Println(conn.NewConns())
}

func (this *Callback) OnMessage(c *gotcp.Conn, p gotcp.Packet) bool {
	cpd_pkg := p.(*pkg.Charging_Pile_Packet)
	switch cpd_pkg.Type {
	case protocol.PROTOCOL_REQ_LOGIN:
		log.Println("on login")
		event_handler_login(c, cpd_pkg)
	case protocol.PROTOCOL_REQ_HEART:
		log.Println("on heart")
		event_handler_heart(c, cpd_pkg)
	case protocol.PROTOCOL_REQ_SETTING:
		log.Println("on setting")
		event_handler_setting(c, cpd_pkg)
	case protocol.PROTOCOL_REQ_TIME:
		log.Println("on time")
		event_handler_time(c, cpd_pkg)
	case protocol.PROTOCOL_REP_CHARGING:
		log.Println("on charging ")
		event_handler_rep_charging(c, cpd_pkg)
	case protocol.PROTOCOL_REP_STOP_CHARGING:
		log.Println("on stop charging ")
		event_handler_rep_stop_charging(c, cpd_pkg)
	case protocol.PROTOCOL_REP_NSQ_NOTIFY_SET_PRICE:
		log.Println("on nsq notify set price")
		event_handler_rep_notify_set_price(c, cpd_pkg)
	case protocol.PROTOCOL_REP_CHARGING_STARTED:
		log.Println("on charging started")
		event_handler_rep_charging_started(c, cpd_pkg)
	case protocol.PROTOCOL_REP_CHARGING_DATA_UPLOAD:
		log.Println("on charging upload")
		event_handler_rep_charging_upload(c, cpd_pkg)
	case protocol.PROTOCOL_REQ_PRICE:
		log.Println("on req price")
		event_handler_price(c, cpd_pkg)
	case protocol.PROTOCOL_REQ_THREE_PHASE_MODE:
		log.Println("on req three phase mode")
		event_handler_three_phase_mode(c, cpd_pkg)
	case protocol.PROTOCOL_REP_CHARGING_STOPPED:
		log.Println("on rep stop charging complete")
		event_handler_rep_charging_stopped(c, cpd_pkg)
	case protocol.PROTOCOL_REP_PIN:
		log.Println("on rep pin")
		event_handler_rep_pin(c, cpd_pkg)
	case protocol.PROTOCOL_REP_GUN_STATUS:
		log.Println("on rep gun status")
		event_handler_rep_gun_status(c, cpd_pkg)
	case protocol.PROTOCOL_REP_UPLOAD_OFFLINE_PACAKGE:
		log.Println("on rep upload offline ")
		event_handler_rep_upload_offline_package(c, cpd_pkg)
	}

	return true
}
