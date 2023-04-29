import { router } from "../router"
import { useSelectIndexStore } from "../stores/modules/selectIndex"
import { SessionStore } from "../stores/modules/sessionList"

/**
 * 联系人列表创建会话，需跳转页面
 * @param uid 用户UID 
 * @param nickname 用户昵称
 * @param headImg  用户头像
 */
export const CreateSession = (uid: string, nickname: string, headImg: string) => {
    SessionStore().addSession(uid, nickname, headImg)//增加会话
    SessionStore().setSelectSession(uid, nickname, headImg, 1)
    useSelectIndexStore().setSessionIndex(uid)//设置会话索引
    useSelectIndexStore().setAsideMenuIndex('/session')//更新左侧菜单索引
    router.push("/session")
}

/**
 * 删除会话
 * @param uid 
 */
export const DelSession = (uid: string) => {
    SessionStore().delSession(uid)
}

/**
 * 消息置顶
 * @param uid 
 */
export const MessageToping = (uid: string) => {
    SessionStore().setSessionTop(uid)
}

/**
 * 取消消息置顶
 * @param uid 
 */
export const CancelMessageToping = (uid: string) => {
    SessionStore().cancelSessionTop(uid)
}