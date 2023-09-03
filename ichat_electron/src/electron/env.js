const isDev = () => {
    return process.env['NODE_ENV'] == 'dev' ? true : false
}

/**
 * 运行环境
 * @returns bool
 */
const isMacos = () => {
    return process.platform === 'darwin' ? true : false
}
const isWinos = () => {
    return process.platform === 'win32' ? true : false
}

/**
 * 获取 icon
 * @returns string
 */
const iconUrl = isMacos ? '../../public/icon.png' : '../../public/icon_2x_256x256.ico'

module.exports = {
    isDev,
    isMacos,
    isWinos,
    iconUrl,
}