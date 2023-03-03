package websocket

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
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
}

// 读
func (c *Connect) readloop() {
	for {
		msgType, msgData, err := c.Conn.ReadMessage()
		if err != nil {
			goto ERROR
		}
		msg := BuildWSMessage(msgType, msgData)
		select {
		case c.readChan <- msg:
		case <-c.closeChan:
			goto CLOSED
		}
	}
ERROR:
	c.Close()
CLOSED:
}

func (c *Connect) writeLoop() {

	var message *WSMessage

	for {
		select {
		case message = <-c.writeChan:
			if err := c.Conn.WriteMessage(message.MsgType, message.MsgData); err != nil {
				goto ERROR
			}
		case <-c.closeChan:
			goto CLOSED
		}
	}

ERROR:
	c.Close()
CLOSED:
}

// 初始化websocket连接 启动读写协程
func HttpUpgraderToWebsocket(ID string, conn *websocket.Conn) *Connect {
	connect := &Connect{
		ID:                ID,
		Conn:              conn,
		LastHeartbeatTime: time.Now(),

		readChan:  make(chan *WSMessage, GlobalConfig.WsReadChannelSize),
		writeChan: make(chan *WSMessage, GlobalConfig.WsWriteChannelSize),
		closeChan: make(chan byte),
	}

	GlobalConnManager.AddConn(connect) //把当前连接加入到连接管理

	go connect.readloop()
	go connect.writeLoop()
	go connect.heartbeatCheck()

	return connect
}

// Push 推送消息
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
		ReadMessageNumber_DEC()
	case <-c.closeChan:
		ReadMessageFailNumber_INC()
	}
	return
}

// Close  线程安全可重入关闭
func (c *Connect) Close() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	logrus.Infof("close connect id: %v", c.ID)

	GlobalConnManager.RemoveConn(c)
	c.Conn.Close()
	if !c.isClosed {
		GlobalServer.Disconnect(c.ID)
		c.isClosed = true
		close(c.closeChan)
	}
}

// 检测连接是否存活
func (c *Connect) IsLive() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isClosed || time.Now().Sub(c.LastHeartbeatTime) > time.Duration(GlobalConfig.WsHeartbeatTimeout)*time.Second {
		logrus.Infof("connect no live id:%v ", c.ID)
		return false
	}
	return true
}

// 更新心跳时间
func (c *Connect) Saveheartbeat() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.LastHeartbeatTime = time.Now()
}

func (c *Connect) heartbeatCheck() {
	timer := time.NewTimer(time.Duration(GlobalConfig.WsHeartbeatTimeout) * time.Second)

	for {
		select {
		case <-timer.C:
			if !c.IsLive() {
				c.Close()
				return
			}
			timer.Reset(time.Duration(GlobalConfig.WsHeartbeatTimeout) * time.Second)
		case <-c.closeChan:
			timer.Stop()
			return
		}
	}
}

func (c *Connect) Pong() (err error) {
	if !c.IsLive() {
		return
	}
	c.Saveheartbeat()
	var buf []byte
	pong := &PONGMessage{Type: PONGTYPE}

	if buf, err = json.Marshal(pong); err != nil {
		return
	}
	if err = c.Push(&WSMessage{MsgType: websocket.TextMessage, MsgData: buf}); err != nil {
		return
	}
	return
}
