package message

import (
	"context"
	"gorm.io/gorm"
	"ichat/pkg/ichat_model"
	"ichat/pkg/ichat_tool/idgen"
	"ichat/pkg/protocol/pb"
	"time"
)

type talk struct{}

var Talk = new(talk)

func (*talk) ToUser(c context.Context, r *pb.TalkToUserRequest) (*pb.TalkToUserResponse, error) {
	var (
		sendTime = time.Now().UnixNano()
		msgId    = idgen.Gen.Node.Generate().Int64()
	)
	if err := ichat_model.DB.Transaction(func(tx *gorm.DB) (err error) {
		ichat_model.ChatMsgContentModel.Add(msgId, int(r.Type), r.Content, r.Extra, r.SendTime)
		ichat_model.ChatMsgIndexModel.Add(r.From, r.To, 1, msgId, sendTime)
		ichat_model.ChatMsgIndexModel.Add(r.To, r.From, 0, msgId, sendTime)
		return
	}); err != nil {
		return nil, err
	}

	return &pb.TalkToUserResponse{MsgId: msgId}, nil
}
