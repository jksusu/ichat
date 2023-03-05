<template>
    <div id="message_input_container" class="message_input_container">
        <div class="tool_list">
            <span id="emoji-chat">
                <discord-picker style="" gif-format="md" @emoji="setEmoji" />
            </span>
            <span>
                <el-upload style="padding-top: 6px" accept="image/*" :show-file-list="false">
                    <svg t="1639381072726" class="icon" viewBox="0 0 1024 1024" version="1.1"
                        xmlns="http://www.w3.org/2000/svg" p-id="7287" width="20" height="20">
                        <path
                            d="M76.8 327.8848V187.7504a25.6 25.6 0 0 1 25.6-25.6h307.2a25.6 25.6 0 0 1 25.6 25.6v64H512v-64a102.4 102.4 0 0 0-102.4-102.4H102.4a102.4 102.4 0 0 0-102.4 102.4v153.6l76.8-13.4656z m0 0V187.7504a25.6 25.6 0 0 1 25.6-25.6h307.2a25.6 25.6 0 0 1 25.6 25.6v64H512v-64a102.4 102.4 0 0 0-102.4-102.4H102.4a102.4 102.4 0 0 0-102.4 102.4v153.6l76.8-13.4656z"
                            p-id="7288" fill="#4C4C4C" />
                        <path
                            d="M102.4 315.7504a25.6 25.6 0 0 0-25.6 25.6v496.2304a25.6 25.6 0 0 0 25.6 25.6h819.2a25.6 25.6 0 0 0 25.6-25.6V341.3504a25.6 25.6 0 0 0-25.6-25.6H102.4z m0-76.8h819.2a102.4 102.4 0 0 1 102.4 102.4v496.2304a102.4 102.4 0 0 1-102.4 102.4H102.4a102.4 102.4 0 0 1-102.4-102.4V341.3504a102.4 102.4 0 0 1 102.4-102.4z"
                            p-id="7289" fill="#4C4C4C" />
                    </svg>
                </el-upload>
            </span>
            <span>
                <svg t="1673446765797" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
                    p-id="11844" width="24" height="24">
                    <path
                        d="M779.7 588.6c-23.2 0-45.3 4.4-65.7 12.5L282.9 108.3c-10.2-11.6-27.9-12.8-39.5-2.6-11.6 10.2-12.8 27.9-2.6 39.5l424.4 485.1c-39.3 32.8-64.3 82.2-64.3 137.2 0 98.6 80.2 178.8 178.8 178.8 98.6 0 178.8-80.2 178.8-178.8s-80.2-178.9-178.8-178.9z m0 301.7c-67.7 0-122.8-55.1-122.8-122.8 0-67.7 55.1-122.8 122.8-122.8 67.7 0 122.8 55.1 122.8 122.8 0 67.7-55.1 122.8-122.8 122.8z"
                        fill="#999999" p-id="11845"></path>
                    <path
                        d="M779.7 105.6c-11.6-10.2-29.3-9-39.5 2.6L309 601.1c-20.3-8.1-42.5-12.5-65.7-12.5-98.6 0-178.8 80.2-178.8 178.8s80.2 178.8 178.8 178.8c98.6 0 178.8-80.2 178.8-178.8 0-55.1-25-104.4-64.3-137.2l424.4-485.1c10.3-11.6 9.1-29.3-2.5-39.5zM243.4 890.3c-67.7 0-122.8-55.1-122.8-122.8 0-67.7 55.1-122.8 122.8-122.8 67.7 0 122.8 55.1 122.8 122.8 0 67.7-55.1 122.8-122.8 122.8z"
                        fill="#999999" p-id="11846"></path>
                </svg>
            </span>
        </div>
        <div class="text">
            <!-- <div>
                                                <el-scrollbar height="86px">
                                                    <div id="text_input" title="Enter发送，Ctrl+Enter换行" contenteditable="true" @paste="handlePaste"
                                                        @keyup.enter.native="send">
                                                    </div>
                                                </el-scrollbar>
                                            </div> -->
            <textarea id="text_input" placeholder="Enter" v-model="content" @keyup="send"></textarea>
            <div class="btn_container">
                <div class="send_btn" @click="send({ keyCode: 13 })">发送</div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue"
import DiscordPicker from "vue3-discordpicker"

const emits = defineEmits(["sendMessage"])

const state = reactive({ contentType: 1 })
const content = ref("")

const send = async (e: any) => {
    if (e.keyCode === 13 && content.value.length) {
        messageSend('2', content.value)
        content.value = "";
    }


    //支持的节点类型
    //const nodess = []
    //判断节点，分段发送
    // const nodes = document.getElementById('text_input').childNodes
    // if (nodes.length > 1) {
    //     nodes.forEach(function (item, index, nodes) {
    //         //多节点数据处理
    //         if (item.nodeType == 1) {
    //             //目前支持图片，如果不是图片直接抛弃
    //             if (document.getElementById('text_input').querySelector('img')) {
    //                 const src = document.getElementById('text_input').querySelector('img').src
    //                 //发送图片类型
    //                 messageSend('2', src)
    //             }
    //         }
    //         //文本
    //         if (item.nodeType == 3) {
    //             messageSend('1', item.nodeValue)
    //         }
    //     })
    // }
}

const messageSend = (messageType: string, content: string) => {
    emits('sendMessage', {
        content: content,
        dateTime: new Date(),
        fromUid: 1,
        targetUid: 2,
        isSelf: true,
        msgType: messageType,
        msgUID: 2121323232,
        source: 2121323232,
        targetType: 2121323232,
    })
    //清空输入框

}

const setEmoji = (emoji) => {

}


// 粘贴图片并自动提交至接口
const handlePaste = async (e: any) => {
    const { items } = e.clipboardData; // 获取粘贴板文件对象

    var text = document.getElementById('#text_input').innerText

    if ((e.clipboardData || e.originalEvent.clipboardData) && (e.clipboardData || e.originalEvent.clipboardData).getData) {
        var pastedText = (e.originalEvent || e).clipboardData.getData('text/plain')

        console.log(pastedText)
    }


    // if (items.length) {
    //     for (const item of items) {
    //         if (item.kind === 'string') {
    //         }
    //         if (item.kind === 'file') {
    //             //图片还是视频还是其他文件
    //             const file = item.getAsFile(); // 获取图片文件
    //             if (file) {
    //                 if (item.type.substr(0, 5) === 'video') {

    //                 }
    //                 if (item.type.substr(0, 5) === 'image') {

    //                 }
    //             }
    //         }
    //     }
    // }
}


</script>

<style lang="scss">
#emoji-chat {
    .vue3-emojipicker {
        .right-0 {
            right: -200px !important;
        }

        .relative {
            .mt-4 {
                /* margin-top: 1rem; */
                margin-top: 0 !important;
            }
        }

        .vue3-discord-emojipicker {
            height: 340px !important;
            width: 380px !important;
        }

        .vue3-discord-emojipicker {
            background: #ffffff !important;
        }

        .bg-grey-400 {
            --tw-bg-opacity: 1;
            background-color: #ffffff !important;
        }

        .bg-grey-700 {
            background-color: #ffffff !important;
        }

        .bg-grey-600 {
            background-color: #ffffff !important;
        }

        .vue3-discordpicker__container {
            height: 100%;
        }

        .px-5 {
            display: none;
        }

        .overflow-auto {
            &::-webkit-scrollbar {
                width: 5px;
            }

            &::-webkit-scrollbar-thumb {
                height: 3px;
                border-radius: 5px;
                background: #ffffff;
            }
        }
    }
}

#text_input {
    display: flex;
    flex-direction: row;

    img {
        max-width: 100px;
        max-height: 100px;
    }
}

#message_input_container {
    min-width: 1080px;
    min-height: 206px;
    background: #F5F6F7;
    position: relative;

    .tool_list {
        width: 189px;
        height: 60px;
        display: flex;
        align-items: center;


        span {
            margin-left: 21px;
        }
    }

    #text_input {
        min-height: 86px;
        max-height: 86px;
        width: 95%;
        margin: 0 21px 0 20px;
        outline: none;
        font-size: 15px;
        word-break: break-all;
    }

    .btn_container {
        .send_btn {
            position: absolute;
            right: 24px;
            top: 148px;
            width: 92px;
            height: 36px;
            text-align: center;
            line-height: 36px;
            background: #1560FA;
            border-radius: 4px 4px 4px 4px;
            font-size: 14px;
            font-weight: 400;
            color: #FFFFFF;
        }
    }

    svg:hover {
        fill: black;
    }

    textarea {
        font-size: 15px;
        border: none;
        outline: none;
        resize: none;
        font-family: "Micrsofot Yahei";
    }
}
</style>