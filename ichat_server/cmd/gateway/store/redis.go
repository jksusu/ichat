package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"ichat/pkg/db"
)

var ctx = context.Background()

const (
	ConnToUid   = "conn:idToUid:"
	UidToConn   = "conn:uidToId:"
	CacheToken  = "token:"
	NeverExpire = 0
)

// 获取连接id根据uid
func GetConnIdByUid(uid string) (val string, err error) {
	val, err = db.Redis.Get(ctx, fmt.Sprintf("%s%s", UidToConn, uid)).Result()
	if err == redis.Nil {
		err = errors.New(fmt.Sprintf("getBindingConnByConnId no data uid:%s", uid))
		return
	} else if err != nil {
		return
	}
	return
}
