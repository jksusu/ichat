package relation

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"ichat/gateway/message"
	"ichat/pkg/ichat_cache/connect"
	"ichat/pkg/ichat_model"
	"ichat/pkg/ichat_model/relation_model/friends_relation"
	"ichat/pkg/ichat_model/unread_message"
	"ichat/pkg/ichat_service/relation_service"
	"ichat/pkg/ichat_service/user_service"
	"ichat/pkg/ichat_tool/idgen"
	"strconv"
	"time"
)

// 好友申请数据
type AddFriendData struct {
	To        string    `json:"to"`
	Remark    string    `json:"remark"`
	ApplyTime time.Time `json:"applytime"`
	Status    int       `json:"status"`
}

func HandlerFriendsApply(r *message.RequestMessage) {
	var (
		data     *AddFriendData
		msgId    = idgen.Gen.Node.Generate().Int64()
		relation *relation_service.RelationFriends
		connId   string
		err      error
	)
	p, _ := message.BuildDecode(r.Data)
	json.Unmarshal(p, &data)

	data.ApplyTime = time.Now()
	data.Status = 1

	//好友申请入库
	relation = &relation_service.RelationFriends{
		MsgId:     msgId,
		UID:       r.From,
		TO:        data.To,
		Remarks:   data.Remark,
		ApplyTime: data.ApplyTime,
	}
	if err := relation.AddFriendApply(); err != nil {
		logrus.Error(err)
		return
	}
	message.Response(&message.ResponseMessage{ReqId: r.ReqId, From: r.From})
	//查询对面是否在线，如果在线则直接推送
	if connId, err = connect.GetConnIdByUID(data.To); err != nil {
		logrus.Error(err)
		//写入离线消
		var u *unread_message.UnReadMessageList
		u.Add(&ichat_model.UnreadMessageList{
			MsgId:     msgId,
			UID:       r.From,
			MsgType:   1,
			CreatedAt: time.Now(),
		})
		return
	}
	user, err := user_service.UserService().GetUserByUid(r.From)
	if err != nil {
		logrus.Error(err)
		return
	}
	if err = message.Notice(connId, &message.NoticeMessage{
		Route: r.Route,
		MsgId: msgId,
		From:  r.From,
		FromInfo: &message.FromInfo{
			UID:      r.From,
			Nickname: user.Nickname,
			Avatar:   user.HeadPortrait,
		},
		Data: &data,
	}); err != nil {
		logrus.Error(err)
	}
}

type Reply struct {
	MsgId  string `json:"msgId"`
	Status int    `json:"status"`
}

func HandlerFriendsApplyReply(request *message.RequestMessage) {
	message.Responses(request)

	var (
		reply     *Reply
		MsgId     int64
		applyInfo *friends_relation.RelationFriendsApply
		err       error
	)
	p, _ := message.BuildDecode(request.Data)
	json.Unmarshal(p, &reply)

	MsgId, err = strconv.ParseInt(reply.MsgId, 10, 64)

	var relation = &friends_relation.RelationFriendsApply{&ichat_model.RelationFriendsApply{MsgId: MsgId, Status: reply.Status}}
	if applyInfo, err = relation.FirstByMsgId(); err != nil {
		logrus.Error(err)
		return
	}
	if err = relation.UpdateProcessStatus(); err != nil {
		logrus.Error(err)
	}

	//好友表新增
	if err = relation_service.AddRelationFriends(applyInfo.UID, applyInfo.TO); err != nil {
		logrus.Error(err)
	}
	connId, err := connect.GetConnIdByUID(applyInfo.UID)
	if err != nil {
		logrus.Error(err)
		//写入离线消
		var u *unread_message.UnReadMessageList
		u.Add(&ichat_model.UnreadMessageList{
			MsgId:     relation.MsgId,
			UID:       request.From,
			MsgType:   1,
			CreatedAt: time.Now(),
		})
		return
	}

	//推送好友添加成功消息
	if err = message.Notice(connId, &message.NoticeMessage{
		Route: request.Route,
		MsgId: relation.MsgId,
		From:  request.From,
		Data:  &reply,
	}); err != nil {
		logrus.Error(err)
	}
}

func HandlerFriendsEdit(message *message.RequestMessage) {

}

func HandlerFriendsDel(message *message.RequestMessage) {

}
