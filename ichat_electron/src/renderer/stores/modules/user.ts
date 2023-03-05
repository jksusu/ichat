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
        }
    },
    actions: {
        setToken(token: string) {
            this.token = token
        },
        deleteToken() {
            this.token = ''
        },
    }
})