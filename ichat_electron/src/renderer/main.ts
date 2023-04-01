import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from '@/renderer/App.vue'
import { router } from '@/renderer/router'
import { setupStore } from './stores'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import log from 'loglevel-es';
import { directive, Vue3Menus } from 'vue3-menus'
log.setLevel("DEBUG")
const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.directive('menus', directive).component('vue3-menus', Vue3Menus).use(createPinia()).use(setupStore).use(router).use(ElementPlus).mount('#app')