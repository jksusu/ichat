<template>
    <div>
        <SecondLevelContainer :name="titleName">
            <template #list>
                <el-scrollbar max-height="823px">
                    <div class="group_lists">
                        <div v-for="item in groupLists" @click="select(item.groupNumber)">
                            <Item :class="{ select: index == item.groupNumber }"
                                style="height: 62px; width: 356px; margin-top: 5px;" :nickname="item.groupName"
                                :headPortraitUrl="item.headPortraitUrl" :account="item.groupNumber" v-bind="$attrs">
                            </Item>
                        </div>
                    </div>
                </el-scrollbar>
            </template>
            <template #detail>
                <div class="detail_container">
                    <div class="box_card" v-if="index != ''">
                        <div style="margin: 0 18px 0 18px;">
                            <div class="title">
                                <el-avatar shape="square" :size="42" :src="state.headPortraitUrl" />
                                <span>{{ state.groupName }}</span>
                            </div>
                            <hr />
                            <div class="lists">
                                <div class="info" v-for="item in state.groupMember">
                                    <img :src="item.headPortraitUrl" />
                                    <div><span>{{ item.nickname }}</span></div>
                                </div>
                            </div>
                            <hr />
                            <div class="buttom">
                                <div class="btn"><span>发消息</span></div>
                            </div>
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
import SecondLevelContainer from "../layout/SecondLevelContainer.vue";
import DefaultNullValue from "@/renderer/components/DefaultNullValue.vue";
import { useSelectIndexStore } from '@/renderer/stores/modules/selectIndex'

const indexStore = useSelectIndexStore()
const index = computed(() => indexStore.getGroupsIndex)
const select = (groupNumber: string) => {
    indexStore.setGroupsIndex(groupNumber)
}

const titleName = '群聊';
const groupLists = [
    {
        'groupNumber': '12365',
        'groupName': '群聊名',
        'message': '接收来自新的好友的申请',
        'headPortraitUrl': '',
        'groupMember': [
            {
                uid: '',
                nickname: '你好',
                headPortraitUrl: '',
            }
        ]
    }
]

const state = reactive({
    headPortraitUrl: '',
    groupName: 'NIEK官方旗舰店',
    groupMember: [
        {
            uid: '',
            nickname: '你好',
            headPortraitUrl: '',
        }
    ]
})

</script>
<style lang="scss" scoped>
.group_lists {
    margin: 8px 7px 0 8px;
    position: relative;

    .items_container:not(:first-child) {
        margin-top: 4px;
    }
}

.detail_container {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;

    .box_card {
        width: 354px;
        height: 369px;
        background: #FFFFFF;
        box-shadow: 0px 8px 30px 1px rgba(0, 0, 0, 0.08);
        border-radius: 8px 8px 8px 8px;

        hr {
            size: 1rem;
            color: #EAEAEA;
        }

        .title {
            display: flex;
            margin-top: 20px;

            img {
                width: 42px;
                height: 42px;
                background: #F5F6F7;
                border-radius: 4px 4px 4px 4px;
            }

            span {
                padding-left: 10px;
                max-width: 200px;
                height: 42px;
                font-size: 14px;
                font-weight: 500;
                color: #333333;
                line-height: 42px;
            }
        }

        .lists {
            min-height: 215px;
            display: flex;
            flex-wrap: wrap;

            .info {
                width: 42px;
                height: 42px;
                background: #E0E0E0;
                border-radius: 4px 4px 4px 4px;
                text-align: center;



                span {
                    width: 46px;
                    height: 17px;
                    font-size: 12px;
                    font-weight: 400;
                    color: #333333;
                    line-height: 14px;
                }
            }
        }

        .btn {
            width: 74px;
            height: 36px;
            background: #1560FA;
            border-radius: 4px 4px 4px 4px;
            text-align: center;
            line-height: 36px;
            margin: 18px 0 18px 244px;

            span {
                width: 42px;
                height: 20px;
                font-size: 14px;
                font-weight: 400;
                color: #FFFFFF;
                line-height: 16px;
            }
        }
    }
}
</style>