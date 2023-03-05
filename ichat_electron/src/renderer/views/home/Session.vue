<template>
    <div class="session_container">
        <LIstItems titlename="全部会话" @clickEvent="sessionItemSelect" :list="state.list"></LIstItems>
        <div>
            <SessionTitle :name="nickname" @clickEvent="open()">
                <template #titleRight>
                    <span id="diandiandian" @click="open()">
                        <SvgIcon color="#808899" name="dian"></SvgIcon>
                    </span>
                </template>
            </SessionTitle>
            <div>
                <el-scrollbar id="chat_message_container" ref="scrollbarRef">
                    <div ref="innerRef">
                        <ChatBubble style="margin: 0 23px 0 23px;" v-for="item in record" :content="item.content"
                            :is-self="item.isSelf"></ChatBubble>
                    </div>
                </el-scrollbar>
                <ChatInputContaineri @sendMessage="sendMessage"></ChatInputContaineri>
            </div>
            <div>
                <Drawer @close="close" :status="state.status" type="person"></Drawer>
            </div>
        </div>
</div>
</template>

<script lang="ts" setup>
import { computed } from "@vue/reactivity"
import Drawer from '@/renderer/components/Drawer.vue'
import SvgIcon from '@/renderer/components/SvgIcon.vue'
import { ElScrollbar as ElScrollbarType } from 'element-plus';
import { getSessionList, getChatRecord } from "@/renderer/api/chat"
import { reactive, ref, onMounted, onUpdated, nextTick } from "vue"
import SessionTitle from "@/renderer/components/title/SessionTitle.vue"
import LIstItems from '@/renderer/components/conversation/LIstItems.vue'
import ChatBubble from '@/renderer/components/conversation/ChatBubble.vue'
import { useSelectIndexStore } from '@/renderer/stores/modules/selectIndex'
import ChatInputContaineri from "@/renderer/components/conversation/ChatInputContaineri.vue"

//控制滚动条在最下面
onMounted(() => {
    updateScrollTop()
})

const state = reactive({
    status: false,
    list: [],
})

const record = ref([
    {
        'fromUid': '123',
        'targetUid': '好啊有',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '功能层面上，我们将通过提供语音的暂停和进度拖拽能力，并可视化音量，以满足语音接收者的使用效率需求。在体验层面上，语音作为用户的高频沟通操作，其设计必须满足QQ8.0中精致这一设计原则，给用户带来极致体验。',
        'dateTime': '123',
        'source': '123',
        'isSelf': true,
    }, {
        'fromUid': '123',
        'targetUid': '123',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '你今天吃饭了吗',
        'dateTime': '123',
        'source': '123',
        'isSelf': false,
    }, {
        'fromUid': '123',
        'targetUid': '好啊有',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '功能层面上，我们将通过提供语音的暂停和进度拖拽能力，并可视化音量，以满足语音接收者的使用效率需求。在体验层面上，语音作为用户的高频沟通操作，其设计必须满足QQ8.0中精致这一设计原则，给用户带来极致体验。',
        'dateTime': '123',
        'source': '123',
        'isSelf': true,
    }, {
        'fromUid': '123',
        'targetUid': '123',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '你今天吃饭了吗',
        'dateTime': '123',
        'source': '123',
        'isSelf': false,
    }, {
        'fromUid': '123',
        'targetUid': '好啊有',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '功能层面上，我们将通过提供语音的暂停和进度拖拽能力，并可视化音量，以满足语音接收者的使用效率需求。在体验层面上，语音作为用户的高频沟通操作，其设计必须满足QQ8.0中精致这一设计原则，给用户带来极致体验。',
        'dateTime': '123',
        'source': '123',
        'isSelf': true,
    }, {
        'fromUid': '123',
        'targetUid': '123',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '你今天吃饭了吗',
        'dateTime': '123',
        'source': '123',
        'isSelf': false,
    }, {
        'fromUid': '123',
        'targetUid': '好啊有',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '功能层面上，我们将通过提供语音的暂停和进度拖拽能力，并可视化音量，以满足语音接收者的使用效率需求。在体验层面上，语音作为用户的高频沟通操作，其设计必须满足QQ8.0中精致这一设计原则，给用户带来极致体验。',
        'dateTime': '123',
        'source': '123',
        'isSelf': true,
    }, {
        'fromUid': '123',
        'targetUid': '123',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '你今天吃饭了吗',
        'dateTime': '123',
        'source': '123',
        'isSelf': false,
    }, {
        'fromUid': '123',
        'targetUid': '好啊有',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '功能层面上，我们将通过提供语音的暂停和进度拖拽能力，并可视化音量，以满足语音接收者的使用效率需求。在体验层面上，语音作为用户的高频沟通操作，其设计必须满足QQ8.0中精致这一设计原则，给用户带来极致体验。',
        'dateTime': '123',
        'source': '123',
        'isSelf': true,
    }, {
        'fromUid': '123',
        'targetUid': '123',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '你今天吃饭了吗',
        'dateTime': '123',
        'source': '123',
        'isSelf': false,
    }, {
        'fromUid': '123',
        'targetUid': '好啊有',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '功能层面上，我们将通过提供语音的暂停和进度拖拽能力，并可视化音量，以满足语音接收者的使用效率需求。在体验层面上，语音作为用户的高频沟通操作，其设计必须满足QQ8.0中精致这一设计原则，给用户带来极致体验。',
        'dateTime': '123',
        'source': '123',
        'isSelf': true,
    }, {
        'fromUid': '123',
        'targetUid': '123',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '你今天吃饭了吗',
        'dateTime': '123',
        'source': '123',
        'isSelf': false,
    }, {
        'fromUid': '123',
        'targetUid': '好啊有',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '功能层面上，我们将通过提供语音的暂停和进度拖拽能力，并可视化音量，以满足语音接收者的使用效率需求。在体验层面上，语音作为用户的高频沟通操作，其设计必须满足QQ8.0中精致这一设计原则，给用户带来极致体验。',
        'dateTime': '123',
        'source': '123',
        'isSelf': true,
    }, {
        'fromUid': '123',
        'targetUid': '123',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '你今天吃饭了吗',
        'dateTime': '123',
        'source': '123',
        'isSelf': false,
    }, {
        'fromUid': '123',
        'targetUid': '好啊有',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '功能层面上，我们将通过提供语音的暂停和进度拖拽能力，并可视化音量，以满足语音接收者的使用效率需求。在体验层面上，语音作为用户的高频沟通操作，其设计必须满足QQ8.0中精致这一设计原则，给用户带来极致体验。',
        'dateTime': '123',
        'source': '123',
        'isSelf': true,
    }, {
        'fromUid': '123',
        'targetUid': '123',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '你今天吃饭了吗',
        'dateTime': '123',
        'source': '123',
        'isSelf': false,
    }, {
        'fromUid': '123',
        'targetUid': '好啊有',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '功能层面上，我们将通过提供语音的暂停和进度拖拽能力，并可视化音量，以满足语音接收者的使用效率需求。在体验层面上，语音作为用户的高频沟通操作，其设计必须满足QQ8.0中精致这一设计原则，给用户带来极致体验。',
        'dateTime': '123',
        'source': '123',
        'isSelf': true,
    }, {
        'fromUid': '123',
        'targetUid': '123',
        'targetType': '123',//接收者类型
        'msgType': '1',//消息内容类型
        'msgUID': '123',
        'content': '你今天吃饭了吗',
        'dateTime': '123',
        'source': '123',
        'isSelf': false,
    },
])

const nickname = ref('')

getSessionList({ 'username': '' }).then((res) => {
    console.log('sessionlist')
    if (res.data.code == 200) {
        state.list = res.data.data
        init()
    }
})

const index = computed(() => useSelectIndexStore().getSessionIndex)

//初始化
const init = () => {
    state.list.filter((item) => {
        if (item.uid == index.value) {
            nickname.value = item.nickname
            //加载聊天记录
            sessionItemSelect({ uid: item.uid, nickname: item.nickname })
        }
    })
}

const open = () => { state.status = true }
const close = () => { state.status = false }

const sessionItemSelect = (item: any) => {
    nickname.value = item.nickname
    useSelectIndexStore().setSessionIndex(item.uid)
    getChatRecord({ 'uid': item.uid }).then((res) => {
        if (res.data.code == 200) {
            record.value = res.data.data
        }
    })
}

const sendMessage = (content: any) => {
    record.value.push(content)
    updateScrollTop()
}

//滚动条在最底下
const innerRef = ref<HTMLDivElement>()
const scrollbarRef = ref<InstanceType<typeof ElScrollbarType>>()
const updateScrollTop = () => {
    nextTick(() => {
        scrollbarRef.value!.setScrollTop(innerRef.value!.clientHeight)
        // if (innerRef.value!.clientHeight > 200) {
        // }
    })
}

</script>

<style lang="scss">
.session_container {
    display: flex;

    #diandiandian {
        position: absolute;
        margin-left: 100px;
    }
}

#chat_message_container {
    min-width: 1080px;
    min-height: 626px;
    height: calc(100vh - 274px);
    background: #F5F6F7;
    border-bottom: 1px solid #EAEAEA;
    box-sizing: border-box;

    .message_container {
        height: 100%;
        padding: 5px 0 5px 0;
    }
}
</style>
