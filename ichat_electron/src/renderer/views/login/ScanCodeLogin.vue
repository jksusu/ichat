<template>
    <WindowContainer>
        <template v-slot:container-content>
            <LogoRegion route="/login" :icon-url="iconUrl" :trip-url="tripUrl"></LogoRegion>
            <div id="qrcode-login">
                <div v-if="state.status != 3" class="qrcode">
                    <div v-if="state.status == 2" class="location index">
                        <img @click="refresh" :src="qccodeInvalidUrl" />
                        <div class="qccode_invalid_message">
                            <p>二维码失效点击刷新</p>
                        </div>
                    </div>
                    <qrcode-vue v-else class="location" :value="state.code" :size="size" level="H" />
                </div>
                <div v-if="state.status != 3" class="qrcode-login-message">请使用APP扫码登录</div>

                <div v-if="state.status == 3" class="qrcode_success">
                    <img width="76" height="76" :src="scanningCodeSuccessUrl" />
                    <div class="qrcode_success_message">
                        <div>扫码成功</div>
                        <div>请在app上确定登录!</div>
                    </div>
                </div>
                <div class="bottom-options">
                    <label v-on:click="jump('/login')">账号登录</label>
                    <label style="margin-left: 8px;">|</label>
                    <label v-on:click="jump('/retrieve_password')" style="margin-left: 8px;">忘记密码</label>
                </div>
            </div>
        </template>
        <template v-slot:container-bottom>
            <LoginBottonMessage :status="false">
                <template v-slot:message-p>
                    <p>{{ message }}</p>
                </template>
            </LoginBottonMessage>
        </template>
    </WindowContainer>
</template>
  
<script lang="ts" setup>
import QrcodeVue from 'qrcode.vue'
import { jump } from '@/renderer/tool/tool';
import { reactive, onBeforeUnmount } from "vue"
import LogoRegion from '@/renderer/views/login/components/LogoRegion.vue';
import WindowContainer from '@/renderer/views/login/components/WindowContainer.vue';
import LoginBottonMessage from '@/renderer/views/login/components/BottomMessage.vue';
import iconUrl from '@/renderer/assets/images/login_method_icon_account.png'
import tripUrl from '@/renderer/assets/images/login_method_trip_account.png'
import qccodeInvalidUrl from '@/renderer/assets/images/qccode_invalid.png'
import scanningCodeSuccessUrl from '@/renderer/assets/images/scanning_code_success.png'

const message = '';
const size = 160;


const state = reactive({
    timeoutTime: 60,//二维码失效时间
    code: "abcsddfd",//扫码登录值
    status: 1,//1显示二维码 2.二维码过期 3.扫码成功手机上点确定 4登录中
})

const refresh = () => {
    state.status = 1
    //刷新二维码
    sacnLoginCode
}


//二维码失效定时器
let timer = null

const sacnLoginCode = () => {
    //请求api

    //成功后设置定时器，失效后加上蒙层
    timer = setInterval(() => {
        state.timeoutTime--
        if (state.timeoutTime <= 1) {
            clearInterval(timer);
            state.timeoutTime = 60;
            state.status = 2;
        }
    }, 1000);
}

sacnLoginCode()

onBeforeUnmount(() => {
    clearInterval(timer)
})
</script>
  
































<style lang="scss" src="@/renderer/assets/scss/login/scanCodeLogin.scss" />