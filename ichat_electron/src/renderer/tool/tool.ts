import { router } from "@/renderer/router"

export const jump = (route: string) => {
    router.push(route)
}
//是否 win环境
export const ifWin = () => {
    return import.meta.env.VITE_TERMINAl == 'win' ? true : false
}