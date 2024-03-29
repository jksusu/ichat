package connect

import (
	"fmt"
	"ichat/pkg/cache"
)

func getKey(key, val string) string {
	return fmt.Sprintf("%s%s", key, val)
}

func GetConnIdByUid(uid string) (string, error) {
	return cache.DB.Get(cache.Ctx, getKey(cache.CONNECT_UID_TO_CONNID, uid)).Result()
}

func GetUidByConnId(connId string) (string, error) {
	return cache.DB.Get(cache.Ctx, getKey(cache.CONNECT_CONNID_TO_UID, connId)).Result()
}

func GetUidByToken(token string) (string, error) {
	return cache.DB.HGet(cache.Ctx, getKey(cache.USER_TOKEN, token), "uid").Result()
}

// AddConnBidirectionalBindingUid  初始化连接双向绑定uid和connId
func AddConnBidirectionalBindingUid(connId, uid string) {
	cache.DB.Set(cache.Ctx, getKey(cache.CONNECT_CONNID_TO_UID, connId), uid, 0)
	cache.DB.Set(cache.Ctx, getKey(cache.CONNECT_UID_TO_CONNID, uid), connId, 0)
}

// DeleteBindingRelationship 关闭连接 删除connId 和 uid 的双向绑定关系
func DeleteBindingRelationship(connId string) error {
	uid, err := GetUidByConnId(connId)
	if err != nil {
		return err
	}
	a := cache.DB.Del(cache.Ctx, getKey(cache.CONNECT_CONNID_TO_UID, connId))
	b := cache.DB.Del(cache.Ctx, getKey(cache.CONNECT_UID_TO_CONNID, uid))
	if a.Err() != nil || b.Err() != nil {
		return err
	}
	return nil
}
