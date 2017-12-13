package base

const (
	STATUS_IDLE     uint8 = 0
	STATUS_CHARGING uint8 = 1
)

type Charging_Pile struct {
	ID              uint64
	BoxVersion      byte
	ProtocolVersion byte
	DB_ID           uint32
	Station_ID      uint32
	PhaseMode       uint8
	AuthMode        uint8
	LockMode        uint8

	UserID            string
	TransactionID     string
	StartTime         uint32
	StartMeterReading uint32

	Status          uint8 // for offline 1 means re online 0 mean normal
	EndMeterReading uint32
	StopTime        uint32
	Timestamp       uint64

	StopSendTime uint32

	StatusEx uint8
}
