import { IChatClient } from "@/sdk/sdk"
import { defineStore } from "pinia"

export const SocketStore = defineStore({
    id: 'socket',
    state: () => {
        return {
            status: false,
            sokcet: null,
        }
    },
    getters: {
        getSocket(): IChatClient {
            return this.sokcet
        },
        getStatus(): boolean {
            return this.status
        }
    },
    actions: {
        setSocket(sokcet: IChatClient) {
            this.sokcet = sokcet
        },
        setStatus(status: boolean) {
            this.status = status
        }
    },
})