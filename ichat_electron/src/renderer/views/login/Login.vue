<template>
  <WindowContainer>
    <template v-slot:container-content>
      <LogoRegion route="/scan_code_login" :icon-url="iconUrl" :trip-url="tripUrl"></LogoRegion>
      <div id="login-page">
        <el-form ref="ruleFormRef" :model="loginFrom" :rules="rules">
          <el-form-item prop="username">
            <el-input v-model="loginFrom.username" placeholder="请输入用户名" />
          </el-form-item>
          <el-form-item prop="password">
            <el-input type="password" show-password v-model="loginFrom.password" placeholder="请输入密码" />
          </el-form-item>
        </el-form>
        <div class="login-option">
          <el-checkbox v-model="loginFrom.automaticLogin" label="自动登录" size="large" />
          <el-checkbox v-model="loginFrom.rememberPassword" label="记住密码" size="large" />
          <label class="password" @click="jump('/retrieve_password')">
            <span class="msg-font">找回密码</span>
            <span style="margin-left: 3px;"><img width="12" height="12" :src="detail" /></span>
          </label>
        </div>
        <el-button type="primary" @click="login(ruleFormRef)">{{ loginFrom.loginButtonName }}</el-button>
      </div>
    </template>
    <template #container-bottom>
      <LoginBottonMessage :status="loginFrom.status">
        <template #message-p>
          <p>{{ loginFrom.message }}</p>
        </template>
      </LoginBottonMessage>
    </template>
  </WindowContainer>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { jump } from '@/renderer/tool/tool';
import { router } from '@/renderer/router';
//import { signin } from '@/renderer/api/signins';

import { showHomeWindow } from '@/common/ipcRenderer';
import LogoRegion from '@/renderer/views/login/components/LogoRegion.vue';
import { optionConfigStore } from '@/renderer/stores/modules/optionConfig'
import WindowContainer from '@/renderer/views/login/components/WindowContainer.vue';
import LoginBottonMessage from '@/renderer/views/login/components/BottomMessage.vue';
import { UserStore } from '@/renderer/stores/modules/user';
import iconUrl from '@/renderer/assets/images/login_method_icon_qccode.png'
import tripUrl from '@/renderer/assets/images/login_method_trip_qccode.png'
import detail from '@/renderer/assets/images/detail.png'
import { signin } from '@/renderer/api/signins';
import type { FormInstance, FormRules } from 'element-plus'


const ruleFormRef = ref<FormInstance>()
const rules = reactive<FormRules>({
  username: [{ required: true, message: '请输入账号', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
})
const loginFrom = reactive({
  username: 'admin',
  password: '123456',
  automaticLogin: false,
  rememberPassword: false,
  message: '',
  status: false,
  loginButtonName: "登录",
})

const login = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      //写入勾选 option
      optionConfigStore().setAutomaticLogin(loginFrom.automaticLogin)
      optionConfigStore().setRememberPassword(loginFrom.rememberPassword)

      console.log('开始登录')

      loginFrom.loginButtonName = '登陆中'
      signin({ 'username': loginFrom.username, password: loginFrom.password }).then((res) => {
        res.code =200
        if (res.code == 200) {
          UserStore().setToken('fdafdasfds')
          showHomeWindow()
          router.push('/session')
        } else {
          loginFrom.status = true
          loginFrom.message = '错误'
          loginFrom.loginButtonName = '登陆'
        }
      })
    } else {
      return false
    }
  })
}
</script>


<style lang="scss" src="@/renderer/assets/scss/login/login.scss" />