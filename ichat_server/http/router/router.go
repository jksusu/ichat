package router

import (
	"github.com/gin-gonic/gin"
	"ichat/http/controller"
	"ichat/http/middleware"
)

func Routers(r *gin.Engine) {
	//全局中间件
	r.Use(middleware.Cors())

	//v1版本
	v1 := r.Group("/v1")
	{
		v1.POST("/register", controller.Register)
		v1.POST("/login", controller.Login)

		////需要鉴权的api
		user := v1.Group("/user").Use(middleware.TokenCheck())
		{
			user.POST("/retrieve_password", controller.RetrievePassword)
		}
		relation := v1.Group("/relation").Use(middleware.TokenCheck())
		{
			relation.POST("/contacts", controller.ContactsLists)
			relation.POST("/group", controller.GroupLists)
		}
		chat := v1.Group("/chat").Use(middleware.TokenCheck())
		{
			chat.POST("/getSingleChatRecord", controller.GetSingleChatRecord)
		}
	}
}
