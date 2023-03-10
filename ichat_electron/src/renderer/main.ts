import { createApp } from 'vue'
import { createPinia } from 'pinia'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from '@/renderer/App.vue'
import { router } from '@/renderer/router'
import { setupStore } from './stores'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import { IChatClient } from '@/sdk/sdk'
import log from 'loglevel-es';


const gateway = "ws://127.0.0.1:9501/ws?token=123456"

log.setLevel("DEBUG")
let client = new IChatClient(gateway)

client.connect()
client.chatUserMsg(666666,"dd",1)

const app = createApp(App)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.use(createPinia()).use(setupStore).use(router).use(ElementPlus).mount('#app')