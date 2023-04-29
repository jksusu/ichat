<template>
    <div id="dropdown">
        <el-dropdown placement="bottom-end">
            <div class="add-btn">
                <SvgIcon name="plus_sign"></SvgIcon>
            </div>
            <template #dropdown>
                <el-dropdown-menu>
                    <el-dropdown-item @click="state.dialogTableVisible = true">添加好友</el-dropdown-item>
                </el-dropdown-menu>
            </template>
        </el-dropdown>
        <el-dialog style="border-radius: 10px;" v-model="state.dialogTableVisible" draggable destroy-on-close
            :show-close="false" width="400" top="30vh" center>
            <div v-if="!state.showAddFriendPage" class="friend-search">
                <el-radio-group id="add-frined-radio" v-model="state.selectRadio" size="large">
                    <el-radio-button label="联系人" />
                    <el-radio-button label="群聊" />
                </el-radio-group>
                <el-input size="large" v-model="state.number" class="w-50 m-2 friend-search-input" placeholder="请输入id" />
                <div class="friend-search-button">
                    <el-button size="large" round @click="state.dialogTableVisible = false">关闭</el-button>
                    <el-button size="large" round @click="fsearch">搜索</el-button>
                </div>
            </div>
            <div v-else>
                <div class="account_info">
                    <div class="account">
                        <div class="nickanme"><span>{{ targetInfo.nickname }}</span></div>
                        <div class="userid"><span>{{ targetInfo.uid }}</span></div>
                    </div>
                    <el-avatar shape="square" :size="60" :src="targetInfo.avatar" />
                </div>
                <div class="friend-search-button">
                    <el-button size="large" round @click="state.dialogTableVisible = false">关闭</el-button>
                    <el-button size="large" round @click="addTargetApply">添加</el-button>
                </div>
            </div>
        </el-dialog>
    </div>
</template>
  
<script setup>
import { ref, reactive, h } from "vue"
import { ElMessage } from 'element-plus'
import { Socket } from '@/renderer/module/socket'
import SvgIcon from '@/renderer/components/SvgIcon.vue'

const state = reactive({
    selectRadio: '联系人',
    dialogTableVisible: false,
    showAddFriendPage: false,
    number: '505389',
})
const targetInfo = reactive({
    uid: '',
    nickname: '',
    avatar: '',
})

const searchF = ref("");

const fsearch = async () => {
    if (state.number == '') { return }
    let { status, res } = await Socket().targetSearch(state.number)
    if (!status && res.message.data == null) {
        //未查询到
        ElMessage({
            message: '用户不存在',
            grouping: true,
            type: 'success',
        })
        return
    }
    let { data } = res.message
    targetInfo.uid = data.uid
    targetInfo.nickname = data.nickname
    targetInfo.avatar = data.head_portrait

    state.showAddFriendPage = true
}

//好友申请
const addTargetApply = async () => {
    if (targetInfo.uid == '') { return }
    let { status, res } = await Socket().contactsApply(targetInfo.uid, 'hello')
    if (status) { closeAll }
}

const closeAll = () => {
    targetInfo = {
        uid: '',
        nickname: '',
        avatar: '',
    }
    state.dialogTableVisible = false
    state.showAddFriendPage = false

}

</script>

<style lang="scss">
#dropdown {

    .el-dialog.is-draggable .el-dialog__header {
        display: none;
    }

    --el-color-primary: #1560FA;

    #add-frined-radio .el-radio-button--large .el-radio-button__inner {
        width: 175px;
    }

    .el-radio-button__inner {
        padding: 10px 20px;
        font-size: 14px;
        width: 332px;
    }

    .friend-search-input {
        width: 350px !important;
        padding-top: 30px;
    }

    .friend-search {
        display: flex;
        flex-direction: column;

    }

    .friend-search-button {
        margin-top: 30px;
        display: flex;
        justify-content: space-around;
    }

    .el-button--large.is-round {
        padding: 18px 40px;
    }

    .add-btn {
        border-radius: 4px;
        width: 24px;
        height: 24px;
        background-color: #e2e2e2;
        display: flex;
        justify-content: center;
        align-items: center;

        // :hover {
        //     background-color: #d1d1d1;
        // }
    }


    .account_info {
        display: flex;
        width: 292px;
        height: 60px;
        margin: 30px 0 0 30px;

        .userid {
            max-width: 94px;
            height: 20px;
            overflow: hidden;
            margin-top: 9px;
            font-size: 14px;
            font-weight: 400;
            color: #666666;
            line-height: 16px;
        }

        .nickanme {
            width: 216px;
            height: 24px;
            font-size: 18px;
            font-weight: 500;
            color: #333333;
            line-height: 24px;
            overflow: hidden;
        }

        img {
            margin-left: 16px;
            width: 60px;
            height: 60px;
            background: #E0E0E0;
            border-radius: 4px 4px 4px 4px;
        }
    }
}
</style>