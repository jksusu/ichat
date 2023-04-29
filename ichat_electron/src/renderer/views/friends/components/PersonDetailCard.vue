<template>
    <div id="person-deta-card-container">
        <div class="account_info">
            <div class="account">
                <div class="nickanme"><span>{{ props.nickname }}</span></div>
                <div class="userid"><span>{{ props.account }}</span></div>
            </div>
            <el-avatar shape="square" :size="60" :src="props.headPortraitUrl" />
        </div>
        <div>
            <slot name="bottom-button"></slot>
        </div>
    </div>
</template>
<script lang="ts" setup>
const props = withDefaults(defineProps<{
    account?: string,
    nickname?: string,
    headPortraitUrl?: string,
}>(), {
    account: "",
    nickname: "",
    headPortrait: "",
})
</script>

<style lang="scss" scoped>
#person-deta-card-container {
    min-width: 352px;
    min-height: 214px;


    .btn {
        width: 74px;
        height: 36px;
        background: #1560FA;
        border-radius: 4px 4px 4px 4px;
        text-align: center;
        line-height: 36px;
        margin: 166px 0 18px 244px;

        span {
            width: 42px;
            height: 20px;
            font-size: 14px;
            font-weight: 400;
            color: #FFFFFF;
            line-height: 16px;
        }
    }

    .account_info {
        display: flex;
        width: 292px;
        height: 60px;
        margin: 30px 0 0 30px;

        .userid {
            max-width: 94px;
            height: 20px;
            overflow: hidden;
            margin-top: 9px;
            font-size: 14px;
            font-weight: 400;
            color: #666666;
            line-height: 16px;
        }

        .nickanme {
            width: 216px;
            height: 24px;
            font-size: 18px;
            font-weight: 500;
            color: #333333;
            line-height: 24px;
            overflow: hidden;
        }

        img {
            margin-left: 16px;
            width: 60px;
            height: 60px;
            background: #E0E0E0;
            border-radius: 4px 4px 4px 4px;
        }
    }
}
</style>