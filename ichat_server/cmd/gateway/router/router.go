package router

import (
	"github.com/sirupsen/logrus"
	"ichat/cmd/gateway/message"
	"ichat/cmd/gateway/router/chat"
	"ichat/cmd/gateway/router/route"
	"ichat/cmd/gateway/router/session"
)

// funcTree 方法树
var funcTree = map[string]func(message *message.RequestMessage){
	route.RouteChatUserMessage:  chat.HandlerSingleChat,
	route.RouteChatGroupMessage: chat.HandlerGroupChat,
	route.RouteChatMessageAck:   chat.HandlerMessageAck,

	//会话管理
	route.RouteSessionLists: session.HandlerSessionLists,
	route.RouteSessionAdd:   session.HandlerSessionAdd,
	route.RouteSessionDel:   session.HandlerSessionDel,

	//消息ack功能
	route.RouteChatMessageArriveAck: chat.HandlerArriveAck,
	route.RouteChatMessageReadAck:   chat.HandlerReadAck,
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
