import { computed, createApp, watch } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from '@/renderer/App.vue'
import { router } from '@/renderer/router'
import { setupStore } from './stores'
import log from 'loglevel-es'
import Menus from 'vue3-menus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

log.setLevel("DEBUG")
const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.use(createPinia()).use(setupStore).use(router).use(ElementPlus).use(Menus, {}).mount('#app')
