package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type UploadOfflineDataPacket struct {
	Uuid            string
	Tid             uint64
	StopReason      uint8
	EndMeterReading uint32
	UserID          string
	StopTime        uint32
	TransactionID   string
	Timestamp       uint64
	DBID            uint32
	StationID       uint32
}

func (p *UploadOfflineDataPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0, PROTOCOL_REP_UPLOAD_OFFLINE_PACAKGE_FEEDBACK, p.Tid)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func (p *UploadOfflineDataPacket) SerializeTss() []byte {
	status := &Report.ChargingPileStatus{
		DasUuid:            p.Uuid,
		Cpid:               p.Tid,
		Status:             uint32(PROTOCOL_CHARGE_PILE_STATUS_STOPPED),
		Timestamp:          p.Timestamp,
		EndMeterReading:    float32(p.EndMeterReading) / 10.0,
		ChargingDuration:   uint32(p.Timestamp),
		ChargingCapacity:   float32(p.EndMeterReading) / 10.0,
		CurrentOrderNumber: p.TransactionID,
		EndTime:            uint64(p.StopTime),
		Id:                 p.DBID,
		StationId:          p.StationID,
	}

	data, _ := proto.Marshal(status)

	return data
}

func ParseUploadOfflineData(buffer []byte, station_id uint32, id uint32) *UploadOfflineDataPacket {
	reader, _, _, tid := ParseHeader(buffer)
	stop_reason, _ := reader.ReadByte()
	end_meter_readging := base.ReadDWord(reader)
	userid := base.ReadString(reader, PROTOCOL_USERID_LEN)
	stop_time := base.ReadDWord(reader)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	time_stamp := base.ReadBcdTime(reader)

	return &UploadOfflineDataPacket{
		Uuid:            conf.GetConf().Uuid,
		Tid:             tid,
		StopReason:      stop_reason,
		EndMeterReading: end_meter_readging,
		UserID:          userid,
		StopTime:        stop_time,
		TransactionID:   transaction_id,
		Timestamp:       time_stamp,

		DBID:      id,
		StationID: station_id,
	}
}
