package unread

import (
	"fmt"
	"ichat/pkg/db"
	"ichat/pkg/ichat_cache"
	"strconv"
	"time"
)

type UnReadMessage struct {
	MessageBelongToUid string    `redis:"messageBelongToUid"`
	MessageType        string    `redis:"messageType"`
	MessageContent     string    `redis:"messageContent"`
	MessageSendTime    time.Time `redis:"messageSendTime"`
}

func getKye(uid string) string {
	return fmt.Sprintf("%s:%s", ichat_cache.UNREAD_MSG_ALL, uid)
}

func SetContactsLastMsg(uid, contactsUid string, msgId int64) error {
	k := fmt.Sprintf("%s%s", ichat_cache.CONTACTS_UNREAD, uid)
	return db.Redis.HSet(ichat_cache.Ctx, k, contactsUid, msgId).Err()
}

func GetContactsLastMsg(uid, contactsUid string) (int64, error) {
	k := fmt.Sprintf("%s%s", ichat_cache.CONTACTS_UNREAD, uid)
	s, err := db.Redis.HGet(ichat_cache.Ctx, k, contactsUid).Result()
	if err != nil {
		return 0, err
	}
	n, _ := strconv.Atoi(s)

	return int64(n), err
}

// 添加未读消息
func AddUnReadMsg(unread *UnReadMessage) error {
	return db.Redis.HSet(ichat_cache.Ctx, getKye(unread.MessageBelongToUid), unread).Err()
}

// 获取所有未读消息
func GetAllReadMsg(uid string) (unread *UnReadMessage, err error) {
	err = db.Redis.HGetAll(ichat_cache.Ctx, getKye(uid)).Scan(&unread)

	return
}

func DelUnReadMsgByMsgId(uid string, msgId int64) {
	db.Redis.Del(ichat_cache.Ctx, getKye(uid))
}
