package model

// 消息内容
type ChatMessageContent struct {
	ID       int64  `gorm:"primarykey"`
	Type     int    `gorm:"default:1"`
	Content  string `gorm:"size:3000;not null"`
	Extra    string `gorm:"size:500"`
	SendTime int64  `gorm:"index"`
}

// 消息索引
type ChatMessageIndex struct {
	ID       int64  `gorm:"primarykey"`
	A        string `gorm:"index;size:60;not null;comment:账户a"`
	B        string `gorm:"size:60;not null;comment:账户b"`
	ISend    int    `gorm:"default:0;not null;comment:是否为账户a发送"`
	MsgId    int64  `gorm:"not null;comment:关联消息内容表中的ID"`
	GroupId  string `gorm:"size:30;comment:群ID，单聊情况为空"`
	SendTime int64  `gorm:"index;not null;comment:消息发送时间"`
}

type msg struct{}

var MessageModel = new(msg)

// GetContentByIds 获取内容根据ids
func (*msg) GetContentByIds(ids []int64) (content []*ChatMessageContent, err error) {
	err = DB.Where("id IN (?)", ids).Find(&content).Error
	return
}

func (*msg) InsertMessage(c *ChatMessageContent, i *ChatMessageIndex) (bool, error) {
	if res := DB.Create(c); res.Error != nil || res.RowsAffected == 0 {
		return false, res.Error
	}
	if res := DB.Create(i); res.Error != nil || res.RowsAffected == 0 {
		return false, res.Error
	}
	return true, nil
}

func (c *msg) BatchInsert(data []*ChatMessageIndex) (bool, error) {
	if tx := DB.Model(&c).Create(data); tx.Error != nil || tx.RowsAffected == 0 {
		return false, tx.Error
	}
	return true, nil
}

func (*msg) GetRecord(form, to string, msgId int64, limit int) []*ChatMessageIndex {
	list := make([]*ChatMessageIndex, 0)
	if msgId < 1 {
		//第一页
		DB.Where("a = ? AND b = ? ", form, to).Limit(limit).Offset(0).Find(&list)
	} else {
		DB.Where("a = ?", form).Where("b = ?", to).Where("msg_id <= ?", msgId).Find(&list)
	}
	return list
}
