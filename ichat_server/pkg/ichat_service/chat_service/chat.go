package chat_service

import (
	"gorm.io/gorm"
	"ichat/gateway/message"
	"ichat/pkg/ichat_model"
	"ichat/pkg/ichat_tool/idgen"
)

func AddMessageContent(from string, messageId int64, chat *message.ChatMessage) (err error) {
	var content = &ichat_model.ChatMessageContent{
		ID:       messageId,
		Type:     chat.Type,
		Content:  chat.Content,
		Extra:    chat.Extra,
		SendTime: chat.SendTime,
	}
	index := make([]ichat_model.ChatMessageIndex, 2)
	index[0] = ichat_model.ChatMessageIndex{
		ID:       idgen.Gen.Node.Generate().Int64(),
		A:        from,
		B:        chat.To,
		ISend:    1,
		MsgId:    messageId,
		SendTime: chat.SendTime,
	}
	index[1] = ichat_model.ChatMessageIndex{
		ID:       idgen.Gen.Node.Generate().Int64(),
		A:        chat.To,
		B:        from,
		ISend:    0,
		MsgId:    messageId,
		SendTime: chat.SendTime,
	}
	//开启事务
	return ichat_model.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Create(content).Error; err != nil {
			return
		}
		err = tx.Create(&index).Error
		return
	})
}
