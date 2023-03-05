<template>
    <div id="login-window-container" class="window-container">
        <div class="header_title">
            <LeftTitle></LeftTitle>
            <RightTitle :border-radius="true" :hover-color="true"></RightTitle>
        </div>
        <slot name="container-content"></slot>
    </div>
    <slot name="container-bottom"></slot>
</template>
<script lang="ts" setup>
import LeftTitle from '@/renderer/components/title/LeftTitle.vue';
import RightTitle from '@/renderer/components/title/RightTitle.vue';
</script>


<style lang="scss" src="@/renderer/assets/scss/login/loginContainer.scss" />