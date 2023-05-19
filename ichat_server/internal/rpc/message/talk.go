package message

import (
	"context"
	"gorm.io/gorm"
	"ichat/pkg/ichat_tool/idgen"
	"ichat/pkg/protocol/pb"
	"time"
)

type talk struct{}

var Talk = new(talk)

func (*talk) ToUser(c context.Context, r *pb.TalkToUserReq) (*pb.TalkToUserResp, error) {
	var (
		sendTime = time.Now().UnixNano()
		msgId    = idgen.Gen.Node.Generate().Int64()
	)
	if err := model.DB.Transaction(func(tx *gorm.DB) (err error) {
		model.ChatMsgContentModel.Add(msgId, int(r.Type), r.Content, r.Extra, r.SendTime)
		model.ChatMsgIndexModel.Add(r.From, r.To, 1, msgId, sendTime)
		model.ChatMsgIndexModel.Add(r.To, r.From, 0, msgId, sendTime)
		return
	}); err != nil {
		return nil, err
	}

	return &pb.TalkToUserResp{MsgId: msgId}, nil
}
