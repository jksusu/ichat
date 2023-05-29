package talk

import (
	"context"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"ichat/internal/format"
	"ichat/pkg/model"
	"ichat/pkg/tools/idgen"
	"time"
)

//type
//文本 语音 图片 gif图片 图文 文件 位置 小视频 引用 合并转发

type talk struct {
	MsgId    int64  `json:"msgId"`
	Content  string `json:"content"`
	Type     int    `json:"type"`
	SendTime int64  `json:"sendTime"`
	State    int    `json:"state"` //消息状态
	Extra    string `json:"extra"`
}

func ToUser(ctx context.Context, r *format.R) (*format.R, error) {
	var (
		data *talk
		err  error
	)
	if err = format.Decode(r.Data, &data); err != nil {
		glog.Errorf("talk talk to user data decode fail %v", err)
	}
	r.MsgId = idgen.GenChatMsgId()
	data.SendTime = time.Now().Unix()
	data.MsgId = r.MsgId
	if err = model.DB.Transaction(func(tx *gorm.DB) (err error) {
		_, err = model.ChatMsgIndexModel.BatchInsert(
			map[int]*model.ChatMessageIndex{
				0: {Seq: 0, A: r.From, B: r.To, ISend: 1, MsgId: r.MsgId, SendTime: data.SendTime},
				1: {Seq: 0, A: r.To, B: r.From, MsgId: r.MsgId, SendTime: data.SendTime},
			},
		)
		_, err = model.ChatMsgContentModel.Insert(&model.ChatMessageContent{
			ID: r.MsgId, Type: data.Type, Content: data.Content, Extra: data.Extra, SendTime: data.SendTime,
		})
		return err
	}); err != nil {
		glog.Errorf("talk talk to user insert fail %s", err)
	}
	return r, err
}

func ToGroup(ctx context.Context, r *format.R) {

}

func ToRoom(ctx context.Context, r *format.R) {

}
