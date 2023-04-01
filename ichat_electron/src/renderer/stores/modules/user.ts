import { defineStore } from "pinia";

export const UserStore = defineStore({
    id: 'user',
    state: () => {
        return {
            token: '',
            menus: '',
            userInfo: {
                userNumber: '',
                nickname: '',
                headPortrait: '',
                signature: '',
            }
        }
    },
    getters: {
        getToken(): string {
            return this.token
        },
        getUserNumber(): string {
            return this.userInfo.userNumber
        },
        getHeadPortrait(): string {
            return this.headPortrait
        },
        getNickname(): string {
            return this.nickname
        },
        getUserInfo() {
            return this.userInfo
        }
    },
    actions: {
        setToken(token: string) {
            this.token = token
        },
        deleteToken() {
            this.token = ''
        },
        setUserInfo(uid: string, nickname: string, headPortrait: string, signature: string) {
            this.userInfo.userNumber = uid
            this.userInfo.nickname = nickname
            this.userInfo.headPortrait = headPortrait
            this.userInfo.signature = signature
        }
    },
    persist: true,
})