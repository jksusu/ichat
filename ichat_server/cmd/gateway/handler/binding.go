package handler

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"ichat/pkg/db"
	"strconv"
)

var ctx = context.Background()

const (
	ConnToUid   = "conn:idToUid:"
	UidToConn   = "conn:uidToId:"
	CacheToken  = "token:"
	NeverExpire = 0
)

// GetUidByToken 根据token获取uid
func GetUidByToken(token string) (uint64, error) {
	val, err := db.Redis.Get(ctx, fmt.Sprintf("%s%v", CacheToken, token)).Result()
	if err == redis.Nil || err != nil {
		return 0, err
	}

	s, _ := strconv.ParseUint(val, 10, 64)

	return s, nil
}

// AddConnBidirectionalBindingUid 连接和uid双向绑定
func AddConnBidirectionalBindingUid(connId string, uid uint64) {
	db.Redis.Set(ctx, fmt.Sprintf("%s%s", ConnToUid, connId), uid, NeverExpire)
	db.Redis.Set(ctx, fmt.Sprintf("%s%d", UidToConn, uid), connId, NeverExpire)
}

// DeleteBindingRelationship 删除connId 和 uid 的双向绑定关系
func DeleteBindingRelationship(connId string) {
	uid := GetUidByConnId(connId)
	db.Redis.Del(ctx, fmt.Sprintf("%s%s", ConnToUid, connId))
	db.Redis.Del(ctx, fmt.Sprintf("%s%v", UidToConn, uid))
}

// GetUidByConnId 根据连接id connId 获取uid
func GetUidByConnId(connId string) string {
	val, err := db.Redis.Get(ctx, fmt.Sprintf("%s%s", ConnToUid, connId)).Result()
	if err == redis.Nil {
		logrus.Infof("getBindingConnByConnId no data connId:%s", connId)
	} else if err != nil {
		logrus.Error(err)
	}
	return val
}
