package user_cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"ichat/pkg/db"
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

var UserCache = new(User)

func (cache *User) TokenCache(token string, user *User, expire time.Duration) bool {
	if err := db.RDB.HSet(context.TODO(), fmt.Sprintf("user:token:%s", token), user).Err(); err != nil {
		return false
	}
	return true
}

// GetUserByToken 获取缓存
func (cache *User) GetUserByToken(key string) (*User, error) {
	var u User
	err := db.RDB.HGetAll(context.TODO(), fmt.Sprintf("user:token:%s", key)).Scan(&u)
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
func (cache *User) DeleteUserCacheByToken(token string) bool {
	_, err := db.RDB.Del(context.TODO(), fmt.Sprintf("user:token:%s", token)).Result()
	if err != nil {
		return false
	}
	return true
}
