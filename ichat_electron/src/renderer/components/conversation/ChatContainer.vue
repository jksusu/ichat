<template>
    <div id="chat_message_container">
        <el-scrollbar class="message_container">
            <ChatBubble style="margin: 0 23px 0 23px;" v-for="item in props.list" :content="item.content"
                :is-self="item.isSelf"></ChatBubble>
        </el-scrollbar>
</div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import ChatBubble from './ChatBubble.vue';
const props = withDefaults(defineProps<{ list?: any }>(), { list: [] })
</script>

<style lang="scss">
#chat_message_container {
    min-width: 1080px;
    min-height: 626px;
    height: calc(100vh - 274px);
    background: #F5F6F7;
    border-bottom: 1px solid #EAEAEA;
    box-sizing: border-box;

    .message_container {
        height: 100%;
        padding: 5px 0 5px 0;
    }
}
</style>