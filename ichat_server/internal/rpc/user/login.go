package user

import (
	"errors"
	"ichat/pkg/db"
	"ichat/pkg/ichat_cache/user_cache"
	"ichat/pkg/ichat_tool/md5"
	"ichat/pkg/ichat_tool/rand"
	"ichat/pkg/protocol/pb"
	"time"
)

var User = new(user)

type user struct{}

func (*user) Login(username string, password string) (*pb.LoginResp, error) {
	list := model.UserModel.ListByUsername(username)
	if list == nil {
		return nil, errors.New("账号不存在")
	}
	if md5.PwdConfound(password, list.Salt) != list.Password {
		return nil, errors.New("密码错误，请重试")
	}
	data, _ := user_cache.UserCache.GetUserByToken(list.Token)
	if data.Token != "" {
		return &pb.LoginResp{Uid: data.Uid, Nickname: list.Nickname, Token: data.Token}, nil
	}
	//生成token
	token := rand.GenToken()
	userInfo := &user_cache.User{
		Token:        token,
		Id:           data.Id,
		Uid:          data.Uid,
		Username:     username,
		Nickname:     data.Nickname,
		Status:       data.Status,
		HeadPortrait: data.HeadPortrait,
	}
	//缓存token
	if ok := user_cache.UserCache.TokenCache(token, userInfo, time.Hour*24); !ok {
		return &pb.LoginResp{Uid: data.Uid, Nickname: list.Nickname, Token: token}, errors.New("登录失败，请重试")
	}
	list.Token = token
	db.DB.Save(&list)

	return &pb.LoginResp{Uid: data.Uid, Nickname: list.Nickname, Token: token}, nil
}
