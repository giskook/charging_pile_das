package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/pb"
)

type NsqNotifySetPricePacket struct {
	Tid    uint64
	Serial uint32
	Prices []*Price
	Cpids  string
}

func (p *NsqNotifySetPricePacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_NSQ_NOTIFY_SET_PRICE, p.Tid)
	base.WriteDWord(&writer, p.Serial)
	writer.WriteByte(byte(len(p.Prices)))
	for _, price := range p.Prices {
		writer.WriteByte(price.Start_hour)
		writer.WriteByte(price.Start_min)
		writer.WriteByte(price.End_hour)
		writer.WriteByte(price.End_min)
		base.WriteWord(&writer, price.Elec_unit_price)
		base.WriteWord(&writer, price.Service_price)
	}
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqNotifySetPrice(cpid uint64, serial uint32, param []*Report.Param) *NsqNotifySetPricePacket {
	var prices []*Price
	for i := 0; i < len(param)/6; i++ {
		prices = append(prices, &Price{
			Start_hour:      uint8(param[1+i*6].Npara),
			Start_min:       uint8(param[1+i*6+1].Npara),
			End_hour:        uint8(param[1+i*6+2].Npara),
			End_min:         uint8(param[1+i*6+3].Npara),
			Elec_unit_price: uint16(param[1+i*6+4].Npara),
			Service_price:   uint16(param[1+i*6+5].Npara),
		})
	}

	return &NsqNotifySetPricePacket{
		Tid:    cpid,
		Serial: serial,
		Prices: prices,
	}
}
