/**
 * 包json编码解码
 */
 export class pack {
    route: string
    sequence: number = 0
    status: number = 1
    unpacking(data: any) { return JSON.stringify(data) }
    balePack(data: any) { return JSON.stringify(data) }
}
