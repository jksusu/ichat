/**
 * 将微妙时间戳转换成微信显示时间规则
 * @param {number} timestamp 微妙时间戳
 * @return {string} 微信显示时间字符串
 */
export const formatWechatTime = (timeStamp) => {
    const millisecondTimestamp = Math.floor(timeStamp / 1000000);
    const date = new Date(millisecondTimestamp);
    const now = new Date();
    const sameDay = date.getDate() === now.getDate() && date.getMonth() === now.getMonth() && date.getFullYear() === now.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();
    const hour = date.getHours() < 10 ? '0' + date.getHours() : date.getHours();
    const minute = date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes();

    if (sameDay) {
        return `${hour}:${minute}`;
    } else if (date.getFullYear() === now.getFullYear()) {
        return `${month}月${day}日 ${hour}:${minute}`;
    } else {
        return `${month}-${day}-${date.getFullYear().toString().substr(-2)}:${hour}:${minute}`;
    }
}

/**
 * 列表最后一条消息时间展示规则
 * @param timestamp 时间展示
 * @returns 
 */
export const LastMessageTimeShow = (timestamp?: number): string => {
    if (!timestamp) {
        return ""
    }
    const now = new Date(timestamp)
    const hour = now.getHours()
    const minute = now.getMinutes()

    let result: string;

    if (hour < 10) {
        result = `0${hour}:`
    } else {
        result = `${hour}:`
    }
    if (minute < 10) {
        result += `0${minute}`
    } else {
        result += `${minute}`
    }
    return result
}