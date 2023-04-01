package lists

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
	err := Add("11111", "22222")
	t.Log(err)
	Add("11111", "33333")
	Add("11111", "55555")
	Add("11111", "44444")
}

func TestGetAll(t *testing.T) {
	Init()
	a, _ := GetAll("11111")
	t.Log(a)
}
