package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/pb"
)

const PROTOCOL_REP_MODE_LEN uint16 = PROTOCOL_COMMON_LEN + 1

type ModeNsqPacket struct {
	Tid  uint64
	Mode uint8
}

func (p *ModeNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	writer.WriteByte(PROTOCOL_START_FLAG)
	base.WriteWord(&writer, PROTOCOL_REP_MODE_LEN)
	base.WriteWord(&writer, PROTOCOL_REP_MODE)
	base.WriteQuaWord(&writer, p.Tid)
	writer.WriteByte(p.Mode)
	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqMode(cpid uint64, param []*Report.Param) *ModeNsqPacket {
	return &ModeNsqPacket{
		Tid:  cpid,
		Mode: uint8(param[0].Npara),
	}

}
