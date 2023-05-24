package main

import (
	"github.com/gin-gonic/gin"
	"ichat/http/router"
)

func main() {
	r := gin.Default()

	//初始化路由
	router.Routers(r)
	//初始化ichat模型
	//model.DB = db.DB

	r.Run("127.0.0.1:80")
}
