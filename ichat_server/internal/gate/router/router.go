package router

import (
	"github.com/sirupsen/logrus"
	"ichat/internal/gate/message"
	chat2 "ichat/internal/gate/router/chat"
	"ichat/internal/gate/router/route"
	"ichat/internal/gate/router/session"
)

var funcTree = map[string]func(message *message.RequestMessage){
	//消息管理
	route.RouteIchatTalkToUser:     chat2.HandlerSingleChat,
	route.RouteIchatTalkToGroup:    chat2.HandlerGroupChat,
	route.RouteIchatTalkAckReceive: chat2.HandlerAckReceive,
	route.RouteChatMessageAck:      chat2.HandlerMessageAck,

	route.RouteChatMessageArriveAck: chat2.HandlerArriveAck,
	route.RouteChatMessageReadAck:   chat2.HandlerReadAck,

	//会话管理
	route.RouteSessionLists: session.HandlerSessionLists,
	route.RouteSessionAdd:   session.HandlerSessionAdd,
	route.RouteSessionDel:   session.HandlerSessionDel,
}

func MessageRoute(message *message.RequestMessage) {
	logrus.Infof("message route fromId:%v", message.From)
	fuc, ok := funcTree[message.Route]
	if ok {
		fuc(message)
		return
	}
	logrus.Error("暂不支持" + message.Route)
}
