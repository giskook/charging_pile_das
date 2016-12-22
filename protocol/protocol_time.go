package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
	"time"
)

const (
	PROTOCOL_TIME_REP_LEN uint16 = PROTOCOL_COMMON_LEN + 8
)

type TimePacket struct {
	Uuid string
	Tid  uint64
}

func (p *TimePacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, PROTOCOL_TIME_REP_LEN,
		PROTOCOL_REP_TIME, p.Tid)
	base.WriteQuaWord(&writer, uint64(time.Now().Unix()))
	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func (p *TimePacket) SerializeTss() []byte {
	status := &Report.ChargingPileStatus{
		DasUuid: p.Uuid,
		Cpid:    p.Tid,
	}

	data, _ := proto.Marshal(status)

	return data
}

func ParseTime(buffer []byte) *TimePacket {
	_, _, _, tid := ParseHeader(buffer)

	return &TimePacket{
		Uuid: conf.GetConf().Uuid,
		Tid:  tid,
	}
}
