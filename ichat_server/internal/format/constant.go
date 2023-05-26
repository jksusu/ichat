package format

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
	RouteChatGroupCreate        = "ic.group.create"
	RouteChatGroupDetail        = "ic.group.detail"
	RouteChatGroupDetailSave    = "ic.group.detail.save"
	RouteChatGroupProhibit      = "ic.group.prohibit" //禁言
	RouteChatGroupUserJoin      = "ic.group.user.join"
	RouteChatGroupUserExit      = "ic.group.user.exit"
	RouteChatGroupInviteMember  = "ic.group.member.invite" //邀请
	RouteChatGroupDeleteFriends = "ic.group.member.delete" //删除

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
