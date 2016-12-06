package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type ChargingPreparePacket struct {
	Uuid    string
	Tid     uint64
	Serial  uint32
	Result  uint8
	PinCode uint16
}

func (p *ChargingPreparePacket) Serialize() []byte {
	command := &Report.Command{
		Type: Report.Command_CMT_REP_CHARGING_PREPARE,
		Uuid: p.Uuid,
		Tid:  p.Tid,
		Paras: []*Report.Param{
			&Report.Param{
				Type:  Report.Param_UINT32,
				Npara: uint64(p.Serial),
			},
			&Report.Param{
				Type:  Report.Param_UINT8,
				Npara: uint64(p.Result),
			},
			&Report.Param{
				Type:  Report.Param_UINT16,
				Npara: uint64(p.PinCode),
			},
		},
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseChargingPrepare(buffer []byte) *ChargingPreparePacket {
	reader, _, _, tid := ParseHeader(buffer)
	serial := base.ReadDWord(reader)
	pincode := base.ReadWord(reader)
	result, _ := reader.ReadByte()

	return &ChargingPreparePacket{
		Uuid:    conf.GetConf().Uuid,
		Tid:     tid,
		Serial:  serial,
		Result:  result,
		PinCode: pincode,
	}
}
