package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type StopChargingPacket struct {
	Uuid      string
	Tid       uint64
	Result    uint8
	Timestamp uint64
}

func (p *StopChargingPacket) Serialize() []byte {
	paras := []*Report.Param{
		&Report.Param{
			Type:  Report.Param_UINT8,
			Npara: uint64(p.Result),
		},
		&Report.Param{
			Type:  Report.Param_UINT64,
			Npara: p.Timestamp,
		},
	}
	command := &Report.Command{
		Type:  Report.Command_CMT_REP_STOP_CHARGING,
		Uuid:  p.Uuid,
		Tid:   p.Tid,
		Paras: paras,
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseStopCharging(buffer []byte) *StopChargingPacket {
	reader, _, _, tid := ParseHeader(buffer)
	result, _ := reader.ReadByte()
	time_stamp := base.ReadBcdTime(reader)

	return &StopChargingPacket{
		Uuid:      conf.GetConf().Uuid,
		Tid:       tid,
		Result:    result,
		Timestamp: time_stamp,
	}
}
