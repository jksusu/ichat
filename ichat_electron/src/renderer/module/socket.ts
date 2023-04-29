import log from 'loglevel-es'
import { IChatClient } from "@/sdk/sdk"
import { UserStore } from "../stores/modules/user"
import { SocketStore } from "../stores/modules/socket"
import { eventAll } from "@/renderer/module/event"

let gateway = 'ws://127.0.0.1:9501/ws?token='

const open = () => {
    //log.info('我是绑定的打开事件')
}

const message = (message) => {
    //log.info('我是绑定的消息事件')
}

const close = () => {
    //log.info("close")
}
/**
 * 连接socket
 * @returns
 */
export const Connect = async () => {
    let token = UserStore().getToken
    if (token != '') {
        let socket = new IChatClient(gateway + token)
        //绑定回调事件
        socket.bindOnOpenCallback(open)
        socket.bindOnMessageCallback(message)
        socket.bindOnCloseCallback(close)
        socket.registerEventAll(eventAll)//批量注册执行器

        //开始登录
        let { ok, err } = await socket.auth()
        if (!ok) {
            log.debug('IchatClient connect error;' + err)
            return
        }
        SocketStore().setSocket(socket)
        return
    }
    log.info("token is null")
    return
}

/**
 * 获取 IchatClient对象
 * @returns 
 */
export const Socket = (): IChatClient => {
    if (SocketStore().getSocket == null) {
        log.error('socket is null')
        return
    }
    return SocketStore().getSocket
}

let event = {}
