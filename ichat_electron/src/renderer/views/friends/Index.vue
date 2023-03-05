<template>
    <div style="display: flex;">
        <div class="mail_container">
            <Information>
                <template #default>
                    <div class="friends_top_feature">
                        <div class="friends_top_title"><span>通讯录</span></div>
                    </div>
                </template>
            </Information>
            <div class="mail_list_container">
                <div class="items_container" v-for="item in featureButtonList" @click="select(item.route)"
                    :class="{ selectStyle: index == item.route }">
                    <div class="svg" :style="item.style"><img height="24" width="24" :src="item.iconUrl" /></div>
                    <div class="info">
                        <div>{{ item.name }}</div>
                        <div>{{ item.message }}</div>
                    </div>
                </div>
            </div>
        </div>

        <RouterView v-if="index"></RouterView>
        <div v-else>
            <SessionTitle></SessionTitle>
            <div class="second_leavel_container">
                <DefaultNullValue></DefaultNullValue>
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { computed } from "@vue/reactivity";
import { jump } from "@/renderer/tool/tool";
import SessionTitle from "@/renderer/components/title/SessionTitle.vue";
import Information from '@/renderer/components/personal/Information.vue';
import DefaultNullValue from "@/renderer/components/DefaultNullValue.vue";
import { useSelectIndexStore } from '@/renderer/stores/modules/selectIndex'
import newFriends from '@/renderer/assets/images/new_friends.png'
import groupChat from '@/renderer/assets/images/group_chat.png'
import friends from '@/renderer/assets/images/friends.png'

const index = computed(() => useSelectIndexStore().getMaillistIndex)
const select = (route: string) => {
    useSelectIndexStore().setMaillistIndex(route)
    jump(route)
}

const featureButtonList = [
    {
        'name': '新的朋友',
        'message': '接收来自新的好友的申请',
        'iconUrl': newFriends,
        'route': '/maillist/new_friends',
        'style': 'background: rgba(5,184,66,0.1);',
    }, {
        'name': '群聊',
        'message': '多人聊天、沟通',
        'iconUrl': groupChat,
        'route': '/maillist/groups',
        'style': 'background: rgba(21,96,250,0.1);',
    }, {
        'name': '联系人',
        'message': '记录同事、朋友的联系方式',
        'iconUrl': friends,
        'route': '/maillist/friends',
        'style': 'background: rgba(253,133,3,0.1);',
    },
]
</script>
<style lang="scss" scoped>
.second_leavel_container {
    display: flex;
    width: 1080px;
    height: 832px;
    background: #FFFFFF;
    align-items: center;
    justify-content: center;
}

.selectStyle {
    background-color: #EAEAEA;
}

.friends_top_feature {
    width: 296px;
    display: flex;
    background: #FFFFFF;
    border-bottom: 1px solid #EAEAEA;

    span {
        padding-left: 18px;
        height: 20px;
        font-size: 14px;
        font-weight: 500;
        color: #333333;
        line-height: 44px;
    }
}

.mail_container {
    display: flex;
    flex-direction: column;
}

.mail_list_container {
    margin-top: 1px;
    width: 296px;
    min-height: 761px;
    height: calc(100vh - 139px);
    background: #FFFFFF;

    .items_container:first-child {
        margin-top: 12px;
    }

    .items_container:nth-child(2) {
        margin-top: 6px;
    }

    .items_container:last-child {
        margin-top: 6px;
    }

    .items_container {
        margin-left: 8px;
        margin-right: 8px;

        .svg {
            margin: 10px 0 10px 10px;
            display: flex;
            justify-content: center;
            align-items: center;
            width: 42px;
            height: 42px;
            border-radius: 4px 4px 4px 4px;
        }

        .info {
            margin-left: 6px;

            :first-child {
                height: 20px;
                font-size: 14px;
                font-weight: 500;
                color: #333333;
                line-height: 16px;
            }

            :last-child {
                margin-top: 4px;
                height: 17px;
                font-size: 12px;
                font-weight: 400;
                color: #999999;
                line-height: 14px;
            }
        }
    }
}
</style>