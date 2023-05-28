package ichat_ws

import (
	"sync"
)

var GlobalConnManager = new(Manager)

type Manager struct {
	*sync.Map
}

func (m *Manager) Add(conn *Connect) {
	m.Store(conn.ID, conn)
}

func (m *Manager) Remove(conn *Connect) {
	m.Delete(conn.ID)
}

func (m *Manager) Get(ID any) (conn *Connect, ok bool) {
	val, ok := m.Load(ID)
	if !ok {
		return nil, false
	}
	return val.(*Connect), true
}

func (m *Manager) GetAll() []*Connect {
	arr := make([]*Connect, 0)
	m.Range(func(key, val interface{}) bool {
		arr = append(arr, val.(*Connect))
		return true
	})
	return arr
}
