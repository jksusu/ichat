package lists

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"ichat/pkg/db"
	"ichat/pkg/ichat_cache"
	"time"
)

type SessionList struct {
	UID                 string `json:"uid"`
	UnreadMessageNumber int    `json:"unread_message_number"` //未读消息数量
}

var ctx = context.Background()

func getKey(UID string) string {
	return fmt.Sprintf("%v%v", ichat_cache.SESSION_LIST, UID)
}

func Add(UID, toUID string) error {
	d := redis.Z{
		Score:  float64(time.Now().UnixNano()),
		Member: toUID,
	}
	return db.RDB.ZAdd(ctx, getKey(UID), d).Err()
}

func Del(UID, toUID string) error {
	return db.RDB.Del(ctx, getKey(UID), toUID).Err()
}

func GetAll(UID string) ([]string, error) {
	return db.RDB.ZRange(ichat_cache.Ctx, getKey(UID), 0, -1).Result()
}
