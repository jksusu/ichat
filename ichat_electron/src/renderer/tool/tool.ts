import { router } from "@/renderer/router"

export const jump = (route: string) => {
    router.push(route)
}