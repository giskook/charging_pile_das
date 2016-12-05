package protocol

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

const (
	PROTOCOL_MODE_REP_LEN uint16 = PROTOCOL_COMMON_LEN + 1
)

type ModePacket struct {
	Uuid string
	Tid  uint64
}

func (p *ModePacket) Serialize() []byte {
	command := &Report.Command{
		Type: Report.Command_CMT_REQ_MODE,
		Uuid: p.Uuid,
		Tid:  p.Tid,
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseMode(buffer []byte) *ModePacket {
	_, _, _, tid := ParseHeader(buffer)

	return &ModePacket{
		Uuid: conf.GetConf().Uuid,
		Tid:  tid,
	}
}
