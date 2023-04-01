package friend_model

import (
	"ichat/pkg/ichat_model"
)

type Friends struct {
	*ichat_model.Friends
}

// PageQuery 分页
func (f *Friends) PageQuery(userUid string, page, pageSize int) ([]Friends, int64, error) {
	var (
		friends []Friends
		total   int64
		err     error
	)
	offset := (page - 1) * pageSize
	//查询总数
	if err := ichat_model.DB.Model(&Friends{}).Where("user_uid = ?", userUid).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = ichat_model.DB.Model(&Friends{}).Where("user_uid= ?", userUid).Limit(pageSize).Offset(offset).Find(&friends).Error

	return friends, total, err
}

// AddFriend 添加一条数据
func (f *Friends) AddFriend(list *Friends) bool {
	if ichat_model.DB.Create(list).RowsAffected > 0 {
		return true
	}
	return false
}

// StatusUpdate 好友状态修改
func (f *Friends) StatusUpdate(userUid, friendUid string, status int) bool {
	result := ichat_model.DB.First(f, "user_uid = ? AND friend_uid = ?", userUid, friendUid)
	if result.RowsAffected > 0 {
		f.Status = status
		result.Save(f)
		return true
	}
	return false
}
