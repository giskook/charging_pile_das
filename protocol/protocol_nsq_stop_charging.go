package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/pb"
)

type StopChargingNsqPacket struct {
	Tid           uint64
	Serial        uint32
	Userid        string
	TransactionID string
}

func (p *StopChargingNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_STOP_CHARGING, p.Tid)
	base.WriteDWord(&writer, p.Serial)
	writer.WriteByte(byte(len(p.Userid)))
	base.WriteString(&writer, p.Userid)
	base.WriteBcdString(&writer, p.TransactionID)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqStopCharging(cpid uint64, serial uint32, param []*Report.Param) *StopChargingNsqPacket {
	return &StopChargingNsqPacket{
		Tid:           cpid,
		Serial:        serial,
		Userid:        param[0].Strpara,
		TransactionID: param[1].Strpara,
	}
}
