import { defineStore } from "pinia";

export const optionConfigStore = defineStore({
    id: 'optionConfig',
    state: () => {
        return {
            'automaticLogin': false,
            'rememberPassword': false,
        }
    },
    getters: {
        getAutomaticLogin(): boolean {
            return this.automaticLogin
        },
        getRememberPassword(): boolean {
            return this.rememberPassword
        },
    },
    actions: {
        setAutomaticLogin(bool: boolean) {
            this.automaticLogin = bool
        },
        setRememberPassword(bool: boolean) {
            this.rememberPassword = bool
        },
    }
})