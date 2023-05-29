package contacts

import (
	"context"
	"github.com/golang/glog"
	"ichat/internal/format"
	"ichat/pkg/model"
	"ichat/pkg/tools/idgen"
	"time"
)

type friend struct {
	Remarks   string    `json:"remarks"`
	ApplyTime time.Time `json:"applyTime"`
	Status    int       `json:"status"`
}

func FriendsApply(ctx context.Context, r *format.R) (*format.R, error) {
	var (
		data *friend
		err  error
	)
	if err = format.Decode(r, &data); err != nil {
		glog.Errorf("contacts friend apply data decode fail %s", err)
	}
	r.MsgId = idgen.GenChatMsgId()
	if _, err = model.RelationFriendsApplyModel.Add(&model.RelationFriendsApply{
		MsgId:     r.MsgId,
		UID:       r.From,
		TO:        r.To,
		Remarks:   data.Remarks,
		ApplyTime: data.ApplyTime,
	}); err != nil {
		glog.Error(err)
		return r, err
	}
	return r, err
}

func FriendsApplyReply(ctx context.Context, r *format.R) (*format.R, error) {
	return r, nil
}

func FriendsEdit(ctx context.Context, r *format.R) {

}

func FriendsDel(ctx context.Context, r *format.R) {

}
