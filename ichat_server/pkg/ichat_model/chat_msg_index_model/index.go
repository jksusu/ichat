package chat_msg_index_model

import "ichat/pkg/ichat_model"

type ChatMessageIndex struct {
	*ichat_model.ChatMessageIndex
}

// GetARecord 查询A的聊天记录
func (mi *ChatMessageIndex) GetARecord(a, b string, page, pageSize int) (index []*ChatMessageIndex, total int64, err error) {
	offset := (page - 1) * pageSize
	if err = ichat_model.DB.Model(&mi).Where("a = ? AND b = ?", a, b).Count(&total).Error; err != nil {
		return
	}
	err = ichat_model.DB.Model(&mi).Where("a = ? AND b = ?", a, b).Limit(pageSize).Offset(offset).Find(&index).Error
	return
}
