import { MockMethod } from 'vite-plugin-mock'

export default [
    {
        url: '/signin',
        method: 'post',
        response: () => {
            return {
                code: 200,
                msg: '执行成功',
                data: {
                    token: '212121',
                }
            }
        }
    },
    {
        url: '/register',
        method: 'post',
        response: () => {
            return {
                code: 200,
                msg: '执行成功',
                data: {
                    token: '212121',
                }
            }
        }
    }
] as MockMethod[] // 这里其实就是定义数据格式的，不了解的同学可以参考typescript的官方文档
