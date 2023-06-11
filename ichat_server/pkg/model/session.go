package model

import "time"

type SessionList struct {
	Id          int64     `json:"id" gorm:"primary_key"`
	Uid         string    `json:"uid"`
	To          string    `json:"to"`
	SessionType int       `json:"session_type"`
	LastMsgId   int64     `json:"last_msg_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var (
	SessionListModel = new(SessionList)

	SESSION_TYPE_CONTACTS         = 1
	SESSION_TYPE_GROUP            = 2
	SESSION_TYPE_ROOM             = 3
	SESSION_TYPE_CUSTOMER_SERVICE = 4
)

func (*SessionList) GetAll(uid string) []*SessionList {
	list := make([]*SessionList, 0)
	DB.Where("uid = ?", uid).Find(list)
	return list
}

func (s *SessionList) Insert(list *SessionList) (int64, error) {
	tx := DB.Create(list)

	return tx.RowsAffected, tx.Error
}

func (s *SessionList) Delete(to string) (int64, error) {
	tx := DB.Where("to = ?", to).Delete(s)
	return tx.RowsAffected, tx.Error
}
