import axios from "axios";
import { UserStore } from '@/renderer/stores/modules/user';

const config = {
    //baseURL: import.meta.env.VITE_API_URL,
    baseURL: 'http://localhost',
    timeout: 10000,
}

/**
 * 公共请求方法
 * @param api api地址
 * @param params 请求参数
 * @param method http方法,默认POST
 * @returns 
 */
export const http = (api: string, params: any, method: string = 'POST') => {
    if (method == 'POST') {
        return axios.post(api, params, config).then((res) => {
            return res.data;
        }).catch((err) => {
            return err;
        })
    }
    if (method == 'GET') {
        return axios({ method: method, url: api, params: params }).then((res) => {
            return res.data;
        }).catch((err) => {
            console.log(err)
            return err;
        })
    }
}

// 添加请求拦截器
axios.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    if (UserStore().getToken) {
        const token = UserStore().getToken
        config.headers.Authorization = `${token}`
    }

    return config;
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});

// 添加响应拦截器
axios.interceptors.response.use(function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    return response;
}, function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    return Promise.reject(error);
});
