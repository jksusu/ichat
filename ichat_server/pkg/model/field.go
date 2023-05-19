package model

import (
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

type GLOBALMODEL struct {
	ID        int            `json:"id" gorm:"primarykey;comment:主键id"`
	CreatedAt time.Time      `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"comment:删除时间"`
}

// 未读消息列表
type UnreadMessageList struct {
	MsgId     int64     `json:"msg_id"`
	UID       string    `json:"uid"`
	MsgType   int       `json:"msg_type"`
	CreatedAt time.Time `json:"created_at"`
}
