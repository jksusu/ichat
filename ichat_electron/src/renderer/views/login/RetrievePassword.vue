<template>
    <WindowContainer>
        <template v-slot:container-content>
            <div id="retrieve-password-from">
                <div v-if="state.show" class="retrieve_password_from">
                    <div class="title">
                        <div class="left">
                            <span>忘记密码</span>
                        </div>
                        <div class="right" v-on:click="jump('/login')">
                            <span>返回登录</span>
                            <img :src="detail" />
                        </div>
                    </div>
                    <el-form :label-position="state.labelPosition" :model="state">
                        <el-form-item>
                            <el-input v-model="state.phone" placeholder="请输入您的手机号" />
                        </el-form-item>
                        <el-form-item>
                            <el-input v-model="state.code" placeholder="请输入验证码">
                                <template #suffix>
                                    <div v-on:click="checkCode()" class="code_button" :class="checkCodeClass">
                                        <button>{{ buttonName }}</button>
                                    </div>
                                </template>
                            </el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-input type="password" show-password v-model="state.newpassword1" placeholder="请设置新密码" />
                        </el-form-item>
                        <el-form-item>
                            <el-input type="password" show-password v-model="state.newpassword2" placeholder="请确认新密码" />
                        </el-form-item>
                    </el-form>
                    <div class="retrieve_trip">
                        <text>密码必须是6-20个英文字母、数字（除空格），且字母、数字至少包含两种，区分大小写字母；</text>
                    </div>
                    <el-button type="primary" v-on:click="passwordUpdate()">确定修改</el-button>
                </div>
                <div v-else class="retrieve_password_from">
                    <img :src="complete">
                    <p>设置密码成功!</p>
                    <el-button style="background: #1560FA;" type="primary" v-on:click="jump('/login')">去登陆</el-button>
                </div>
            </div>
        </template>
        <template v-slot:container-bottom>
            <LoginBottonMessage :status="state.status">
                <template v-slot:message-p>
                    <p>{{ state.message }}</p>
                </template>
            </LoginBottonMessage>
        </template>
    </WindowContainer>
</template>

<script lang="ts" setup>
import { reactive, ref, onBeforeUnmount } from 'vue'
import { jump } from '@/renderer/tool/tool';
import WindowContainer from '@/renderer/views/login/components/WindowContainer.vue';
import LoginBottonMessage from '@/renderer/views/login/components/BottomMessage.vue';
import detail from '@/renderer/assets/images/detail.png'
import complete from '@/renderer/assets/images/complete.png'

const state = reactive({
    phone: '',
    code: '',
    newpassword1: '',
    newpassword2: '',
    show: true,
    message: "",
    status: false,
    labelPosition: 'right',
})

const passwordUpdate = () => {
    state.message = "修改成功"
    state.status = true
}


//验证码
const duration = ref(60)
const buttonName = ref("发送验证码")
const checkCodeClass = ref("")

let timer = null

const checkCode = () => {
    buttonName.value = duration.value + "s后重新发送"
    checkCodeClass.value = 'code_get_success'
    timers()
}

function timers() {
    if (!timer) {
        timer = setInterval(() => {
            buttonName.value = duration.value-- + "s后重新发送"
            console.log(duration.value);
            if (duration.value <= 0) {
                buttonName.value = '重新发送'
                clearInterval(timer);
                duration.value = 60;
            }
        }, 1000);
    }
}

onBeforeUnmount(() => {
    clearInterval(timer)
})

</script>


<style lang="scss" src="@/renderer/assets/scss/login/retrievePassword.scss" />
