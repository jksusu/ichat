package ichat_ws

import (
	"encoding/json"
	"errors"
	"github.com/golang/glog"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Connect 连接管理
type Connect struct {
	mutex             sync.Mutex      //互斥锁
	ID                string          //连接id
	Conn              *websocket.Conn //连接对象
	LastHeartbeatTime time.Time       //最近一次心跳时间

	readChan  chan *WSMessage //读
	writeChan chan *WSMessage //写
	closeChan chan byte       //关闭chan
	isClosed  bool            //关闭状态机
	*Config
}

// 读
func (c *Connect) readLoop(s *Server) {
	var (
		messageType int
		p           []byte
		err         error
		m           *WSMessage
	)
	for {
		if messageType, p, err = c.Conn.ReadMessage(); err != nil {
			c.Close(s)
		}

		m = BuildWSMessage(messageType, p)
		select {
		case c.readChan <- m:
		case <-c.closeChan:
			return
		}
	}
	return
}

func (c *Connect) writeLoop(s *Server) {
	var message *WSMessage
	for {
		select {
		case message = <-c.writeChan:
			if err := c.Conn.WriteMessage(message.MsgType, message.MsgData); err != nil {
				c.Close(s)
			}
		case <-c.closeChan:
			return
		}
	}
	return
}

func HttpUpgraderToWebsocket(s *Server, ID string, conn *websocket.Conn) *Connect {
	connect := &Connect{
		ID:                ID,
		Conn:              conn,
		LastHeartbeatTime: time.Now(),
		readChan:          make(chan *WSMessage, s.Config.WsReadChannelSize),
		writeChan:         make(chan *WSMessage, s.Config.WsWriteChannelSize),
		closeChan:         make(chan byte),
		Config:            s.Config,
	}
	s.Add(connect)
	go connect.readLoop(s)
	go connect.writeLoop(s)
	go connect.heartbeatCheck(s)
	return connect
}

func (c *Connect) Push(message *WSMessage) (err error) {
	select {
	case c.writeChan <- message:
	case <-c.closeChan:
		return errors.New("push 消息连接关闭")
	default:
		return errors.New("ERR_CONNECTION_LOSS")
	}
	return
}

func (c *Connect) ReadMsg() (message *WSMessage, err error) {
	select {
	case message = <-c.readChan:
	case <-c.closeChan:
		err = errors.New("connect close")
	}
	return
}

func (c *Connect) Close(s *Server) {
	c.Conn.Close()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	GlobalConnManager.Remove(c)
	if !c.isClosed {
		glog.Infof("connect close id:%v", c.ID)
		s.Disconnect(c.ID)
		c.isClosed = true
		close(c.closeChan)
	}
	return
}

func (c *Connect) IsLive() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.isClosed || time.Now().Sub(c.LastHeartbeatTime) > time.Duration(c.Config.WsHeartbeatTimeout)*time.Second {
		glog.Infof("connect no live id:%v ", c.ID)
		return false
	}
	return true
}

func (c *Connect) Saveheartbeat() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.LastHeartbeatTime = time.Now()
}

func (c *Connect) heartbeatCheck(s *Server) {
	timer := time.NewTimer(time.Duration(c.Config.WsHeartbeatTimeout) * time.Second)
	for {
		select {
		case <-timer.C:
			if !c.IsLive() {
				c.Close(s)
				return
			}
			timer.Reset(time.Duration(c.Config.WsHeartbeatTimeout) * time.Second)
		case <-c.closeChan:
			timer.Stop()
			return
		}
	}
}

func (c *Connect) Pong(m *BusinessMessage) (err error) {
	if !c.IsLive() {
		return
	}
	c.Saveheartbeat()
	var buf []byte
	m.Type = RESPONSEPOMG
	if buf, err = json.Marshal(m); err != nil {
		return
	}
	glog.Infoln(m.From + "pong")
	return c.Push(&WSMessage{MsgType: websocket.TextMessage, MsgData: buf})
}
