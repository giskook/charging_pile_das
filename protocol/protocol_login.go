package protocol

import (
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
}

func (p *LoginPacket) Serialize() []byte {
	command := &Report.Command{
		Type: 1,
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

	return &LoginPacket{
		Uuid:            conf.GetConf().Uuid,
		Tid:             tid,
		ProtocolVersion: protocol_version,
		HardwareVersion: hardware_version,
	}

}
