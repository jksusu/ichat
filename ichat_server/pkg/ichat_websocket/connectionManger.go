package ichat_websocket

import (
	"sync"
)

var GlobalConnManager *ConnManager

type ConnManager struct {
	onlines *sync.Map
}

// 初始化连接管理器
func InitConnManager() (err error) {
	GlobalConnManager = &ConnManager{
		onlines: &sync.Map{},
	}
	return
}

// api 加入连接
func (cm *ConnManager) AddConn(conn *Connect) {
	cm.onlines.Store(conn.ID, conn)
	OnlinesConnectNumber_INC()
}

func (cm *ConnManager) RemoveConn(conn *Connect) {
	cm.onlines.Delete(conn.ID)
	OnlinesConnectNumber_DEC()
}

func (cm *ConnManager) GetConn(ID any) (conn *Connect, ok bool) {
	val, ok := cm.onlines.Load(ID)
	if !ok {
		return nil, false
	}
	return val.(*Connect), true
}

// 获取所有连接
func (cm *ConnManager) GetAllConn() []*Connect {
	arr := make([]*Connect, 0)
	cm.onlines.Range(func(key, val interface{}) bool {
		arr = append(arr, val.(*Connect))
		return true
	})
	return arr
}
