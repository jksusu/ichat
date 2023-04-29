import { http } from "../axios/axios"


const Apis = {
    login: '/v1/login',
    register: '/v1/register',
    contacts: '/v1/relation/contacts',
    singleChatRecord: '/v1/chat/getSingleChatRecord',
}

export const getSessionList = (params: any) => {
    return http('/getSessionList', params);
}

export const singleChatRecord = (params: any) => {
    return http(Apis.singleChatRecord, params)
}

export const login = (params: any) => {
    return http(Apis.login, params)
}

export const register = (params: any) => {
    return http(Apis.register, params)
}

export const contactsList = (params: any) => {
    return http(Apis.contacts, params)
}