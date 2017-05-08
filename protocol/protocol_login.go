package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type LoginPacket struct {
	Uuid            string
	Tid             uint64
	ProtocolVersion uint8
	HardwareVersion uint8
	Status          uint8
	StopReason      uint8
	EndMeterReading uint32
	UserID          string
	StopTime        uint32
	TransactionID   string
	Timestamp       uint64
}

func (p *LoginPacket) Serialize() []byte {
	command := &Report.Command{
		Type: Report.Command_CMT_REQ_LOGIN,
		Uuid: p.Uuid,
		Tid:  p.Tid,
		Paras: []*Report.Param{
			&Report.Param{
				Type:  Report.Param_UINT8,
				Npara: uint64(p.ProtocolVersion),
			},
			&Report.Param{
				Type:  Report.Param_UINT8,
				Npara: uint64(p.HardwareVersion),
			},
		},
	}

	data, _ := proto.Marshal(command)

	return data
}

//func (p *LoginPacket) SerialTss() []byte {
//	status := &Report.ChargingPileStatus{
//		DasUuid:            p.Uuid,
//		Cpid:               p.Tid,
//		Status:             uint32(PROTOCOL_CHARGE_PILE_STATUS_STOPPED),
//		Timestamp:          p.Timestamp,
//		EndMeterReading:    float32(p.EndMeterReading) / 10.0,
//		ChargingDuration:   uint32(p.StopTime),
//		ChargingCapacity:   float32(p.EndMeterReading) / 10.0,
//		CurrentOrderNumber: p.TransactionID,
//		EndTime:            uint64(p.StopTime),
//		Id:                 p.DBID,
//		StationId:          p.StationID,
//	}
//
//	data, _ := proto.Marshal(status)
//
//	return data
//}

func ParseLogin(buffer []byte) *LoginPacket {
	reader, _, _, tid := ParseHeader(buffer)
	protocol_version, _ := reader.ReadByte()
	hardware_version, _ := reader.ReadByte()
	status, _ := reader.ReadByte()
	stop_reason, _ := reader.ReadByte()
	end_meter_reading := base.ReadDWord(reader)
	user_id := base.ReadString(reader, PROTOCOL_USERID_LEN)
	stop_time := base.ReadDWord(reader)
	transcation_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	time_stamp := base.ReadBcdTime(reader)

	return &LoginPacket{
		Uuid:            conf.GetConf().Uuid,
		Tid:             tid,
		ProtocolVersion: protocol_version,
		HardwareVersion: hardware_version,
		Status:          status,
		StopReason:      stop_reason,
		EndMeterReading: end_meter_reading,
		UserID:          user_id,
		StopTime:        stop_time,
		TransactionID:   transcation_id,
		Timestamp:       time_stamp,
	}
}
