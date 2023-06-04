package message

import (
	"context"
	"fmt"
	"ichat/pkg/cache"
)

type lastMsg struct {
	to    string `redis:"to" json:"to"`
	msgId int64  `redis:"msgId" json:"msgId"`
}

// 标记会话最后一条未读消息
func MarkTalkLastReadMessage(uid, to string, msgId int64) error {
	return cache.DB.HSet(context.TODO(), fmt.Sprintf("last:aync:%s", uid), lastMsg{to: to, msgId: msgId}).Err()
}

func GetTalkLastReadMessageId(uid, to string) int64 {
	s := cache.DB.HGet(context.TODO(), fmt.Sprintf("last:aync:%s", uid), "msgId")
	msgId, _ := s.Int64()
	return msgId
}
