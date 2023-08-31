<template>
    <div id="login-window-container" class="window-container">
        <div class="header_title">
            <div></div>
            <RightTitle v-if="ifWin()" :border-radius="true" :hover-color="true"></RightTitle>
        </div>
        <slot name="container-content"></slot>
    </div>
    <slot name="container-bottom"></slot>
</template>
<style lang="scss" src="@/renderer/assets/scss/login/loginContainer.scss" />
<script lang="ts" setup>
import { ifWin } from '@/renderer/tool/tool';
import RightTitle from '@/renderer/components/title/RightTitle.vue';
</script>