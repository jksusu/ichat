import { Message } from "../format/base"

/**
 * 构建请求消息
 * from 当前登录用户uid
 * route 要执行的操作
 * message route操作需要的数据
 */
export class Request {
    sendTime: number
    callback: (response: Message) => void
    constructor(callback: (response: Message) => void) {
        this.sendTime = Date.now()
        this.callback = callback
    }
}