package user

import (
	"ichat/pkg/db"
	"ichat/pkg/ichat_model/user_model"
	"ichat/pkg/ichat_tool/md5"
	"ichat/pkg/ichat_tool/rand"
	"ichat/pkg/response"
	"time"
)

func UserService() *User {
	return &User{}
}

type User struct {
	*user_model.Users
}

// Register 账号注册
func (u *User) Register(username, password, ip string) *response.Response {
	if u.WhereUsernameIsExists(username) {
		return response.Fail.WithMsg("账号" + username + "已经存在")
	}
	slot := rand.GenSlot()
	encryptionPass := md5.PwdConfound(password, slot)
	if u.Users.Register(rand.GenUid(), username, encryptionPass, slot, ip) {
		return response.Ok.WithMsg("注册成功")
	}

	return response.Fail.WithMsg("注册失败，请稍后再试")
}

// Login 登录接口
func (u *User) Login(username, password string) *response.Response {
	user := u.ListByUsername(username)
	if user != nil {
		if md5.PwdConfound(password, user.Salt) != user.Password {
			return response.Fail.WithMsg("密码错误，请重试")
		}
		//如果token存在
		if user.Token != "" {
			data, _ := CreateUserCacheService().GetUserByToken(user.Token)
			if data.Token != "" {
				return response.Ok.WithMsg("登录成功").WithData(data)
			}
		}
		//生成token
		token := rand.GenToken()
		userCacheData := &UserCache{
			Token:        token,
			Id:           user.ID,
			Uid:          user.UID,
			Username:     username,
			Nickname:     user.Nickname,
			Status:       user.Status,
			HeadPortrait: user.HeadPortrait,
		}
		//缓存token
		if ok := CreateUserCacheService().TokenCache(token, userCacheData, time.Hour*24); !ok {
			return response.Fail.WithMsg("登录失败，请重试")
		}
		//更新token
		user.Token = token
		db.DB.Save(&user)
		return response.Ok.WithMsg("登录成功").WithData(userCacheData)
	}
	return response.Fail.WithMsg("账号不存在，请重试")
}

// Logout 退出
func (u *User) Logout(token string) *response.Response {
	userCache, _ := CreateUserCacheService().GetUserByToken(token)
	if userCache.Token != "" {
		u.UpdateToken(userCache.Uid, "")
	}
	CreateUserCacheService().DeleteUserCacheByToken(userCache.Token)
	return response.Ok.WithMsg("账号退出成功")
}

// Detail 用户信息
func (u *User) Detail(token string) *response.Response {
	userCache, _ := CreateUserCacheService().GetUserByToken(token)
	if userCache.Token != "" {
		return response.Ok.WithData(userCache)
	}
	return response.Fail.WithMsg("未查询到用户信息")
}

// RetrievePassword 找回密码
func (u *User) RetrievePassword(code int, uid string, password string) *response.Response {
	//验证code
	if code != 6666 {
		return response.Fail.WithMsg("验证码错误，请重试")
	}

	slot := rand.GenSlot()
	encryptionPass := md5.PwdConfound(password, slot)

	//修改密码
	u.UpdatePassword(uid, encryptionPass, slot)

	return response.Ok.WithMsg("密码重置成功")
}

func (u *User) GetUserInfo(uid string) *response.Response {
	return response.Fail.WithMsg("未查询到用户信息")
}
