export const useIpcRenderer = () => {
    if (electronAPI) {
        return electronAPI.ipcRenderer
    }
    return null
}

export const sendIpc = (enevt, arg) => {
    const ipc = useIpcRenderer()
    if (ipc) ipc.send(enevt, arg)
}

/**
 * 托盘闪烁方法
 */
export const appIconTwinkle = () => {
    useIpcRenderer().send('tray-message-trip')
}

/**
 * 打开主页窗体
 */
export const showHomeWindow = () => {
    useIpcRenderer().send('show-home-window')
}

/**
 * 打开主页窗体
 */
 export const showLoginWindow = () => {
    console.log('登录窗体')
    useIpcRenderer().send('show-login-window')
}