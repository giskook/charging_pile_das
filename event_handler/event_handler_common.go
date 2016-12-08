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
	connection.Close()
	conn.NewConns().Remove(connection)
	log.Println(conn.NewConns())
}

func (this *Callback) OnMessage(c *gotcp.Conn, p gotcp.Packet) bool {
	cpd_pkg := p.(*pkg.Charging_Pile_Packet)
	switch cpd_pkg.Type {
	case protocol.PROTOCOL_REQ_LOGIN:
		event_handler_login(c, cpd_pkg)
	case protocol.PROTOCOL_REQ_HEART:
		event_handler_heart(c, cpd_pkg)
	case protocol.PROTOCOL_REQ_SETTING:
		log.Println("on setting")
		event_handler_setting(c, cpd_pkg)
	case protocol.PROTOCOL_REQ_TIME:
		log.Println("on time")
		event_handler_time(c, cpd_pkg)
	case protocol.PROTOCOL_REQ_MODE:
		log.Println("on mode")
		event_handler_mode(c, cpd_pkg)
	case protocol.PROTOCOL_REQ_MAX_CURRENT:
		log.Println("on max current")
		event_handler_max_current(c, cpd_pkg)
	case protocol.PROTOCOL_REP_CHARGING_PREPARE:
		log.Println("on charging prepare")
		event_handler_rep_charging_prepare(c, cpd_pkg)
	case protocol.PROTOCOL_REP_CHARGING:
		log.Println("on charging ")
		event_handler_rep_charging(c, cpd_pkg)
	case protocol.PROTOCOL_REP_STOP_CHARGING:
		log.Println("on stop charging ")
		event_handler_rep_stop_charging(c, cpd_pkg)
	case protocol.PROTOCOL_REP_NSQ_NOTIFY_SET_PRICE:
		log.Println("on nsq notify set price")
		event_handler_rep_notify_set_price(c, cpd_pkg)

	}

	return true
}
