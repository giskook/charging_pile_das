package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
	"strconv"
)

type ChargingCost struct {
	StartTime        string
	EndTime          string
	Cost             uint32
	ChargingCapacity uint32
}

type StopChargingPacket struct {
	Uuid             string
	Tid              uint64
	Serial           uint32
	Userid           string
	TransactionID    string
	StopReason       uint8
	MeterReading     uint32
	ChargingDuration uint32
	ChargingCapacity uint32
	ChargingPrice    uint32
	CurrentTime      string
	Timestamp        uint64
	Costs            []*ChargingCost
}

func (p *StopChargingPacket) Serialize() []byte {
	paras := []*Report.Param{
		&Report.Param{
			Type:    Report.Param_STRING,
			Strpara: p.Userid,
		},
		&Report.Param{
			Type:    Report.Param_STRING,
			Strpara: p.TransactionID,
		},
		&Report.Param{
			Type:  Report.Param_UINT8,
			Npara: uint64(p.StopReason),
		},
		&Report.Param{
			Type:  Report.Param_UINT32,
			Npara: uint64(p.MeterReading),
		},
		&Report.Param{
			Type:  Report.Param_UINT32,
			Npara: uint64(p.ChargingDuration),
		},
		&Report.Param{
			Type:  Report.Param_UINT32,
			Npara: uint64(p.ChargingCapacity),
		},
		&Report.Param{
			Type:  Report.Param_UINT32,
			Npara: uint64(p.ChargingPrice),
		},
		&Report.Param{
			Type:    Report.Param_STRING,
			Strpara: p.CurrentTime,
		},
	}
	paras = append(paras, &Report.Param{
		Type:  Report.Param_UINT8,
		Npara: uint64(len(p.Costs)),
	})

	for _, cost := range p.Costs {
		cost_paras := []*Report.Param{
			&Report.Param{
				Type:    Report.Param_STRING,
				Strpara: cost.StartTime,
			},
			&Report.Param{
				Type:    Report.Param_STRING,
				Strpara: cost.EndTime,
			},
			&Report.Param{
				Type:  Report.Param_UINT32,
				Npara: uint64(cost.Cost),
			},
			&Report.Param{
				Type:  Report.Param_UINT32,
				Npara: uint64(cost.ChargingCapacity),
			},
		}
		paras = append(paras, cost_paras...)
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
	serial := base.ReadDWord(reader)
	userid_len, _ := reader.ReadByte()
	userid := base.ReadString(reader, userid_len)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	stop_reason, _ := reader.ReadByte()
	meter_reading := base.ReadDWord(reader)
	charging_duration := base.ReadDWord(reader)
	charging_capacity := base.ReadDWord(reader)
	charging_cost := base.ReadDWord(reader)

	var costs []*ChargingCost
	costs_count, _ := reader.ReadByte()
	for i := byte(0); i < costs_count; i++ {
		start_time := base.ReadBcdString(reader, PROTOCOL_TIME_BCD_LEN)
		end_time := base.ReadBcdString(reader, PROTOCOL_TIME_BCD_LEN)
		cost := base.ReadDWord(reader)
		capacity := base.ReadDWord(reader)
		costs = append(costs, &ChargingCost{
			StartTime:        start_time,
			EndTime:          end_time,
			Cost:             cost,
			ChargingCapacity: capacity,
		})

	}
	current_time := base.ReadBcdString(reader, PROTOCOL_TIME_BCD_LEN)
	time_stamp, _ := strconv.ParseUint(current_time, 10, 64)

	return &StopChargingPacket{
		Uuid:             conf.GetConf().Uuid,
		Tid:              tid,
		Serial:           serial,
		Userid:           userid,
		TransactionID:    transaction_id,
		StopReason:       stop_reason,
		MeterReading:     meter_reading,
		ChargingDuration: charging_duration,
		ChargingCapacity: charging_capacity,
		ChargingPrice:    charging_cost,
		CurrentTime:      current_time,
		Costs:            costs,
		Timestamp:        time_stamp,
	}
}
