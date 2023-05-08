package ichat_model

type ChatMessageContent struct {
	ID       int64  `gorm:"primarykey"`
	Type     int    `gorm:"default:1"`
	Content  string `gorm:"size:3000;not null"`
	Extra    string `gorm:"size:500"`
	SendTime int64  `gorm:"index"`
}

var ChatMsgContentModel = new(ChatMessageContent)

// GetContentByIds 获取内容根据ids
func (c *ChatMessageContent) GetContentByIds(ids []int64) (content []*ChatMessageContent, err error) {
	err = DB.Model(&c).Where("id IN (?)", ids).Find(&content).Error
	return
}
func (c *ChatMessageContent) Add(msgId int64, types int, content, extra string, sendTime int64) error {
	res := DB.Create(&ChatMessageContent{
		ID:       msgId,
		Type:     types,
		Content:  content,
		Extra:    extra,
		SendTime: sendTime,
	})
	return res.Error
}
