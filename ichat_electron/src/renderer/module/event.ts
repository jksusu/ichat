//事件
import { RouteHandler } from '@/sdk/enum/enum'
import { ChatRecordStore } from '../stores/modules/chatRecord'
import { NewFriendsStore } from '../stores/modules/newFriends'
import { SessionStore } from '../stores/modules/sessionList'



export function noticeFriendsApply(message) {
    NewFriendsStore().addMessage({
        msgId: message.msgId,
        from: message.from,
        nickname: message.fromInfo.nickname ?? '',
        avatar: message.fromInfo.avatar ?? '',
        remark: message.data.remark ?? '',
        applytime: message.data.applytime ?? '',
        status: message.data.status ?? '',
    })
}

export function noticeFriendsApplyReply(message) {
    console.log(message)
}

export const noticeTalkToUser = (message) => {
    let { data } = message
    //写入聊天列表
    let msg = {
        content: data.content ?? '',
        from: message.from,
        i_send: false,
        msgId: data.msgId ?? '',
        msg_type: data.msg_type,
        send_time: data.send_time,
        status: data.status,
        to: data.to,
    }
    console.log(msg)
    //sender 
    ChatRecordStore().addMessage(message.from, msg)
    // //更新聊天会话 更新信息和时间
    //SessionStore().addSession(message.fromInfo.nickname, message.from, '', data.content)
}


export let eventAll = {
    [RouteHandler.RouteRelationFriendsApply]: noticeFriendsApply,
    [RouteHandler.RouteRelationFriendsApplyReply]: noticeFriendsApply,
    [RouteHandler.RouteIchatTalkToUser]: noticeTalkToUser,
}