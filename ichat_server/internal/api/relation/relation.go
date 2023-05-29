package relation

import (
	"context"
	"github.com/golang/glog"
	"ichat/internal/area/contacts"
	"ichat/internal/area/push"
	"ichat/internal/format"
)

func FriendsApply(ctx context.Context, r *format.R) {
	res, err := contacts.FriendsApply(ctx, r)
	if err != nil {
		glog.Error(err)
	}
	push.Server.Resp(res).Notice(res)
}

func FriendsApplyReply(ctx context.Context, r *format.R) {
	res, err := contacts.FriendsApplyReply(ctx, r)
	if err != nil {
		glog.Error(err)
	}
	push.Server.Resp(res).Notice(res)
}

func FriendsDel(ctx context.Context, r *format.R) {

}

func FriendsEdit(ctx context.Context, r *format.R) {

}
