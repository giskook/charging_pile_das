package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

const (
	PROTOCOL_MAX_CURRENT_REP_LEN uint16 = PROTOCOL_COMMON_LEN + 1
)

type MaxCurrentPacket struct {
	Uuid       string
	Tid        uint64
	MaxCurrent uint8
}

func (p *MaxCurrentPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, PROTOCOL_MAX_CURRENT_REP_LEN,
		PROTOCOL_REP_MAX_CURRENT, p.Tid)
	writer.WriteByte(p.MaxCurrent)
	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func (p *MaxCurrentPacket) SerializeTss() []byte {
	command := &Report.Command{
		Type: Report.Command_CMT_REQ_MAX_CURRENT,
		Uuid: p.Uuid,
		Tid:  p.Tid,
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseMaxCurrent(buffer []byte) *MaxCurrentPacket {
	_, _, _, tid := ParseHeader(buffer)

	return &MaxCurrentPacket{
		Uuid: conf.GetConf().Uuid,
		Tid:  tid,
	}
}
