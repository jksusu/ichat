package relation_service

import (
	"gorm.io/gorm"
	"ichat/pkg/ichat_model"
	"ichat/pkg/ichat_model/relation_model/friends_relation"
	"ichat/pkg/ichat_tool/idgen"
	"time"
)

type RelationFriends struct {
	*friends_relation.RelationFriendsApply
	MsgId     int64     `json:"msgId"`
	UID       string    `json:"uid"`
	TO        string    `json:"to"`
	Remarks   string    `json:"remarks"`
	ApplyTime time.Time `json:"applytime"`
}

func (r *RelationFriends) AddFriendApply() error {
	_, err := r.Add(&ichat_model.RelationFriendsApply{
		MsgId:     idgen.Gen.Node.Generate().Int64(),
		UID:       r.UID,
		TO:        r.TO,
		Status:    1,
		Remarks:   r.Remarks,
		ApplyTime: r.ApplyTime,
	})
	return err
}

func AddRelationFriends(from, to string) error {
	friends := []*ichat_model.Friends{
		{
			UserUid:   from,
			FriendUid: to,
			Status:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserUid:   to,
			FriendUid: from,
			Status:    1,
		},
	}
	friends[1].CreatedAt = friends[0].CreatedAt
	friends[1].UpdatedAt = friends[0].UpdatedAt

	return ichat_model.DB.Transaction(func(tx *gorm.DB) error {
		return tx.CreateInBatches(&friends, len(friends)).Error
	})
}
