import { fa, tr } from 'element-plus/es/locale';
import log from 'loglevel-es';
import { resolve } from 'path';

//日志处理
// const log = (log: any) => {
//     console.log(log)
// }

enum ICahtClientStatus {
    INIT,
    CONNECTING,
    CONNECTED,
    RECONNECTING,
    CLOSEING,
    CLOSED,
}

// enum MsgType {
//     TextMessage = 1,
//     BinaryMessage = 2,
//     CloseMessage = 8,
//     PingMessage = 9,
//     PongMessag = 10,
// }

enum ContentType {
    text = 1,
    image = 2,
    video = 3,
}

//路由处理器
enum RouteHandler {
    RouteChatUserMessage = 'chat.user.message',
    RouteChatGroupMessage = 'chat.group.message',
    RouteChatMessageAck = 'chat.message.ack',
}

enum BusinessMessageType {
    PINGTYPE = 'PING',
    CHATTYPE = 'CHAT',
}

//基础消息格式
// const Message = {
//     MsgType: MsgType,
//     MsgData: ""
// }

//业务消息格式
interface BusinessMessage {
    type: BusinessMessageType,
    date: ChatMessage
}

//消息格式
interface ChatMessage {
    msgv: string;
    msgId: number;
    seq: number;
    route: RouteHandler;
    from: number;
    to: number;
    content: string;
    type: ContentType;
    send_time: number;
    extra: string;
}


export class Message {
    sequence: number = 0;
    type: number = 1;
    message?: string;
    from?: string; // sender
    constructor(message?: string) {
        this.message = message;
        this.sequence = 1
    }
}
export class Request {
    sendTime: number
    callback: (response: Message) => void
    constructor(callback: (response: Message) => void) {
        this.sendTime = Date.now()
        this.callback = callback
    }
}
export class Response {
    success: boolean = false
    message?: Message
    constructor(success: boolean, message?: Message) {
        this.success = success;
        this.message = message;
    }
}



export class IChatClient {
    conn?: WebSocket//websocket实例
    wsAgainNumbe: number//重连次数
    wsUrl: string
    account: number
    status: ICahtClientStatus
    timerMap = {}//定时器
    private sendq = new Map<number, Request>()
    constructor(wsurl: string) {
        this.status = ICahtClientStatus.INIT//初始画实例状态
        this.wsUrl = wsurl
        if (this.conn == null) {
            this.conn = this.connect()
        }
    }

    async request(data: Message): Promise<Response> {
        return new Promise((resolve, _) => {
            let seq = data.sequence

            // asynchronous wait ack from server
            let callback = (msg: Message) => {
                // remove from sendq
                this.sendq.delete(seq)
                resolve(new Response(true, msg))
            }

            this.sendq.set(seq, new Request(callback))

            if (!this.send(JSON.stringify(data))) {
                resolve(new Response(false))
            }
        })
    }

    send(data: string): boolean {
        try {
            if (this.conn == null) {
                return false
            }
            this.conn.send(data)
            return true
        } catch (error) {
            log.error(error)
        }
    }

    connect(): WebSocket {
        if (this.status == ICahtClientStatus.CONNECTED) {
            log.info("已经是连接成功状态了")
            return
        }
        this.status = ICahtClientStatus.CONNECTING
        let conn = null
        try {
            conn = new WebSocket(this.wsUrl)
            conn.onopen = () => {
                log.info("websocket 连接成功")
                this.status = ICahtClientStatus.CONNECTED
                //心跳开始
                this.heartbeat()
            }
            conn.onmessage = (evt) => {
                try {
                    let msg = new Message();
                    Object.assign(msg, JSON.parse(<string>evt.data))
                    if (msg.type == 2) {
                        let req = this.sendq.get(msg.sequence)
                        if (req) {
                            req.callback(msg)
                        }
                    } else if (msg.type == 3) {
                        console.log(msg.message, msg.from)
                    }
                } catch (error) {
                    console.error(evt.data, error)
                }
            }
            conn.onerror = (error) => {
                log.error(error)
            }
            log.info(this.conn)

        } catch (error) {
            log.error(error)
        }
        return conn
    }
    setTimer = (id: string, callback: Function, time: number) => {
        if (this.timerMap[id]) { return }
        var timerId = setInterval(function () { callback() }, time)
        this.timerMap[id] = timerId
    }
    chatUserMsg(to: number, content: string, type: ContentType) {
        const message: ChatMessage = {
            msgv: "0.1",
            seq: 0,
            msgId: 1234,
            route: RouteHandler.RouteChatUserMessage,
            from: this.account,
            to: to,
            content: content,
            type: type,
            send_time: new Date().getTime(),
            extra: "",
        }
        this.to(message)
    }
    chatGroupMsg() {

    }
    to(chat: ChatMessage) {
        if (this.conn.readyState == 1) {
            const message = { type: BusinessMessageType.CHATTYPE, date: chat } as BusinessMessage
            const m = (new pack).balePack(message)
            log.info(m)
            this.conn.send(m)
        }
        log.info("为主吧内耗")
    }
    heartbeat() {
        this.setTimer("heartbeat", () => {
            log.info("heartbeat")
            this.conn.send('{"type":"PING"}')
        }, 20000)
    }
    logout() {
        if (this.status == ICahtClientStatus.CLOSEING) {
            return
        }
        this.status = ICahtClientStatus.CLOSEING
        if (!this.conn) { return }

        try {
            log.info("websocket closeing")
            this.conn.close()
        } catch (error) {
            log.error(error)
        }
    }
}

//包处理
export class pack {
    unpacking(data: any) {
        return JSON.stringify(data)
    }
    balePack(data: any) {
        return JSON.stringify(data)
    }
}