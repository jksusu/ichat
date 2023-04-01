package route

const (
	RoutedLoginSignOut = "login.signout"

	RouteChatUserMessage  = "ichat.user.message"
	RouteChatGroupMessage = "ichat.group.message"
	RouteChatMessageAck   = "ichat.message.ack"

	//ack
	RouteChatMessageServerAck = "ichat.message.server.ack" //服务端确定
	RouteChatMessageArriveAck = "ichat.message.arrive.ack" //已经送达ack
	RouteChatMessageReadAck   = "ichat.message.read.ack"   //已读ack

	//group
	RouteChatGroupCreate       = "ichat.group.create"
	RouteChatGroupDetail       = "ichat.group.detail"
	RouteChatGroupDetailSave   = "ichat.group.detail.save"
	RouteChatGroupProhibitChat = "ichat.group.prohibit.chat" //禁言

	//群用户
	RouteChatGroupJoin          = "ichat.group.join"
	RouteChatGroupExit          = "ichat.group.exit"
	RouteChatGroupInviteMember  = "ichat.group.member.invite" //邀请
	RouteChatGroupDeleteFriends = "ichat.group.member.delete" //删除

	//会话列表
	RouteSessionLists  = "ichat.session.lists"
	RouteSessionAdd    = "ichat.session.list.add"
	RouteSessionDel    = "ichat.session.list.del"
	RouteSessionUpdate = "ichat.session.list.update"

	//通知
	RouteUserInfoUpdate = "ichat..user.info.update"
)

const (
	//通知消息
	RouteNoticeMessageImportIng = "ichat.notice.message.importing" //通知对方正在输入
)
