package ichat_cache

import (
	"context"
)

var Ctx = context.Background()

const (
	SESSION_LIST         = "session:list:"         //会话列表
	SESSION_UNREAD       = "session:unread:"       //未读
	SESSION_COUNT_UNREAD = "session:count_unread:" //总未读

	//连接
	CONNECT_CONNID_TO_UID = "connect:connid_to_uid:"
	CONNECT_UID_TO_CONNID = "connect:uid_to_connid:"

	//用户
	USER_TOKEN = "user:token:"

	//未读消息
	CONTACTS_UNREAD = "unread:contacts:"
	GROUP_UNREAD    = "unread:group:"
	UNREAD_MSG_ALL  = "unreadmsg:all"

	MESSAGE_STATUS = "message:status:"
)
