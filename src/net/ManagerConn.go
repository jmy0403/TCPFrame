package net

import (
	"GameOfTCP/src/iface"
	"errors"
	"sync"
)

type ManagerConn struct {
	connMap map[uint32]iface.IConn
	lock    sync.RWMutex
}

func NewManagerConn() iface.IManageConn {
	return &ManagerConn{
		connMap: make(map[uint32]iface.IConn),
	}
}
func (m *ManagerConn) Add(conn iface.IConn) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.connMap[conn.GetConnID()] = conn
}

func (m *ManagerConn) Remove(conn iface.IConn) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.connMap, conn.GetConnID())
}

func (m *ManagerConn) Get(id uint32) (iface.IConn, error) {
	m.lock.RLock()
	defer m.lock.Unlock()
	if conn, ok := m.connMap[id]; ok {
		return conn, nil
	} else {

		return nil, errors.New("通过这个id没有找到所对应的连接")
	}

}

func (m *ManagerConn) Length() int {
	return len(m.connMap)
}

func (m *ManagerConn) CloseConn() {
	m.lock.Lock()
	defer m.lock.Unlock()
	for id, conn := range m.connMap {
		conn.Stop()
		delete(m.connMap, id)
	}
}
