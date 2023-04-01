package count_unread

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
	Add("11111", 100)
}

func TestSub(t *testing.T) {
	Init()
	Sub("11111", 50)
}

func TestGet(t *testing.T) {
	Init()
	n := Get("11111")
	t.Log(n)
}
