package charging_pile_das

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/gotcp"
	"log"
	"sync"
)

type Charging_Pile_Protocol struct {
}

func (this *Charging_Pile_Protocol) ReadPacket(c *gotcp.Conn) (gotcp.Packet, error) {
	smconn := c.GetExtraData().(*conn.Conn)
	var once sync.Once
	once.Do(smconn.UpdateReadflag)

	buffer := smconn.GetBuffer()

	conn := c.GetRawConn()
	for {
		if smconn.ReadMore {
			data := make([]byte, 2048)
			readLengh, err := conn.Read(data)
			log.Printf("<IN>    %x\n", data[0:readLengh])
			if err != nil {
				return nil, err
			}

			if readLengh == 0 {
				return nil, gotcp.ErrConnClosing
			}
			buffer.Write(data[0:readLengh])
		}

		cmdid, pkglen := protocol.CheckProtocol(buffer)
		log.Printf("protocol id %d\n", cmdid)

		pkgbyte := make([]byte, pkglen)
		buffer.Read(pkgbyte)
		switch cmdid {
		case protocol.PROTOCOL_REQ_LOGIN:
			p := protocol.ParseLogin(pkgbyte)
			smconn.ReadMore = false
			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REQ_LOGIN, p), nil
		case protocol.PROTOCOL_REQ_HEART:
			p := protocol.ParseHeart(pkgbyte)
			smconn.ReadMore = false
			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REQ_HEART, p), nil

		case protocol.PROTOCOL_ILLEGAL:
			smconn.ReadMore = true
		case protocol.PROTOCOL_HALF_PACK:
			smconn.ReadMore = true
		}
	}

}
