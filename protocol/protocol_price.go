package protocol

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type PricePacket struct {
	Uuid      string
	Tid       uint64
	StationID uint32
}

func (p *PricePacket) Serialize() []byte {
	command := &Report.Command{
		Type: Report.Command_CMT_REQ_PRICE,
		Uuid: p.Uuid,
		Tid:  p.Tid,
		Paras: []*Report.Param{
			&Report.Param{
				Type:  Report.Param_UINT32,
				Npara: uint64(p.StationID),
			},
		},
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParsePrice(buffer []byte, station_id uint32) *PricePacket {
	_, _, _, tid := ParseHeader(buffer)

	return &PricePacket{
		Uuid:      conf.GetConf().Uuid,
		Tid:       tid,
		StationID: station_id,
	}
}
