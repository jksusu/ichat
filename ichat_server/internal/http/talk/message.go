package talk

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"ichat/internal/area/talk"
	"ichat/pkg/tools/response"
)

type (
	record struct {
		MsgId int64  `json:"msgId" binding:"required"`
		To    string `json:"to" binding:"required"`
	}
)

func GetMessageRecord(c *gin.Context) {
	p := &record{}
	if err := c.ShouldBindWith(&p, binding.JSON); err != nil {
		response.ValidatorError.Json(c)
		return
	}
	info, err := talk.MessageRecord(c, c.Value("uid").(string), p.To, p.MsgId)
	if err != nil {
		response.Fail.WithMsg(err.Error()).Json(c)
		return
	}
	response.Ok.WithData(info).Json(c)
}
