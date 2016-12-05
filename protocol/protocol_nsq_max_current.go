package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/pb"
)

const PROTOCOL_REP_MAX_CURRENT_LEN uint16 = PROTOCOL_COMMON_LEN + 1

type MaxCurrentNsqPacket struct {
	Tid        uint64
	MaxCurrent uint8
}

func (p *MaxCurrentNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	writer.WriteByte(PROTOCOL_START_FLAG)
	base.WriteWord(&writer, PROTOCOL_REP_MAX_CURRENT_LEN)
	base.WriteWord(&writer, PROTOCOL_REP_MAX_CURRENT)
	base.WriteQuaWord(&writer, p.Tid)
	writer.WriteByte(p.MaxCurrent)
	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqMaxCurrent(cpid uint64, param []*Report.Param) *MaxCurrentNsqPacket {
	return &MaxCurrentNsqPacket{
		Tid:        cpid,
		MaxCurrent: uint8(param[0].Npara),
	}

}
