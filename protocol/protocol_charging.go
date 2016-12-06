package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type ChargingPacket struct {
	Uuid          string
	Tid           uint64
	Serial        uint32
	Userid        string
	TransactionID string
	Result        uint8
}

func (p *ChargingPacket) Serialize() []byte {
	command := &Report.Command{
		Type: Report.Command_CMT_REP_CHARGING_PREPARE,
		Uuid: p.Uuid,
		Tid:  p.Tid,
		Paras: []*Report.Param{
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
				Npara: uint64(p.Result),
			},
		},
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseCharging(buffer []byte) *ChargingPacket {
	reader, _, _, tid := ParseHeader(buffer)
	serial := base.ReadDWord(reader)
	userid_len, _ := reader.ReadByte()
	userid := base.ReadString(reader, userid_len)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	result, _ := reader.ReadByte()

	return &ChargingPacket{
		Uuid:          conf.GetConf().Uuid,
		Tid:           tid,
		Serial:        serial,
		Userid:        userid,
		TransactionID: transaction_id,
		Result:        result,
	}
}
