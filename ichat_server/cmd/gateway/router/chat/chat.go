package chat

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ichat/cmd/gateway/message"
	"ichat/cmd/gateway/store"
	"ichat/pkg/db"
	"ichat/pkg/ichat_cache/connect"
	"ichat/pkg/ichat_cache/unread"
	"time"
)

func HandlerSingleChat(request *message.RequestMessage) {
	var (
		chatMsg   *message.ChatMessage
		err       error
		messageId int64
		connId    string
	)
	if chatMsg, err = message.BuildDecodeChatMessage(request.Data); err != nil {
		logrus.Error(err)
		return
	}
	sendTime := time.Now().UnixNano()
	chatMsg.SendTime = sendTime

	if messageId, err = AddMessageContent(request.From, chatMsg); err != nil {
		logrus.Error(err)
		return
	}
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
	if err = message.Push(connId, &message.PushMessage{
		From:  request.From,
		Route: request.Route,
		Data:  request.Data,
	}); err != nil {
		logrus.Error(err)
		return
	}
}

func HandlerGroupChat(message *message.RequestMessage) {

}

func HandlerMessageAck(message *message.RequestMessage) {

}

func AddMessageContent(from string, msg *message.ChatMessage) (msgId int64, err error) {
	msgId = store.Gen.Node.Generate().Int64()
	content := store.ChatMessageContent{
		ID:       msgId,
		Type:     msg.Type,
		Content:  msg.Content,
		SendTime: msg.SendTime,
		Extra:    msg.Extra,
	}
	index := make([]store.ChatMessageIndex, 2)
	index[0] = store.ChatMessageIndex{
		ID:       store.Gen.Node.Generate().Int64(),
		A:        from,
		B:        msg.To,
		ISend:    1,
		MsgId:    msgId,
		SendTime: msg.SendTime,
	}
	index[1] = store.ChatMessageIndex{
		ID:       store.Gen.Node.Generate().Int64(),
		A:        msg.To,
		B:        from,
		ISend:    0,
		MsgId:    msgId,
		SendTime: msg.SendTime,
	}

	err = db.DB.Transaction(func(d *gorm.DB) (err error) {
		if err = d.Table(store.MESSAGE_CONTENT).Create(&content).Error; err != nil {
			return
		}
		if err = d.Table(store.MESSAGE_INDEX).Create(&index).Error; err != nil {
			return
		}
		return
	})
	if err != nil {
		logrus.Error(err)
	}
	return
}
