package unread

import (
	"fmt"
	"ichat/pkg/db"
	"ichat/pkg/ichat_cache"
	"strconv"
)

type Unread struct {
}

func getKey(UID string) string {
	return fmt.Sprintf("%s%s", ichat_cache.SESSION_UNREAD, UID)
}

func Add(UID, toUID string, number int64) error {
	return db.Redis.HSet(ichat_cache.Ctx, getKey(UID), toUID, number).Err()
}

func Sub(UID, toUID string, number int) error {
	val := db.Redis.HGet(ichat_cache.Ctx, getKey(UID), toUID).Val()
	count, _ := strconv.Atoi(val)
	n := count - number
	if n < 0 {
		n = 0
	}
	return db.Redis.HSet(ichat_cache.Ctx, getKey(UID), toUID, n).Err()
}

func Get(UID string, toUID string) int {
	v := db.Redis.HGet(ichat_cache.Ctx, getKey(UID), toUID).Val()

	n, _ := strconv.Atoi(v)

	return n
}
