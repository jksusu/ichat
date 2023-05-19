package ichat_websocket

import (
	"sync"
)

var GlobalConnManager *ConnManager

type ConnManager struct {
	onlines *sync.Map
}

func InitConnManager() (err error) {
	GlobalConnManager = &ConnManager{
		onlines: &sync.Map{},
	}
	return
}

func (cm *ConnManager) AddConn(conn *Connect) {
	cm.onlines.Store(conn.ID, conn)
}

func (cm *ConnManager) RemoveConn(conn *Connect) {
	cm.onlines.Delete(conn.ID)
}

func (cm *ConnManager) GetConn(ID any) (conn *Connect, ok bool) {
	val, ok := cm.onlines.Load(ID)
	if !ok {
		return nil, false
	}
	return val.(*Connect), true
}

func (cm *ConnManager) GetAllConn() []*Connect {
	arr := make([]*Connect, 0)
	cm.onlines.Range(func(key, val interface{}) bool {
		arr = append(arr, val.(*Connect))
		return true
	})
	return arr
}
