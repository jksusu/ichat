import { defineStore } from "pinia";

export const SessionStore = defineStore({
    id: 'sessionList',
    state: () => {
        return {
            selectSession: {
                uid: '',
                nickname: '',
                headPortraitUrl: '',
                type: ''
            },
            sessionList: [
                {
                    uid: '',
                    nickname: '',
                    headPortraitUrl: '',
                    lastMessage: '',
                    lastMessageTime: '',
                    unreadMessageNumber: 0,
                    lable: '',
                    type: ''
                }
            ]
        }
    },
    getters: {
        getSessionLists() {
            return this.sessionList
        },
        getSelectSession() {
            return this.selectSession
        },
        getSelectSessionUID(): boolean {
            if (this.selectSession.uid) {
                return true
            }
            return false
        }
    },
    actions: {
        addSession(uid: string, nickname?: string, headPortraitUrl?: string, lastMessage?: string) {
            const isExist = this.sessionList.find(el => el.uid === uid)
            if (!isExist) {
                this.sessionList.push({
                    uid: uid,
                    nickname: nickname,
                    headPortraitUrl: headPortraitUrl,
                    lastMessage: '',
                    lastMessageTime: Date.now(),
                    unreadMessageNumber: 0,
                    lable: ''
                })
            }
        },
        delSession(uid: string) {
            this.sessionList = this.sessionList.filter(el => el.uid != uid)
            if (this.selectSession.uid == uid) {
                Object.keys(this.selectSession).forEach(key => {
                    this.selectSession[key] = '';
                })
            }
        },
        delSelectSession() {
            Object.keys(this.selectSession).forEach(key => {
                this.selectSession[key] = '';
            })
        },
        setSelectSession(uid: string, nickname: string, headPortraitUrl: string, type) {
            this.selectSession.uid = uid
            this.selectSession.nickname = nickname
            this.selectSession.headPortraitUrl = headPortraitUrl
            this.selectSession.type = type
        },
        setSessionTop(uid: string) {
            const index = this.sessionList.findIndex((el) => el.uid == uid)
            if (index !== -1) {
                const [session] = this.sessionList.splice(index, 1);
                this.sessionList.unshift(session)
            }
        },
        cancelSessionTop(uid: string) {
            console.log("取消消息置顶，根据加入session时间排序")
        }
    },
    persist: true,
})