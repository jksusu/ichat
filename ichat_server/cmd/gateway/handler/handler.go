package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"ichat/cmd/gateway/router"
	"ichat/internal/websocket"
	"net/http"
	"net/url"
)

type HandlerImpl struct{}

func (h *HandlerImpl) BeforeAccept(r *http.Request) error {
	var token string
	if token = UrlDecodeToken(r); token == "" {
		return errors.New("certification failed")
	}
	if _, err := GetUidByToken(token); err != nil {
		return errors.New("certification failed")
	}
	return nil
}

func (h *HandlerImpl) Accept(conn *websocket.Connect, r *http.Request) (uint64, error) {
	uid, _ := GetUidByToken(UrlDecodeToken(r))
	AddConnBidirectionalBindingUid(conn.ID, uid)

	fmt.Println("连接开始")
	return 1, nil
}

func (h *HandlerImpl) Receive(conn *websocket.Connect, msg *websocket.ChatMessage) {

	router.MessageRoute(msg)

	//websocket.SendMsg(conn.ID, []byte("哈哈哈哈哈"))
	//logrus.Infof("当前连接id:%v ,接收到消息:%v", conn.ID, MsgDecode(msg))

	//conn.Push(&websocket.WSMessage{MsgType: 1, MsgData: []byte("你好，你好")})
}

func (h *HandlerImpl) Disconnect(id string) error {

	DeleteBindingRelationship(id)

	fmt.Println("关机了")
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
