package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
)

type GetGunStatusPacket struct {
	Tid uint64
}

func (p *GetGunStatusPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_GUN_STATUS, p.Tid)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqGunStatus(cpid uint64) *GetGunStatusPacket {
	return &GetGunStatusPacket{
		Tid: cpid,
	}
}
