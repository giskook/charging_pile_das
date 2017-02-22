package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type ChargingStoppedPacket struct {
	Uuid            string
	Tid             uint64
	StopReason      uint8
	EndMeterReading uint32
	UserID          string
	StopTime        uint32
	TransactionID   string
	Timestamp       uint64

	StartTime         uint32
	StartMeterReading uint32

	DBID      uint32
	StationID uint32
}

func (p *ChargingStoppedPacket) Serialize() []byte {
	status := &Report.ChargingPileStatus{
		DasUuid:          p.Uuid,
		Cpid:             p.Tid,
		Status:           uint32(PROTOCOL_CHARGING_PILE_IDLE),
		Timestamp:        p.Timestamp,
		EndMeterReading:  float32(p.EndMeterReading) / 10.0,
		ChargingCapacity: float32(p.EndMeterReading-p.StartMeterReading) / 10.0,
		ChargingDuration: uint32(p.Timestamp) - p.StartTime,
		Id:               p.DBID,
		StationId:        p.StationID,
	}

	data, _ := proto.Marshal(status)

	return data
}

func (p *ChargingStoppedPacket) SerializeWeChat() []byte {
	paras := []*Report.Param{
		&Report.Param{
			Type:    Report.Param_STRING,
			Strpara: p.TransactionID,
		},
	}
	command := &Report.Command{
		Type:  Report.Command_CMT_REP_CHARGING_STOPPED,
		Uuid:  p.Uuid,
		Tid:   p.Tid,
		Paras: paras,
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseChargingStopped(buffer []byte, station_id uint32, id uint32, start_time uint32, start_meter_reading uint32) *ChargingStoppedPacket {
	reader, _, _, tid := ParseHeader(buffer)
	stop_reason, _ := reader.ReadByte()
	end_meter_readging := base.ReadDWord(reader)
	userid := base.ReadString(reader, PROTOCOL_USERID_LEN)
	stop_time := base.ReadDWord(reader)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	time_stamp := base.ReadBcdTime(reader)

	return &ChargingStoppedPacket{
		Uuid:            conf.GetConf().Uuid,
		Tid:             tid,
		StopReason:      stop_reason,
		EndMeterReading: end_meter_readging,
		UserID:          userid,
		StopTime:        stop_time,
		TransactionID:   transaction_id,
		Timestamp:       time_stamp,

		StartTime:         start_time,
		StartMeterReading: start_meter_reading,

		DBID:      id,
		StationID: station_id,
	}
}
