package contacts

import (
	"context"
	"github.com/golang/glog"
	"ichat/internal/area/gate"
	"ichat/internal/format"
	"ichat/pkg/ichat_tool/idgen"
	"ichat/pkg/model"
	"time"
)

type friend struct {
	Remarks   string    `json:"remarks"`
	ApplyTime time.Time `json:"applyTime"`
	Status    int       `json:"status"`
}

var Friend = new(friend)

func (f *friend) FriendsApply(ctx context.Context, r *format.R) {
	format.Decode(r.Data, &f)
	msgId := idgen.Gen.Node.Generate().Int64()
	if _, err := model.RelationFriendsApplyModel.Add(&model.RelationFriendsApply{
		MsgId:     msgId,
		UID:       r.From,
		TO:        r.To,
		Remarks:   f.Remarks,
		ApplyTime: f.ApplyTime,
	}); err != nil {
		glog.Error(err)
		return
	}
	gate.Gateway.Resp(r).Notice(r)
}

func (*friend) FriendsApplyReply(ctx context.Context, r *format.R) {
}

func (*friend) FriendsEdit(ctx context.Context, r *format.R) {

}

func (*friend) FriendsDel(ctx context.Context, r *format.R) {

}
