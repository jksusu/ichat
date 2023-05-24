package count_unread

import (
	"fmt"
	"ichat/pkg/cache"
	"strconv"
)

func getKey(UID string) string {
	return fmt.Sprintf("%v%v", cache.SESSION_COUNT_UNREAD, UID)
}

func Add(UID string, number int64) error {
	return cache.DB.IncrBy(cache.Ctx, getKey(UID), number).Err()
}

func Get(UID string) int {
	v, err := cache.DB.Get(cache.Ctx, getKey(UID)).Result()
	if err != nil {
		return 0
	}
	n, _ := strconv.Atoi(v)
	return n
}

func Sub(UID string, number int64) error {
	return cache.DB.DecrBy(cache.Ctx, getKey(UID), number).Err()
}
