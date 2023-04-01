package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ichat/cmd/http/app/service/contacts"
	"ichat/pkg/response"
)

type Paging struct {
	Page int `json:"page"`
}

// ContactsLists 联系人列表
func ContactsLists(c *gin.Context) {
	uid := c.GetString("uid")

	var contactsPaging Paging
	if err := c.ShouldBindJSON(&contactsPaging); err != nil {
		response.ValidatorError.WithMsg(err.Error()).Json(c)
		return
	}
	validate := validator.New()
	if err := validate.Struct(contactsPaging); err != nil {
		response.ValidatorError.WithMsg(err.Error()).Json(c)
		return
	}

	contacts.ContactsService().FriendList(uid, contactsPaging.Page, 20).Json(c)
}

func GroupLists(c *gin.Context) {

}
