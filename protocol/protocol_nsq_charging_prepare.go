package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/pb"
)

type ChargingPrepareNsqPacket struct {
	Tid     uint64
	Serial  uint32
	Userid  string
	PinCode uint16
}

func (p *ChargingPrepareNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	writer.WriteByte(PROTOCOL_START_FLAG)
	base.WriteWord(&writer, 0)
	base.WriteWord(&writer, PROTOCOL_REQ_CHARGING_PREPARE)
	base.WriteQuaWord(&writer, p.Tid)
	base.WriteDWord(&writer, p.Serial)
	base.WriteWord(&writer, p.PinCode)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqChargingPrepare(cpid uint64, serial uint32, param []*Report.Param) *ChargingPrepareNsqPacket {
	return &ChargingPrepareNsqPacket{
		Tid:     cpid,
		Serial:  serial,
		PinCode: uint16(param[0].Npara),
	}
}
