package conn

import (
	"sync/atomic"
)

type Conns struct {
	connsindex map[uint32]*Conn
	connsuid   map[uint64]*Conn
	index      uint32
}

var connsInstance *Conns

func NewConns() *Conns {
	if connsInstance == nil {
		connsInstance = &Conns{
			connsindex: make(map[uint32]*Conn),
			connsuid:   make(map[uint64]*Conn),
			index:      0,
		}
	}

	return connsInstance
}

func (cs *Conns) Add(conn *Conn) {
	conn.index = atomic.AddUint32(&cs.index, 1)
	cs.connsindex[conn.index] = conn
}

func (cs *Conns) SetID(gatewayid uint64, conn *Conn) {
	cs.connsuid[gatewayid] = conn
}

func (cs *Conns) SetStationIDAndDbID(cpid uint64, id uint32, station_id uint32) {
	cs.connsuid[cpid].Charging_Pile.DB_ID = id
	cs.connsuid[cpid].Charging_Pile.Station_ID = station_id
}

func (cs *Conns) GetConn(uid uint64) *Conn {
	return cs.connsuid[uid]
}

func (cs *Conns) Remove(c *Conn) {
	delete(cs.connsindex, c.index)

	connuid, ok := cs.connsuid[c.ID]
	if ok && c.index == connuid.index {
		delete(cs.connsuid, c.ID)
	}
}

func (cs *Conns) Check(uid uint64) bool {
	conn, ok := cs.connsuid[uid]
	if ok {
		_, realok := cs.connsindex[conn.index]

		return realok
	}
	return ok
}

func (cs *Conns) GetCount() int {
	return len(cs.connsindex)
}
