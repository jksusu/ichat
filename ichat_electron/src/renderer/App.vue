<template>
  <div class="common-layout" v-if="$route.meta.auth">
    <el-container class="container">
      <el-aside>
        <AsideMenu></AsideMenu>
      </el-aside>
      <el-main class="el_mian" style="background-color: #ffffff;">
        <RouterView />
      </el-main>
    </el-container>
  </div>
  <div v-else>
    <RouterView />
  </div>
</template>
<script setup lang="ts">
import { useRouter } from 'vue-router';
import { UserStore } from './stores/modules/user';
import { showLoginWindow } from '@/common/ipcRenderer';
import AsideMenu from '@/renderer/components/AsideMenu.vue';

let router = useRouter()
router.beforeEach((to, from) => {
  const token = UserStore().getToken
  //读取本地token
  if (to.meta.auth && token == '') {
    showLoginWindow()
    return {
      path: '/login',
      // 保存我们所在的位置，以便以后再来
      query: { redirect: to.fullPath },
    }
  }
})

// if (!router) {
//   jump('/login')
// }
</script>
<style lang="scss">
:root {
  --itemHoverColor: #333333
}

body {
  margin: 0;
}

.common-layout {
  border-radius: 12px;
  overflow: hidden;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.container {
  min-width: 1440px;
  min-height: 900px;
  width: 100vw;
  height: 100vh;
  position: relative;
}

.el-aside {
  width: 64px;
}

.el-main {
  --el-main-padding: 0;
}
</style>