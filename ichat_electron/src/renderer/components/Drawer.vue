<template>
    <div class="drawer_container">
        <el-drawer v-model="state.status" @close="close()" :with-header="false" :size="360">
            <div v-if="state.show == 'group'" class="session_container_drawer_group">
                <div class="margin_">
                    <div class="image_and_nickname">
                        <div><el-avatar shape="square" :size="42" /></div>
                        <div class="nickname"><span>群名字</span></div>
                    </div>
                    <div>
                        <el-divider />
                    </div>
                    <div>
                        <div style="margin-top: 19px;"><span class="trip">我在本群的昵称</span></div>
                        <div style="margin-top: 8px;"><span class="value">力气秒</span></div>
                    </div>
                    <div>
                        <div style="margin-top: 19px;"><span class="trip">群公告</span></div>
                        <div class="setting_item" style="margin-top: 8px;">
                            <div><span class="value" style="line-height: 32px;">点击设置群公告</span></div>
                            <div style="display: flex;">
                                <div style="line-height: 29px;"><span class="trip" style="font-size: 12px;">仅群主可编辑</span>
                                </div>
                                <el-icon style="margin-top: 9px;">
                                    <ArrowRightBold color="#C4C4C4" size="4" />
                                </el-icon>
                            </div>
                        </div>
                    </div>
                    <div class="setting_item" style="margin-top: 8px;">
                        <div><span class="value" style="line-height: 32px;">聊天记录</span></div>
                        <el-icon style="margin-top: 9px;">
                            <ArrowRightBold color="#C4C4C4" size="4" />
                        </el-icon>
                    </div>
                    <el-tabs v-model="state.activeName" type="card" style="margin-top: 19px;" @tab-click="handleClick">
                        <hr style="margin-top: 18px;border-color: rgb(255 255 255 / 24%);" />
                        <el-tab-pane class="groups" label="群成员管理" name="first">
                            <div class="session_container_search">
                                <el-input v-model="state.searchNickname" placeholder="输入搜索的会员名">
                                    <template #prepend>
                                        <el-button :icon="Search" />
                                    </template>
                                </el-input>
                            </div>
                            <div style="display: flex;margin-top: 18px;">
                                <div class="group_icon">
                                    <el-icon color="#C4C4C4" size="22">
                                        <Plus />
                                    </el-icon>
                                </div>
                                <div class="group_icon" style="margin-left: 20px;">
                                    <el-icon color="#C4C4C4" size="22">
                                        <Minus />
                                    </el-icon>
                                </div>
                            </div>
                            <div class="group_person_list">
                                <el-scrollbar height="350px" v-if="state.groupList">
                                    <div style="width: 318px;" v-infinite-scroll="state.groupList" class="items"
                                        v-for="item in state.groupList">
                                        <Item :nickname="item.nickname" :headPortraitUrl="item.headPortraitUrl"
                                            :account="item.uid" pattern="show" v-bind="$attrs">
                                        </Item>
                                    </div>
                                </el-scrollbar>
                            </div>
                        </el-tab-pane>
                        <el-tab-pane label="群聊设置" name="second">
                            <div style="margin-top: 18px;">
                                <span class="trip">聊天设置</span>
                            </div>
                            <div class="setting_item">
                                <div><span class="value" style="line-height: 32px;">消息免打扰</span></div>
                                <el-switch v-model="state.messageImmunity" />
                            </div>
                            <div class="setting_item">
                                <div><span class="value" style="line-height: 32px;">群聊置顶</span></div>
                                <el-switch v-model="state.top" />
                            </div>
                            <div style="margin-top: 24px;">
                                <span class="trip">群设置</span>
                            </div>
                            <div class="setting_item">
                                <div><span class="value" style="line-height: 32px;">全员禁言</span></div>
                                <el-switch v-model="state.estoppel" />
                            </div>
                            <div class="setting_item">
                                <div><span class="value" style="line-height: 32px;">入群审核</span></div>
                                <el-switch v-model="state.joinGroupToExamine" />
                            </div>
                            <div class="setting_item">
                                <div><span class="value" style="line-height: 32px;">群成员邀请好友入群</span></div>
                                <el-switch v-model="state.InviteFriends" />
                            </div>
                            <div class="setting_item">
                                <div><span class="value" style="line-height: 32px;">入群申请</span></div>
                                <div style="line-height: 32px;text-align: center;"><el-icon>
                                        <ArrowRightBold color="#C4C4C4" size="12" />
                                    </el-icon></div>
                            </div>
                            <div class="setting_item">
                                <div><span class="value" style="line-height: 32px;">群管理员设置</span></div>
                                <div style="line-height: 32px;text-align: center;"><el-icon>
                                        <ArrowRightBold color="#C4C4C4" size="12" />
                                    </el-icon></div>
                            </div>
                        </el-tab-pane>
                    </el-tabs>
                </div>
                <div class="jiesanqunliao">
                    <hr style="border-color: rgb(255 255 255 / 24%);" />
                    <span style="margin-top: 15px;">解散群聊</span>
                </div>
            </div>
            <div v-else class="session_container_drawer_friend">
                <div class="margin_">
                    <div class="person_image_and_nickname">
                        <div><el-avatar shape="square" :size="42" /></div>
                        <div class="person_nickname"><span>{{ personInfo.nickname }}</span></div>
                    </div>
                    <hr style="border-color: rgb(255 255 255 / 24%);" />
                    <div class="setting_item" style="margin-top: 8px;">
                        <div><span class="value" style="line-height: 32px;">聊天记录</span></div>
                        <el-icon style="margin-top: 9px;">
                            <ArrowRightBold color="#C4C4C4" size="4" />
                        </el-icon>
                    </div>
                    <hr style="border-color: rgb(255 255 255 / 24%);" />
                    <div class="setting_item">
                        <div><span class="value" style="line-height: 32px;">消息免打扰</span></div>
                        <el-switch v-model="state.messageImmunity" />
                    </div>
                    <div class="setting_item">
                        <div><span class="value" style="line-height: 32px;">聊天置顶</span></div>
                        <el-switch v-model="state.top" />
                    </div>
                </div>
                <div class="qingkongliaotianjilu">
                    <hr style="border-color: rgb(255 255 255 / 24%);" />
                    <span>清空聊天记录</span>
                </div>
            </div>
        </el-drawer>
    </div>
</template>

<script lang="ts" setup>
import { Search } from '@element-plus/icons-vue'
import type { TabsPaneContext } from 'element-plus'
import { reactive, watch, ref, computed } from "vue"
import Item from '@/renderer/components/conversation/Item.vue';
import { SessionStore } from '../stores/modules/sessionList';

//要传给父组件的方法
const emits = defineEmits(["close"])
const props = withDefaults(defineProps<{ status: boolean, type: string }>(), { status: false, type: 'person' })
watch(props, (newValue, oldValue) => {
    state.status = newValue.status
    if (props.type == 'person') {
        state.show = 'person'
    } else {
        state.show = 'group'
    }
})
//当前选择数据
const personInfo = computed(() => SessionStore().getSelectSession)


// const personInfo = reactive({
//     uid: "",
//     nickname: "",
//     headPortraitUrl: '',
//     lable: '',
//     type: ''
// })

const state = reactive({
    status: false,
    show: "person",
    searchNickname: "",
    activeName: "first",
    messageImmunity: false,
    top: false,
    estoppel: false,
    joinGroupToExamine: false,
    InviteFriends: false,
    groupList: [{
        uid: '12345',
        nickname: '芝士蛋糕',
        headPortraitUrl: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
        lastMessage: '',
        lastMessageTime: '15:20',
        unreadMessageNumber: 10,
        lable: ''
    }],
})

//关闭调用父组件的关闭方法
const close = () => { emits("close") }

const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event)
}

</script>

<style lang='scss'>
.infinite-list {
    height: 300px;
    padding: 0;
    margin: 0;
    list-style: none;
}

.infinite-list .infinite-list-item {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 50px;
    background: var(--el-color-primary-light-9);
    margin: 10px;
    color: var(--el-color-primary);
}

.infinite-list .infinite-list-item+.list-item {
    margin-top: 10px;
}

.drawer_container {

    .el-tabs__item {
        padding: 0 8px;
        height: var(--el-tabs-header-height);
        box-sizing: border-box;
        line-height: var(--el-tabs-header-height);
        display: inline-block;
        list-style: none;
        font-size: var(--el-font-size-base);
        font-weight: 500;
        color: var(--el-text-color-primary);
        position: relative;
    }

    .el-tabs__header {
        padding: 0;
        position: relative;
        margin: 0 0 0 0;
    }

    .el-tabs__item :hover {
        color: none;
    }

    .el-tabs__item:hover {
        color: none;
        cursor: pointer;
    }


    .el-tabs--card>.el-tabs__header .el-tabs__nav {
        border: none
    }

    .el-tabs__item.is-active {
        //color: #1560FA;
        //--el-bg-color: #F5F6F7;
        //--el-bg-color: none;
        --el-color-primary: #1560FA;
        background: #F5F6F7;
    }

    .el-tabs--card>.el-tabs__header .el-tabs__item.is-active {
        border-bottom-color: none !important;
    }

    .el-tabs--card>.el-tabs__header .el-tabs__item {
        border-left: none
    }

    .el-tabs--card>.el-tabs__header {
        border-bottom: none
    }

    .el-overlay {
        position: absolute !important;
        background-color: rgb(0 0 0 / 20%);
        border-radius: 12px 12px 12px 12px;
    }

    .el-drawer__body {
        flex: 1;
        padding: 0;
        overflow: auto;
    }

    .el-divider--horizontal {
        margin: 18px 0;
    }

    .session_container_drawer_group {
        display: flex;
        flex-direction: column;

        .image_and_nickname {
            margin-top: 18px;
            height: 42px;
            display: flex;
            align-items: center;

            .nickname {
                height: 32px;
                max-width: 235px;

                span {
                    line-height: 32px;
                    margin-left: 12px;
                    height: 20px;
                    font-size: 14px;
                    font-weight: 500;
                    color: #333333;
                }
            }
        }

        .session_container_search {
            margin-top: 18px;

            .el-input {
                width: 318px;
                height: 30px;
                background: #F5F6F7;
                border-radius: 4px 4px 4px 4px;
                opacity: 1;
                border: 1px solid #EBEBEB;
            }

            .el-input-group__prepend {
                padding: 0 10px !important;
                box-shadow: none !important;

                .el-icon {
                    height: 2em !important;
                }
            }


        }

        .group_icon {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 36px;
            height: 36px;
            border-radius: 4px 4px 4px 4px;
            opacity: 1;
            border: 1px solid #EBEBEB;
            margin-left: 10px;
        }

        .jiesanqunliao {
            height: 72px;
            position: absolute;
            bottom: 0;
            width: 100%;
            padding: 0px;

            span {
                position: absolute;
                left: 40%;
                width: 56px;
                height: 20px;
                font-size: 14px;
                font-weight: 400;
                color: #F43E24;
                line-height: 16px;
            }
        }
    }

    .setting_item {
        display: flex;
        justify-content: space-between;
        margin-top: 16px;

        .el-switch {
            --el-switch-on-color: #1560fa;
        }
    }

    .group_person_list {
        margin-top: 18px;
        width: 318px;
    }


    .session_container_drawer_friend {
        display: flex;
        flex-direction: column;

        .person_image_and_nickname {
            margin-top: 18px;
            display: flex;
            align-items: center;

            .person_nickname {
                margin-left: 8px;
                width: 103px;
                height: 20px;
                font-size: 14px;
                font-weight: 500;
                color: #333333;
                line-height: 16px;
            }
        }

        .qingkongliaotianjilu {
            text-align: center;

            span {
                margin-top: 20px;
                font-size: 14px;
                font-weight: 400;
                color: #F43E24;
                line-height: 16px;
            }
        }
    }
}



.margin_ {
    margin: 0 24px 0 18px;
}

.trip {
    width: 98px;
    height: 20px;
    font-size: 14px;
    font-weight: 400;
    color: #999999;
    line-height: 16px;
}

.value {
    width: 42px;
    height: 20px;
    font-size: 14px;
    font-weight: 400;
    color: #333333;
    line-height: 16px;
}
</style>


<style lang="scss" ></style>