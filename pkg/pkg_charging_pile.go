package pkg

import (
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/gotcp"
)

type Charging_Pile_Packet struct {
	Type   uint16
	Packet gotcp.Packet
}

func (this *Charging_Pile_Packet) Serialize() []byte {
	switch this.Type {
	case protocol.PROTOCOL_LOGIN:
		return this.Packet.(*protocol.LoginPacket).Serialize()
	}

	return nil
}

func New_Charging_Pile_Packet(Type uint16, Packet gotcp.Packet) *Charging_Pile_Packet {
	return &Charging_Pile_Packet{
		Type:   Type,
		Packet: Packet,
	}
}
