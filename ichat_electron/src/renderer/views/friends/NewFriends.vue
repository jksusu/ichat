<template>
    <div>
        <SecondLevelContainer name="新的好友">
            <template #list>
                <el-scrollbar max-height="calc(100vh - 68px)">
                    <div style="margin: 8px 7px 0 8px;">
                        <div class="fried_container" v-for="item in newFriendsList" @click="select(item)">
                            <Item :class="{ select: item.uid == index }"
                                style="height: 62px; width: 356px; margin-top: 5px;" :nickname="item.nickname"
                                :headPortraitUrl="item.avatar" :account="item.from" :last-message="item.remark">
                            </Item>
                            <div v-if="item.status == 1">
                                <div @click="applyReplay(false, item.from)" class="button toreject"><span>拒绝</span></div>
                                <div @click="applyReplay(true, item.from)" class="button agree"><span>同意</span></div>
                            </div>
                        </div>
                    </div>
                </el-scrollbar>
            </template>
            <template #detail>
                <div class="detail_container">
                    <div class="card_container" v-if="index != ''">
                        <PersonalDataCard :nickname="state.nickname" :account="state.uid" :message="state.message"
                            :headPortraitUrl='state.avatar'>
                        </PersonalDataCard>
                        <div v-if="state.status == 1">
                            <div @click="applyReplay(false, state.uid)" class="button toreject"><span>拒绝</span></div>
                            <div @click="applyReplay(true, state.uid)" class="button agree"><span>同意</span></div>
                        </div>
                    </div>
                    <DefaultNullValue v-else></DefaultNullValue>
                </div>
            </template>
        </SecondLevelContainer>
    </div>
</template>
<script lang="ts" setup>
import { reactive } from "vue"
import { computed } from "@vue/reactivity"
import Item from '@/renderer/components/conversation/Item.vue'
import SecondLevelContainer from "../layout/SecondLevelContainer.vue"
import DefaultNullValue from "@/renderer/components/DefaultNullValue.vue"
import { useSelectIndexStore } from '@/renderer/stores/modules/selectIndex'
import PersonalDataCard from "@/renderer/components/personal/PersonalDataCard.vue"
import { NewFriendsStore } from "@/renderer/stores/modules/newFriends"
import { Socket } from "@/renderer/module/socket"
import { ElMessage } from 'element-plus'

const indexStore = useSelectIndexStore()
const index = computed(() => indexStore.getNewFriendsIndex)
const select = (list) => {
    indexStore.setNewFriendsIndex(list.uid)
    setSelectUser(list)
}

const state = reactive({
    uid: '',
    nickname: '',
    message: '',
    avatar: '',
    status: 0,
    detailContainerShow: false,
})

const newFriendsList = NewFriendsStore().getList

const setSelectUser = (list) => {
    state.uid = list.from
    state.nickname = list.nickname
    state.avatar = list.avatar
    state.status = list.status
    state.detailContainerShow = true
}


const applyReplay = async (bool: boolean, to: string) => {
    let states = bool ? 2 : 3
    let { status, res } = await Socket().contactsApplyReply(to, states)
    if (!status) {
        ElMessage({ showClose: true, message: '失败', type: 'error' })
        return
    }
    ElMessage({ showClose: true, message: '处理成功', type: 'success' })
}

</script>
<style lang="scss" scoped>
.selectStyle {
    background-color: #EAEAEA;
}

.fried_container {
    position: relative;

    &:not(:first-child) {
        margin-top: 4px;
    }

    .toreject {
        position: absolute;
        right: 60px;
        top: 16px;
    }

    .agree {
        right: 10px;
        position: absolute;
        top: 16px;
    }
}

.button {
    width: 44px;
    height: 28px;
    background: #FFFFFF;
    border-radius: 4px 4px 4px 4px;
    text-align: center;
    line-height: 28px;
    border: 1px solid #DEDEDE;

    &:hover {
        color: #1560FA;
        border: 1px solid #1560FA;
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

        .toreject {
            position: absolute;
            right: 86px;
            top: 252px;
        }

        .agree {
            position: absolute;
            right: 30px;
            top: 252px;
        }
    }
}
</style>