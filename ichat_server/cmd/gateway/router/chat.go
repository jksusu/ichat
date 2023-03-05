package router

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ichat/cmd/gateway/store"
	"ichat/internal/websocket"
	"ichat/pkg/db"
	"time"
)

func handlerSingleChat(message *websocket.ChatMessage) {
	//消息入库
	msgId, err := AddMessageContent(message)
	if err != nil {
		logrus.Error(err)
		return
	}
	messageAck(message.To, message.Seq, message.MsgId, message.Route) //通知发送端已经发送成功了

	message.MsgId = msgId

	//如果在线则发送给对方
	connId, err := store.GetConnIdByUid(message.To)
	if err := websocket.SendChatMessage(connId, message); err != nil {
		return
	}
}

func handlerGroupChat(message *websocket.ChatMessage) {

}

func handlerMessageAck(message *websocket.ChatMessage) {

}

func AddMessageContent(message *websocket.ChatMessage) (msgId int64, err error) {
	msgId = store.Gen.Node.Generate().Int64()
	sendTime := time.Now().UnixNano()
	content := store.ChatMessageContent{
		ID:       msgId,
		Type:     message.Type,
		Content:  message.Content,
		SendTime: sendTime,
		Extra:    message.Extra,
	}
	index := make([]store.ChatMessageIndex, 2)
	index[0] = store.ChatMessageIndex{
		ID:       store.Gen.Node.Generate().Int64(),
		A:        message.From,
		B:        message.To,
		ISend:    1,
		MsgId:    msgId,
		SendTime: sendTime,
	}
	index[1] = store.ChatMessageIndex{
		ID:       store.Gen.Node.Generate().Int64(),
		A:        message.To,
		B:        message.From,
		ISend:    0,
		MsgId:    msgId,
		SendTime: sendTime,
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

	return
}

// 消息ack
func messageAck(to string, seq int64, msgId int64, route string) error {

	chat := websocket.ChatMessage{Seq: seq, Route: route, MsgId: msgId}

	connId, err := store.GetConnIdByUid(to)
	if err != nil {
		return err
	}
	if err = websocket.SendChatMessage(connId, &chat); err != nil {
		return err
	}
	return nil
}
