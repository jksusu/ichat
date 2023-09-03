<template>
    <div class="title-case-container">
        <div>
            <information></information>
        </div>
        <session-title class="session-title"></session-title>
    </div>
</template>
<script setup>
import SessionTitle from '@/renderer/components/title/SessionTitle.vue';
import Information from '@/renderer/components/personal/Information.vue';
</script>

<style>
.title-case-container {
    display: flex;
}
.session-title {
    width: 100%;
}
</style>