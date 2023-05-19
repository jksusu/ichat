package chat

import (
	"github.com/sirupsen/logrus"
	"ichat/internal/gate/message"
	"ichat/pkg/ichat_cache/connect"
	"ichat/pkg/ichat_cache/unread"
	"ichat/pkg/ichat_tool/idgen"
	"time"
)

func HandlerSingleChat(request *message.RequestMessage) {
	var (
		chatMsg   *message.ChatMessage
		err       error
		messageId int64
		connId    string
		sendTime  = time.Now().UnixNano()
	)
	if chatMsg, err = message.BuildDecodeChatMessage(request.Data); err != nil {
		logrus.Error(err)
		return
	}
	chatMsg.SendTime = sendTime
	messageId = idgen.Gen.Node.Generate().Int64()

	//if err = chat_service.AddMessageContent(request.From, messageId, chatMsg); err != nil {
	//	logrus.Error(err)
	//	return
	//}

	chatMsg.MsgId = messageId
	//通知发送端已经发送成功了
	if err = message.Response(&message.ResponseMessage{
		ReqId: request.ReqId,
		From:  request.From,
		Data:  chatMsg,
	}); err != nil {
		logrus.Error(err)
		return
	}
	//更新消息状态
	//if err = message_status.UpdateMsgHasBeenSent(request.From, chatMsg.To, messageId); err != nil {
	//	logrus.Error(err)
	//}
	////更新接收方未读消息
	if err = unread.SetContactsLastMsg(chatMsg.To, request.From, chatMsg.MsgId); err != nil {
		logrus.Error(err)
	}
	//转发给to
	if connId, err = connect.GetConnIdByUID(chatMsg.To); err != nil {
		logrus.Error(err)
		return
	}
	//获取发送者的详情

	if err = message.Notice(connId, &message.NoticeMessage{
		MsgId: messageId,
		Route: request.Route,
		From:  request.From,
		FromInfo: &message.FromInfo{
			UID: request.From,
		},
		Data: chatMsg,
	}); err != nil {
		logrus.Error(err)
		return
	}
}

func HandlerGroupChat(message *message.RequestMessage) {

}

func HandlerMessageAck(message *message.RequestMessage) {

}
