package unread

import (
	"fmt"
	"ichat/pkg/db"
	"ichat/pkg/ichat_cache"
	"strconv"
)

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
