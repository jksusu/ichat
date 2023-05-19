package model

import "time"

type Friends struct {
	ID        int64     `json:"id" gorm:"comment:主键id"`
	UserUid   string    `json:"user_uid" gorm:"comment:我的uid"`
	FriendUid string    `json:"friend_uid" gorm:"uniqueIndex;comment:联系人uid"`
	Remarks   string    `json:"remarks" gorm:"comment:备注"`
	Status    int       `json:"status" gorm:"comment:状态，1：正常，2：拉黑"`
	Extra     string    `json:"extra" gorm:"comment:附加字段"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新时间"`
}

var FriendModel = new(Friends)

// PageQuery 分页
func (f *Friends) PageQuery(userUid string, page, pageSize int) ([]Friends, int64, error) {
	var (
		friends []Friends
		total   int64
		err     error
	)
	offset := (page - 1) * pageSize
	//查询总数
	if err = DB.Model(&Friends{}).Where("user_uid = ?", userUid).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = DB.Model(&Friends{}).Where("user_uid= ?", userUid).Limit(pageSize).Offset(offset).Find(&friends).Error
	return friends, total, err
}

// AddFriend 添加一条数据
func (f *Friends) AddFriend(list *Friends) bool {
	if DB.Create(list).RowsAffected > 0 {
		return true
	}
	return false
}

// StatusUpdate 好友状态修改
func (f *Friends) StatusUpdate(userUid, friendUid string, status int) bool {
	result := DB.First(f, "user_uid = ? AND friend_uid = ?", userUid, friendUid)
	if result.RowsAffected > 0 {
		f.Status = status
		result.Save(f)
		return true
	}
	return false
}
