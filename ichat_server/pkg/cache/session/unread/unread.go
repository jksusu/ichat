package unread

import (
	"fmt"
	"ichat/pkg/cache"
	"strconv"
)

type Unread struct {
}

func getKey(UID string) string {
	return fmt.Sprintf("%s%s", cache.SESSION_UNREAD, UID)
}

func Add(UID, toUID string, number int64) error {
	return cache.DB.HSet(cache.Ctx, getKey(UID), toUID, number).Err()
}

func Sub(UID, toUID string, number int) error {
	val := cache.DB.HGet(cache.Ctx, getKey(UID), toUID).Val()
	count, _ := strconv.Atoi(val)
	n := count - number
	if n < 0 {
		n = 0
	}
	return cache.DB.HSet(cache.Ctx, getKey(UID), toUID, n).Err()
}

func Get(UID string, toUID string) int {
	v := cache.DB.HGet(cache.Ctx, getKey(UID), toUID).Val()

	n, _ := strconv.Atoi(v)

	return n
}
