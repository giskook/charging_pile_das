package base

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
}
