<template>
    <div class="items_container" @click="$emit('clickEvent', props)" :uid="props.account">
        <el-avatar class="head_image" shape="square" :size="42" :src="props.headPortraitUrl" />
        <div class="info">
            <div class="top">
                <span class="username">{{ props.nickname }}</span>
                <span class="last_time">{{ props.lastMessageTime }}</span>
            </div>
            <div class="bottom" v-if="props.pattern == 'message'">
                <span class="last_message">{{ props.lastMessage }}</span>
                <el-badge style="margin-left: 20px;" v-if="props.unreadMessageNumber > 0"
                    :value="props.unreadMessageNumber" :max="99" />
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
/**
 * pattern 模式
 * 1.message 消息模式 会完全展示
 * 2.show 展示模式 昵称居中显示 其他配置项不显示
 */
const props = withDefaults(defineProps<{
    pattern?: string,
    topping?: boolean,
    headPortraitUrl?: string,
    account?: string,
    nickname?: string,
    lastMessage?: string,
    lastMessageTime?: string,
    unreadMessageNumber?: number,
}>(), {
    pattern: 'message',
    topping: false,
    headPortraitUrl: "https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg",
    account: "",
    nickname: "",
    lastMessage: "",
    lastMessageTime: "",
    unreadMessageNumber: 0,
})

</script>

<style lang="scss">
.el-badge__content {
    --el-color-danger: #F43E24;
    --el-bg-color: none;
    --el-badge-font-size: 8px;
}

.head_image {
    margin: 10px 0 10px 10px;
    border-radius: 4px 4px 4px 4px;
}

.items_container:hover {
    //background-color: #D9D8D8;
    background-color: #EAEAEA;
}

.items_container {
    display: flex;
    align-items: center;
    width: 280px;
    height: 62px;
    border-radius: 8px 8px 8px 8px;

    .info {
        width: 199px;
        margin-left: 6px;

        .top {
            height: 20px;
            line-height: 18px;

            .username {
                max-width: 160px;
                font-size: 14px;
                font-weight: 500;
                color: #333333;
                line-height: 16px;
            }

            .last_time {
                max-width: 70px;
                float: right;
                height: 17px;
                font-size: 12px;
                font-weight: 400;
                color: #999999;
                line-height: 17px;
            }
        }

        .bottom {
            display: flex;
            height: 17px;
            margin-top: 9px;

            .last_message {
                width: 166px;
                height: 17px;
                font-size: 12px;
                font-weight: 400;
                color: #999999;
                line-height: 14px;
            }
        }
    }
}
</style>