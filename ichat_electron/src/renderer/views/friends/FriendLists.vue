<template>
    <div>
        <SecondLevelContainer :name="titleName">
            <template #list>
                <el-scrollbar max-height="823px">
                    <div class="friends_container">
                        <Items style="height: 62px; width: 356px;" v-for="item in friendList" :username="item.name"
                            :isNameCenter="true" @click="select(item)"
                            :class="{ friends_select_style: item.friendUid == state.account }" />
                    </div>
                </el-scrollbar>
            </template>
            <template #detail>
                <div class="detail_container">
                    <div class="card_container" v-if="state.detailContainerShow">
                        <PersonalDataCard :nickname="state.nickname" :account="state.account"
                            :headPortraitUrl="state.headPortraitUrl">
                        </PersonalDataCard>
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
import DefaultNullValue from "@/renderer/components/DefaultNullValue.vue";
import { useSelectIndexStore } from '@/renderer/stores/modules/selectIndex'
import PersonalDataCard from "@/renderer/components/personal/PersonalDataCard.vue";
import SecondLevelContainer from "@/renderer/views/layout/SecondLevelContainer.vue";

const indexStore = useSelectIndexStore()
const index = computed(() => indexStore.getFriendsIndex)
const select = (item) => {
    indexStore.setFriendsIndex(item.friendUid)
    setSelectUser(item)
}

const titleName = '联系人';

const friendList = [
    {
        'friendUid': '1212342232',
        'name': '老bady',
        'headPortraitUrl': '',
    },
    {
        'friendUid': '1212342232443fd',
        'name': 'halo',
        'headPortraitUrl': '',
    }
]

const state = reactive({
    nickname: '',
    account: '',
    headPortraitUrl: '',
    detailContainerShow: false,
})

const setSelectUser = (list) => {
    state.nickname = list.name
    state.account = list.friendUid
    state.headPortraitUrl = list.headPortraitUrl
    state.detailContainerShow = true
}
if (index.value) {
    friendList.forEach((items) => {
        if (items.friendUid == index.value) {
            setSelectUser(items)
        }
    })
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
</style>