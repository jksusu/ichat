package model

type Users struct {
	GLOBALMODEL
	UID         string `json:"uid" gorm:"index;comment:用户uid"`
	Nickname    string `json:"nickname" gorm:"comment:用户昵称"`
	Username    string `json:"username" gorm:"uniqueIndex;comment:用户名"`
	Password    string `json:"-" gorm:"comment:登录密码"`
	Salt        string `json:"-" gorm:"comment:密码盐"`
	Avatar      string `json:"avatar" gorm:"comment:头像"`
	Status      int    `json:"status" gorm:"default:1;comment:账号状态 1:正常 2:封号中 3:拉黑"`
	LastLoginIp string `json:"last_login_ip" gorm:"column:last_login_ip;comment:最后登录ip地址"`
	Token       string `json:"token" gorm:"column:token;comment:token"`
}

var UserModel = new(Users)

func (*Users) GetUserByUid(uid string) (user *Users, err error) {
	result := DB.First(&user, "uid = ?", uid)
	if result.RowsAffected < 1 {
		return nil, result.Error
	}
	return user, nil
}

func (*Users) GetUserByUsername(username string) (user *Users, err error) {
	result := DB.First(&user, "username = ?", username)
	if result.RowsAffected < 1 {
		return nil, result.Error
	}
	return user, nil
}

func (u *Users) Insert(users *Users) bool {
	result := DB.Create(users)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func (u *Users) FindInIds(ids []string) ([]Users, error) {
	var user []Users
	err := DB.Model(&u).Where("uid IN (?)", ids).Find(&user).Error
	return user, err
}

// 查询用户名是否存在
func (u *Users) WhereUsernameIsExists(username string) bool {
	result := DB.First(&u, "username = ?", username)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

// 用户注册
func (u *Users) Register(uid, username, password, slot, ip string) bool {
	result := DB.Create(&Users{
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

// 更新token
func (u *Users) UpdateToken(uid, token string) bool {
	result := DB.First(&u, "uid=?", uid)
	if result.RowsAffected > 0 {
		result.Update("token", token)
		return true
	}
	return true
}

// 更新token
func (u *Users) UpdatePassword(uid, password, salt string) bool {
	result := DB.First(&u, "uid=?", uid)
	if result.RowsAffected > 0 {
		u.Token = ""
		u.Password = password
		u.Salt = salt
		DB.Save(&u)
	}
	return true
}
