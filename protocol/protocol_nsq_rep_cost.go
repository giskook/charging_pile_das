package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/pb"
)

type RepCostNsqPacket struct {
	Tid  uint64
	Cost uint32
}

func (p *RepCostNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_CHARGING_COST, p.Tid)
	base.WriteDWord(&writer, p.Cost)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqRepCost(cpid uint64, param []*Report.Param) *RepCostNsqPacket {
	return &RepCostNsqPacket{
		Tid:  cpid,
		Cost: uint32(param[0].Npara),
	}
}
