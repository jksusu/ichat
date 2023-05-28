package user

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"ichat/pkg/cache"
	"time"
)

type User struct {
	Token        string `redis:"token"json:"token"`
	Id           int    `redis:"id"json:"id"`
	Uid          string `redis:"uid"json:"uid"`
	Username     string `redis:"username"json:"username"`
	Nickname     string `redis:"nickname"json:"nickname"`
	HeadPortrait string `redis:"head_portrait"json:"head_portrait"`
	Status       int    `redis:"status"json:"status"`
}

var Rdb = new(User)

func (c *User) TokenCache(token string, user *User, expire time.Duration) bool {
	if err := cache.DB.HSet(context.TODO(), fmt.Sprintf("user:token:%s", token), user).Err(); err != nil {
		return false
	}
	return true
}

// GetUserByToken 获取缓存
func (c *User) GetUserByToken(key string) (*User, error) {
	var u User
	err := cache.DB.HGetAll(context.TODO(), fmt.Sprintf("user:token:%s", key)).Scan(&u)
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
func (c *User) DeleteUserCacheByToken(token string) bool {
	_, err := cache.DB.Del(context.TODO(), fmt.Sprintf("user:token:%s", token)).Result()
	if err != nil {
		return false
	}
	return true
}
