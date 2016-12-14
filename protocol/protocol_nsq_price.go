package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/pb"
)

type Price struct {
	Start_hour      uint8
	Start_min       uint8
	End_hour        uint8
	End_min         uint8
	Elec_unit_price uint16
	Service_price   uint16
}

type PriceNsqPacket struct {
	Tid    uint64
	Prices []*Price
}

func (p *PriceNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_PRICE, p.Tid)
	for _, price := range p.Prices {
		writer.WriteByte(price.Start_hour)
		writer.WriteByte(price.Start_min)
		writer.WriteByte(price.End_hour)
		writer.WriteByte(price.End_min)
		base.WriteWord(&writer, price.Elec_unit_price)
		base.WriteWord(&writer, price.Service_price)
		//writer.WriteByte(p.Prices[i].Start_hour)
		//writer.WriteByte(p.Prices[i].Start_min)
		//writer.WriteByte(p.Prices[i].End_hour)
		//writer.WriteByte(p.Prices[i].End_min)
		//base.WriteWord(&writer, p.Prices[i].Elec_unit_price)
		//base.WriteWord(&writer, p.Prices[i].Service_price)
	}
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqPrice(cpid uint64, param []*Report.Param) *PriceNsqPacket {
	var prices []*Price
	for i := 0; i < len(param)/6; i++ {
		prices = append(prices, &Price{
			Start_hour:      uint8(param[i*6].Npara),
			Start_min:       uint8(param[i*6+1].Npara),
			End_hour:        uint8(param[i*6+2].Npara),
			End_min:         uint8(param[i*6+3].Npara),
			Elec_unit_price: uint16(param[i*6+4].Npara),
			Service_price:   uint16(param[i*6+5].Npara),
		})
	}

	return &PriceNsqPacket{
		Tid:    cpid,
		Prices: prices,
	}
}
