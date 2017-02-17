package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/pb"
)

type SettingNsqPacket struct {
	Tid                     uint64
	Mode                    uint8
	BaudRateOrInterfaceType uint8
	WifiAndPasswd           []byte
}

func (p *SettingNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_SETTING, p.Tid)
	writer.WriteByte(p.Mode)
	if p.Mode == 1 || p.Mode == 3 {
		writer.WriteByte(p.BaudRateOrInterfaceType)
	} else if p.Mode == 2 {
		writer.Write(p.WifiAndPasswd)
	}
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqSetting(cpid uint64, param []*Report.Param) *SettingNsqPacket {
	if len(param) != 2 {
		return nil
	}

	if uint8(param[0].Npara) == 1 || uint8(param[0].Npara) == 3 {
		return &SettingNsqPacket{
			Tid:  cpid,
			Mode: uint8(param[0].Npara),
			BaudRateOrInterfaceType: uint8(param[1].Npara),
		}
	}

	return &SettingNsqPacket{
		Tid:           cpid,
		Mode:          uint8(param[0].Npara),
		WifiAndPasswd: param[1].Bpara,
	}

}
