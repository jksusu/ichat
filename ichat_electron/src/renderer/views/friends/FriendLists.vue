<template>
    <div>
        <SecondLevelContainer :name="titleName">
            <template #list>
                <el-scrollbar max-height="823px">
                    <div class="friends_container">
                        <div v-for="item in friendList" @click="select(item)">
                            <Item :class="{ select: item.friendUid == state.indexUid }"
                                style="height: 62px; width: 356px; margin-top: 5px;" :nickname="item.nickname"
                                :headPortraitUrl="item.headPortraitUrl" :account="item.friendUid" :pattern="pattern"
                                v-bind="$attrs" :data="item" v-menus:right="friend_menus">
                            </Item>
                        </div>
                    </div>
                </el-scrollbar>
            </template>
            <template #detail>
                <div class="detail_container">
                    <div class="card_container" v-if="state.detailContainerShow">
                        <PersonDetailCard @createSession="createSession" :nickname="state.nickname"
                            :account="state.indexUid" :headPortraitUrl="state.headPortraitUrl">
                            <template v-slot:bottom-button>
                                <div class="buttom">
                                    <div class="btn" @click="createSession(state)"><span>发消息</span></div>
                                </div>
                            </template>
                        </PersonDetailCard>
                    </div>
                    <DefaultNullValue v-else></DefaultNullValue>
                </div>
            </template>
        </SecondLevelContainer>
    </div>
</template>
<script lang="ts" setup>
import { onMounted, reactive, ref } from "vue"
import Item from '@/renderer/components/conversation/Item.vue'
import DefaultNullValue from "@/renderer/components/DefaultNullValue.vue"
import PersonDetailCard from "./components/PersonDetailCard.vue"
import SecondLevelContainer from "@/renderer/views/layout/SecondLevelContainer.vue"
import { contactsList } from '@/renderer/api/apis'
import { sendMessage, setRemarks, delFriend } from "@/renderer/module/menu"
import { CreateSession } from "@/renderer/module/sessionEffect"

//右键
const friend_menus = ref([sendMessage, setRemarks, delFriend])

//const indexStore = useSelectIndexStore()
//const index = computed(() => indexStore.getFriendsIndex)
const select = (item: any) => {
    // if (item.friendUid == state.indexUid) {
    //     return
    // }
    console.log(item)
    //indexStore.setFriendsIndex(item.friendUid)
    setSelectUser(item)
}
const titleName = '联系人';
const friendList = ref([])
const pattern = 'show'
const state = reactive({
    nickname: '',
    indexUid: '',//当前选择uid
    headPortraitUrl: '',
    detailContainerShow: false,
    isInfiniteScroll: false,
    page: 1,//
    pageSize: 0,//
    total: 0,//总数
})

onMounted(() => {
    ApiContactsList(1)//获取第一页数据
})

const setSelectUser = (list: any) => {
    state.nickname = list.nickname
    state.indexUid = list.friendUid
    state.headPortraitUrl = list.headPortraitUrl
    state.detailContainerShow = true
    console.log(state)
}

/**
 * 获取联系人列表
 * @param page 
 */
const ApiContactsList = (page: number) => {
    state.isInfiniteScroll = true
    contactsList({ "page": state.page }).then((res) => {
        if (res.code == 0) {
            state.page++
            state.total = res.data.total
            state.page = res.data.page
            state.pageSize = res.data.pageSize
            res.data.list.forEach(e => {
                friendList.value.push({
                    'friendUid': e.uid,
                    'nickname': e.nickname,
                    'headPortraitUrl': e.head_portrait,
                    'status': e.status,
                })
            })
            state.isInfiniteScroll = false
        }
    })
}

//创建会话
const createSession = (user: any) => {
    CreateSession(user.indexUid, user.nickname, user.headPortraitUrl)
}
</script>
<style lang="scss" scoped>
.friends_container {
    margin: 8px 7px 0 8px;

    .items_container:not(:first-child) {
        margin-top: 4px;
    }
}

.detail_container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;

    .card_container {
        position: relative;
        width: 352px;
        height: 310px;
        background: #FFFFFF;
        box-shadow: 0px 8px 30px 1px rgba(0, 0, 0, 0.08);
        border-radius: 8px 8px 8px 8px;
    }
}

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
</style>