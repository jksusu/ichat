package message

import (
	"encoding/json"
	"errors"
	"ichat/pkg/ichat_cache/connect"
	"ichat/pkg/ichat_websocket"
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

// 通知消息
type NoticeMessage struct {
	MsgId    int64       `json:"msgId"`
	Route    string      `json:"route"`
	Type     int         `json:"type"`
	From     string      `json:"from"`
	FromInfo *FromInfo   `json:"fromInfo"`
	Data     interface{} `json:"data"`
}

type FromInfo struct {
	UID      string `json:"uid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
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

func DataUnpack(data interface{}, add interface{}) error {
	p, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(p, add); err != nil {
		return err
	}
	return nil
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
	return ichat_websocket.Response(connId, res)
}

// 回复
func Responses(request *RequestMessage) (err error) {
	if len(request.From) == 0 {
		return errors.New("from uid non-existent ")
	}
	var connId string
	if connId, err = connect.GetConnIdByUID(request.From); err != nil {
		return err
	}
	return ichat_websocket.Response(connId, &ResponseMessage{
		ReqId: request.ReqId,
		Type:  RESPONSE,
		From:  request.From,
		Route: request.Route,
	})
}

func Push(connId string, p *PushMessage) error {
	p.Type = NOTICE
	return ichat_websocket.Response(connId, p)
}

func Notice(connId string, p *NoticeMessage) error {
	p.Type = NOTICE
	return ichat_websocket.Response(connId, p)
}
