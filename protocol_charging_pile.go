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

	tcp_conn := c.GetRawConn()
	for {
		if smconn.ReadMore {
			data := make([]byte, 2048)
			readLengh, err := tcp_conn.Read(data)
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
		if smconn.Status == conn.ConnUnauth {
			if cmdid != protocol.PROTOCOL_REQ_LOGIN &&
				cmdid != protocol.PROTOCOL_HALF_PACK &&
				cmdid != protocol.PROTOCOL_ILLEGAL {
				cmdid = protocol.PROTOCOL_SWALLOW
			}
		}

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
		case protocol.PROTOCOL_REQ_SETTING:
			p := protocol.ParseSetting(pkgbyte)
			smconn.ReadMore = false

			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REQ_SETTING, p), nil

		case protocol.PROTOCOL_REQ_PRICE:
			p := protocol.ParsePrice(pkgbyte)
			smconn.ReadMore = false

			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REQ_PRICE, p), nil
		case protocol.PROTOCOL_REQ_TIME:
			p := protocol.ParseTime(pkgbyte)
			smconn.ReadMore = false

			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REQ_TIME, p), nil
		case protocol.PROTOCOL_REP_STOP_CHARGING:
			p := protocol.ParseStopCharging(pkgbyte)
			smconn.ReadMore = false

			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REP_STOP_CHARGING, p), nil
		case protocol.PROTOCOL_REP_NSQ_NOTIFY_SET_PRICE:
			p := protocol.ParseNotifySetPrice(pkgbyte)
			smconn.ReadMore = false

			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REP_NSQ_NOTIFY_SET_PRICE, p), nil

		case protocol.PROTOCOL_REP_CHARGING_STARTED:
			p := protocol.ParseChargingStarted(pkgbyte, smconn.Charging_Pile.Station_ID, smconn.Charging_Pile.DB_ID)
			smconn.ReadMore = false

			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REP_CHARGING_STARTED, p), nil

		case protocol.PROTOCOL_REP_CHARGING_DATA_UPLOAD:
			p := protocol.ParseChargingUpload(pkgbyte, smconn.Charging_Pile.Station_ID, smconn.Charging_Pile.DB_ID, smconn.Charging_Pile.TransactionID)
			smconn.ReadMore = false

			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REP_CHARGING_DATA_UPLOAD, p), nil
		case protocol.PROTOCOL_REQ_THREE_PHASE_MODE:
			p := protocol.ParseThreePhaseMode(pkgbyte, smconn.Charging_Pile.PhaseMode, smconn.Charging_Pile.AuthMode, smconn.Charging_Pile.LockMode)
			smconn.ReadMore = false
			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REQ_THREE_PHASE_MODE, p), nil
		case protocol.PROTOCOL_REP_PIN:
			p := protocol.ParseRepPin(pkgbyte, smconn.Charging_Pile.Station_ID, smconn.Charging_Pile.DB_ID)

			smconn.ReadMore = false
			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REP_PIN, p), nil
		case protocol.PROTOCOL_REP_OFFLINE_DATA:
			p := protocol.ParseUploadOfflineData(pkgbyte)
			smconn.ReadMore = false

			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REP_OFFLINE_DATA, p), nil
		case protocol.PROTOCOL_REP_GUN_STATUS:
			p := protocol.ParseRepGunStatus(pkgbyte)
			smconn.ReadMore = false

			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REP_GUN_STATUS, p), nil
		case protocol.PROTOCOL_REP_CHARGING:
			p := protocol.ParseRepCharging(pkgbyte)
			smconn.ReadMore = false

			return pkg.New_Charging_Pile_Packet(protocol.PROTOCOL_REP_CHARGING, p), nil

		case protocol.PROTOCOL_ILLEGAL:
			smconn.ReadMore = true
		case protocol.PROTOCOL_HALF_PACK:
			smconn.ReadMore = true
		case protocol.PROTOCOL_SWALLOW:
			smconn.ReadMore = false
		}
	}

}
