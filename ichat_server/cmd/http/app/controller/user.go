package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ichat/cmd/http/app/service/user"
	"ichat/pkg/response"
)

type UserController struct{}

type UserRegister struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Register 用户注册
func Register(c *gin.Context) {
	var userRegister UserRegister
	if err := c.ShouldBindJSON(&userRegister); err != nil {
		response.ValidatorError.WithMsg(err.Error()).Json(c)
		return
	}
	validate := validator.New()
	if err := validate.Struct(userRegister); err != nil {
		response.ValidatorError.WithMsg(err.Error()).Json(c)
		return
	}
	user.UserService().Register(userRegister.Username, userRegister.Password, c.RemoteIP()).Json(c)
}

// Login 登录账号
func Login(c *gin.Context) {
	var userRegister UserRegister
	if err := c.ShouldBindJSON(&userRegister); err != nil {
		response.ValidatorError.WithMsg(err.Error()).Json(c)
		return
	}
	validate := validator.New()
	if err := validate.Struct(userRegister); err != nil {
		response.ValidatorError.WithMsg(err.Error()).Json(c)
		return
	}
	user.UserService().Login(userRegister.Username, userRegister.Password).Json(c)
}

// RetrievePassword 找回密码
func RetrievePassword(c *gin.Context) {
	var Pwd struct {
		Uid      string `json:"uid" binding:"required"`
		Password string `json:"password" binding:"required"`
		Code     int    `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&Pwd); err != nil {
		response.ValidatorError.WithMsg(err.Error()).Json(c)
		return
	}
	validate := validator.New()
	if err := validate.Struct(Pwd); err != nil {
		response.ValidatorError.WithMsg(err.Error()).Json(c)
		return
	}
	user.UserService().RetrievePassword(Pwd.Code, Pwd.Uid, Pwd.Password).Json(c)
}
