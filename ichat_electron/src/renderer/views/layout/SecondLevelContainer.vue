<template>
    <SessionTitle :name="name"></SessionTitle>
    <div class="second_leavel_container">
        <div class="left_item">
            <slot name="list"></slot>
        </div>
        <div class="right_item">
            <slot name="detail"></slot>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { computed } from "@vue/reactivity";
import SessionTitle from '@/renderer/components/title/SessionTitle.vue';
import { useSelectIndexStore } from '@/renderer/stores/modules/selectIndex'

const props = withDefaults(defineProps<{ name?: string }>(), { name: "" })
const name = props.name

const indexStore = useSelectIndexStore()
const index = computed(() => indexStore.getMaillistIndex)

</script>
<style lang="scss" scoped>
.second_leavel_container {
    display: flex;
    //min-width: 1080px;
    //min-height: 832px;
    height: calc(100vh - 68px);
    background: #FFFFFF;
    border-left: 1px solid #EAEAEA;
    box-sizing: border-box;

    .left_item {
        //width: 371px;
        background: #FFFFFF;
    }

    .right_item {
        //width: calc(100vw - 732px);
        background: #F5F6F7;
    }
}
</style>