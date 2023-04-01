package session_service

import (
	"ichat"
	"ichat/pkg/db"
	"ichat/pkg/ichat_model"
	"testing"
)

func Init() {
	ichat.InitConf("../../../conf.yaml")
	db.InitMysql(ichat.MysqlConf)
	db.InitRedis(ichat.RedisConf)

	ichat_model.DB = db.DB
}

func TestGetSessionAll(t *testing.T) {
	Init()
	GetSessionAll("11111")
}
