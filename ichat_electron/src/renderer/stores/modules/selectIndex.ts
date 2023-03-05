import { defineStore } from "pinia";

export const useSelectIndexStore = defineStore({
    id: 'userSelectIndex',
    state: () => {
        return {
            'asideMenuIndex': '/session',
            'sessionIndex': '',
            'sessionIndexInfo': {},
            'maillistIndex': '/maillist/new_friends',
            'newFriendsIndex': '',
            'friendsIndex': '',
            'groupsIndex': '',
        }
    },
    getters: {
        getMaillistIndex(): string {
            return this.maillistIndex
        },
        getNewFriendsIndex(): string {
            return this.newFriendsIndex
        },
        getFriendsIndex(): string {
            return this.friendsIndex
        },
        getGroupsIndex(): string {
            return this.groupsIndex
        },
        getAsideMenuIndex(): string {
            return this.asideMenuIndex
        },
        getSessionIndex(): string {
            return this.sessionIndex
        },
        getSessionIndexInfo(): any {
            return this.sessionIndexInfo
        }
    },
    actions: {
        setMaillistIndex(index: string) {
            this.maillistIndex = index
        },
        setNewFriendsIndex(index: string) {
            this.newFriendsIndex = index
        },
        setFriendsIndex(index: string) {
            this.friendsIndex = index
        },
        setGroupsIndex(index: string) {
            this.groupsIndex = index
        },
        setAsideMenuIndex(index: string) {
            this.asideMenuIndex = index
        },
        setSessionIndex(index: string) {
            this.sessionIndex = index
        },
        setSessionIndexInfo(index: any) {
            this.sessionIndexInfo = index
        },
    },
    persist: true,
})