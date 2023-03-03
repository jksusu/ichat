package db

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"ichat"
)

var Redis *redis.Client

func InitRedis(conf *ichat.RedisConfig) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.Database,
	})

	logrus.Infof("redis database:%s:%d connection succeeded", conf.Host, conf.Port)

	Redis = rdb
}
