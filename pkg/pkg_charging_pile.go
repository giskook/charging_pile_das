package pkg

import (
	//"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/gotcp"
)

type Charging_Pile_Packet struct {
	Type   uint16
	Packet gotcp.Packet
}

func (this *Charging_Pile_Packet) Serialize() []byte {
	return this.Packet.Serialize()
	//	switch this.Type {
	//	case protocol.PROTOCOL_REQ_LOGIN:
	//		return this.Packet.(*protocol.LoginPacket).Serialize()
	//	case protocol.PROTOCOL_REQ_HEART:
	//		return this.Packet.(*protocol.HeartPacket).Serialize()
	//	case protocol.PROTOCOL_REQ_SETTING:
	//		return this.Packet.(*protocol.SettingPacket).Serialize()
	//	case protocol.PROTOCOL_REQ_PRICE:
	//		return this.Packet.(*protocol.PricePacket).Serialize()
	//	case protocol.PROTOCOL_REQ_TIME:
	//		return this.Packet.(*protocol.TimePacket).Serialize()
	//	case protocol.PROTOCOL_REP_CHARGING:
	//		return this.Packet.(*protocol.ChargingPacket).Serialize()
	//	case protocol.PROTOCOL_REP_STOP_CHARGING:
	//		return this.Packet.(*protocol.StopChargingPacket).Serialize()
	//	case protocol.PROTOCOL_REP_NSQ_NOTIFY_SET_PRICE:
	//		return this.Packet.(*protocol.NotifySetPricePacket).Serialize()
	//	case protocol.PROTOCOL_REP_CHARGING_STARTED:
	//		return this.Packet.(*protocol.ChargingStartedPacket).Serialize()
	//	case protocol.PROTOCOL_REP_CHARGING_DATA_UPLOAD:
	//		return this.Packet.(*protocol.ChargingUploadPacket).Serialize()
	//	case protocol.PROTOCOL_REQ_THREE_PHASE_MODE:
	//		return this.Packet.(*protocol.ThreePhaseModePacket).Serialize()
	//	case protocol.PROTOCOL_REP_PIN:
	//		return this.Packet.(*protocol.RepPinPacket).Serialize()
	//	case protocol.PROTOCOL_REP_OFFLINE_DATA:
	//		return this.Packet.(*protocol.UploadOfflineDataPacket).Serialize()
	//	case protocol.PROTOCOL_REP_CHARGING_STOPPED:
	//		return this.Packet.(*protocol.ChargingStoppedPacket).Serialize()
	//	case protocol.PROTOCOL_REP_GUN_STATUS:
	//		return this.Packet.(*protocol.RepGunStatusPacket).Serialize()
	//	}
	//
	//	return nil
}

func New_Charging_Pile_Packet(Type uint16, Packet gotcp.Packet) *Charging_Pile_Packet {
	return &Charging_Pile_Packet{
		Type:   Type,
		Packet: Packet,
	}
}
