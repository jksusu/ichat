package router

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"ichat/internal/websocket"
)

const (
	RoutedLoginSignOut = "login.signout"

	RouteChatUserMessage  = "chat.user.message"
	RouteChatGroupMessage = "chat.group.message"
	RouteChatMessageAck   = "chat.message.ack"

	//ack
	RouteChatMessageServerAck = "chat.message.server.ack" //服务端确定
	RouteChatMessageArriveAck = "chat.message.arrive.ack" //已经送达ack
	RouteChatMessageReadAck   = "chat.message.read.ack"   //已读ack

	//group
	RouteChatGroupCreate       = "chat.group.create"
	RouteChatGroupDetail       = "chat.group.detail"
	RouteChatGroupDetailSave   = "chat.group.detail.save"
	RouteChatGroupProhibitChat = "chat.group.prohibit.chat" //禁言

	//群用户
	RouteChatGroupJoin          = "chat.group.join"
	RouteChatGroupExit          = "chat.group.exit"
	RouteChatGroupInviteMember  = "chat.group.invite.member" //邀请
	RouteChatGroupDeleteFriends = "chat.group.delete.member" //删除
)

const (
	//通知消息
	RouteNoticeMessageImportIng = "notice.message.importing" //通知对方正在输入
)

// funcTree 方法树
var funcTree = map[string]func(message *websocket.ChatMessage){
	RouteChatUserMessage:  handlerSingleChat,
	RouteChatGroupMessage: handlerGroupChat,
	RouteChatMessageAck:   handlerMessageAck,
}

func MessageRoute(message *websocket.ChatMessage) {
	logrus.Infof("message route fromId:%v", message.From)

	fuc, ok := funcTree[message.Route]
	if ok {
		fuc(message)
		return
	}

}

func msgDecode(p []byte) map[string]interface{} {
	var resMap map[string]interface{}

	if err := json.Unmarshal([]byte(p), &resMap); err != nil {
		fmt.Println("byte转map失败", err)
	}
	//fmt.Println("args取值", resMap["contents"], reflect.TypeOf(resMap["content"]))
	return resMap
}
