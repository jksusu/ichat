import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import { fileURLToPath, URL } from 'node:url'
import { viteMockServe } from 'vite-plugin-mock'
import { svgBuilder } from './src/renderer/plugins/svgBuilder';

// https://vitejs.dev/config/
export default defineConfig({
  base: './',
  plugins: [
    vue(),
    vueJsx(),
    viteMockServe({
      mockPath: "./src/renderer/mock/_modules/", // 解析，路径可根据实际变动
      localEnabled: true // 此处可以手动设置为true，也可以根据官方文档格式
    }),
    svgBuilder("./src/renderer/assets/svgs/"),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
})
