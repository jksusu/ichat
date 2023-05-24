package user

import (
	"errors"
	"ichat/pkg/cache/user_cache"
	"ichat/pkg/ichat_tool/md5"
	"ichat/pkg/ichat_tool/rand"
	"ichat/pkg/model"
	"time"
)

var User = new(user)

type user struct {
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
	Avatar   string `json:"avatar"`
}

func (*user) Login(username string, password string) (*user, error) {
	list := model.UserModel.ListByUsername(username)
	if list == nil {
		return nil, errors.New("账号不存在")
	}
	if md5.PwdConfound(password, list.Salt) != list.Password {
		return nil, errors.New("密码错误，请重试")
	}
	data, _ := user_cache.UserCache.GetUserByToken(list.Token)
	if data.Token != "" {
		return &user{Uid: list.UID, Nickname: list.Nickname, Token: data.Token, Avatar: list.HeadPortrait}, nil
	}
	//生成token
	token := rand.GenToken()
	if ok := user_cache.UserCache.TokenCache(token, &user_cache.User{
		Token:        token,
		Id:           data.Id,
		Uid:          data.Uid,
		Username:     username,
		Nickname:     data.Nickname,
		Status:       data.Status,
		HeadPortrait: data.HeadPortrait,
	}, time.Hour*24); !ok {
		return nil, errors.New("登录失败，请重试")
	}

	list.Token = token
	model.DB.Save(&list)

	return &user{Uid: list.UID, Nickname: list.Nickname, Token: data.Token, Avatar: list.HeadPortrait}, nil
}
