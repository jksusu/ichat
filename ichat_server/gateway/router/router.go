package router

import (
	"github.com/sirupsen/logrus"
	"ichat/gateway/message"
	"ichat/gateway/router/chat"
	"ichat/gateway/router/relation"
	"ichat/gateway/router/route"
	"ichat/gateway/router/session"
	"ichat/gateway/router/user"
)

var funcTree = map[string]func(message *message.RequestMessage){
	//消息管理
	route.RouteIchatTalkToUser:     chat.HandlerSingleChat,
	route.RouteIchatTalkToGroup:    chat.HandlerGroupChat,
	route.RouteIchatTalkAckReceive: chat.HandlerAckReceive,
	route.RouteChatMessageAck:      chat.HandlerMessageAck,

	route.RouteChatMessageArriveAck: chat.HandlerArriveAck,
	route.RouteChatMessageReadAck:   chat.HandlerReadAck,

	//会话管理
	route.RouteSessionLists: session.HandlerSessionLists,
	route.RouteSessionAdd:   session.HandlerSessionAdd,
	route.RouteSessionDel:   session.HandlerSessionDel,

	//关系管理
	route.RouteRelationFriendsApply:      relation.HandlerFriendsApply,
	route.RouteRelationFriendsApplyReply: relation.HandlerFriendsApplyReply,
	route.RouteRelationFriendsEdit:       relation.HandlerFriendsEdit,
	route.RouteRelationFriendsDel:        relation.HandlerFriendsDel,

	//用户查询
	route.RouteUserInfoSearch: user.HandlerUserSearch,
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
