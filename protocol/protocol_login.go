package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

const CMD_LOGIN_LEN uint16 = 0x12

type LoginPacket struct {
	Uuid            string
	Tid             uint64
	ProtocolVersion uint8
	HardwareVersion uint8
	PinCode         string
	Status          uint8
	UserID          string
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

func ParseLogin(buffer []byte) *LoginPacket {
	reader, _, _, tid := ParseHeader(buffer)
	protocol_version, _ := reader.ReadByte()
	hardware_version, _ := reader.ReadByte()
	pin_code := base.ReadString(reader, PROTOCOL_PINCODE_LEN)
	status, _ := reader.ReadByte()
	user_id := base.ReadString(reader, PROTOCOL_USERID_LEN)
	transcation_id := base.ReadString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	time_stamp := base.ReadBcdTime(reader)

	return &LoginPacket{
		Uuid:            conf.GetConf().Uuid,
		Tid:             tid,
		ProtocolVersion: protocol_version,
		HardwareVersion: hardware_version,
		PinCode:         pin_code,
		Status:          status,
		UserID:          user_id,
		TransactionID:   transcation_id,
		Timestamp:       time_stamp,
	}
}
