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
	case protocol.PROTOCOL_LOGIN:
		event_handler_login(c, cpd_pkg)
	}

	return true
}
