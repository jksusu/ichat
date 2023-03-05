<template>
    <div class="chat_bubble_container" :class="direction">
        <div class="img_container">
            <img :src="headPortraitUrl" />
        </div>
        <div class="message_container">
            <!-- <div class="name_role" :class="direction" v-if="isShowNicknameRole">
                <span>飞天瞎子</span>
                <div><span>群主</span></div>
            </div> -->
            <div class="message_content">
                <div class="font_message">
                    <div class="font_content">
                        <span>{{ content }}</span>
                    </div>
                    <!-- <div class="quote_message">
                        <span>{{ content }}</span>
                    </div> -->
                </div>
                <!-- <div class="file_message">
                    <div>
                        <img />
                    </div>
                    <div class="file_message_font">
                        <span>文件名</span>
                        <img />
                        <span>350kb</span>
                    </div>
                </div> -->
                <!-- <div class="video_message"></div> -->
                <!-- <div v-if="!isShowNicknameRole" class="message_status" :style="messageStatus"><span>已读</span></div> -->

                <!-- <div class="record_message">
                    <div class="record_message_title">与“芝士蛋糕”的聊天记录<span></span></div>
                    <div class="record_message_content">
                        ：导航菜单是一个网站的灵魂，用户依赖导航在各个页面中进行跳转。一般分为顶部导航和侧边导航，顶部导航提供全局性的类目和功能，侧边导航提供多级结构来收纳和排列...<span></span>
                    </div>
                    <el-divider />
                    <div class="record_message_end"><span>聊天记录</span></div>
                </div> -->
            </div>
            <!-- <div class="withdraw_message"><span>“芝士蛋糕”撤回了一条信息</span></div> -->
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, toRefs } from 'vue'

/**
 * 定义支持的类型
 * 文字 语音消息 图片 视频消息 文件 链接
 */
const contentType = [1, 2, 3, 4, 5, 6]

const props = withDefaults(defineProps<{
    headPortraitUrl?: string,
    nickname?: string,
    role?: string,
    status?: string,
    content?: string,
    messageType?: string,
    isSelf?: boolean,
    quote?: string,//引用内容
}>(), {
    headPortrait: "",
    nickname: "",
    role: "",
    status: "",
    content: "",
    messageType: "",
    isSelf: false,
    quote: "",
})


const state = reactive({
    direction: "",
    isShowNicknameRole: true,
    content: "",
    messageStatus: ""
})

//赋值
state.content = props.content


if (props.isSelf) {
    state.direction = 'right'
    state.isShowNicknameRole = false//不显示昵称和角色
} else {
    state.direction = 'left'
    state.isShowNicknameRole = true//不显示昵称和角色
}

//是否已读
if (props.status == "1") {
    state.messageStatus = 'color: #999999;'
} else {
    state.messageStatus = 'color: #1560FA;'
}



const { direction, isShowNicknameRole, messageStatus, content } = toRefs(state)
</script>

<style lang="scss">
:root {
    --chatMessageMaxWidth: 464px;
    --chatNameMaxWidth: 464px;
}

.file_message {
    display: flex;
    width: 270px;
    height: 72px;
    background: #FFFFFF;
    border-radius: 4px 4px 4px 4px;
    opacity: 1;

    img {
        width: 48px;
        height: 48px;
        margin: 12px 0 12px 12px;
    }

    .file_message_font {
        margin: 12px 0 0 12px;
        width: 196px;
        height: 41px;
        display: flex;
        flex-direction: column;
        position: relative;

        //复制按钮
        img {
            height: 24px;
            width: 24px;
            position: absolute;
            right: 10px;
            bottom: 7px;
        }

        :nth-child(1) {
            width: 56px;
            height: 20px;
            font-size: 14px;
            font-weight: 500;
            color: #333333;
            line-height: 16px;
        }

        :last-child {
            padding-top: 6px;
            max-width: 182px;
            height: 17px;
            font-size: 12px;
            font-weight: 400;
            color: #999999;
            line-height: 14px;
        }
    }
}

.video_message {
    width: 261px;
    height: 147px;
    background: rgba(0, 0, 0, 0.4);
    border-radius: 4px 4px 4px 4px;
    opacity: 1;
}

.quote_message {
    background: #EAEAEA;
    border-radius: 4px 4px 4px 4px;
    margin-top: 12px;
    min-height: 36px;

    span {
        padding: 8px;
        text-align: center;
        font-size: 12px;
        font-weight: 400;
        color: #666666;
        line-height: 36px;
    }
}

.record_message {
    width: 270px;
    height: 149px;
    background: #FFFFFF;
    border-radius: 4px 4px 4px 4px;
    opacity: 1;

    .record_message_title {
        width: 246px;
        height: 20px;
        font-size: 14px;
        font-weight: 500;
        color: #333333;
        line-height: 16px;
        margin: 12px 12px 0 12px;
    }

    .record_message_content {
        width: 246px;
        height: 70px;
        font-size: 12px;
        font-weight: 400;
        color: #999999;
        line-height: 14px;
        margin: 6px 12px 0 12px;
    }

    .el-divider--horizontal {
        display: block;
        width: 246px;
        margin: 0;
        margin: 6px 12px 0 12px;
    }

    .record_message_end {
        width: 48px;
        height: 17px;
        font-size: 12px;
        font-weight: 400;
        color: #999999;
        line-height: 14px;
        margin: 6px 210px 12px 12px;
    }
}

.withdraw_message {
    max-width: 200px;
    height: 24px;
    font-size: 12px;
    font-weight: 400;
    color: #999999;
    line-height: 24px;
}

.left {
    flex-direction: row;

    .message_container {
        margin-left: 10px;
    }

    .font_content {
        background: #FFFFFF;
        word-break: break-all;
        :before {
            right: 100%;
            border-right: 6px solid #FFFFFF;
        }
    }
}

.right {
    flex-direction: row-reverse;

    .message_container {
        margin-right: 14px;
    }

    .message_content {
        flex-direction: row-reverse
    }

    .font_content {
        background: #0652EE;
        word-break: break-all;
        span {
            color: #FFFFFF !important;
        }

        :before {
            left: 100%;
            border-left: 6px solid #0652EE;
        }
    }
}

.chat_bubble_container {
    display: flex;

    .img_container {
        width: 44px;
        height: 44px;

        img {
            width: 36px;
            height: 36px;
            padding: 2px 4px 6px 4px;
        }

    }

    .message_container {
        .name_role {
            margin-top: 10px;
            display: flex;
            align-items: center;
            height: 17px;
            font-size: 12px;

            :nth-child(1) {
                max-width: var(--chatNameMaxWidth);
                font-size: 12px;
                font-weight: 400;
                color: #666666;
                line-height: 14px;
            }

            div {
                width: 28px;
                background: #EAEAEA;
                border-radius: 2px 2px 2px 2px;
                opacity: 1;
                text-align: center;
                margin-left: 6px;

                span {
                    max-width: 20px;
                    height: 14px;
                    font-size: 10px;
                    font-weight: 400;
                    color: #999999;
                    line-height: 12px;
                }
            }
        }

        .message_content {
            display: flex;

            .font_message {
                margin-top: 6px;
                max-width: var(--chatMessageMaxWidth);
                min-height: 36px;
                opacity: 1;
                position: relative;


                .font_content {
                    min-height: 20px;
                    padding: 8px;
                    border-radius: 4px 4px 4px 4px;

                    span {
                        font-size: 14px;
                        font-weight: 400;
                        color: #333333;
                        line-height: 16px;
                    }

                    :before {
                        content: "";
                        position: absolute;
                        top: 10px;
                        border-top: 6px solid transparent;
                        border-bottom: 11px solid transparent;
                    }
                }

            }

            .message_status {
                width: 24px;
                height: 17px;
                font-size: 12px;
                font-weight: 400;
                color: #1560FA;
                line-height: 14px;
                margin: 26px 6px 0 0;
            }
        }
    }
}
</style>