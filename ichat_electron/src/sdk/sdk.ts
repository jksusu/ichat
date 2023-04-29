import log from 'loglevel-es';
import { pack } from './pack/pack';
import { Request } from './request/request';
import { Response } from './request/response';
import { Message, ChatMessage } from './format/base';
import { RouteHandler, ContentType, ICahtClientStatus, MessageType } from './enum/enum';

const sendTimeout = 5 * 1000 // 5 seconds 没有回复则超时

export class Seq {
    static num: number = 0
    static Next() {
        Seq.num++
        Seq.num = Seq.num % 65536
        return Seq.num
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
    private sendq = new Map<number, Request>()//请求队列
    eventTree = {}//事件树

    private onMessageCallback: (e: Message) => void
    private onOpenCallback: () => void
    private onCloseCallback?: () => void
    constructor(gateway: string) {
        this.gateway = gateway
        this.registerEvent(RouteHandler.RouteUserInfoUpdate, this.noticeUpdateUserInfo)
        //this.registerEvent(RouteHandler.RouteIchatTalkToUser, this.noticeTalkToUser)
    }

    bindOnMessageCallback(callback: (message: Message) => void) { this.onMessageCallback = callback }
    bindOnOpenCallback(callback: () => void) { this.onOpenCallback = callback }
    bindOnCloseCallback(callback: () => void) { this.onCloseCallback = callback }

    /**
     * 注册事件
     * @param eventName 
     * @param callback 
     */
    registerEvent(eventName: string, callback: Function): { status: boolean, err?: string } {
        if (this.eventTree[eventName]) {
            return { status: false, err: '事件已经存在' }
        }
        this.eventTree[eventName] = callback

        log.info(this.eventTree)

        return { status: true }
    }

    /**
     * 批量注册
     * @param eventAll 
     * @returns 
     */
    registerEventAll = (eventAll): { status: boolean, err?: string } => {
        for (const key in eventAll) {
            if (Object.hasOwnProperty.call(eventAll, key)) {
                let { status, err } = this.registerEvent(key, eventAll[key])
                if (!status) {
                    return { status, err }
                }
            }
        }
        return { status: true }
    }

    /**
     * 分发服务端下发消息
     * @param message 
     * @returns 
     */
    private route(message: Message) {
        log.info(this.eventTree)
        let f = this.eventTree[message.route]
        if (!f) {
            log.debug('DEBUG:暂不支持此功能:' + message.route)
            return
        }
        f(message)
    }
    /**
     * 开始websocket连接
     * @returns 
     */
    async auth(): Promise<{ ok: boolean, err?: Error }> {
        if (this.status == ICahtClientStatus.CONNECTED) {
            return { ok: false, err: new Error("connection has been created") }
        }
        this.status = ICahtClientStatus.CONNECTING
        let { ok, err, conn } = await this.connect()
        if (!ok) {
            this.status = ICahtClientStatus.INIT
            return { ok: false, err: err }
        }
        conn.onmessage = (evt) => {
            let message = new Message()
            Object.assign(message, JSON.parse(<string>evt.data))
            log.info(JSON.parse(<string>evt.data))
            if (this.onMessageCallback) { this.onMessageCallback(message) }
            if (message.type == MessageType.RESPONSE || message.type == MessageType.RESPONSEPOMG) {
                let req = this.sendq.get(message.reqId)
                if (req) { req.callback(message); return }
            } else if (message.type == MessageType.NOTICE) {
                //通知消息，需要路由分发
                this.route(message)
            }
        }
        conn.onclose = (e) => {
            log.info("close")
            this.clearTimer("heartbeat")
            if (this.status == ICahtClientStatus.CLOSEING) {
                //调用关闭方法
                this.errHandler()
                return
            }
        }
        this.status = ICahtClientStatus.CONNECTED
        this.conn = conn

        //增加定时器
        this.setTimer("heartbeat", () => {
            this.heartbeat()
        }, 30000)
        log.info("INFO: ichat connect success")
        return { ok: true, err }
    }
    /**
     * ichat连接方法
     * @returns 
     */
    private async connect(): Promise<{ ok: boolean, err?: Error, conn?: WebSocket }> {
        return new Promise((resolve, _) => {
            let conn = new WebSocket(this.gateway)
            this.setTimer("connect", () => {
                this.status = ICahtClientStatus.INIT
                resolve({ ok: false, err: new Error("connect timer out") })
            }, 10000)
            conn.onerror = (error) => {
                log.error("ERROR:", error)
                this.clearTimer("connect")
                resolve({ ok: false, err: new Error("connect close"), conn: conn })
            }
            conn.onopen = () => {
                if (this.onOpenCallback) { this.onOpenCallback() }
                this.status = ICahtClientStatus.CONNECTED
                this.clearTimer("connect")
                resolve({ ok: true, conn: conn })
            }
        })
    }

    /**
     * 设置定时器
     * @param id 定时器ID
     * @param callback 定时器执行回调函数
     * @param time 定时器执行间隔时间，格式传毫秒 如5秒执行一次 5000
     * @returns 
     */
    setTimer = (id: string, callback: Function, time: number) => { if (this.timerMap[id]) { return }; this.timerMap[id] = setInterval(function () { callback() }, time) }
    /**
     * 清除定时器
     * @param id 定时器ID
     * @returns 
     */
    clearTimer = (id: string) => { if (!this.timerMap[id]) { log.info("INFO:timer not exist"); return }; clearInterval(this.timerMap[id]) }

    /**
     * 获取会话列表
     * @returns 
     */
    getSessionList = async (): Promise<Response> => { return await this.request(new Message(this.account, RouteHandler.RouteSessionLists)) }

    //更新缓存数据
    noticeUpdateUserInfo = async (message: Message) => {
        log.info('INFO: save account ' + message.data.uid)
        const data = message.data
        if (data.uid && data.connId) {
            this.account = data.uid
            this.connId = data.connId
            return
        }
        log.error("ERROR:" + message.route + " data is null")
    }
    //收到好友添加申请信息
    noticeContactsApply = async (message: Message) => {
        //增加一条消息，回复ack
        const data = message.data
    }
    noticeTalkToUser = async (message: Message) => {
        //更新最后一条收到的消息ID
        //ack已收到
        this.talkReceive(message.data.msgId, 2, message.route)
    }
    /**
     * 用户搜索
     * @param to 所有用户id
     * @param target 类型，用户|群聊 默认用户 1用户 2群聊
     * @returns 
     */
    async targetSearch(to: string, target?: number): Promise<{ status: boolean, res?: Response }> {
        let route = target == 1 ? RouteHandler.RouteUserInfoSearch : RouteHandler.RouteUserInfoSearch
        const message = new Message(this.account, route, { 'to': to })
        let resp = await this.request(message)
        return resp.success ? { status: true, res: resp } : { status: false, res: resp }
    }

    /**
     * 联系人添加
     * @param to 
     * @param remark 
     * @returns 
     */
    async contactsApply(to: string, remark?: string): Promise<{ status: boolean, res?: Response }> {
        const message = new Message(this.account, RouteHandler.RouteRelationFriendsApply, { 'to': to, 'remark': remark })
        let resp = await this.request(message)
        return resp.success ? { status: true, res: resp } : { status: false, res: resp }
    }

    async contactsEdit(to: string, remark?: string): Promise<{ status: boolean, res?: Response }> {
        return
    }

    /**
     * 好友申请答复
     * @param to 
     * @param status 
     * @param remark 
     */
    async contactsApplyReply(to: string, status: number, remark?: string): Promise<{ status: boolean, res?: Response }> {
        const message = new Message(this.account, RouteHandler.RouteRelationFriendsApplyReply, { 'to': to, 'status': status, 'remark': remark ?? '' })
        let resp = await this.request(message)
        return resp.success ? { status: true, res: resp } : { status: false, res: resp }
    }

    /**
     * 给USER发消息
     * @param to 用户UID
     * @param content 内容
     * @param type 内容类型
     */
    async talkToUser(to: string, content: string, type: ContentType): Promise<{ status: boolean, res?: Response }> {
        const chatMessage: ChatMessage = { to: to, content: content, type: type }
        const message = new Message(this.account, RouteHandler.RouteIchatTalkToUser, chatMessage)
        let resp = await this.request(message)
        return resp.success ? { status: true, res: resp } : { status: false, res: resp }
    }


    /**
     * 给群聊发消息
     * @param to 群聊ID
     * @param content 内容
     * @param type 内容类型
     */
    // async talkToGroup(to: string, content: string, type: ContentType): Promise<{ status: boolean, res?: Response }> {

    // }

    async talkReceive(msgId: number, status: number, type?: string): Promise<{ status: boolean, res?: Response }> {
        const message = new Message(this.account, RouteHandler.RouteIchatTalkAckReceive, { msgId: msgId, status: status, type: type })
        let resp = await this.request(message)
        return resp.success ? { status: true, res: resp } : { status: false, res: resp }
    }

    //关闭方法
    private ichatClose(info: string) {
        if (this.status == ICahtClientStatus.CLOSED) { return }
        this.status = ICahtClientStatus.CLOSED
        log.info("INFO: conn close" + info)
        this.conn = null
        if (this.onCloseCallback) { this.onCloseCallback() }
    }
    /**
     * Ichat websocket 请求方法
     * @param data
     * @returns 
     */
    private async request(data: Message): Promise<Response> {
        return new Promise((resolve, _) => {
            let tr = setTimeout(() => {
                log.info("INFO:" + "request reqId:" + data.reqId + " timeout")
                this.sendq.delete(data.reqId)
                resolve(new Response(false))
            }, sendTimeout)
            let callback = (msg: Message) => {
                clearTimeout(tr)
                this.sendq.delete(data.reqId)
                resolve(new Response(true, msg))
            }
            //请求队列
            this.sendq.set(data.reqId, new Request(callback))
            let balepacked = (new pack).balePack(data)
            log.info(balepacked)
            try {
                if (this.conn == null) { resolve(new Response(false)) }
                this.conn.send(balepacked)
            } catch (error) {
                resolve(new Response(false))
            }
        })
    }
    /**
     * 发送心跳包
     * @returns 
     */
    private async heartbeat() {
        let message = new Message()
        message.type = MessageType.REQUESTPING
        let resp = await this.request(message)
        if (!resp.success) {
            log.info("INFO:heartbeat ping error")
            this.reconnectionNumber++
            return
        }
        log.info("INFO:heartbeat-" + this.heartbeatNumber++)
    }
    //短线重连
    private async errHandler() {
        log.info("errhandler")
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

        this.ichatClose("timer out close")
    }
    /**
     * 退出登录，关闭websocket连接
     * @returns 
     */
    logout() {
        if (this.status == ICahtClientStatus.CLOSEING) { return }
        this.status = ICahtClientStatus.CLOSEING
        if (!this.conn) { return }
        try {
            log.info("INFO: websocket closeing")
            this.conn.close()
        } catch (error) {
            log.error("ERROR:", error)
        }
    }
}


export let sleep = async (second: number, Unit: 1000): Promise<void> => {
    return new Promise((resolve, _) => {
        setTimeout(() => {
            resolve()
        }, second * Unit)
    })
}