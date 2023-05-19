package contacts

import (
	"github.com/golang/glog"
	"ichat/internal/area/gate"
	"ichat/internal/gate/message"
	"ichat/pkg/ichat_tool/idgen"
	"ichat/pkg/model"
	"time"
)

type friend struct{}

var Friend = new(friend)

func (*friend) FriendsApply(to, from, remark string, applyTime time.Time) {
	if _, err := model.RelationFriendsApplyModel.Add(&model.RelationFriendsApply{
		MsgId:     idgen.Gen.Node.Generate().Int64(),
		UID:       from,
		TO:        to,
		Remarks:   remark,
		ApplyTime: applyTime,
	}); err != nil {
		glog.Error(err)
	}
	gate.Gateway.Resp()
	gate.Gateway.Notice(to)
}

type Reply struct {
	MsgId  string `json:"msgId"`
	Status int    `json:"status"`
}

func (*friend) FriendsApplyReply() {
	//ichat.RpcClient.Gateway.Response(&pb.ResponseReq{ReqId: 0, From: "", Route: ""})
}

func (*friend) FriendsEdit(message *message.RequestMessage) {

}

func (*friend) FriendsDel(message *message.RequestMessage) {

}
