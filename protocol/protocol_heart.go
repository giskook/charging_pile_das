package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
)

const (
	PROTOCOL_HEART_REP_LEN uint16 = PROTOCOL_COMMON_LEN
)

type HeartPacket struct {
	Uuid      string
	Tid       uint64
	Status    uint8
	Timestamp uint64
}

func (p *HeartPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, PROTOCOL_HEART_REP_LEN,
		PROTOCOL_REP_HEART, p.Tid)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

//func (p *HeartPacket) SerializeTss() []byte {
//	status := &Report.ChargingPileStatus{
//		DasUuid:   p.Uuid,
//		Cpid:      p.Tid,
//		Status:    Report.ChargingPileStatus_ChargingPileStatusType(p.Status),
//		Timestamp: p.Timestamp,
//	}
//
//	data, _ := proto.Marshal(status)
//
//	return data
//}

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
