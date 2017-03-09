package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
)

type ReqPinCodeNsqPacket struct {
	Tid    uint64
	Serial uint32
}

func (p *ReqPinCodeNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_PIN, p.Tid)
	base.WriteDWord(&writer, p.Serial)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqReqPinCode(cpid uint64, serial uint32) *ReqPinCodeNsqPacket {
	return &ReqPinCodeNsqPacket{
		Tid:    cpid,
		Serial: serial,
	}
}
