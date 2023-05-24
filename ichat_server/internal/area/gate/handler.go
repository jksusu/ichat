package gate

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"ichat/internal/format"
	"ichat/internal/gate/router/route"
	"ichat/pkg/ichat_cache/connect"
	"ichat/pkg/ichat_ws"
	"net/http"
	"net/url"
)

type HandlerImpl struct{}

type UserInfo struct {
	ConnID   string `json:"connId"`
	UID      string `json:"uid"`
	Nickname string `json:"nickname"`
}

func (h *HandlerImpl) BeforeAccept(r *http.Request) error {
	var token string
	if token = UrlDecodeToken(r); token == "" {
		return errors.New("certification failed")
	}
	if _, err := connect.GetUidByToken(token); err != nil {
		return errors.New("certification failed")
	}
	return nil
}

func (h *HandlerImpl) Accept(conn *ichat_ws.Connect, r *http.Request) (uint64, error) {
	uid, err := connect.GetUidByToken(UrlDecodeToken(r))
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	//双向绑定
	connect.AddConnBidirectionalBindingUid(conn.ID, uid)
	//推送第一个包是用户信息
	conn.Conn.WriteJSON(&format.R{
		Route: route.RouteUserInfoUpdate,
		Type:  format.RESPONSE,
		Data: UserInfo{
			ConnID:   conn.ID,
			UID:      uid,
			Nickname: "",
		},
	})

	return 1, nil
}

func (h *HandlerImpl) Receive(conn *ichat_ws.Connect, msg *ichat_ws.WSMessage) {
	if msg.MsgData != nil {
		r := new(format.R)
		j, _ := json.Marshal(msg.MsgData)
		json.Unmarshal(j, &r)
		ctx := context.Background()
		context.WithValue(ctx, "fcId", conn.ID) //formConnId
		context.WithValue(ctx, "type", r.Type)
		context.WithValue(ctx, "from", r.From)
		context.WithValue(ctx, "to", r.To)
		Router(ctx, r)
	}
}

func (h *HandlerImpl) Disconnect(connId string) error {
	if err := connect.DeleteBindingRelationship(connId); err != nil {
		logrus.Error(err)
	}
	return nil
}

func UrlDecodeToken(r *http.Request) string {
	raw := r.URL.RawQuery
	values, _ := url.ParseQuery(raw)
	return values.Get("token")
}
