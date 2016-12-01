package protocol

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

const CMD_SETTING_LEN uint16 = PROTOCOL_COMMON_LEN + 1

type SettingPacket struct {
	Uuid        string
	Tid         uint64
	SettingMode uint8
}

func (p *SettingPacket) Serialize() []byte {
	command := &Report.Command{
		Type: Report.Command_CMT_REQ_SETTING,
		Uuid: p.Uuid,
		Tid:  p.Tid,
		Paras: []*Report.Param{
			&Report.Param{
				Type:  Report.Param_UINT8,
				Npara: uint64(p.SettingMode),
			},
		},
	}

	data, _ := proto.Marshal(command)

	return data
}

func ParseSetting(buffer []byte) *SettingPacket {
	reader, _, _, tid := ParseHeader(buffer)
	setting_mode, _ := reader.ReadByte()

	return &SettingPacket{
		Uuid:        conf.GetConf().Uuid,
		Tid:         tid,
		SettingMode: setting_mode,
	}

}
