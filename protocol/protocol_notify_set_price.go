package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type NotifySetPricePacket struct {
	Uuid   string
	Tid    uint64
	Serial uint32
	Result uint8
}

func (p *NotifySetPricePacket) Serialize() []byte {
	command := &Report.Command{
		Type:         Report.Command_CMT_REP_NOTIFY_SET_PRICE,
		Uuid:         p.Uuid,
		Tid:          p.Tid,
		SerialNumber: p.Serial,
		Paras: []*Report.Param{
			&Report.Param{
				Type:  Report.Param_UINT8,
				Npara: uint64(p.Result),
			},
		},
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseNotifySetPrice(buffer []byte) *NotifySetPricePacket {
	reader, _, _, tid := ParseHeader(buffer)
	serial := base.ReadDWord(reader)
	result, _ := reader.ReadByte()

	return &NotifySetPricePacket{
		Uuid:   conf.GetConf().Uuid,
		Tid:    tid,
		Serial: serial,
		Result: result,
	}
}
