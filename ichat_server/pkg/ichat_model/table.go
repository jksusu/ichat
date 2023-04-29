package ichat_model

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

type ChatMessageContent struct {
	ID       int64  `gorm:"primarykey"`
	Type     int    `gorm:"default:1"`
	Content  string `gorm:"size:3000;not null"`
	Extra    string `gorm:"size:500"`
	SendTime int64  `gorm:"index"`
}

type ChatMessageIndex struct {
	ID       int64  `gorm:"primarykey"`
	A        string `gorm:"index;size:60;not null;comment:账户a"`
	B        string `gorm:"size:60;not null;comment:账户b"`
	ISend    int    `gorm:"default:0;not null;comment:是否为账户a发送"`
	MsgId    int64  `gorm:"not null;comment:关联消息内容表中的ID"`
	GroupId  string `gorm:"size:30;comment:群ID，单聊情况为空"`
	SendTime int64  `gorm:"index;not null;comment:消息发送时间"`
}

type Users struct {
	GLOBALMODEL
	UID          string `json:"uid" gorm:"index;comment:用户uid"`
	Nickname     string `json:"nickname" gorm:"comment:用户昵称"`
	Username     string `json:"username" gorm:"uniqueIndex;comment:用户名"`
	Password     string `json:"-" gorm:"comment:登录密码"`
	Salt         string `json:"-" gorm:"comment:密码盐"`
	HeadPortrait string `json:"head_portrait" gorm:"comment:头像"`
	Status       int    `json:"status" gorm:"default:1;comment:账号状态 1:正常 2:封号中 3:拉黑"`
	LastLoginIp  string `json:"last_login_ip" gorm:"column:last_login_ip;comment:最后登录ip地址"`
	Token        string `json:"token" gorm:"column:token;comment:token"`
}

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

// 好友申请表
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

// 未读消息列表
type UnreadMessageList struct {
	MsgId     int64     `json:"msg_id"`
	UID       string    `json:"uid"`
	MsgType   int       `json:"msg_type"`
	CreatedAt time.Time `json:"created_at"`
}
