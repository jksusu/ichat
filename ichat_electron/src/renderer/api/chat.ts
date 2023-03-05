import { http } from "../axios/axios"

export const getSessionList = (params: any) => {
    return http('/getSessionList', params);
}

export const getChatRecord = (params: any) => {
    return http('/getChatRecord', params);
}