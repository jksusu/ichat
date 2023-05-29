package http

import (
	"github.com/gin-gonic/gin"
	"ichat"
	"ichat/pkg/cache/user"
	"ichat/pkg/tools/response"
	"net/http"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Authorization,User-Agent, Keep-Alive, Content-Type, X-Requested-With,X-CSRF-Token,AccessToken")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusAccepted)
		}
		c.Next()
	}
}

func tokenInspect() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			response.Fail.WithCodeMessage(http.StatusUnauthorized, "token不存在，请登录后再试").Json(c)
			c.Abort()
		}
		info, _ := user.Rdb.GetUserByToken(token)
		if len(info.Token) == 0 {
			response.Fail.WithCodeMessage(http.StatusUnauthorized, "请登录后在试").Json(c)
			c.Abort()
		}
		c.Set("uid", info.Uid)
		c.Next()
	}
}

func RegisterApi() error {
	router := gin.Default()
	router.Use(cors())
	v1 := router.Group("/v1")
	{
		v1.POST("/login")
		v1.POST("/register")
		v1.GET("/code")
		v1.POST("/scan_code_login")

		contacts := v1.Group("contacts").Use(tokenInspect())
		{
			contacts.POST("/lists")
			contacts.POST("/del")
			contacts.POST("/edit")
		}
		group := v1.Group("group").Use(tokenInspect())
		{
			group.POST("/lists")
			group.POST("/del")
			group.POST("/edit")
			group.POST("/member/lists")
			group.POST("/member/remove")
		}
		session := v1.Group("session").Use(tokenInspect())
		{
			session.POST("/lists")
			session.POST("/del")
			session.POST("/edit")
		}
		user := v1.Group("user").Use(tokenInspect())
		{
			user.POST("/info") //用户信息
		}
	}

	return router.Run(ichat.GlobalConf.Other.HttpAddr)
}
