package unread_message

import "ichat/pkg/ichat_model"

type UnReadMessageList struct {
	*ichat_model.UnreadMessageList
}

func (u *UnReadMessageList) Add(unread *ichat_model.UnreadMessageList) (bool, error) {
	if result := ichat_model.DB.Create(&unread); result.RowsAffected > 0 {
		return true, nil
	} else {
		return false, result.Error
	}
}

func (u *UnReadMessageList) DelUnReadMessage(msgId int64) error {
	return ichat_model.DB.Where("msg_id = ?", msgId).Delete(&u).Error
}
