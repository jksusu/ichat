package contacts

import (
	"context"
	"ichat"
	"ichat/internal/gate/message"
	"ichat/pkg/ichat_model"
	"ichat/pkg/ichat_tool/idgen"
	"ichat/pkg/protocol/pb"
	"time"
)

type Friend struct{}

func (*Friend) FriendsApply(ctx context.Context, to, from, remark string, applyTime time.Time) {
	apply := &ichat_model.RelationFriendsApply{MsgId: idgen.Gen.Node.Generate().Int64(), UID: from, TO: to, Remarks: remark, ApplyTime: applyTime}
	ichat_model.RelationFriendsApplyModel.Add(apply)
	ichat.RpcClient.Gateway.Response(ctx, &pb.ResponseReq{ReqId: 0, From: from, Route: ""})
	ichat.RpcClient.Message.NoticeToUser(ctx, &pb.NoticeToUserRequest{})
}

type Reply struct {
	MsgId  string `json:"msgId"`
	Status int    `json:"status"`
}

func (*Friend) FriendsApplyReply(ctx context.Context, request *message.RequestMessage) {
	ichat.RpcClient.Gateway.Response(ctx, &pb.ResponseReq{ReqId: 0, From: "", Route: ""})
}

func (*Friend) FriendsEdit(message *message.RequestMessage) {

}

func (*Friend) FriendsDel(message *message.RequestMessage) {

}
