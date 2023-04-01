<template>
    <div class="items_container" :class="state.selectStyle" @click="$emit('clickEvent')" v-on:click="selectEvent(props.account)">
        <img :height="headPortraitSize" :width="headPortraitSize"
            style="margin: 10px 0 10px 10px;border-radius: 4px 4px 4px 4px;" :src="state.headPortraitUrl" />
        <div class="info">
            <div class="top">
                <span class="username">{{ state.username }}</span>
                <span class="last_time is_hiddle" :class="state.isHiddle">{{ state.lastTime }}</span>
            </div>
            <div class="bottom" :class="state.isHiddle">
                <span class="last_message">{{ state.lastMessage }}</span>
                <div class="message_number" :class="state.isHiddle"><sup style="--num:100" :style="state.messageNumber">{{
                    state.messageNumber
                }}</sup></div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { reactive, toRefs } from 'vue'

const props = withDefaults(defineProps<{
    isTopping?: boolean,
    headPortraitUrl?: string,
    account?: string,
    username?: string,
    isNameCenter?: boolean,
    messageNumber?: number,
    lastMessage?: string,
    lastTime?: string,
    isSelect?: boolean,
    headPortraitSize?: number,
}>(), {
    isTopping: false,
    headPortrait: "",
    account: "",
    username: "",
    isNameCenter: false,
    messageNumber: 0,
    lastMessage: "",
    lastTime: "",
    isSelect: false,
    headPortraitSize: 42,
})

const state = reactive({
    username: '',
    lastTime: '',
    lastMessage: "",
    messageNumber: "",
    selectStyle: "",
    selctIndex:"",
    isHiddle: "",
    headPortraitSize: 0,
    headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
})

//复制
state.username = props.username
state.lastTime = props.lastTime
state.lastMessage = props.lastMessage
state.headPortraitSize = props.headPortraitSize

//姓名居中
if (props.isNameCenter) {
    state.isHiddle = 'hiddle'
}

if (props.isSelect) {
    state.selectStyle = "select"
}
if (props.messageNumber > 0) {
    state.messageNumber = "--num:" + props.messageNumber
}

const selectEvent = (index)=> {
    state.selctIndex = index
    state.selectStyle = 'select'
}

</script>

<style lang="scss" scoped>
// :root {
//     --num: 0
// }

.select {
    background-color: #EAEAEA;
}

.friends_select_style {
    background-color: #EAEAEA;

}

.hiddle {
    display: none !important;
}

.items_container:hover {
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
                font-family: PingFang SC-Medium, PingFang SC;
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

            .message_number {
                margin-left: 23px;
                height: 14px;
                max-width: 27px;
                margin-right: 30px;
                background: #F43E24;
                border-radius: 10px 10px 10px 10px;
            }
        }
    }
}

sup {
    position: absolute;
    box-sizing: border-box;
    min-width: 1rem;
    padding: 0 0.1875rem;
    color: #fff;
    font-size: min(.75rem, calc(10000px - var(--num) * 100px));
    line-height: 1.2;
    text-align: center;
    background-color: #eb4646;
    border-radius: 1rem;
    opacity: var(--num);
}

sup::before {
    content: '';
    font-size: min(.75rem, calc(var(--num) * 100px - 9900px));
}
</style>