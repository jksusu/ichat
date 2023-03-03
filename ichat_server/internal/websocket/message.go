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

type WSMessage struct {
	MsgType int
	MsgData []byte
}

// 业务消息体
type BusinessMessage struct {
	Type string          `json:"type"` // ping pong
	Body json.RawMessage `json:"body"`
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

// 编码消息体 方便后续替换
func EncodeMessage(msg *BusinessMessage) (buf []byte, err error) {
	return json.Marshal(msg)
}

// 聊天消息
type ChatMessage struct {
	Msgv    string            `json:"msgv"`
	MsgId   string            `json:"msgId"`
	Seq     int64             `json:"seq"`
	Route   string            `json:"route"` //消息处理 action
	From    string            `json:"from"`
	To      string            `json:"to"`
	Content string            `json:"data"`
	Extra   map[string]string `json:"extra"`
}

// 消息ack
type AckMessage struct {
	Type   string `json:"type"`
	MsgId  string `json:"msgId"`
	Status int    `json:"status"`
}
