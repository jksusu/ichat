package model

import (
	"ichat/pkg/tools/idgen"
)

type ChatMessageIndex struct {
	ID       int64  `gorm:"primarykey"`
	A        string `gorm:"index;size:60;not null;comment:账户a"`
	B        string `gorm:"size:60;not null;comment:账户b"`
	ISend    int    `gorm:"default:0;not null;comment:是否为账户a发送"`
	MsgId    int64  `gorm:"not null;comment:关联消息内容表中的ID"`
	GroupId  string `gorm:"size:30;comment:群ID，单聊情况为空"`
	SendTime int64  `gorm:"index;not null;comment:消息发送时间"`
}

var ChatMsgIndexModel = new(ChatMessageIndex)

// GetARecord 查询A的聊天记录
func (mi *ChatMessageIndex) GetARecord(a, b string, page, pageSize int) (index []*ChatMessageIndex, total int64, err error) {
	offset := (page - 1) * pageSize
	if err = DB.Model(&mi).Where("a = ? AND b = ?", a, b).Count(&total).Error; err != nil {
		return
	}
	err = DB.Model(&mi).Where("a = ? AND b = ?", a, b).Limit(pageSize).Offset(offset).Find(&index).Error
	return
}

func (*ChatMessageIndex) Add(from, to string, isend int, msgId int64, sendTime int64) {
	DB.Create(ChatMessageIndex{
		ID:       idgen.Gen.Node.Generate().Int64(),
		A:        from,
		B:        to,
		ISend:    isend,
		MsgId:    msgId,
		SendTime: sendTime,
	})
}
