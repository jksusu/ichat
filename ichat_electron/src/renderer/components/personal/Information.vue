<template>
    <div id="infomation_container">
        <div class="name">
            {{ "Hi!"+ props.name }}
        </div>
        <div class="autograph">
            {{ props.autograph }}
        </div>
    </div>
    <div class="slot">
        <slot></slot>
    </div>
</template>

<script lang="ts" setup>
const props = withDefaults(defineProps<{
    name?: string,
    autograph?: string,
}>(), {
    name: 'im',
    autograph: '祝你有一个愉快的一天!'
})
</script>
<style lang="scss">
.slot {
    height: 44px;
    width: 296px;
    background: #FFFFFF;
}

#infomation_container {
    height: 94px;
    width: 296px;
    background: #FFFFFF;
    -webkit-app-region: drag;

    div {
        margin-left: 18px;
    }

    .name {
        padding-top: 30px;
        max-width: 100px;
        height: 25px;
        font-size: 18px;
        font-weight: 600;
        color: #333333;
        line-height: 21px;
    }

    .autograph {
        margin-top: 34px;
        max-width: 260px;
        height: 17px;
        font-size: 12px;
        font-weight: 400;
        color: #666666;
        line-height: 14px;
    }
}
</style>