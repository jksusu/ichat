
import { ElMessageBox } from "element-plus";
import { CreateSession, DelSession, MessageToping } from "./sessionEffect";

export const sendMessage = {
    label: "发送消息",
    click: (menu, arg) => {
        CreateSession(arg.data.friendUid, arg.data.nickname, arg.data.headPortraitUrl)
    }
}
export const setRemarks = {
    label: "设置备注",
    click: (menu, arg) => {
        ElMessageBox.prompt('请输入备注', { confirmButtonText: '确定', cancelButtonText: '取消' }).then(({ value }) => {
            let data = { 'id': arg.data.id, 'type': 'remarks', 'accept_type': arg.data.accept_type, 'acceptCode': arg.data.code, configValue: value }
        })
    }
}
export const delFriend = {
    label: "删除好友",
    click: (menu, arg) => {
    }
}
export const quitGroup = {
    label: "退出群聊",
    click: (menu, arg) => {
        console.log('退出群聊');
    }
}
export const msgTopping = {
    label: "消息置顶",
    click: (menu, arg) => {
        MessageToping(arg.uid)
    }
}
export const cancelTopping = {
    label: "取消置顶",
    click: (menu, arg) => {
        console.log(arg)
    }
}

export const nodisturb = {
    label: "消息免打扰",
    click: (menu, arg) => {
        console.log(arg)
    }
}
export const disturb = {
    label: "开启消息提醒",
    click: (menu, arg) => {
        console.log(arg)
    }
}

export const delSession = {
    label: "删除会话",
    click: (menu, arg) => {
        console.log(arg)
        DelSession(arg.uid)
    }
}