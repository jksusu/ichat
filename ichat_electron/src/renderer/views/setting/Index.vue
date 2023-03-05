<template>
    <div id="setting-page-container">
        <SessionTitle></SessionTitle>
        <button @click="showAboutMe =true">显示</button>
        <AboutMe v-if="showAboutMe"></AboutMe>
</div>
</template>
<script lang="ts" setup>
import SessionTitle from '@/renderer/components/title/SessionTitle.vue';
import { ref } from 'vue';
import AboutMe from './components/AboutMe.vue';

const showAboutMe = ref(false)


</script>

<style lang="scss">
#setting-page-container {
    .session_title {
        width: 100%;
    }
}
</style>