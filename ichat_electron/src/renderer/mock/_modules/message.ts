import { MockMethod } from 'vite-plugin-mock'

export default [
    {
        url: '/getSessionList',
        method: 'post',
        response: () => {
            return {
                code: 200,
                msg: '执行成功',
                data: [
                    {
                        uid: '12345',
                        nickname: '吴妮基瓦',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },{
                        uid: '12345fdafd',
                        nickname: '八个亚克',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },{
                        uid: '123fdas45',
                        nickname: '股东模拟',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },{
                        uid: '123fda45',
                        nickname: '牛逼普拉斯',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },{
                        uid: '12fdfgfs345',
                        nickname: '真好也',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },{
                        uid: '12345hgdhjg',
                        nickname: '却尼玛的',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },{
                        uid: '1jhjh2345',
                        nickname: '嗯嗯嗯额',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },{
                        uid: '12fdafe345',
                        nickname: '大哥大',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },{
                        uid: '123czvc45',
                        nickname: '哟你你',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },{
                        uid: '123wew45',
                        nickname: '东城',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },{
                        uid: '1234cxa5',
                        nickname: '西就',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },{
                        uid: '1234gfhg5',
                        nickname: '老伯',
                        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
                        lastMessage: '',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 10,
                        lable: ''
                    },
                    {
                        uid: '123321hghg45',
                        nickname: '牛腩肉',
                        headPortraitUrl: '',
                        lastMessage: '你好吗',
                        lastMessageTime: '35:21',
                        unreadMessageNumber: 120,
                        lable: ''
                    }, {
                        uid: '12332gfgf23245',
                        nickname: '空你吉瓦',
                        headPortraitUrl: '',
                        lastMessage: '你好吗',
                        lastMessageTime: '15:20',
                        unreadMessageNumber: 2,
                        lable: ''
                    }, {
                        uid: '123323jhjh3245',
                        nickname: '睡觉了',
                        headPortraitUrl: '',
                        lastMessage: '你好吗',
                        lastMessageTime: '',
                        unreadMessageNumber: 33,
                        lable: ''
                    }, {
                        uid: '12334jhjh23245',
                        nickname: 'ceshi',
                        headPortraitUrl: '',
                        lastMessage: '你好吗',
                        lastTilastMessageTimeme: '',
                        unreadMessageNumber: 0,
                        lable: ''
                    }
                ]
            }
        }
    },
    {
        url: '/getChatRecord',
        method: 'post',
        response: () => {
            return {
                code: 200,
                msg: '执行成功',
                data: [
                    {
                        'fromUid': '123',
                        'targetUid': '好啊有',
                        'targetType': '123',//接收者类型
                        'msgType': '1',//消息内容类型
                        'msgUID': '123',
                        'content': '功能层面上，我们将通过提供语音的暂停和进度拖拽能力，并可视化音量，以满足语音接收者的使用效率需求。在体验层面上，语音作为用户的高频沟通操作，其设计必须满足QQ8.0中精致这一设计原则，给用户带来极致体验。',
                        'dateTime': '123',
                        'source': '123',
                        'isSelf': true,
                    }, {
                        'fromUid': '123',
                        'targetUid': '123',
                        'targetType': '123',//接收者类型
                        'msgType': '1',//消息内容类型
                        'msgUID': '123',
                        'content': '你今天吃饭了吗',
                        'dateTime': '123',
                        'source': '123',
                        'isSelf': false,
                    },
                ]
            }
        }
    }
] as MockMethod[]
