package middleware

import (
	"github.com/gin-gonic/gin"
	"ichat/http/service/user"
	"ichat/pkg/response"
	"net/http"
)

func TokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			response.Fail.WithCodeMessage(http.StatusUnauthorized, "token不存在，请登录后再试").Json(c)
			c.Abort()
		}
		//获取 user
		info, _ := user.CreateUserCacheService().GetUserByToken(token)
		if len(info.Token) == 0 {
			response.Fail.WithCodeMessage(http.StatusUnauthorized, "请登录后在试").Json(c)
			c.Abort()
		}
		c.Set("uid", info.Uid)
		c.Next()
	}
}
