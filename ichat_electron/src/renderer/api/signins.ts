import { http } from "../axios/axios"
import { Apis } from "./urls";

export const signin = (params: any) => {
    console.log('请求页面')
    return http(Apis.signin, params);
}

export const register = (params: any) => {
    return http(Apis.register, params);
}