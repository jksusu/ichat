const path = require('path')
const { iconUrl, isDev, isMacos } = require('./env')
const ElectronTrayTips = require('electron-tray-tips');
const { app, BrowserWindow, ipcMain, Menu, Tray, nativeImage } = require('electron');
const console = require('console');

let mainWindow
let trayTwinkleIng = false
let appIcon = null//托盘图标
let mouseMoveIng = null//鼠标移动到托盘事件

const createWindow = () => {
    const mainWindow = new BrowserWindow({
        titleBarStyle: 'hidden',
        transparent: true,
        fullscreen: false,
        resizable: false,//禁止改变窗口尺寸
        backgroundColor: '#00000000',
        webPreferences: {
            preload: path.join(__dirname, 'preload.js'),
            nodeIntegration: true,
            webSecurity: false,
        }
    })
    console.log(process.env['NODE_ENV'])
    if (process.env['NODE_ENV'] == 'dev') {
        mainWindow.loadURL('http://localhost:5173/#/login')
        mainWindow.webContents.openDevTools({ mode: 'detach' })
    } else {
        console.log('正式环境')
        mainWindow.loadFile('dist/index.html')
        mainWindow.webContents.openDevTools({ mode: 'detach' })
    }
    return mainWindow
}
const createLoginWindow = () => { mainWindow.setContentSize(360, 560) }
const createMainWindow = () => { mainWindow.setContentSize(1440, 900), mainWindow.center() }


app.whenReady().then(() => {
    mainWindow = createWindow()
    createLoginWindow()
    app.on('activate', () => {
        if (BrowserWindow.getAllWindows().length === 0) createLoginWindow()
    }).on('window-all-closed', () => {
        if (isMacos) app.quit()
    })

    appIcon = new Tray(path.join(__dirname, iconUrl))
    appIcon.setContextMenu(menus)
    appIcon.setToolTip("悬停事件")
    appIcon.on('click', () => {
        mainWindow.show()
    })
    const ElectronTrayTipsClass = new ElectronTrayTips(appIcon);
    appIcon.addListener('mouse-move', ((event, position) => {
        if (!mouseMoveIng) {
            // ElectronTrayTipsClass.showTrayTips(`http://localhost:5173/#/messageTrip`, () => {
            //     mouseMoveIng = false
            // })
        }
    }))
})

//注册事件
const registryListeners = () => {
    ipcMain.on('app-quit', () => {
        console.log('app-quit')
        appIcon.destroy()
        app.quit()
    }).on('app-max', () => {
        //如果已经全屏则缩小
        if (mainWindow.isMaximized()) {
            mainWindow.unmaximize()
        } else {
            mainWindow.maximize()
        }
    }).on('app-min', () => {
        if (mainWindow.isFullScreen()) mainWindow.setFullScreen(false)
        if (mainWindow.isMaximized() || mainWindow.isNormal()) mainWindow.minimize()
    }).on('app-login', () => {
        console.log('app-login')
    }).on('show-home-window', () => {
        createMainWindow()//创建主窗体
    }).on('show-login-window', () => {
        createLoginWindow()
    }).on('tray-message-trip', () => {
        if (!trayTwinkleIng) {
            trayTwinkleIng = setInterval(() => {
                trayTwinkleIng = !trayTwinkleIng
                trayTwinkleIng ? appIcon.setImage(nativeImage.createEmpty()) : appIcon.setImage(path.join(__dirname, iconUrl))
                appIcon.setToolTip('您有一条新消息')
            }, 500)
        }
    }).on('cancel-tray-message-trip', () => {
        trayTwinkleIng = false
        clearInterval(trayTwinkleIng)
    })
}
registryListeners()


const toggleWinOSFullScreen = () => {
    const isMax = mainWindow.isMaximized()
    if (isMax) {
        mainWindow.unmaximize()
    } else {
        mainWindow.maximize()
    }
    return !isMax
}

/**
 * 菜单配置
 */
const menus = Menu.buildFromTemplate([
    {
        label: "意见反馈",
        click: function () {
            console.log("意见反馈")
        }
    },
    {
        label: "设置",
        click: function () {
            console.log("设置")
        }
    },
    {
        label: "升级",
        click: function () {
            console.log("升级")
        }
    },
    {
        label: "退出im",
        click: function () {
            app.quit()
        }
    }
])

const twinkle = () => {

}