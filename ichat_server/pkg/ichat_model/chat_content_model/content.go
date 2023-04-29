package chat_content_model

import "ichat/pkg/ichat_model"

type ChatMessageContent struct {
	*ichat_model.ChatMessageContent
}

// GetContentByIds 获取内容根据ids
func (c *ChatMessageContent) GetContentByIds(ids []int64) (content []*ChatMessageContent, err error) {
	err = ichat_model.DB.Model(&c).Where("id IN (?)", ids).Find(&content).Error

	return
}

func (c *ChatMessageContent) AddMessageContent() {
}
