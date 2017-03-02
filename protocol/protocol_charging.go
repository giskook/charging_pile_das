package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
	"strconv"
)

type ChargingPacket struct {
	Uuid            string
	Tid             uint64
	StationID       uint32
	DBID            uint32
	Serial          uint32
	Userid          string
	TransactionID   string
	Result          uint8
	TimestampString string
	Timestamp       uint64
}

func (p *ChargingPacket) Serialize() []byte {
	command := &Report.Command{
		Type: Report.Command_CMT_REP_CHARGING,
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
			&Report.Param{
				Type:    Report.Param_STRING,
				Strpara: p.TimestampString,
			},
		},
	}

	data, _ := proto.Marshal(command)

	return data
}

func (p *ChargingPacket) SerializeTss() []byte {
	status := &Report.ChargingPileStatus{
		DasUuid:            p.Uuid,
		Cpid:               p.Tid,
		Status:             uint32(PROTOCOL_CHARGE_PILE_STATUS_CHARGING),
		Timestamp:          p.Timestamp,
		Id:                 p.DBID,
		StationId:          p.StationID,
		CurrentOrderNumber: p.TransactionID,
	}

	data, _ := proto.Marshal(status)

	return data
}

func ParseCharging(buffer []byte, station_id uint32, db_id uint32) *ChargingPacket {
	reader, _, _, tid := ParseHeader(buffer)
	serial := base.ReadDWord(reader)
	userid_len, _ := reader.ReadByte()
	userid := base.ReadString(reader, userid_len)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	result, _ := reader.ReadByte()
	time_stamp_string := base.ReadBcdString(reader, PROTOCOL_TIME_BCD_LEN)
	time_stamp, _ := strconv.ParseUint(time_stamp_string, 10, 64)

	return &ChargingPacket{
		Uuid:            conf.GetConf().Uuid,
		Tid:             tid,
		StationID:       station_id,
		DBID:            db_id,
		Serial:          serial,
		Userid:          userid,
		TransactionID:   transaction_id,
		Result:          result,
		TimestampString: time_stamp_string,
		Timestamp:       time_stamp,
	}
}
