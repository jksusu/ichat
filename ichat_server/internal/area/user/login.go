package user

import (
	"errors"
	"ichat/pkg/cache/user"
	cache "ichat/pkg/cache/user"
	"ichat/pkg/model"
	"ichat/pkg/tools/md5"
	"ichat/pkg/tools/rand"
	"time"
)

type info struct {
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
	Avatar   string `json:"avatar"`
}

func Login(username string, password string) (*info, error) {
	list := model.UserModel.ListByUsername(username)
	if list == nil {
		return nil, errors.New("账号不存在")
	}
	if md5.PwdConfound(password, list.Salt) != list.Password {
		return nil, errors.New("密码错误，请重试")
	}

	data, _ := cache.Rdb.GetUserByToken(list.Token)
	if data.Token != "" {
		return &info{Uid: list.UID, Nickname: list.Nickname, Token: data.Token, Avatar: list.HeadPortrait}, nil
	}
	//生成token
	token := rand.GenToken()
	if ok := cache.Rdb.TokenCache(token, &user.User{
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

	return &info{Uid: list.UID, Nickname: list.Nickname, Token: data.Token, Avatar: list.HeadPortrait}, nil
}
