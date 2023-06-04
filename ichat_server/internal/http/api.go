package http

import (
	"github.com/gin-gonic/gin"
	"ichat/internal/http/contacts"
	"ichat/internal/http/group"
	"ichat/internal/http/session"
	"ichat/internal/http/talk"
	"ichat/internal/http/user"
	cache "ichat/pkg/cache/user"
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
		info, _ := cache.Rdb.GetUserByToken(token)
		if len(info.Token) == 0 {
			response.Fail.WithCodeMessage(http.StatusUnauthorized, "请登录后在试").Json(c)
			c.Abort()
		}
		c.Set("uid", info.Uid)
		c.Next()
	}
}

func RegisterApi() *gin.Engine {
	router := gin.Default()
	router.Use(cors())
	v1 := router.Group("/v1")
	{
		v1.POST("/login", user.Login)
		v1.POST("/register", user.Register)
		v1.GET("/code", user.GetCode)
		v1.POST("/scan_code_login")

		v1contacts := v1.Group("contacts").Use(tokenInspect())
		{
			v1contacts.POST("/lists", contacts.Lists)
			v1contacts.POST("/del", contacts.Del)
			v1contacts.POST("/edit", contacts.Edit)
		}
		v1group := v1.Group("group").Use(tokenInspect())
		{
			v1group.POST("/lists", group.Lists)
			v1group.POST("/del", group.Del)
			v1group.POST("/edit", group.Edit)
			v1group.POST("/member/lists", group.MemberLists)
			v1group.POST("/member/remove", group.MemberRemove)
		}
		v1session := v1.Group("session").Use(tokenInspect())
		{
			v1session.POST("/lists", session.Lists)
			v1session.POST("/del", session.Del)
			v1session.POST("/edit", session.Edit)
		}
		v1user := v1.Group("user").Use(tokenInspect())
		{
			v1user.POST("/info", user.GetUserInfo) //用户信息
		}
		v1message := v1.Group("message").Use(tokenInspect())
		{
			v1message.POST("/record", talk.GetMessageRecord) //用户信息
		}
	}

	return router
}
