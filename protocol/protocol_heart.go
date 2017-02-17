package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
)

type HeartPacket struct {
	Uuid      string
	Tid       uint64
	Status    uint8
	Timestamp uint64
}

func (p *HeartPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_HEART, p.Tid)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseHeart(buffer []byte) *HeartPacket {
	reader, _, _, tid := ParseHeader(buffer)
	status, _ := reader.ReadByte()
	_time := base.ReadBcdTime(reader)

	return &HeartPacket{
		Uuid:      conf.GetConf().Uuid,
		Tid:       tid,
		Status:    status,
		Timestamp: _time,
	}

}
