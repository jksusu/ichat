import { defineStore } from "pinia";

export const NewFriendsStore = defineStore({
    id: 'newFriends',
    state: () => {
        return {
            'lists': [],
        }
    },
    getters: {
        getList() {
            return this.lists
        }
    },
    actions: {
        addMessage(message) {
            const isExist = this.lists.find(el => el.from === message.from)
            if (!isExist) {
                this.lists.push(message)
            }
        }
    },
    persist: true,
})

