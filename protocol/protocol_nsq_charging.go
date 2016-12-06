package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/pb"
)

type ChargingNsqPacket struct {
	Tid              uint64
	Serial           uint32
	Userid           string
	PinCode          uint16
	TransactionID    string
	TranscationValue uint32
}

func (p *ChargingNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	writer.WriteByte(PROTOCOL_START_FLAG)
	base.WriteWord(&writer, 0)
	base.WriteWord(&writer, PROTOCOL_REQ_CHARGING_PREPARE)
	base.WriteQuaWord(&writer, p.Tid)
	base.WriteDWord(&writer, p.Serial)
	base.WriteWord(&writer, p.PinCode)
	writer.WriteByte(byte(len(p.Userid)))
	base.WriteString(&writer, p.Userid)
	base.WriteBcdString(&writer, p.TransactionID)
	base.WriteDWord(&writer, p.TranscationValue)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqCharging(cpid uint64, serial uint32, param []*Report.Param) *ChargingNsqPacket {
	return &ChargingNsqPacket{
		Tid:              cpid,
		Serial:           serial,
		Userid:           param[0].Strpara,
		PinCode:          uint16(param[1].Npara),
		TransactionID:    param[2].Strpara,
		TranscationValue: uint32(param[3].Npara),
	}
}
