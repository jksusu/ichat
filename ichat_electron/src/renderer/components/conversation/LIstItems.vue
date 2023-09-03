<template>
    <div class="list_container">
        <Information>
            <template #default>
                <div class="top_feature">
                    <span>{{ props.titlename }}</span>
                </div>
            </template>
        </Information>
        <div class="item_container">
            <el-scrollbar v-if="props.list">
                <div class="items" v-for="item in props.list" @click="select(item.uid)">
                    <Item :class="{ select: item.uid == index }" :nickname="item.nickname" :uid="item.uid"
                        :headPortraitUrl="item.headPortraitUrl" :account="item.uid" :lastMessage="item.lastMessage"
                        :lastMessageTime="LastMessageTimeShow(item.lastMessageTime)" :pattern="props.pattern"
                        :unreadMessageNumber="item.unreadMessageNumber" v-bind="$attrs" :data="item"
                        @contextmenu.prevent="handleRightClick(item)" 
                        @contextmenu="sessionListRightClick">
                    </Item>
                </div>
            </el-scrollbar>
            <div style="display:flex;justify-content: center" v-else>
                <DefaultNullValue></DefaultNullValue>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { computed } from "@vue/reactivity";
import { ref, reactive } from "vue"
import Item from '@/renderer/components/conversation/Item.vue';
import Information from '@/renderer/components/personal/Information.vue';
import { useSelectIndexStore } from '@/renderer/stores/modules/selectIndex'
import DefaultNullValue from "../DefaultNullValue.vue";
import { LastMessageTimeShow } from "@/renderer/module/tool";
import { delSession, nodisturb, msgTopping, disturb, cancelTopping } from "@/renderer/module/menu"
import { menusEvent } from 'vue3-menus';

const index = computed(() => useSelectIndexStore().getSessionIndex)
const select = (uid: string) => {
    useSelectIndexStore().setSessionIndex(uid)
}

//右键菜单
const rightSelect = reactive({
    item: ''
})
const handleRightClick = (item: any) => {
    rightSelect.item = item
}
//组装按钮
const sessionListRightClick = (event) => {
    if (rightSelect.item) {
        menusEvent(event, [delSession, nodisturb, msgTopping], rightSelect.item)
    } else {
        menusEvent(event, [delSession, disturb, cancelTopping], rightSelect.item)
    }
}

/**
 * pattern 模式
 * 1.message 消息模式 会完全展示
 * 2.show 展示模式 昵称居中显示 其他配置项不显示
 */
const props = withDefaults(defineProps<{
    titlename?: string,
    pattern?: string,
    list?: any[],
}>(), {})


</script>

<style lang="scss" >
.select {
    background-color: #EAEAEA;
}

.list_container {
    width: 296px;
    height: 100%;
    background: #FFFFFF;

    .item_container {
        margin: 0 8px 0 8px;
        //height: calc(100vh - 139px);
        .items {
            padding-top: 4px;

            &:first-child {
                margin-top: 12px;
            }
        }
    }
}

.top_feature {
    width: 296px;
    height: 44px;
    background: #FFFFFF;
    border-bottom: 1px solid #EAEAEA;

    span {
        padding-left: 18px;
        width: 56px;
        height: 20px;
        font-size: 14px;
        font-weight: 500;
        color: #333333;
        line-height: 44px;
    }
}
</style>