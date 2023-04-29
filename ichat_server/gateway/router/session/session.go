package session

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"ichat/gateway/message"
	"ichat/pkg/ichat_cache/session/lists"
	"ichat/pkg/ichat_service/session_service"
)

// 获取会话列表
func HandlerSessionLists(request *message.RequestMessage) {
	sessionList := session_service.GetSessionAll(request.From)

	if err := message.Response(&message.ResponseMessage{
		ReqId: request.ReqId,
		From:  request.From,
		Route: request.Route,
		Data:  sessionList,
	}); err != nil {
		logrus.Error(err)
	}
}

type SessionMessage struct {
	To string `json:"to"`
}

// bindSessionMessage 绑定消息到 SessionMessage
func bindSessionMessage(request *message.RequestMessage) (session *SessionMessage, err error) {
	var b []byte
	if b, err = message.BuildDecode(request.Data); err != nil {
		return
	}
	err = json.Unmarshal(b, &session)

	return
}

func HandlerSessionAdd(request *message.RequestMessage) {
	session, err := bindSessionMessage(request)
	if err != nil {
		logrus.Error(err)
		return
	}
	//增加
	if err = lists.Add(request.From, session.To); err != nil {
		logrus.Error(err)
	}
	return
}

func HandlerSessionDel(request *message.RequestMessage) {
	session, err := bindSessionMessage(request)
	if err != nil {
		logrus.Error(err)
		return
	}
	//增加
	if err = lists.Del(request.From, session.To); err != nil {
		logrus.Error(err)
	}
	return
}
