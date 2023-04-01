package connect

import (
	"fmt"
	"ichat/pkg/db"
	"ichat/pkg/ichat_cache"
)

func getKey(key, val string) string {
	return fmt.Sprintf("%s%s", key, val)
}

func GetConnIdByUID(uid string) (string, error) {
	return db.Redis.Get(ichat_cache.Ctx, getKey(ichat_cache.CONNECT_UID_TO_CONNID, uid)).Result()
}

func GetUIDbyConnId(connId string) (string, error) {
	return db.Redis.Get(ichat_cache.Ctx, getKey(ichat_cache.CONNECT_CONNID_TO_UID, connId)).Result()
}

func GetUIDbyToken(token string) (string, error) {
	return db.Redis.HGet(ichat_cache.Ctx, getKey(ichat_cache.USER_TOKEN, token), "uid").Result()
}

// 初始化连接双向绑定uid和connid
func AddConnBidirectionalBindingUid(connId, uid string) {
	db.Redis.Set(ichat_cache.Ctx, getKey(ichat_cache.CONNECT_CONNID_TO_UID, connId), uid, 0)
	db.Redis.Set(ichat_cache.Ctx, getKey(ichat_cache.CONNECT_UID_TO_CONNID, uid), connId, 0)
}

// 关闭连接 删除connId 和 uid 的双向绑定关系
func DeleteBindingRelationship(connId string) error {
	uid, err := GetUIDbyConnId(connId)
	if err != nil {
		return err
	}
	a := db.Redis.Del(ichat_cache.Ctx, getKey(ichat_cache.CONNECT_CONNID_TO_UID, connId))
	b := db.Redis.Del(ichat_cache.Ctx, getKey(ichat_cache.CONNECT_UID_TO_CONNID, uid))
	if a.Err() != nil || b.Err() != nil {
		return err
	}
	return nil
}
