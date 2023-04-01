package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ichat/cmd/http/app/service/chat"
	"ichat/pkg/response"
)

type SingleRecord struct {
	Page int    `json:"page"`
	To   string `json:"to"`
}

func GetSingleChatRecord(c *gin.Context) {
	uid := c.GetString("uid")
	var record SingleRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		response.ValidatorError.WithMsg(err.Error()).Json(c)
		return
	}
	validate := validator.New()
	if err := validate.Struct(record); err != nil {
		response.ValidatorError.WithMsg(err.Error()).Json(c)
		return
	}
	chat.RecordService().SingleChatRecord(uid, record.To, record.Page, 20).Json(c)
}
