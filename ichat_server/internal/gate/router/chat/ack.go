package chat

import (
	"github.com/sirupsen/logrus"
	"ichat/internal/gate/message"
)

type Ack struct {
	MsgId  int64 `json:"msgId"`
	Status int   `json:"status"` //ack状态 1已发送 2已收到 3已读
	Type   int   `json:"type"`   //ack类型
}

func HandlerAckReceive(message *message.RequestMessage) {

}

// 已送达 收到确认
func HandlerArriveAck(request *message.RequestMessage) {
	var err error
	//修改消息发送人最后一条未读消息
	if err = message.Response(&message.ResponseMessage{
		ReqId: request.ReqId,
		From:  request.From,
		//Route: router.RouteChatMessageArriveAck,
	}); err != nil {
		logrus.Error(err)
		return
	}

	//修改会话最后确认送达的消息id

	//if err = message.Push("fda", &message.PushMessage{From: request.From, Route: router.RouteChatMessageArriveAck}); err != nil {
	//	logrus.Error(err)
	//	return
	//}
}

// 消息已读回执
func HandlerReadAck(request *message.RequestMessage) {
	var err error
	//修改消息发送人最后一条未读消息
	if err = message.Response(&message.ResponseMessage{ReqId: request.ReqId, From: request.From}); err != nil {
		logrus.Error(err)
		return
	}
	//if err = message.Push("fda", &message.PushMessage{From: request.From, Route: router.RouteChatMessageReadAck}); err != nil {
	//	logrus.Error(err)
	//	return
	//}
}
