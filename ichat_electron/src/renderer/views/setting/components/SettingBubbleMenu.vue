<template>
    <div class="setting_bubble_container">
        <div class="btn font" style="color:#FFFFFF;margin-top: 6px;margin-bottom: 6px;">声音设置</div>
        <div class="btn font" style="color:#FFFFFF;margin-top: 6px;margin-bottom: 6px;">关于我们</div>
        <div class="btn" style="color:#FF6161;">退出登录</div>
    </div>
</template>
<script lang="ts" setup>
</script>

<style lang="scss">
.setting_bubble_container {
    width: 84px;
    height: 111px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    .btn {
        text-align: center;
        width: 72px;
        height: 29px;
        font-size: 12px;
        font-family: PingFang HK-Regular, PingFang HK;
        font-weight: 400;
        line-height: 29px;
    }

    .btn:hover {
        background: #333333;
        border-radius: 4px 4px 4px 4px;
    }
}
</style>