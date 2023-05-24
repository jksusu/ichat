package ichat_ws

import (
	"sync"
)

var GlobalConnManager = new(manager)

type manager struct {
	online *sync.Map
}

func (m *manager) Add(conn *Connect) {
	m.online.Store(conn.ID, conn)
}

func (m *manager) Remove(conn *Connect) {
	m.online.Delete(conn.ID)
}

func (m *manager) Get(ID any) (conn *Connect, ok bool) {
	val, ok := m.online.Load(ID)
	if !ok {
		return nil, false
	}
	return val.(*Connect), true
}

func (m *manager) GetAll() []*Connect {
	arr := make([]*Connect, 0)
	m.online.Range(func(key, val interface{}) bool {
		arr = append(arr, val.(*Connect))
		return true
	})
	return arr
}
