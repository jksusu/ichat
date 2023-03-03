package router

import (
	"ichat/cmd/gateway/store"
	"ichat/internal/websocket"
)

func handlerSingleChat(message *websocket.ChatMessage) {
	//消息入库

	//通知发送端已经收到消息了
	handlerMessageAck(message)
}

func handlerGroupChat(message *websocket.ChatMessage) {

}

func handlerMessageAck(message *websocket.ChatMessage) {

}

// 消息ack
func messageAck(to string, seq int64, msgId, route string) error {

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
