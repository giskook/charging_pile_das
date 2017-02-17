package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
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
	WriteHeader(&writer, 0,
		PROTOCOL_REP_TIME, p.Tid)
	cur_time := time.Now()
	cur_time_str := cur_time.Format("20060102150405")
	base.WriteString(&writer, cur_time_str)
	base.WriteDWord(&writer, uint32(cur_time.Unix()))
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseTime(buffer []byte) *TimePacket {
	_, _, _, tid := ParseHeader(buffer)

	return &TimePacket{
		Uuid: conf.GetConf().Uuid,
		Tid:  tid,
	}
}
