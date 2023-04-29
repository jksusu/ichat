import { defineStore } from "pinia";

export const ChatRecordStore = defineStore({
    id: 'chatRecord',
    state: () => {
        return {
            'lists': {},
        }
    },
    getters: {
        getChatById: (state) => (id: string) => {
            return state.lists[id] || []
        }
    },
    actions: {
        addMessage(id: string, message) {
            if (!this.lists[id]) {
                this.lists[id] = []
            }
            this.lists[id].push(message)
        }
    },
    persist: true,
})