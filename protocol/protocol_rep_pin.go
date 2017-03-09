package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type RepPinPacket struct {
	Uuid    string
	Tid     uint64
	Serial  uint32
	PinCode string
}

func (p *RepPinPacket) Serialize() []byte {
	command := &Report.Command{
		Type:         Report.Command_CMT_REP_PIN,
		Uuid:         p.Uuid,
		SerialNumber: p.Serial,
		Tid:          p.Tid,
		Paras: []*Report.Param{
			&Report.Param{
				Type:  Report.Param_UINT8,
				Npara: 0,
			},
			&Report.Param{
				Type:    Report.Param_STRING,
				Strpara: p.PinCode,
			},
		},
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseRepPin(buffer []byte, station_id uint32, id uint32) *RepPinPacket {
	reader, _, _, tid := ParseHeader(buffer)
	serial := base.ReadDWord(reader)
	pin_code := base.ReadString(reader, PROTOCOL_PINCODE_LEN)

	return &RepPinPacket{
		Uuid:    conf.GetConf().Uuid,
		Tid:     tid,
		Serial:  serial,
		PinCode: pin_code,
	}
}
