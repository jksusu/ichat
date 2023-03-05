<template>
    <div id="aside_container">
        <div class="title_menu">
            <LeftTitle />
        </div>
        <div class="aside_logo">
            <img width="40" height="40" :src="logo" />
        </div>
        <div style="height: 40px;"></div>
        <div v-for="item in state" @click="select(item.route)">
            <el-badge :value="item.tipsNumber">
                <div :id="item.iconName" class="aside_menu" :class="{ selects: index == item.route }">
                    <SvgIcon :color="index == item.route ? '#ffffff' : item.color" :name="item.iconName"></SvgIcon>
                </div>
            </el-badge>
        </div>
    </div>
</template>
  
<script lang="ts" setup>
import { watch, reactive } from 'vue';
import SvgIcon from './SvgIcon.vue';
import { jump } from '@/renderer/tool/tool';
import { computed } from "@vue/reactivity";
import LeftTitle from './title/LeftTitle.vue';
import logo from '@/renderer/assets/images/logo.png'
import { useSelectIndexStore } from '@/renderer/stores/modules/selectIndex'

const index = computed(() => useSelectIndexStore().getAsideMenuIndex)
const select = (route: string) => {
    useSelectIndexStore().setAsideMenuIndex(route)
    if (route == '/maillist') {
        //是否有缓存的路由，
        const i = useSelectIndexStore().getMaillistIndex
        jump(i)
    } else {
        jump(route)
    }
}

watch(index, async (newQuestion, oldQuestion) => {
    console.log('监听到数据变化')
    console.log(newQuestion)
})

const state = reactive({
    session: {
        'menuName': '会话',
        'iconName': 'conversation',
        'route': '/session',
        'color': '#808899',
        'tipsNumber': '',
    },
    friendlist: {
        'menuName': '通讯录',
        'iconName': 'maillist',
        'route': '/maillist',
        'color': '#808899',
        'tipsNumber': '',
    },
    setting: {
        'menuName': '设置',
        'iconName': 'setting',
        'route': '/setting',
        'color': '#808899',
        'tipsNumber': '',
    }
})
</script>
  
<style lang="scss" scoped>
.selects {
    background: #1560FA;
}

#setting {
    position: relative;
    bottom: 0;
}

#aside_container {
    width: 64px;
    height: 100%;
    background: #F5F6F7;
    display: flex;
    flex-direction: column;
    align-items: center;


    .el-badge__content.is-fixed {
        position: absolute;
        top: 12px;
        right: calc(9px + var(--el-badge-size)/ 2);
        transform: translateY(-50%) translateX(100%);
    }

    .title_menu {
        width: 64px;
        height: 32px;
    }

    .aside_logo {
        margin: 48px 12px 0 12px;
    }

    .aside_menu {
        margin-top: 8px;
        display: flex;
        align-items: center;
        justify-content: center;
        width: 40px;
        height: 40px;
        border-radius: 12px 12px 12px 12px;
    }

    .head_portrait {
        margin: 345px 17px 0 15px;
        -webkit-app-region: drag;
    }

}
</style>