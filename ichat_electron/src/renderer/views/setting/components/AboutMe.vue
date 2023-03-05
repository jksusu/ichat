<template>
    <div>
        <el-dialog id="about-me-container" v-model="dialogTableVisible">
            <div class="logo">
                <img :src="logoUrl">
            </div>
            <div class="title-name">
                <span class="title">im即时通讯</span>
                <span class="entitle">im jishi tongxun</span>
            </div>
            <div class="version">
                <span>当前版本：v1.1</span>
            </div>

            <div class="button" :class="{ yes: newVersion > thisVersion }" @click="upgrade"><span>{{ buttonName }}</span>
            </div>
        </el-dialog>
</div>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import logoUrl from "@/renderer/assets/images/logo.png"

const buttonName = ref("当前已经是最新版本")
const dialogTableVisible = ref(true)
const newVersion = 1.2
const thisVersion = 1.2

if (newVersion > thisVersion) {
    buttonName.value = '马上升级'
}

const upgrade = () => {
    if (buttonName.value == '马上升级') {
        console.log('开始升级')
    }
}

</script>

<style lang="scss">
#setting-page-container {
    .session_title {
        width: 100%;
    }
}

#about-me-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 360px;
    height: 453px;
    background: #FFFFFF;
    box-shadow: 0px 8px 30px 1px rgba(0, 0, 0, 0.16);
    border-radius: 12px 12px 12px 12px;


    button:focus {
        outline: none;
        outline: none;
    }


    .logo {
        display: flex;
        justify-content: center;
    }

    .title-name {
        margin-top: 26px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;

        .title {
            height: 36px;
            font-size: 26px;
            font-weight: 600;
            color: #333333;
            line-height: 30px;
        }

        .entitle {
            height: 17px;
            font-size: 12px;
            font-weight: 400;
            color: #999999;
            line-height: 14px;
            letter-spacing: 1px;
        }
    }


    .version {
        text-align: center;
        margin-top: 26px;
        height: 20px;
        font-size: 14px;
        font-weight: 400;
        color: #999999;
        line-height: 16px;
    }

    .button {
        margin-top: 38px;
        width: 300px;
        height: 48px;
        background: #729FFC;
        border-radius: 8px 8px 8px 8px;
        display: flex;
        justify-content: center;
        align-items: center;

        span {
            margin: auto;
            height: 22px;
            font-size: 16px;
            font-weight: 500;
            color: #FFFFFF;
        }
    }

    .yes {
        background: #1560FA !important;
    }
}
</style>