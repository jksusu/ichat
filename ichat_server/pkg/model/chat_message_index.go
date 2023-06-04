package model

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

func (*ChatMessageIndex) Insert(data *ChatMessageIndex) bool {
	if tx := DB.Create(data); tx.Error != nil || tx.RowsAffected == 0 {
		return false
	}
	return true
}

// 批量
func (c *ChatMessageIndex) BatchInsert(data []*ChatMessageIndex) (bool, error) {
	if tx := DB.Model(&c).Create(data); tx.Error != nil || tx.RowsAffected == 0 {
		return false, tx.Error
	}
	return true, nil
}

// 游标查询
func (*ChatMessageIndex) GetRecord(form, to string, msgId int64, limit int) []*ChatMessageIndex {
	list := make([]*ChatMessageIndex, 0)
	if msgId < 1 {
		//第一页
		DB.Where("a = ? AND b = ? ", form, to).Limit(limit).Offset(0).Find(&list)
	} else {
		DB.Where("a = ?", form).Where("b = ?", to).Where("msg_id <= ?", msgId).Find(&list)
	}
	return list
}
