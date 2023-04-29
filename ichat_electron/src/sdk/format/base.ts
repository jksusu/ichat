import { Seq } from "../sdk"
import { MessageType, RouteHandler, ContentType } from "../enum/enum"

/**
 * 基础消息格式
 */
export class Message {
    reqId: number = 0
    type: MessageType = MessageType.REQUEST
    from: string
    route: RouteHandler
    data?: any
    constructor(from?: string, route?: RouteHandler, message?: any) {
        this.data = message
        this.from = from
        this.route = route
        this.reqId = Seq.Next()
    }
}

/**
 * 聊天消息
 */
 export interface ChatMessage {
    msgId?: number;
    to: string;
    content: string;
    type: ContentType;
    send_time?: number;
    extra?: string;
}

/**
 * 添加联系人
 */
export interface Contacts {
    to: string,
    remark: string
}