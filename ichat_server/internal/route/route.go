package route

import (
	"context"
	"github.com/golang/glog"
	"ichat/internal/api/chat"
	"ichat/internal/api/relation"
	"ichat/internal/format"
)

const (
	RouteIcUserSignOut  = "ic.user.signout"
	RouteIcUserLogin    = "ic.user.login"
	RouteIcUserRegister = "ic.user.register"

	RouteIcTalkToUser  = "ic.talk.to.user"
	RouteIcTalkToGroup = "ic.talk.to.group"
	RouteIcTalkToRoom  = "ic.talk.to.room"

	RouteChatMessageAck = "ic.message.ack"

	//ack
	RouteChatMessageArriveAck = "ic.message.arrive.ack" //已经送达ack
	RouteChatMessageReadAck   = "ic.message.read.ack"   //已读ack

	//group
	RouteRelationGroupCreate        = "ic.relation.group.create"
	RouteRelationGroupDetail        = "ic.relation.group.detail"
	RouteRelationGroupDetailSave    = "ic.relation.group.detail.save"
	RouteRelationGroupProhibit      = "ic.relation.group.prohibit" //禁言
	RouteRelationGroupUserJoin      = "ic.relation.group.user.join"
	RouteRelationGroupUserExit      = "ic.relation.group.user.exit"
	RouteRelationGroupInviteMember  = "ic.relation.group.member.invite" //邀请
	RouteRelationGroupDeleteFriends = "ic.relation.group.member.delete" //删除

	//会话列表
	RouteSessionLists  = "ic.session.lists"
	RouteSessionAdd    = "ic.session.add"
	RouteSessionDel    = "ic.session.del"
	RouteSessionUpdate = "ic.session.update"

	//系统通知
	RouteUserInfoUpdate = "ic.user.info.update"
	RouteUserInfoSearch = "ic.user.info.search"

	//关系
	RouteRelationFriendsApply      = "ic.relation.friends.apply"
	RouteRelationFriendsApplyReply = "ic.relation.friends.apply.reply"
	RouteRelationFriendsEdit       = "ic.relation.friends.edit"
	RouteRelationFriendsDel        = "ic.relation.friends.del"

	RouteNoticeMessageImportIng = "ic.notice.message.importing" //通知对方正在输入
)

var Map = map[string]func(ctx context.Context, r *format.R){
	RouteIcTalkToUser:  chat.TalkToUser,
	RouteIcTalkToGroup: chat.TalkToGroup,
	RouteIcTalkToRoom:  chat.TalkToRoom,

	RouteRelationFriendsApply:      relation.FriendsApply,
	RouteRelationFriendsApplyReply: relation.FriendsApplyReply,
	RouteRelationFriendsDel:        relation.FriendsDel,
	RouteRelationFriendsEdit:       relation.FriendsEdit,
}

func Router(ctx context.Context, r *format.R) {
	if fuc, ok := Map[r.Route]; ok {
		fuc(ctx, r)
	} else {
		glog.Errorf("route %s not absent", r.Route)
	}
}
