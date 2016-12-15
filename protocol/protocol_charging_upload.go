package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type ChargingUploadPacket struct {
	Uuid             string
	Tid              uint64
	UserID           string
	ChargingDuration uint32
	ChargingCapacity uint32
	ChargingPrice    uint32
	MeterReading     uint32
	RealTimeCurrent  uint32
	RealTimeVoltage  uint32
	StationID        uint32
	DBID             uint32
	Timestamp        uint64
}

func (p *ChargingUploadPacket) Serialize() []byte {
	status := &Report.ChargingPileStatus{
		DasUuid:          p.Uuid,
		Cpid:             p.Tid,
		Status:           Report.ChargingPileStatus_CHARGING,
		Timestamp:        p.Timestamp,
		Id:               p.DBID,
		StationId:        p.StationID,
		ChargingDuration: p.ChargingDuration,
		ChargingCapacity: p.ChargingCapacity,
		ChargingPrice:    p.ChargingPrice,
		RealTimeCurrent:  p.RealTimeCurrent,
		RealTimeVoltage:  p.RealTimeVoltage,
	}

	data, _ := proto.Marshal(status)

	return data
}

//func (p *HeartPacket) SerializeTss() []byte {
//	status := &Report.ChargingPileStatus{
//		DasUuid:   p.Uuid,
//		Cpid:      p.Tid,
//		Status:    Report.ChargingPileStatus_ChargingPileStatusType(p.Status),
//		Timestamp: p.Timestamp,
//	}
//
//	data, _ := proto.Marshal(status)
//
//	return data
//}

func ParseChargingUpload(buffer []byte, station_id uint32, id uint32) *ChargingUploadPacket {
	reader, _, _, tid := ParseHeader(buffer)
	userid_len, _ := reader.ReadByte()
	userid := base.ReadString(reader, userid_len)
	base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	charging_duration := base.ReadDWord(reader)
	charging_capacity := base.ReadDWord(reader)
	charging_price := base.ReadDWord(reader)
	meter_reading := base.ReadDWord(reader)
	realtime_elec := base.ReadDWord(reader)
	realtime_voltage := base.ReadDWord(reader)
	_time := base.ReadBcdTime(reader)

	return &ChargingUploadPacket{
		Uuid:             conf.GetConf().Uuid,
		Tid:              tid,
		UserID:           userid,
		ChargingDuration: charging_duration,
		ChargingCapacity: charging_capacity,
		ChargingPrice:    charging_price,
		MeterReading:     meter_reading,
		RealTimeCurrent:  realtime_elec,
		RealTimeVoltage:  realtime_voltage,
		StationID:        station_id,
		DBID:             id,
		Timestamp:        _time,
	}
}
