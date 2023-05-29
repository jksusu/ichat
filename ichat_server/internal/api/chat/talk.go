package chat

import (
	"context"
	"github.com/golang/glog"
	"ichat/internal/area/push"
	"ichat/internal/area/talk"
	"ichat/internal/format"
)

func TalkToUser(ctx context.Context, r *format.R) {
	res, err := talk.ToUser(ctx, r)
	if err != nil {
		glog.Error(err)
	}
	if s := push.Server.Resp(res).Notice(res); s.Err != nil {
		glog.Error(s.Err)
	}
}

func TalkToGroup(ctx context.Context, r *format.R) {

}

func TalkToRoom(ctx context.Context, r *format.R) {

}
