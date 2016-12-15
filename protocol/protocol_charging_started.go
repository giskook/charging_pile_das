package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type ChargingStartedPacket struct {
	Uuid          string
	Tid           uint64
	UserID        string
	TransactionID string
	Timestamp     uint64
	ElectriMeter  uint32
	StationID     uint32
	DBID          uint32
}

func (p *ChargingStartedPacket) Serialize() []byte {
	status := &Report.ChargingPileStatus{
		DasUuid:   p.Uuid,
		Cpid:      p.Tid,
		Status:    Report.ChargingPileStatus_CHARGING,
		Timestamp: p.Timestamp,
		Id:        p.DBID,
		StationId: p.StationID,
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

func ParseChargingStarted(buffer []byte, station_id uint32, id uint32) *ChargingStartedPacket {
	reader, _, _, tid := ParseHeader(buffer)
	userid_len, _ := reader.ReadByte()
	userid := base.ReadString(reader, userid_len)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)

	_time := base.ReadBcdTime(reader)
	electric_meter := base.ReadDWord(reader)

	return &ChargingStartedPacket{
		Uuid:          conf.GetConf().Uuid,
		Tid:           tid,
		UserID:        userid,
		TransactionID: transaction_id,
		ElectriMeter:  electric_meter,
		Timestamp:     _time,
		DBID:          id,
		StationID:     station_id,
	}

}
