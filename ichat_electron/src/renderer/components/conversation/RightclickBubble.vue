<template>
    <div class="container">

    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, toRefs } from 'vue'

const props = withDefaults(defineProps<{
    featureList: string,
}>(), {
    featureList: "",
})

const featureList = { "withdraw": "撤回", "delete": "删除", "quote": "引用", "copy": "复制", "paste": "粘贴", "topping": "置顶" }

const state = reactive({
    featureList: '',
})
</script>

<style lang="scss">
.container {}
</style>