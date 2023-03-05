<template>
    <div class="right_title">
        <div class="windons-btn btn" v-on:click="sendIpc('app-min')" :class="hoverColorStyle">
            <svg width="16" height="16" t="1672237314693" class="icon" viewBox="0 0 1024 1024" version="1.1"
                xmlns="http://www.w3.org/2000/svg" p-id="3108">
                <path
                    d="M213.333333 469.333333m21.333334 0l554.666666 0q21.333333 0 21.333334 21.333334l0 42.666666q0 21.333333-21.333334 21.333334l-554.666666 0q-21.333333 0-21.333334-21.333334l0-42.666666q0-21.333333 21.333334-21.333334Z"
                    fill="rgba(0,0,0,0.9);" p-id="3109"></path>
            </svg>
        </div>
        <div v-if="showMax" class="windons-btn btn" v-on:click="sendIpc('app-max')" :class="hoverColorStyle">
            <svg width="16" height="16" t="1672632990570" class="icon" viewBox="0 0 1024 1024" version="1.1"
                xmlns="http://www.w3.org/2000/svg" p-id="11479">
                <path
                    d="M768 170.666667H256c-46.933333 0-85.333333 38.4-85.333333 85.333333v512c0 46.933333 38.4 85.333333 85.333333 85.333333h512c46.933333 0 85.333333-38.4 85.333333-85.333333V256c0-46.933333-38.4-85.333333-85.333333-85.333333zM256 768V256h512v512H256z"
                    fill="#2c2c2c" p-id="11480"></path>
            </svg>
        </div>

        <div class="windons-btn close" v-on:click="sendIpc('app-quit')" :class="borderRadiusClass">
            <svg width="16" height="16" t="1672236963007" class="icon" viewBox="0 0 1024 1024" version="1.1"
                xmlns="http://www.w3.org/2000/svg" p-id="2016">
                <path
                    d="M780.501333 183.168a42.666667 42.666667 0 0 1 63.872 56.32l-3.541333 4.010667-597.333333 597.333333a42.666667 42.666667 0 0 1-63.872-56.32l3.541333-4.010667 597.333333-597.333333z"
                    fill="rgba(0,0,0,0.9);" p-id="2017"></path>
                <path
                    d="M183.168 183.168a42.666667 42.666667 0 0 1 56.32-3.541333l4.010667 3.541333 597.333333 597.333333a42.666667 42.666667 0 0 1-56.32 63.872l-4.010667-3.541333-597.333333-597.333333a42.666667 42.666667 0 0 1 0-60.330667z"
                    fill="rgba(0,0,0,0.9);" p-id="2018"></path>
            </svg>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { ref, toRef } from "vue"
import { sendIpc } from '@/common/ipcRenderer';

const props = withDefaults(defineProps<{ showMax?: boolean, borderRadius?: boolean, hoverColor?: boolean }>(), { showMax: false, borderRadius: false, hoverColor: false })
const showMax = props.showMax//是否显示放大
const borderRadius = props.borderRadius//是否圆角
const hoverColor = props.hoverColor


const borderRadiusClass = ref("")
if (borderRadius) {
    borderRadiusClass.value = 'borderRadiusClass'
}

//选中颜色切换
const hoverColorStyle = ref("")
if (hoverColor) {
    hoverColorStyle.value = 'newHover'
}

</script>
<style lang="scss" >
:root {
    --hoverTitleolor: #EAEAEA
}

.borderRadiusClass {
    border-radius: 0 12px 0 0
}

.newHover {
    --hoverTitleolor: #F5F6F7;
}

.right_title {
    display: flex;
    float: right;
    height: 32px;
    -webkit-app-region: none;
    
    .windons-btn {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 46px;
        height: 32px;
    }

    .btn:hover {
        background-color: var(--hoverTitleolor);
    }

    .close:hover {
        background-color: rgba(244, 62, 36, 0.8);

        svg {
            fill: rgb(255, 255, 255);
        }
    }
}
</style>