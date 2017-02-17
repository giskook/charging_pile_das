package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
)

type ThreePhaseModePacket struct {
	Uuid      string
	Tid       uint64
	PhaseMode uint8
	AuthMode  uint8
	LockMode  uint8
}

func (p *ThreePhaseModePacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0, PROTOCOL_REP_THREE_PHASE_MODE, p.Tid)
	writer.WriteByte(p.PhaseMode)
	writer.WriteByte(p.AuthMode)
	writer.WriteByte(p.LockMode)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseThreePhaseMode(buffer []byte, phase_mode uint8, auth_mode uint8, lock_mode uint8) *ThreePhaseModePacket {
	_, _, _, tid := ParseHeader(buffer)

	return &ThreePhaseModePacket{
		Uuid:      conf.GetConf().Uuid,
		Tid:       tid,
		PhaseMode: phase_mode,
		AuthMode:  auth_mode,
		LockMode:  lock_mode,
	}
}
