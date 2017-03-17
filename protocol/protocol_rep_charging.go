package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type RepChargingPacket struct {
	Uuid          string
	Tid           uint64
	Result        uint8
	Pincode       string
	TransactionID string
	Timestamp     uint64
}

func (p *RepChargingPacket) Serialize() []byte {
	paras := []*Report.Param{
		&Report.Param{
			Type:  Report.Param_UINT8,
			Npara: uint64(p.Result),
		},
		&Report.Param{
			Type:    Report.Param_STRING,
			Strpara: p.TransactionID,
		},
	}

	command := &Report.Command{
		Type:  Report.Command_CMT_REP_CHARGING,
		Uuid:  p.Uuid,
		Tid:   p.Tid,
		Paras: paras,
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseRepCharging(buffer []byte, transcation_id string) *RepChargingPacket {
	reader, _, _, tid := ParseHeader(buffer)
	result, _ := reader.ReadByte()
	pin_code := base.ReadString(reader, PROTOCOL_PINCODE_LEN)
	time_stamp := base.ReadBcdTime(reader)

	return &RepChargingPacket{
		Uuid:          conf.GetConf().Uuid,
		Tid:           tid,
		Result:        result,
		Pincode:       pin_code,
		TransactionID: transcation_id,
		Timestamp:     time_stamp,
	}
}
