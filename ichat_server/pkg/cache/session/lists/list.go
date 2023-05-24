package lists

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"ichat/pkg/cache"
	"time"
)

type SessionList struct {
	UID                 string `json:"uid"`
	UnreadMessageNumber int    `json:"unread_message_number"` //未读消息数量
}

var ctx = context.Background()

func getKey(UID string) string {
	return fmt.Sprintf("%v%v", cache.SESSION_LIST, UID)
}

func Add(UID, toUID string) error {
	d := redis.Z{
		Score:  float64(time.Now().UnixNano()),
		Member: toUID,
	}
	return cache.DB.ZAdd(ctx, getKey(UID), d).Err()
}

func Del(UID, toUID string) error {
	return cache.DB.Del(ctx, getKey(UID), toUID).Err()
}

func GetAll(UID string) ([]string, error) {
	return cache.DB.ZRange(cache.Ctx, getKey(UID), 0, -1).Result()
}
