package event_handler

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/gotcp"
)

func event_handler_three_phase_mode(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	connection := c.GetExtraData().(*conn.Conn)
	if connection != nil {
		connection.SendToTerm(p)
	}
}
