package user

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"ichat/pkg/cache"
	"time"
)

func CreateUserCacheService() *Users {
	return &Users{}
}

type Users struct {
	Token        string `redis:"token"json:"token"`
	Id           int    `redis:"id"json:"id"`
	Uid          string `redis:"uid"json:"uid"`
	Username     string `redis:"username"json:"username"`
	Nickname     string `redis:"nickname"json:"nickname"`
	HeadPortrait string `redis:"head_portrait"json:"head_portrait"`
	Status       int    `redis:"status"json:"status"`
}

var (
	ctx       = context.Background()
	UserCache = new(Users)
)

func (c *Users) TokenCache(token string, user *Users, expire time.Duration) bool {
	if err := cache.DB.HSet(ctx, fmt.Sprintf("user:token:%s", token), user).Err(); err != nil {
		return false
	}
	return true
}

// GetUserByToken 获取缓存
func (c *Users) GetUserByToken(key string) (*Users, error) {
	var u Users
	err := cache.DB.HGetAll(ctx, fmt.Sprintf("user:token:%s", key)).Scan(&u)
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
func (c *Users) DeleteUserCacheByToken(token string) bool {
	_, err := cache.DB.Del(ctx, fmt.Sprintf("user:token:%s", token)).Result()
	if err != nil {
		return false
	}
	return true
}
