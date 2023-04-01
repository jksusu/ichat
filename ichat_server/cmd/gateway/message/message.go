package message

import (
	"encoding/json"
	"errors"
	"ichat/internal/websocket"
	"ichat/pkg/ichat_cache/connect"
)

const (
	REQUEST  = 1
	RESPONSE = 2
	NOTICE   = 3
)

// ChatMessage 公共消息字段
type ChatMessage struct {
	MsgId    int64  `json:"msgId"`
	To       string `json:"to"`
	Content  string `json:"content"`
	Type     int    `json:"type"`
	SendTime int64  `json:"send_time"`
	Extra    string `json:"extra"`
}

type RequestMessage struct {
	From  string      `json:"from"`
	ReqId int         `json:"reqId"`
	Type  int         `json:"type"`
	Route string      `json:"route"`
	Data  interface{} `json:"data"`
}

type ResponseMessage struct {
	ReqId int         `json:"reqId"`
	Type  int         `json:"type"`
	From  string      `json:"from"`
	Route string      `json:"route"`
	Data  interface{} `json:"data"`
}

// 推送消息
type PushMessage struct {
	Type  int         `json:"type"`
	From  string      `json:"from"`
	Route string      `json:"route"`
	Data  interface{} `json:"data"`
}

type Message struct {
}

func BuildDecodeChatMessage(p interface{}) (*ChatMessage, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	var msg *ChatMessage
	if err = json.Unmarshal(b, &msg); err != nil {
		return nil, err
	}
	return msg, nil
}

// 解析data
func BuildDecode(p interface{}) (b []byte, err error) {
	return json.Marshal(p)
}

func Response(res *ResponseMessage) error {
	if len(res.From) == 0 {
		return errors.New("from uid non-existent ")
	}
	connId, err := connect.GetConnIdByUID(res.From)
	if err != nil {
		return err
	}
	res.Type = RESPONSE
	return websocket.Response(connId, res)
}

func Push(connId string, p *PushMessage) error {
	p.Type = NOTICE
	return websocket.Response(connId, p)
}
