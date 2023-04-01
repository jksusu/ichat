package user

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"ichat/pkg/db"
	"time"
)

func CreateUserCacheService() *UserCache {
	return &UserCache{}
}

type UserCache struct {
	Token        string `redis:"token"json:"token"`
	Id           int    `redis:"id"json:"id"`
	Uid          string `redis:"uid"json:"uid"`
	Username     string `redis:"username"json:"username"`
	Nickname     string `redis:"nickname"json:"nickname"`
	HeadPortrait string `redis:"head_portrait"json:"head_portrait"`
	Status       int    `redis:"status"json:"status"`
}

var ctx = context.Background()

func (cache *UserCache) TokenCache(token string, user *UserCache, expire time.Duration) bool {
	if err := db.Redis.HSet(ctx, fmt.Sprintf("user:token:%s", token), user).Err(); err != nil {
		return false
	}
	return true
}

// GetUserByToken 获取缓存
func (cache *UserCache) GetUserByToken(key string) (*UserCache, error) {
	var u UserCache
	err := db.Redis.HGetAll(ctx, fmt.Sprintf("user:token:%s", key)).Scan(&u)
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &u, nil
}

// DeleteUserCacheByToken 删除token
func (cache *UserCache) DeleteUserCacheByToken(token string) bool {
	_, err := db.Redis.Del(ctx, fmt.Sprintf("user:token:%s", token)).Result()
	if err != nil {
		return false
	}
	return true
}
