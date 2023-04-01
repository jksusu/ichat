package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"ichat/cmd/gateway/message"
	"ichat/cmd/gateway/router"
	"ichat/cmd/gateway/router/route"
	"ichat/internal/websocket"
	"ichat/pkg/ichat_cache/connect"
	"net/http"
	"net/url"
)

type HandlerImpl struct{}

func (h *HandlerImpl) BeforeAccept(r *http.Request) error {
	var token string
	if token = UrlDecodeToken(r); token == "" {
		return errors.New("certification failed")
	}
	if _, err := connect.GetUIDbyToken(token); err != nil {
		return errors.New("certification failed")
	}
	return nil
}

func (h *HandlerImpl) Accept(conn *websocket.Connect, r *http.Request) (uint64, error) {
	uid, err := connect.GetUIDbyToken(UrlDecodeToken(r))
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	//双向绑定
	connect.AddConnBidirectionalBindingUid(conn.ID, uid)

	//推送第一个包是用户信息
	m := make(map[string]string)
	m["connId"] = conn.ID
	m["uid"] = uid

	message.Push(conn.ID, &message.PushMessage{Route: route.RouteUserInfoUpdate, Data: m})

	return 1, nil
}

func (h *HandlerImpl) Receive(conn *websocket.Connect, msg *websocket.BusinessMessage) {
	router.MessageRoute(&message.RequestMessage{From: msg.From, ReqId: msg.ReqId, Type: msg.Type, Route: msg.Route, Data: msg.Data})
}

func (h *HandlerImpl) Disconnect(connId string) error {
	if err := connect.DeleteBindingRelationship(connId); err != nil {
		logrus.Error(err)
	}
	return nil
}

// MsgDecode 解析消息到map
func MsgDecode(p []byte) map[string]interface{} {
	var resMap map[string]interface{}

	if err := json.Unmarshal([]byte(p), &resMap); err != nil {
		fmt.Println("byte转map失败", err)
	}
	//fmt.Println("args取值", resMap["contents"], reflect.TypeOf(resMap["content"]))
	return resMap
}

// UrlDecodeToken 从url中解析token
func UrlDecodeToken(r *http.Request) string {
	raw := r.URL.RawQuery
	values, _ := url.ParseQuery(raw)
	return values.Get("token")
}
