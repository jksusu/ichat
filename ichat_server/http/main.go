package main

import (
	"github.com/gin-gonic/gin"
	"ichat"
	"ichat/http/router"
	"ichat/pkg/db"
)

func main() {
	r := gin.Default()

	ichat.InitConf("./conf.yaml")
	db.InitMysql(ichat.MysqlConf)
	db.InitRedis(ichat.RedisConf)

	//初始化路由
	router.Routers(r)
	//初始化ichat模型
	model.DB = db.DB

	r.Run("127.0.0.1:80")
}
