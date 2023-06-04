package talk

import (
	"context"
	"ichat/pkg/cache/message"
	cache "ichat/pkg/cache/user"
	"ichat/pkg/model"
)

type (
	sync   struct{}
	record struct {
		MsgId    int64  `json:"msgId"`
		Isend    int    `json:"i_send"`
		SendTime int64  `json:"send_time"`
		Content  string `json:"content"`
		CType    int    `json:"c_type"`
		Extra    string `json:"extra"`
		FromInfo *info  `json:"f_info"`
		ToInfo   *info  `json:"t_info"`
	}
	info struct {
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	}
)

// 更新会话最后未读
func UpdateLastReadMsg(ctx context.Context, msgId int64) {
	message.MarkTalkLastReadMessage(ctx.Value("from").(string), ctx.Value("to").(string), msgId)
}

func MessageRecord(ctx context.Context, from, to string, msgId int64) (interface{}, error) {
	index := model.ChatMsgIndexModel.GetRecord(from, to, msgId, 20)
	if index == nil {
		return nil, nil
	}
	ids := make([]int64, 0)
	for _, v := range index {
		ids = append(ids, v.MsgId)
	}
	content, err := model.ChatMsgContentModel.GetContentByIds(ids)
	if err != nil {
		return nil, err
	}
	//获取会话头像昵称信息
	finfo, _ := cache.Rdb.GetUserByKey(from)
	tinfo, _ := cache.Rdb.GetUserByKey(to)
	fromInfo := &info{Nickname: finfo.Nickname, Avatar: finfo.Avatar}
	toInfo := &info{Nickname: tinfo.Nickname, Avatar: tinfo.Avatar}

	data := make([]*record, 0)

	for _, v := range index {
		for _, c := range content {
			if v.MsgId == c.ID {
				data = append(data, &record{
					MsgId:    msgId,
					Isend:    v.ISend,
					SendTime: v.SendTime,
					Content:  c.Content,
					CType:    c.Type,
					Extra:    c.Extra,
					FromInfo: fromInfo,
					ToInfo:   toInfo,
				})
			}
		}
	}

	return data, err
}
