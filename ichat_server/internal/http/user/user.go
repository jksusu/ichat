package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"ichat/internal/area/user"
	"ichat/pkg/tools/response"
)

type (
	login struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	register struct {
		login
		Nickname string `json:"nickname" binding:"required"`
	}
)

func Login(c *gin.Context) {
	p := &login{}
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		response.ValidatorError.Json(c)
		return
	}
	info, err := user.Login(p.Username, p.Password)
	if err != nil {
		response.Fail.WithMsg(err.Error()).Json(c)
		return
	}
	response.Ok.WithData(info).Json(c)
}

func Register(c *gin.Context) {
	p := &register{}
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		response.ValidatorError.Json(c)
		return
	}
	if err := user.Register(p.Username, p.Username, p.Password); err != nil {
		response.Fail.WithMsg("注册失败，请重试").Json(c)
		return
	}
	response.Ok.WithMsg("注册成功").Json(c)
}

func GetCode(c *gin.Context) {

}

func GetUserInfo(c *gin.Context) {

}
