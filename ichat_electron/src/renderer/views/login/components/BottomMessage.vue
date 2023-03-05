<template>
    <div v-if="!props.status" class="message" style="visibility: visible;"></div>
    <div v-else class="message message_error">
        <slot name="message-p"></slot>
    </div>
</template>
<script lang="ts" setup>
const props = defineProps<{ status: Boolean }>()
</script>
<style lang="scss">

.message {
    -webkit-app-region: drag;
    height: var(--windonBottomMessageHeight);
    background-color: var(--backgroundColor);
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 0 0 12px 12px;
    visibility: hidden;
}

.message_error {
    background-color: #FEF0EE;
    visibility: visible;
    p {
        color: #F43E24;
        font-weight: 400;
        font-weight: 400;
        line-height: 14px;
        font-size: 12px;
        user-select:none;
    }
}
</style>