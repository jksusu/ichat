package session

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"ichat/internal/area/session"
	"ichat/pkg/tools/response"
)

type (
	SAdd struct {
		To          string `json:"to" binding:"required"`
		SessionType int    `json:"session_type" binding:"required"`
		LastMsgId   int64  `json:"last_msg_id" binding:"required"`
	}
	SEdit struct {
		To          string `json:"to" binding:"required"`
		SessionType int    `json:"session_type" binding:"required"`
		Topping     int    `json:"topping" binding:"required"`
		NoDisturb   int    `json:"no_disturb" binding:"required"`
	}
)

func Lists(c *gin.Context) {
	list, err := session.List(c.GetString("uid"))
	if err != nil {
		response.Fail.WithMsg(err.Error()).Json(c)
		return
	}
	response.Ok.WithData(list).Json(c)
}

func Add(c *gin.Context) {
	p := &SAdd{}
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		response.ValidatorError.Json(c)
		return
	}
	err := session.Add(c.GetString("uid"), p.To, p.SessionType, p.LastMsgId)
	if err != nil {
		response.Fail.WithMsg(err.Error()).Json(c)
		return
	}
	response.Ok.Json(c)
}

func Del(c *gin.Context) {
	if c.GetString("to") == "" {
		response.ValidatorError.Json(c)
		return
	}
	err := session.Remove(c.GetString("uid"))
	if err != nil {
		response.Fail.WithMsg(err.Error()).Json(c)
		return
	}
	response.Ok.Json(c)
}

func Edit(c *gin.Context) {
	p := &SEdit{}
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		response.ValidatorError.Json(c)
		return
	}
	err := session.Options(p.To, p.SessionType, p.Topping, p.NoDisturb)
	if err != nil {
		response.Fail.WithMsg(err.Error()).Json(c)
		return
	}
	response.Ok.Json(c)
}
