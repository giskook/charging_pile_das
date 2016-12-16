package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/pb"
)

const PROTOCOL_NSQ_LOGIN_LEN uint16 = PROTOCOL_COMMON_LEN + 1

type LoginNsqPacket struct {
	Tid    uint64
	Result uint8
}

func (p *LoginNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, PROTOCOL_NSQ_LOGIN_LEN,
		PROTOCOL_NSQ_LOGIN, p.Tid)
	writer.WriteByte(p.Result)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqLogin(cpid uint64, param []*Report.Param) *LoginNsqPacket {
	return &LoginNsqPacket{
		Tid:    cpid,
		Result: uint8(param[0].Npara),
	}

}
