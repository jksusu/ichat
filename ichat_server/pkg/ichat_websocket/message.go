package ichat_websocket

import (
	"encoding/json"
)

// socket收到的消息类型解析
const (
	REQUEST      = 1
	RESPONSE     = 2
	NOTICE       = 3
	REQUESTPING  = 4
	RESPONSEPOMG = 5
)

// ichat_websocket 基础消息格式
type WSMessage struct {
	MsgType int
	MsgData []byte
}

// BuildWSMessage 绑定消息
func BuildWSMessage(msgType int, msgData []byte) (wsMessage *WSMessage) {
	return &WSMessage{
		MsgType: msgType,
		MsgData: msgData,
	}
}

// 业务消息体
type BusinessMessage struct {
	ReqId int         `json:"reqId"`
	Type  int         `json:"type"`
	From  string      `json:"from"`
	Route string      `json:"route"`
	Data  interface{} `json:"data"`
}

// DecodeMessageBody 解码json消息体
func DecodeMessageBody(buf []byte) (msg *BusinessMessage, err error) {
	var message BusinessMessage
	if err = json.Unmarshal(buf, &message); err != nil {
		return
	}
	msg = &message
	return
}
