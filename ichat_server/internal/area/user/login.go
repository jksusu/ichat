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
	list, err := model.UserModel.GetUserByUsername(username)
	if err != nil || list == nil {
		return nil, errors.New("账号不存在")
	}
	if md5.PwdConfound(password, list.Salt) != list.Password {
		return nil, errors.New("密码错误，请重试")
	}
	data, _ := cache.Rdb.GetUserByToken(list.Token)
	if data.Token != "" {
		return &info{Uid: list.UID, Nickname: list.Nickname, Token: data.Token, Avatar: list.Avatar}, nil
	}
	//生成token
	token := rand.GenToken()
	if ok := cache.Rdb.TokenCache(token, &user.User{
		Token:    token,
		Id:       list.ID,
		Uid:      list.UID,
		Username: username,
		Nickname: list.Nickname,
		Status:   list.Status,
		Avatar:   list.Avatar,
	}, time.Hour*24); !ok {
		return nil, errors.New("登录失败，请重试")
	}
	list.Token = token
	model.DB.Save(&list)

	return &info{Uid: list.UID, Nickname: list.Nickname, Token: token, Avatar: list.Avatar}, nil
}

func Register(username, pwd, nickname string) error {
	user, err := model.UserModel.GetUserByUsername(username)
	if err != nil || user.Username != "" {
		return errors.New("注册失败")
	}
	if ok := model.UserModel.Insert(&model.Users{
		UID:      "",
		Nickname: nickname,
		Username: username,
		Password: "",
		Salt:     "",
	}); !ok {
		return errors.New("注册失败")
	}
	return nil
}

func Logout(uid, channel string) {
	//修改redsi状态
	cache.Rdb.DeleteUserCacheByUid(uid)

	//修改mysql user state
}
