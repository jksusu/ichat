package user

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"ichat/gateway/message"
	"ichat/pkg/ichat_service/user_service"
)

type Users struct {
	To string `json:"to"`
}

func HandlerUserSearch(r *message.RequestMessage) {
	p, _ := message.BuildDecode(r.Data)
	var user *Users
	json.Unmarshal(p, &user)

	//查询用户
	data, err := user_service.FindUserByUid(user.To)
	if err != nil {
		logrus.Error(err)
	}
	if err = message.Response(&message.ResponseMessage{ReqId: r.ReqId, From: r.From, Type: message.RESPONSE, Data: data}); err != nil {
		logrus.Error(err)
	}
}
