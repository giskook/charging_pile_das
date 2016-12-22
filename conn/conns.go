package conn

import (
	"log"
	"sync"
	"sync/atomic"
)

type Conns struct {
	//	mutex_index sync.RWMutex
	connsindex map[uint32]*Conn
	//	mutex_id    sync.RWMutex
	connsuid map[uint64]*Conn

	index uint32
}

var connsInstance *Conns
var mutex_conns sync.RWMutex

func NewConns() *Conns {
	defer func() {
		mutex_conns.RUnlock()
	}()
	mutex_conns.RLock()
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
	//	cs.mutex_index.Lock()
	cs.connsindex[conn.index] = conn
	//	cs.mutex_index.Unlock()
}

func (cs *Conns) SetID(gatewayid uint64, conn *Conn) {
	log.Printf("add conn %d\n", conn.ID)
	//	cs.mutex_id.Lock()
	cs.connsuid[gatewayid] = conn
	//	cs.mutex_id.Unlock()
}

func (cs *Conns) SetStationIDAndDbID(cpid uint64, id uint32, station_id uint32) {
	//	cs.mutex_id.Lock()
	cs.connsuid[cpid].Charging_Pile.DB_ID = id
	cs.connsuid[cpid].Charging_Pile.Station_ID = station_id
	//	cs.mutex_id.Unlock()
}

func (cs *Conns) GetConn(uid uint64) *Conn {
	defer func() {
		//		cs.mutex_id.RUnlock()
	}()

	//	cs.mutex_id.RLock()
	return cs.connsuid[uid]
}

func (cs *Conns) Remove(c *Conn) {
	log.Printf("rm conn %d\n", c.ID)
	//	cs.mutex_index.Lock()
	delete(cs.connsindex, c.index)
	//	cs.mutex_index.Unlock()

	//	cs.mutex_id.Lock()
	connuid, ok := cs.connsuid[c.ID]
	if ok && c.index == connuid.index {
		delete(cs.connsuid, c.ID)
	}
	//	cs.mutex_id.Unlock()
}

func (cs *Conns) Check(uid uint64) bool {
	//	cs.mutex_id.RLock()
	conn, ok := cs.connsuid[uid]
	//	cs.mutex_id.RUnlock()
	if ok {
		//		cs.mutex_index.RLock()
		_, realok := cs.connsindex[conn.index]
		//		cs.mutex_index.RUnlock()

		return realok
	}
	return ok
}

func (cs *Conns) GetCount() int {
	defer func() {
		//		cs.mutex_index.RUnlock()
	}()
	//	cs.mutex_index.RLock()

	return len(cs.connsindex)
}
