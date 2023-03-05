<template>
    <div class="default">
        <img :src="src" />
        <span>{{ message }}</span>
    </div>
</template>

<script lang="ts" setup>
import contentEmpty from '@/renderer/assets/images/content_empty.png'

const props = withDefaults(defineProps<{ message?: string }>(), { message: "暂无内容" })

const list = [
    {
        src: contentEmpty,
        message: '暂无内容',
    },
    {
        src: contentEmpty,
        message: '暂无网络',
    },
    {
        src: contentEmpty,
        message: '暂无对话内容',
    }
]

const index = props.message
let src = ''
let message = ''
list.forEach(element => {
    if (index == element.message) {
        src = element.src
        message = element.message
    }
});

</script>

<style lang="scss">
.default {
    width: 200px;
    height: 230px;
    text-align: center;

    span {
        width: 72px;
        height: 17px;
        font-size: 12px;
        font-weight: 400;
        color: #999999;
        line-height: 14px;
    }
}
</style>