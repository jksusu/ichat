export enum MsgType {
    TextMessage = 1,
    BinaryMessage = 2,
    CloseMessage = 8,
    PingMessage = 9,
    PongMessag = 10,
}

//路由处理器
export enum RouteHandler {
    RouteConnectPingMessage = 'connect.ping',
    RouteConnectPongMessage = 'connect.pong',

    //业务
    RouteIchatTalkToUser = 'ichat.talk.to.user',
    RouteIchatTalkToGroup = 'ichat.talk.to.group',
    RouteIchatTalkAckReceive = 'ichat.talk.to.ack.receive',//已收到
    RouteIchatTalkAckRead = 'ichat.talk.to.ack.read',//已读

    //session
    RouteSessionLists = 'ichat.session.lists',

    RouteUserInfoUpdate = 'ichat.user.info.update',//当前登录用户信息修改
    RouteUserInfoSearch = 'ichat.user.info.search',
    RouteGroupInfoSearch = 'ichat.group.info.search',

    //关系
    RouteRelationFriendsApply = 'ichat.relation.friends.apply',
    RouteRelationFriendsApplyReply = 'ichat.relation.friends.apply.reply',
    RouteRelationFriendsEdit = 'ichat.relation.friends.edit',
    RouteRelationFriendsDel = 'ichat.relation.friends.del',
}

export enum ContentType {
    text = 1,
    image = 2,
    video = 3,
}


export enum ICahtClientStatus {
    INIT,
    CONNECTING,
    CONNECTED,
    RECONNECTING,
    CLOSEING,
    CLOSED,
}

//消息包类型
export enum MessageType {
    REQUEST = 1,
    RESPONSE = 2,
    NOTICE = 3,
    REQUESTPING = 4,
    RESPONSEPOMG = 5
}