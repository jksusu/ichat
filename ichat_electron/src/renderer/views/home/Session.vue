<template>
    <div class="session_container">
        <LIstItems titlename="全部会话" @clickEvent="sessionItemSelect" :list="SessionStore().getSessionLists"></LIstItems>
        <div>
            <div v-if="selectSession">
                <SessionTitle :name="nickname" @sessionTitleClickEvent="open()">
                    <template #titleRight>
                        <div id="diandiandian">
                            <SvgIcon @click="open()" color="#808899" name="dian"></SvgIcon>
                        </div>
                    </template>
                </SessionTitle>
                <div>
                    <el-scrollbar id="chat_message_container" ref="scrollbarRef">
                        <div ref="innerRef">
                            <ChatBubble style="margin: 0 23px 0 23px;" v-for="item in record" :content="item.content"
                                :is-self="item.i_send" :send-time="formatWechatTime(item.send_time)"></ChatBubble>

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
import ChatInputContaineri from "@/renderer/components/conversation/ChatInputContaineri.vue"
import { singleChatRecord } from '@/renderer/api/apis'
import DefaultNullValue from "@/renderer/components/DefaultNullValue.vue"
import { formatWechatTime } from "@/renderer/module/tool"

//控制滚动条在最下面
onMounted(() => {
    updateScrollTop()
    Init()
})

const state = reactive({
    drawerStatus: false,
    chatPageShow: false,
    page: 1,
    total: 0,
})

const selectSession = computed(() => SessionStore().getSelectSessionUID)
const index = computed(() => useSelectIndexStore().getSessionIndex)
// watch(index, (newQuestion, oldQuestion) => {
// })


const Init = async () => {
    console.log('开始同步会话列表')

    // let resp = await Connect().sessionList
    // if (resp.success) {
    //     console.log(resp.message)
    // }
}

const record = ref([])
// {
//         'fromUid': '123',
//         'targetUid': '123',
//         'targetType': '123',//接收者类型
//         'msgType': '1',//消息内容类型
//         'msgUID': '123',
//         'content': '你今天吃饭了吗',
//         'dateTime': '123',
//         'source': '123',
//         'isSelf': false,
//     }
const nickname = ref('')
// getSessionList({ 'username': '' }).then((res) => {
//     console.log('sessionlist')
//     if (res.data.code == 200) {
//         state.list = res.data.data
//         init()
//     }
// })



const open = () => {
    console.log("详情开关")
    state.drawerStatus = true
}
const close = () => { state.drawerStatus = false }

//item切换
const sessionItemSelect = (item: any) => {
    nickname.value = item.nickname
    useSelectIndexStore().setSessionIndex(item.uid)
    SessionStore().setSelectSession(item.account, item.nickname, item.headPortraitUrl, item.type)
    chatRecord(item.account)
    state.page = 1
    //更新子组件数据

    // getChatRecord({ 'uid': item.uid }).then((res) => {<span id="diandiandian">
    //     if (res.data.code == 200) {
    //         record.value = res.data.data
    //     }
    // })
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


const chatRecord = (to: string) => {
    console.log("请求聊天记录kaishi ")
    singleChatRecord({ to: to, page: state.page }).then((res) => {
        if (res.code == 0) {
            if (res.data.total > 0) {
                state.page++
            }
            state.total = res.data.total
            record.value = res.data.list
        } else {

        }
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
