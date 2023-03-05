package websocket

import (
	"encoding/json"
)

// socket收到的消息类型解析
const (
	PINGTYPE = "PING" //ping
	PONGTYPE = "PONG" //ping
	CHATTYPE = "CHAT" //chat消息
)

// websocket 基础消息格式
type WSMessage struct {
	MsgType int
	MsgData []byte
}

// 业务消息体
type BusinessMessage struct {
	Type string `json:"type"` // ping pong
	Data *ChatMessage
}

type PONGMessage struct {
	Type string `json:"type"`
}

// BuildWSMessage 绑定消息
func BuildWSMessage(msgType int, msgData []byte) (wsMessage *WSMessage) {
	return &WSMessage{
		MsgType: msgType,
		MsgData: msgData,
	}
}

// 解码json消息体
func DecodeMessageBody(buf []byte) (msg *BusinessMessage, err error) {
	var message BusinessMessage
	if err = json.Unmarshal(buf, &message); err != nil {
		return
	}
	msg = &message
	return
}

// ChatMessage 公共消息字段
type ChatMessage struct {
	Msgv     string `json:"msgv"`
	MsgId    int64  `json:"msgId"`
	Seq      int64  `json:"seq"`
	Route    string `json:"route"`
	From     string `json:"from"`
	To       string `json:"to"`
	Content  string `json:"content"`
	Type     int    `json:"type"`
	SendTime int64  `json:"send_time"`
	Extra    string `json:"extra"`
}
