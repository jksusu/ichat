<template>
    <div class="session_title">
        <div class="title_name">
            <label v-if="route == '/session'" @click="$emit('sessionTitleClickEvent')">{{ select.nickname }}</label>
        </div>
        <div class="menu">
            <RightTitle :show-max="true" />
            <div class="session_slot">
                <slot name="titleRight"></slot>
            </div>
        </div>
    </div>
</template>
  
<script lang="ts" setup>
import { computed } from "vue"
import RightTitle from './RightTitle.vue';
import { SessionStore } from '@/renderer/stores/modules/sessionList'
import { useSelectIndexStore } from '@/renderer/stores/modules/selectIndex'
const select = computed(() => SessionStore().getSelectSession)
const route = computed(() => useSelectIndexStore().getAsideMenuIndex)
</script>
  
<style lang="scss">
.session_title {
    display: flex;
    justify-content: space-between;
    min-width: 1080px;
    width: calc(100vw - 360px);
    height: 68px;
    background: #F5F6F7;
    border-bottom: 1px solid #EAEAEA;
    box-sizing: border-box;
    -webkit-app-region: drag;

    .menu {
        display: flex;
        flex-direction: column;

        .session_slot {
            margin-top: 7px;
            height: 17px;
            width: 132px;
            -webkit-app-region: no-drag
        }
    }
    .title_name {
        height: 22px;
        max-width: 150px;
        margin: 34px 0 12px 24px;
        font-size: 16px;
        font-family: PingFang SC-Medium, PingFang SC;
        font-weight: 500;
        color: #333333;
        line-height: 19px;

        label {
            -webkit-app-region: no-drag
        }
    }
}
</style>