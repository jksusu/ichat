package chat

import (
	"ichat/pkg/ichat_model/chat_content_model"
	"ichat/pkg/ichat_model/chat_msg_index_model"
	"ichat/pkg/response"
)

type Record struct {
	*chat_msg_index_model.ChatMessageIndex
	MsgId       int64  `json:"msgId"`
	MsgType     int    `json:"msg_type"`
	From        string `json:"from"`
	To          string `json:"to"`
	Content     string `json:"content"`
	ContentType int    `json:"content_type"`
	SendTime    int64  `json:"send_time"`
	ISend       int    `json:"i_send"`
}

func RecordService() *Record {
	return &Record{}
}

// 获取聊天记录
func (record *Record) SingleChatRecord(uid, to string, page, pageSize int) *response.Response {
	var (
		contentModel *chat_content_model.ChatMessageContent
		content      []*chat_content_model.ChatMessageContent
		index        []*chat_msg_index_model.ChatMessageIndex
		total        int64
		err          error
		ids          []int64
	)
	if index, total, err = record.GetARecord(uid, to, page, pageSize); err != nil {
		return response.Fail.WithMsg(err.Error())
	}
	for _, item := range index {
		ids = append(ids, item.MsgId)
	}
	//获取消息内容
	content, err = contentModel.GetContentByIds(ids)
	if err != nil {
		return response.Fail.WithMsg(err.Error())
	}

	var list = make([]*Record, len(index))
	for i, item := range index {
		for _, contents := range content {
			if item.MsgId == contents.ID {
				list[i] = &Record{
					MsgId:       item.MsgId,
					From:        item.A,
					To:          item.B,
					ISend:       item.ISend,
					Content:     contents.Content,
					ContentType: contents.Type,
					SendTime:    contents.SendTime,
				}
			}
		}
	}

	return response.Ok.WithPageData(&response.PageData{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}
