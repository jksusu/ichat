package model

import (
	"time"
)

type RelationFriendsApply struct {
	MsgId       int64     `json:"msg_id" gorm:"comment:消息ID"`
	UID         string    `json:"uid" gorm:"comment:我的uid"`
	TO          string    `json:"to"`
	Remarks     string    `json:"remarks" `
	Status      int       `json:"status"`
	Extra       string    `json:"extra" gorm:"comment:附加字段"`
	ApplyTime   time.Time `json:"apply_time" gorm:"comment:发起时间"`
	ProcessTime time.Time `json:"process_time" gorm:"comment:处理时间"`
}

var RelationFriendsApplyModel = new(RelationFriendsApply)

func (r *RelationFriendsApply) Add(data *RelationFriendsApply) (ok bool, err error) {
	result := DB.Create(data)
	return result.RowsAffected > 0, result.Error
}

func (r *RelationFriendsApply) UpdateProcessStatus(msgId int64, status int) error {
	var relation = &RelationFriendsApply{
		Status:      status,
		ProcessTime: time.Now(),
	}
	res := DB.Where("msg_id = ?", msgId).Updates(&relation)
	if res.RowsAffected == 0 {
		return res.Error
	}
	return nil
}

func (r *RelationFriendsApply) FirstByMsgId() (relation *RelationFriendsApply, err error) {
	err = DB.Where("msg_id = ?", r.MsgId).First(&relation).Error
	return
}
