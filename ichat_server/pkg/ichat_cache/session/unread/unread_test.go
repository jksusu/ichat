package unread

import (
	"ichat"
	"ichat/pkg/db"
	"testing"
)

func Init() {
	ichat.InitConf("../../../../conf.yaml")
	db.InitMysql(ichat.MysqlConf)
	db.InitRedis(ichat.RedisConf)
}

func TestAdd(t *testing.T) {
	Init()
	err := Add("11111", "22222", 100)
	if err != nil {
		t.Error(err)
	}
}

func TestSub(t *testing.T) {
	Init()
	err := Sub("11111", "22222", 50)
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	Init()
	res := Get("11111", "22222")

	t.Log(res)
}
