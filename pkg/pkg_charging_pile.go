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
	case protocol.PROTOCOL_REQ_LOGIN:
		return this.Packet.(*protocol.LoginPacket).Serialize()
	case protocol.PROTOCOL_REQ_HEART:
		return this.Packet.(*protocol.HeartPacket).Serialize()
	case protocol.PROTOCOL_REQ_SETTING:
		return this.Packet.(*protocol.SettingPacket).Serialize()
	case protocol.PROTOCOL_REQ_PRICE:
		return this.Packet.(*protocol.PricePacket).Serialize()
	case protocol.PROTOCOL_REQ_TIME:
		return this.Packet.(*protocol.TimePacket).Serialize()
	case protocol.PROTOCOL_REQ_MODE:
		return this.Packet.(*protocol.ModePacket).Serialize()
	case protocol.PROTOCOL_REQ_MAX_CURRENT:
		return this.Packet.(*protocol.MaxCurrentPacket).Serialize()

	}

	return nil
}

func New_Charging_Pile_Packet(Type uint16, Packet gotcp.Packet) *Charging_Pile_Packet {
	return &Charging_Pile_Packet{
		Type:   Type,
		Packet: Packet,
	}
}
