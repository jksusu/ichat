package count_unread

import (
	"fmt"
	"ichat/pkg/db"
	"ichat/pkg/ichat_cache"
	"strconv"
)

func getKey(UID string) string {
	return fmt.Sprintf("%v%v", ichat_cache.SESSION_COUNT_UNREAD, UID)
}

func Add(UID string, number int64) error {
	return db.RDB.IncrBy(ichat_cache.Ctx, getKey(UID), number).Err()
}

func Get(UID string) int {
	v, err := db.RDB.Get(ichat_cache.Ctx, getKey(UID)).Result()
	if err != nil {
		return 0
	}
	n, _ := strconv.Atoi(v)
	return n
}

func Sub(UID string, number int64) error {
	return db.RDB.DecrBy(ichat_cache.Ctx, getKey(UID), number).Err()
}
