<template>
    <div>
        <SecondLevelContainer :name="titleName">
            <template #list>
                <el-scrollbar max-height="calc(100vh - 68px)">
                    <div style="margin: 8px 7px 0 8px;">
                        <div class="fried_container" v-for="item in newFriendsList">
                            <Items style="height: 62px; width: 356px;" :username="item.nickname"
                                :last-message="item.message" @click="select(item)"
                                :class="{ selectStyle: item.uid == index }" />
                            <div v-if="item.status == 1">
                                <div @click="toreject()" class="button toreject"><span>拒绝</span></div>
                                <div @click="agree()" class="button agree"><span>同意</span></div>
                            </div>
                        </div>
                    </div>
                </el-scrollbar>
            </template>
            <template #detail>
                <div class="detail_container">
                    <div class="card_container" v-if="index != ''">
                        <PersonalDataCard :nickname="state.nickname" :account="state.uid" :message="state.message"
                            :headPortraitUrl='state.headPortraitUrl'>
                        </PersonalDataCard>
                        <div v-if="state.status == 1">
                            <div @click="toreject()" class="button toreject"><span>拒绝</span></div>
                            <div @click="agree()" class="button agree"><span>同意</span></div>
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
import { computed } from "@vue/reactivity";
import Items from '@/renderer/components/conversation/Items.vue';
import SecondLevelContainer from "../layout/SecondLevelContainer.vue";
import DefaultNullValue from "@/renderer/components/DefaultNullValue.vue";
import { useSelectIndexStore } from '@/renderer/stores/modules/selectIndex'
import PersonalDataCard from "@/renderer/components/personal/PersonalDataCard.vue";

const titleName = '新的好友';

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
    headPortraitUrl: '',
    status: 0,
    detailContainerShow: false,
})


//status 1:申请中 2:已添加 3:已拒绝 3:已过期
const newFriendsList = [
    {
        'uid': '123214',
        'nickname': '飞天小子',
        'message': '能加一个好友吗？',
        'headPortraitUrl': '',
        'status': 1,
    },
    {
        'uid': '1232114',
        'nickname': '情爱的',
        'message': '能加一个好友吗？',
        'headPortraitUrl': '',
        'status': 2,
    }
]

const setSelectUser = (list) => {
    state.uid = list.uid
    state.nickname = list.nickname
    state.headPortraitUrl = list.headPortraitUrl
    state.status = list.status
    state.detailContainerShow = true
}

//index变化渲染右侧值
if (index.value != '') {
    newFriendsList.forEach(Items => {
        if (Items.uid == index.value) {
            setSelectUser(Items)
        }
    })
}

const agree = () => {
    console.log('同意添加好友')
}

const toreject = () => {
    console.log('拒绝添加好友')
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