package chat

import (
	"context"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"ichat/internal/format"
	"ichat/internal/gate"
	"ichat/pkg/model"
	"ichat/pkg/tools/idgen"
	"time"
)

type talk struct {
	Content string `json:"content"`
	Type    int    `json:"type"`
	Extra   string `json:"extra"`
}

var Talk = new(talk)

func ToUser(ctx context.Context, r *format.R) {
	var (
		msgId    = idgen.GenChatMsgId()
		sendTime = time.Now().Unix()
		data     *talk
	)
	if err := format.Decode(r, &data); err != nil {
		glog.Errorf("chat talk to user data decode fail %s", err)
	}
	if err := model.DB.Transaction(func(tx *gorm.DB) (err error) {
		_, err = model.ChatMsgIndexModel.BatchInsert(
			map[int]*model.ChatMessageIndex{
				0: {Seq: 0, A: r.From, B: r.To, ISend: 1, MsgId: msgId, SendTime: sendTime},
				1: {Seq: 0, A: r.To, B: r.From, MsgId: msgId, SendTime: sendTime},
			},
		)
		_, err = model.ChatMsgContentModel.Insert(&model.ChatMessageContent{
			ID: msgId, Type: data.Type, Content: data.Content, Extra: data.Extra, SendTime: sendTime,
		})
		return err
	}); err != nil {
		glog.Errorf("chat talk to user insert fail %s", err)
	}

	gate.Gateway.Resp(r).Notice(r)
}

func ToGroup(ctx context.Context, r *format.R) {

}

func ToRoom(ctx context.Context, r *format.R) {

}
