package model

type UnReadMessageList struct {
	*UnreadMessageList
}

func (u *UnReadMessageList) Add(unread *UnreadMessageList) (bool, error) {
	if result := DB.Create(&unread); result.RowsAffected > 0 {
		return true, nil
	} else {
		return false, result.Error
	}
}

func (u *UnReadMessageList) DelUnReadMessage(msgId int64) error {
	return DB.Where("msg_id = ?", msgId).Delete(&u).Error
}
