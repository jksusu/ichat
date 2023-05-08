package relation

import (
	"context"
	"ichat"
	"ichat/internal/gate/message"
	"ichat/pkg/protocol/pb"
	"time"
)

type AddFriendData struct {
	To        string    `json:"to"`
	Remark    string    `json:"remark"`
	ApplyTime time.Time `json:"applytime"`
	Status    int       `json:"status"`
}

type Reply struct {
	MsgId  string `json:"msgId"`
	Status int    `json:"status"`
}

func HandlerFriendsApply(r *message.RequestMessage) {
	data := &AddFriendData{}
	message.DataUnpack(r.Data, data)
	ichat.RpcClient.Relation.FriendsApply(context.Background(), &pb.FriendsApplyRequest{To: data.To, Remark: data.Remark})
}

func HandlerFriendsApplyReply(r *message.RequestMessage) {
	message.Responses(r)

	//var (
	//	data      = &Reply{}
	//	MsgId     int64
	//	applyInfo *friends_relation.RelationFriendsApply
	//	err       error
	//)
	//message.DataUnpack(r.Data, data)
	//
	//MsgId, err = strconv.ParseInt(data.MsgId, 10, 64)
	//
	//ichat.RpcClient.Relation.FriendsApplyReply(context.Background(), &pb.FriendsApplyReplyRequest{})
	//
	//var relation = &friends_relation.RelationFriendsApply{&ichat_model.RelationFriendsApply{MsgId: MsgId, Status: reply.Status}}
	//if applyInfo, err = relation.FirstByMsgId(); err != nil {
	//	logrus.Error(err)
	//	return
	//}
	//if err = relation.UpdateProcessStatus(); err != nil {
	//	logrus.Error(err)
	//}
	//
	////好友表新增
	//if err = relation_service.AddRelationFriends(applyInfo.UID, applyInfo.TO); err != nil {
	//	logrus.Error(err)
	//}
	//connId, err := connect.GetConnIdByUID(applyInfo.UID)
	//if err != nil {
	//	logrus.Error(err)
	//	//写入离线消
	//	var u *unread_message.UnReadMessageList
	//	u.Add(&ichat_model.UnreadMessageList{
	//		MsgId:     relation.MsgId,
	//		UID:       request.From,
	//		MsgType:   1,
	//		CreatedAt: time.Now(),
	//	})
	//	return
	//}
	//
	////推送好友添加成功消息
	//if err = message.Notice(connId, &message.NoticeMessage{
	//	Route: request.Route,
	//	MsgId: relation.MsgId,
	//	From:  request.From,
	//	Data:  &reply,
	//}); err != nil {
	//	logrus.Error(err)
	//}
}

func HandlerFriendsEdit(message *message.RequestMessage) {

}

func HandlerFriendsDel(message *message.RequestMessage) {

}
