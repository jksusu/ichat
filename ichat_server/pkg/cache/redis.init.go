package cache

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/redis/go-redis/v9"
	"ichat"
)

var (
	Ctx = context.Background()
	DB  *redis.Client
)

const (
	SESSION_LIST         = "session:list:"         //会话列表
	SESSION_UNREAD       = "session:unread:"       //未读
	SESSION_COUNT_UNREAD = "session:count_unread:" //总未读

	//连接
	CONNECT_CONNID_TO_UID = "connect:connid_to_uid:"
	CONNECT_UID_TO_CONNID = "connect:uid_to_connid:"

	//用户
	USER_TOKEN = "user:token:"

	//未读消息
	CONTACTS_UNREAD = "unread:contacts:"
	GROUP_UNREAD    = "unread:group:"
	UNREAD_MSG_ALL  = "unreadmsg:all"

	MESSAGE_STATUS = "message:status:"
)

func InitRedis(conf *ichat.RedisConfig) {
	DB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.Database,
	})
	glog.Infoln("redis database:%s:%d connection succeeded", conf.Host, conf.Port)
}
