package friends_relation

import (
	"ichat/pkg/ichat_model"
	"time"
)

type RelationFriendsApply struct {
	*ichat_model.RelationFriendsApply
}

func (r *RelationFriendsApply) Add(friendsApply *ichat_model.RelationFriendsApply) (ok bool, err error) {
	if result := ichat_model.DB.Create(friendsApply); result.RowsAffected > 0 {
		return true, nil
	} else {
		return false, result.Error
	}
}

// 修改处理状态
func (r *RelationFriendsApply) UpdateProcessStatus() error {
	var relation = &ichat_model.RelationFriendsApply{
		Status:      r.Status,
		ProcessTime: time.Now(),
	}
	res := ichat_model.DB.Where("msg_id = ?", r.MsgId).Updates(&relation)
	if res.RowsAffected == 0 {
		return res.Error
	}
	return nil
}

func (r *RelationFriendsApply) FirstByMsgId() (relation *RelationFriendsApply, err error) {
	err = ichat_model.DB.Where("msg_id = ?", r.MsgId).First(&relation).Error
	return
}
