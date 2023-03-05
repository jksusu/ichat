<template>
    <div class="message_trip_container">
        <span class="title">新消息({{ countUnreadMessageNumber }})</span>
        <div class="items_container" v-for="item in list">
            <el-avatar class="head_image" shape="square" :size="32" :src="item.headPortraitUrl" />
            <div class="info">
                <span class="username">{{ item.nickname }}</span>
                <el-badge style="margin-left: 10px;margin-top: 4px;" :value="item.unreadMessageNumber" :max="99" />
            </div>
        </div>
        <hr />
        <span style="float: right;">忽略全部</span>
    </div>
</template>
<script lang="ts" setup>

const countUnreadMessageNumber = 20

const list = [
    {
        uid: '12345',
        nickname: '芝士蛋糕',
        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
        unreadMessageNumber: 10,
    },
    {
        uid: '12345',
        nickname: '芝士蛋糕',
        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
        unreadMessageNumber: 10,
    },
]

</script>
<style lang="scss" scoped>
.message_trip_container {
    width: 220px;
    height: 115px;
    background: #333333;

    .title {
        margin-top: 5px;
        margin-left: 5px;
    }

    .items_container {
        margin-top: 5px;
        width: 222px;
        height: 40px;
        display: flex;
        align-items: center;

        .info {
            display: flex;
            justify-content: space-between;

            .username {
                max-width: 160px;
                font-size: 14px;
                font-weight: 500;
                color: #333333;
                line-height: 24px;
            }
        }
    }
}
</style>