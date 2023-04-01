<template>
    <div id="person-deta-card-container">
        <div class="account_info">
            <div class="account">
                <div class="nickanme"><span>{{ props.nickname }}</span></div>
                <div class="userid"><span>{{ props.account }}</span></div>
            </div>
            <img width="60" height="60" :src="props.headPortraitUrl">
        </div>
        <div class="slot">
            <slot name="remarks"></slot>
        </div>

        <div class="buttom">
            <div class="btn"><span>发消息</span></div>
        </div>
        <!-- <div class="sendmessage_icon">
            <svg t="1672919704918" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
                p-id="17436" width="22" height="22">
                <path
                    d="M910.222222 113.777778a56.888889 56.888889 0 0 1 56.888889 56.888889v568.888889a56.888889 56.888889 0 0 1-56.888889 56.888888h-260.892444l-16.668445 16.497778-80.440889 80.782222a57.287111 57.287111 0 0 1-80.440888 0l-80.440889-80.782222-16.668445-16.497778H113.777778a56.888889 56.888889 0 0 1-56.888889-56.888888V170.666667a56.888889 56.888889 0 0 1 56.888889-56.888889h796.444444m0-56.888889H113.777778a113.777778 113.777778 0 0 0-113.777778 113.777778v568.888889a113.777778 113.777778 0 0 0 113.777778 113.777777h237.283555l80.440889 80.213334a113.208889 113.208889 0 0 0 160.881778 0L672.824889 853.333333H910.222222a113.777778 113.777778 0 0 0 113.777778-113.777777V170.666667a113.777778 113.777778 0 0 0-113.777778-113.777778zM341.333333 455.111111a56.888889 56.888889 0 1 1-56.888889-56.888889 56.888889 56.888889 0 0 1 56.888889 56.888889z m227.555556 0a56.888889 56.888889 0 1 1-56.888889-56.888889 56.888889 56.888889 0 0 1 56.888889 56.888889z m227.555555 0a56.888889 56.888889 0 1 1-56.888888-56.888889 56.888889 56.888889 0 0 1 56.888888 56.888889z"
                    fill="#666666" p-id="17437"></path>
            </svg>
        </div> -->
        <!-- <div class="addfriend_icon">
            <svg t="1672919220592" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
                p-id="10239" width="24" height="24">
                <path
                    d="M512 464c88 0 160-72 160-160S600 144 512 144s-160 72-160 160 72 160 160 160z m224 288v-96c0-17.6 14.4-32 32-32s32 14.4 32 32v96h96c17.6 0 32 14.4 32 32s-14.4 32-32 32h-96v96c0 17.6-14.4 32-32 32s-32-14.4-32-32v-96h-96c-17.6 0-32-14.4-32-32s14.4-32 32-32h96z m-16-244.8c9.6 6.4 16 16 16 27.2 0 17.6-14.4 32-32 32-6.4 0-11.2-1.6-16-4.8-43.2-22.4-91.2-33.6-144-33.6h-64c-176 0-320 144-320 320 0 17.6 14.4 32 32 32h384c17.6 0 32 14.4 32 32s-14.4 32-32 32H192c-52.8 0-96-43.2-96-96 0-174.4 116.8-321.6 276.8-368-51.2-41.6-84.8-104-84.8-176 0-123.2 100.8-224 224-224s224 100.8 224 224c0 70.4-33.6 134.4-84.8 176 24 6.4 48 16 68.8 27.2z"
                    fill="#666666" p-id="10240"></path>
            </svg>
        </div> -->
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
            margin: 18px 0 18px 244px;

            span {
                width: 42px;
                height: 20px;
                font-size: 14px;
                font-weight: 400;
                color: #FFFFFF;
                line-height: 16px;
            }
        }

    .slot {
        width: 100%;
        height: 124px;
        border-radius: 0 0 8px 8px;
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