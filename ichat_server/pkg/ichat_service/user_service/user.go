package user_service

import (
	"ichat/pkg/ichat_model/user_model"
)

func UserService() *User {
	return &User{}
}

type User struct {
	*user_model.Users
}

func FindUserByUid(uid string) (user *user_model.Users, err error) {
	var u user_model.Users
	user, err = u.GetUserByUid(uid)
	return
}
