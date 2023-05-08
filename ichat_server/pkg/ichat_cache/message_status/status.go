package message_status

import (
	"fmt"
	"ichat/pkg/db"
	"ichat/pkg/ichat_cache"
)

func getKey(uid, touid string) string {
	return fmt.Sprintf("%s%s:to:%s", ichat_cache.MESSAGE_STATUS, uid, touid)
}

type M struct {
	LastReceiveMsgId     int64 `json:"last_receive_msg_id"`
	LastHasBeenSentMsgId int64 `json:"last_has_been_sent_msg_id"`
	LastReadMsgId        int64 `json:"last_read_msg_id"`
}

// 已发送
func UpdateMsgHasBeenSent(fromUID, toUID string, msgId int64) error {
	return db.RDB.HSet(ichat_cache.Ctx, getKey(fromUID, toUID), "last_has_been_sent_msg_id", msgId).Err()
}

// 已送达
func UpdateMsgReceive(fromUID, toUID string, msgId int64) error {
	return db.RDB.HSet(ichat_cache.Ctx, getKey(fromUID, toUID), "last_receive_msg_id", msgId).Err()
}

// 已读
func UpdateMessageRead(fromUID, toUID string, msgId int64) error {
	return db.RDB.HSet(ichat_cache.Ctx, getKey(fromUID, toUID), "last_read_msg_id", msgId).Err()
}

func GetAll(fromUID, toUID string, msgId int64) {

}
