package gate

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"ichat"
	"ichat/internal/format"
	"ichat/internal/route"
	"ichat/pkg/cache/connect"
	"ichat/pkg/ichat_ws"
	"net/http"
	"net/url"
)

type HandlerImpl struct{}

var Handler = new(HandlerImpl)

type UserInfo struct {
	ConnID   string `json:"connId"`
	UID      string `json:"uid"`
	Nickname string `json:"nickname"`
}

func NewWsServer() *ichat_ws.Server {
	ws := ichat_ws.NewServer()
	ws.Config = ichat.GlobalConf.Gateway
	ws.Acceptor = &HandlerImpl{}
	ws.MessageListener = &HandlerImpl{}
	ws.StateListener = &HandlerImpl{}
	ws.BeforeAcceptor = &HandlerImpl{}
	return ws
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
	connect.AddConnBidirectionalBindingUid(conn.ID, uid)
	conn.Conn.WriteJSON(&format.R{
		Route: route.RouteUserInfoUpdate,
		Type:  format.NOTICE,
		Data:  UserInfo{ConnID: conn.ID, UID: uid},
	})
	return 1, nil
}

func (h *HandlerImpl) Receive(conn *ichat_ws.Connect, msg *ichat_ws.WSMessage) {
	if msg.MsgData != nil {
		r := &format.R{}
		json.Unmarshal(msg.MsgData, &r)
		ctx := context.Background()
		context.WithValue(ctx, "connId", conn.ID) //formConnId
		context.WithValue(ctx, "type", r.Type)
		context.WithValue(ctx, "from", r.From)
		context.WithValue(ctx, "to", r.To)
		route.Router(ctx, r)
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
