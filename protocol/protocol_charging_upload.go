package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type ChargingUploadPacket struct {
	Uuid          string
	Tid           uint64
	MeterReading  uint32
	Power         uint16
	Status        uint8
	RealV         uint32
	RealI         uint32
	StationID     uint32
	DBID          uint32
	Timestamp     uint64
	TransactionID string
}

func (p *ChargingUploadPacket) Serialize() []byte {
	status := &Report.ChargingPileStatus{
		DasUuid:            p.Uuid,
		Cpid:               p.Tid,
		Status:             uint32(PROTOCOL_CHARGING_PILE_CHARGING),
		Timestamp:          p.Timestamp,
		Id:                 p.DBID,
		StationId:          p.StationID,
		RealTimeCurrent:    p.RealI,
		RealTimeVoltage:    p.RealV,
		CurrentOrderNumber: p.TransactionID,
		AmmeterNumber:      p.MeterReading,
	}

	data, _ := proto.Marshal(status)

	return data
}

func ParseChargingUpload(buffer []byte, station_id uint32, id uint32, transaction_id string) *ChargingUploadPacket {
	reader, _, _, tid := ParseHeader(buffer)
	meter_reading := base.ReadDWord(reader)
	power := base.ReadWord(reader)
	status, _ := reader.ReadByte()
	va := base.ReadWord(reader)
	vb := base.ReadWord(reader)
	vc := base.ReadWord(reader)
	ia := base.ReadWord(reader)
	ib := base.ReadWord(reader)
	ic := base.ReadWord(reader)
	_time := base.ReadBcdTime(reader)

	return &ChargingUploadPacket{
		Uuid:          conf.GetConf().Uuid,
		Tid:           tid,
		MeterReading:  meter_reading,
		Power:         power,
		Status:        status,
		RealV:         uint32(va+vb+vc) * 577, // 577 == 1.732/3*1000
		RealI:         uint32(ia+ib+ic) * 577,
		StationID:     station_id,
		DBID:          id,
		Timestamp:     _time,
		TransactionID: transaction_id,
	}
}
