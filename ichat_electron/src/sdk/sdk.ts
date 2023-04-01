import log from 'loglevel-es';

//日志处理
// const log = (log: any) => {
//     console.log(log)
// }
const sendTimeout = 5 * 1000 // 10 seconds

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
    RouteConnectPingMessage = 'connect.ping',
    RouteConnectPongMessage = 'connect.pong',

    //业务
    RouteChatUserMessage = 'chat.user.message',
    RouteChatGroupMessage = 'chat.group.message',
    RouteChatMessageAck = 'chat.message.ack',

    //session
    RouteSessionLists = 'chat.session.lists',

    //信息修改
    RouteUserInfoUpdate = 'chat.user.info.update',
}

//基础消息格式
// const Message = {
//     MsgType: MsgType,
//     MsgData: ""
// }

//业务消息格式
interface BusinessMessage {
    seq: 1
    type: MessageType,
    data: ChatMessage
}

//消息格式
interface ChatMessage {
    msgId: number;
    route: RouteHandler;
    from: string;
    to: string;
    content: string;
    type: ContentType;
    send_time: number;
    extra: string;
}

//消息包类型
enum MessageType {
    REQUEST = 1,
    RESPONSE = 2,
    NOTICE = 3,
    REQUESTPING = 4,
    RESPONSEPOMG = 5
}

export class Seq {
    static num: number = 0
    static Next() {
        Seq.num++
        Seq.num = Seq.num % 65536
        return Seq.num
    }
}

//消息 对应服务端 reqeust数据
export class Message {
    reqId: number = 0
    type: MessageType = MessageType.REQUEST
    from: string
    route: RouteHandler
    data?: any
    constructor(message?: any) {
        this.data = message
        this.reqId = Seq.Next()
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
    connId?: string//连接id
    wsAgainNumbe: number//重连次数
    heartbeatNumber: number = 0//心跳次数
    maxReconnectionNumber: number = 3 //最大重连次数 时间规则为阶梯计算以 10 20 30 递增
    reconnectionNumber: number = 0//心跳次数
    gateway: string
    account: string
    status: ICahtClientStatus = ICahtClientStatus.INIT
    timerMap = {}//定时器
    private sendq = new Map<number, Request>()
    private funcTree//方法树
    private onMessageCallback: (e: Message) => void
    private onOpenCallback: () => void
    private onCloseCallback?: () => void
    constructor(gateway: string) {
        this.gateway = gateway
        this.initFucTree()
    }
    bindOnMessageCallback(callback: (message: Message) => void) { this.onMessageCallback = callback }
    bindOnOpenCallback(callback: () => void) { this.onOpenCallback = callback }
    bindOnCloseCallback(callback: () => void) { this.onCloseCallback = callback }

    async auth(): Promise<{ ok: boolean, err?: Error }> {
        if (this.status == ICahtClientStatus.CONNECTED) {
            return { ok: false, err: new Error("connection has been created") }
        }
        this.status = ICahtClientStatus.CONNECTING

        let { ok, err, conn, connId } = await this.connect()
        if (!ok) {
            this.status = ICahtClientStatus.INIT
            return { ok: false, err: err }
        }
        log.info("connect success connId:" + connId)

        conn.onmessage = (evt) => {
            let msg = new Message();
            Object.assign(msg, JSON.parse(<string>evt.data))

            console.log(msg)

            if (msg.type == MessageType.RESPONSE) {
                let req = this.sendq.get(msg.reqId)
                if (req) {
                    req.callback(msg)
                    return
                }
            } else if (msg.type == MessageType.NOTICE) {
                this.route(msg)
            }
        }
        conn.onclose = (e) => {
            log.info("close")
            this.clearTimer("heartbeat")
            if (this.status == ICahtClientStatus.CLOSEING) {
                //调用关闭方法
                return
            }
        }
        this.status = ICahtClientStatus.CONNECTED
        this.conn = conn
        this.connId = connId

        this.setTimer("heartbeat", () => {
            log.info("heartbeat:" + this.heartbeatNumber++)
            //this.heartbeat()
        }, 20000)
        return { ok: true, err }
    }
    private async connect(): Promise<{ ok: boolean, err?: Error, conn?: WebSocket, connId?: string }> {
        return new Promise((resolve, _) => {
            let conn = new WebSocket(this.gateway)

            //设置定期器，超过时间
            this.setTimer("connect", () => {
                this.status = ICahtClientStatus.INIT
                log.info("connect timer out")
                resolve({ ok: false, err: new Error("connect timer out") })
            }, 10000)

            conn.onerror = (error) => {
                log.error(error)
                this.clearTimer("connect")
                resolve({ ok: false, err: new Error("connect close"), conn: conn })
            }
            conn.onopen = () => {
                log.info("websocket 连接成功")
                this.status = ICahtClientStatus.CONNECTED
                this.clearTimer("connect")

                resolve({ ok: true, conn: conn })
            }
        })
    }
    private route(message: Message) {
        let f = this.funcTree[message.route]
        if (!f) {
            log.debug('方法数 方法:' + message.route + ";不存在")
            return
        }
        f(message)
    }
    private initFucTree() {
        this.funcTree = {
            [RouteHandler.RouteSessionLists]: this.sessionList,
            [RouteHandler.RouteUserInfoUpdate]: this.updateUserInfo,
        }
        console.log(this.funcTree)
    }

    setTimer = (id: string, callback: Function, time: number) => {
        if (this.timerMap[id]) { return }
        var timerId = setInterval(function () { callback() }, time)
        this.timerMap[id] = timerId
    }
    clearTimer = (id: string) => {
        if (!this.timerMap[id]) {
            log.debug(id + " timer not exist")
            return
        }
        clearTimeout(this.timerMap[id])
    }
    async chatUserMsg(to: string, content: string, type: ContentType): Promise<{ status: boolean, res?: Response }> {
        const message: ChatMessage = {
            msgId: 1234,
            route: RouteHandler.RouteChatUserMessage,
            from: this.account,
            to: to,
            content: content,
            type: type,
            send_time: new Date().getTime(),
            extra: "",
        }
        return this.to(message)
    }
    updateUserInfo = async (uid: string, connId: string) => {
        this.account = uid
        this.connId = connId
        log.info('执行成功了')
    }
    chatGroupMsg() {

    }
    async sessionList(): Promise<Response> {
        let m = new Message()
        m.route = RouteHandler.RouteSessionLists
        m.from = this.account
        return await this.request(m)
    }
    private async to(chat: ChatMessage): Promise<{ status: boolean, res?: Response }> {
        if (this.conn.readyState == 1) {
            let m = new Message(chat)
            let resp = await this.request(m)
            if (resp.success) {
                return { status: true, res: resp }
            }
            return { status: false }
        }
        log.info("websocket 未准备好")
        return { status: false }
    }
    //关闭方法
    private connClose(info: string) {
        if (this.status == ICahtClientStatus.CLOSED) {
            return
        }
        this.status = ICahtClientStatus.CLOSED
        log.info("conn close:" + info)
        this.conn = null
        if (this.onCloseCallback) {
            this.onCloseCallback()
        }
    }
    async request(data: Message): Promise<Response> {
        return new Promise((resolve, _) => {
            //设置发送消息超时时间
            let tr = setTimeout(() => {
                log.info("request reqId:" + data.reqId + "timeout")
                this.sendq.delete(data.reqId)
                resolve(new Response(false))
            }, sendTimeout)

            let callback = (msg: Message) => {
                clearTimeout(tr)
                this.sendq.delete(data.reqId)
                resolve(new Response(true, msg))
            }

            log.info("request reqId:" + data.reqId + " data:" + data.data)

            //请求队列
            this.sendq.set(data.reqId, new Request(callback))
            let m = (new pack).balePack(data)
            log.info(m)
            if (!this.send(m)) {
                resolve(new Response(false))
            }
        })
    }
    private send(data: string): boolean {
        try {
            if (this.conn == null) {
                return false
            }
            this.conn.send(data)
        } catch (error) {
            return false
        }
        return true
    }
    async heartbeat() {
        let m = new Message()
        m.type = MessageType.REQUESTPING
        let resp = await this.request(m)
        if (!resp.success) {
            log.info("heartbeat ping error")
            this.reconnectionNumber++
        }
        //this.clearTimer("heartbeat")
    }
    //短线重连
    private async errHandler() {
        if (this.status == ICahtClientStatus.CLOSED || this.status == ICahtClientStatus.CLOSEING) {
            return
        }
        this.status = ICahtClientStatus.RECONNECTING
        for (let index = 0; index <= 3; index++) {
            await sleep(3, 1000)
            try {
                log.info("try to relogin")
                let { ok, err } = await this.auth()
                if (ok) {
                    return
                }
                log.info(err)
            } catch (error) {
                log.warn(error)
            }
        }

        this.connClose("timer out close")
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
    route: string
    sequence: number = 0
    status: number = 1

    unpacking(data: any) {
        return JSON.stringify(data)
    }
    balePack(data: any) {
        return JSON.stringify(data)
    }
}


export let sleep = async (second: number, Unit: 1000): Promise<void> => {
    return new Promise((resolve, _) => {
        setTimeout(() => {
            resolve()
        }, second * Unit)
    })
}