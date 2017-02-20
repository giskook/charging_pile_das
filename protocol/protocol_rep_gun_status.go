package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type RepGunStatusPacket struct {
	Uuid   string
	Tid    uint64
	Status uint8
}

func (p *RepGunStatusPacket) Serialize() []byte {
	command := &Report.Command{
		Type: Report.Command_CMT_REP_GET_GUN_STATUS,
		Uuid: p.Uuid,
		Tid:  p.Tid,
		Paras: []*Report.Param{
			&Report.Param{
				Type:  Report.Param_UINT8,
				Npara: uint64(p.Status),
			},
		},
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseRepGunStatus(buffer []byte) *RepGunStatusPacket {
	reader, _, _, tid := ParseHeader(buffer)
	status, _ := reader.ReadByte()
	base.ReadBcdTime(reader)

	return &RepGunStatusPacket{
		Uuid:   conf.GetConf().Uuid,
		Tid:    tid,
		Status: status,
	}
}
