<template>
    <div class="session_container">
        <LIstItems titlename="全部会话" @clickEvent="sessionItemSelect" :list="SessionStore().getSessionLists"></LIstItems>
        <div>
            <div v-if="selectSession">
                <SessionTitle :name="state.nickname" @sessionTitleClickEvent="open()">
                    <template #titleRight>
                        <div id="diandiandian">
                            <SvgIcon @click="open()" color="#808899" name="dian"></SvgIcon>
                        </div>
                    </template>
                </SessionTitle>
                <div>
                    <el-scrollbar id="chat_message_container" ref="scrollbarRef">
                        <div id="content_container" ref="innerRef">
                            <ChatBubble style="margin: 0 23px 0 23px;" v-for="item in record" :content="item.content"
                                :is-self="Boolean(item.i_send)" :send-time="formatWechatTime(item.send_time)"></ChatBubble>
                        </div>
                    </el-scrollbar>
                    <ChatInputContaineri @sendMessage="sendMessage"></ChatInputContaineri>
                </div>
            </div>
            <div v-else>
                <SessionTitle style="background:rgb(255, 255, 255);border-bottom:none"></SessionTitle>
                <div style="display: flex;justify-content: center;height: calc(100vh - 240px);">
                    <DefaultNullValue
                        style="display: flex;align-items: center;height: 100%;flex-direction: column;justify-content: center;">
                    </DefaultNullValue>
                </div>
            </div>
            <div>
                <Drawer @close="close" :status="state.drawerStatus" type="person"></Drawer>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import log from 'loglevel-es'
import { computed } from "@vue/reactivity"
import Drawer from '@/renderer/components/Drawer.vue'
import SvgIcon from '@/renderer/components/SvgIcon.vue'
import { ElScrollbar as ElScrollbarType } from 'element-plus'
import { SessionStore } from "@/renderer/stores/modules/sessionList"
import SessionTitle from "@/renderer/components/title/SessionTitle.vue"
import LIstItems from '@/renderer/components/conversation/LIstItems.vue'
import ChatBubble from '@/renderer/components/conversation/ChatBubble.vue'
import { reactive, ref, onMounted, nextTick, watch } from "vue"
import { useSelectIndexStore } from '@/renderer/stores/modules/selectIndex'
import { singleChatRecord } from '@/renderer/api/apis'
import DefaultNullValue from "@/renderer/components/DefaultNullValue.vue"
import { formatWechatTime } from "@/renderer/module/tool"
import { Socket } from "@/renderer/module/socket"
import ChatInputContaineri from "@/renderer/components/conversation/ChatInputContaineri.vue"
import { ChatRecordStore } from '@/renderer/stores/modules/chatRecord'
import { UserStore } from '@/renderer/stores/modules/user'
import { fr } from 'element-plus/es/locale'

//控制滚动条在最下面
onMounted(() => {
    updateScrollTop()
})

const state = reactive({
    drawerStatus: false,
    chatPageShow: false,
    nickname: '',
    page: 1,
    total: 0,
})

const selectSession = computed(() => SessionStore().getSelectSessionUID)
const index = computed(() => useSelectIndexStore().getSessionIndex)
// watch(index, (newQuestion, oldQuestion) => {
// })

const record = computed(() => ChatRecordStore().getChatById(index.value))

const open = () => { state.drawerStatus = true }
const close = () => { state.drawerStatus = false }

//item切换
const sessionItemSelect = (item: any) => {
    state.nickname = item.nickname
    useSelectIndexStore().setSessionIndex(item.uid)
    SessionStore().setSelectSession(item.account, item.nickname, item.headPortraitUrl, item.type)
    state.page = 1
    updateScrollTop()
}

const sendMessage = async (content: any) => {
    let from = UserStore().getUid
    let to = useSelectIndexStore().getSessionIndex

    const m = {
        "msgId": "",
        "content": content.content,
        "content_type": content.type,
        "from": from,
        "to": to,
        "i_send": true,
        "msg_type": "",//消息类型，如撤回
        "send_time": '',
        "status": 1,
    }
    //push到本地缓存中
    // ChatRecordStore().addMessage(to, m)
    // updateScrollTop()

    let { status, res } = await Socket().talkToUser(to, content.content, content.type)
    if (!status) {
        log.error(status, 'message send fail')
    }
    m.msgId = res.message.data.msgId ?? ''
    m.send_time = res.message.data.send_time ?? ''
    ChatRecordStore().addMessage(to, m)
    updateScrollTop()
}


//滚动条在最底下
const innerRef = ref<HTMLDivElement>()
const scrollbarRef = ref<InstanceType<typeof ElScrollbarType>>()

const updateScrollTop = () => {
    nextTick(() => {
        if (record.value) {
            scrollbarRef.value!.setScrollTop(innerRef.value!.clientHeight)
        }
    })
}

// const chatRecord = (to: string) => {
//     console.log("请求聊天记录kaishi ")
//     singleChatRecord({ to: to, page: state.page }).then((res) => {
//         if (res.code == 0) {
//             if (res.data.total > 0) {
//                 state.page++
//             }
//             state.total = res.data.total
//             //record = res.data.list
//         } else {

//         }
//     })
// }
</script>

<style lang="scss">
@keyframes spinner {
    to {
        transform: rotate(360deg);
    }
}


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

    #content_container {
        margin-top: 10px;
        margin-bottom: 10px;
    }

}
</style>
