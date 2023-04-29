package user_model

import "ichat/pkg/ichat_model"

type Users struct {
	ichat_model.Users
}

func (u *Users) GetUserByUid(uid string) (user *Users, err error) {
	result := ichat_model.DB.First(&user, "uid = ?", uid)
	if result.RowsAffected < 1 {
		return nil, result.Error
	}
	return user, nil
}

func (u *Users) FindInIds(ids []string) ([]Users, error) {
	var user []Users
	err := ichat_model.DB.Model(&u).Where("uid IN (?)", ids).Find(&user).Error
	return user, err
}

// 查询用户名是否存在
func (u *Users) WhereUsernameIsExists(username string) bool {
	result := ichat_model.DB.First(&u, "username = ?", username)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

// 用户注册
func (u *Users) Register(uid, username, password, slot, ip string) bool {
	result := ichat_model.DB.Create(&ichat_model.Users{
		UID:         uid,
		Username:    username,
		Password:    password,
		Salt:        slot,
		LastLoginIp: ip,
	})
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

// 查询单条记录
func (u *Users) ListByUsername(username string) *Users {
	result := ichat_model.DB.First(&u, "username=?", username)
	if result.Error != nil {
		return nil
	}
	return u
}

// 更新token
func (u *Users) UpdateToken(uid, token string) bool {
	result := ichat_model.DB.First(&u, "uid=?", uid)
	if result.RowsAffected > 0 {
		result.Update("token", token)
		return true
	}
	return true
}

// 更新token
func (u *Users) UpdatePassword(uid, password, salt string) bool {
	result := ichat_model.DB.First(&u, "uid=?", uid)
	if result.RowsAffected > 0 {
		u.Token = ""
		u.Password = password
		u.Salt = salt
		ichat_model.DB.Save(&u)
	}
	return true
}
