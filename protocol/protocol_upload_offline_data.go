package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type UploadOfflineDataPacket struct {
	Uuid          string
	Tid           uint64
	StopReason    uint8
	TransactionID string
	Cost          uint32
}

func (p *UploadOfflineDataPacket) Serialize() []byte {
	command := &Report.Command{
		Type: Report.Command_CMT_REP_OFFLINE_DATA,
		Uuid: p.Uuid,
		Tid:  p.Tid,
		Paras: []*Report.Param{
			&Report.Param{
				Type:  Report.Param_UINT8,
				Npara: uint64(p.StopReason),
			},
			&Report.Param{
				Type:    Report.Param_STRING,
				Strpara: p.TransactionID,
			},
			&Report.Param{
				Type:  Report.Param_UINT32,
				Npara: uint64(p.Cost),
			},
		},
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseUploadOfflineData(buffer []byte) *UploadOfflineDataPacket {
	reader, _, _, tid := ParseHeader(buffer)
	stop_reason, _ := reader.ReadByte()
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	cost := base.ReadDWord(reader)

	return &UploadOfflineDataPacket{
		Uuid:          conf.GetConf().Uuid,
		Tid:           tid,
		StopReason:    stop_reason,
		TransactionID: transaction_id,
		Cost:          cost,
	}
}
