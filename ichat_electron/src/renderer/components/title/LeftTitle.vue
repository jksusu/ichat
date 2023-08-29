<template>
    <div class="left_title">
        <!-- <div class="btn close" v-on:click="sendIpc('app-quit')">
            <svg width="10" height="10" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg">
                <g id="cross">
                    <line class="cls-1" x1="7" x2="25" y1="7" y2="25" />
                    <line class="cls-1" x1="7" x2="25" y1="25" y2="7" />
                </g>
            </svg>
        </div>
        <div class="btn min" v-on:click="sendIpc('app-min')">
            <svg width="8" height="8" viewBox="0 0 256 256" xml:space="preserve" xmlns="http://www.w3.org/2000/svg"
                xmlns:xlink="http://www.w3.org/1999/xlink">
                <path
                    d="M208,134.4H48c-3.534,0-6.4-2.866-6.4-6.4s2.866-6.4,6.4-6.4h160c3.534,0,6.4,2.866,6.4,6.4S211.534,134.4,208,134.4z" />
            </svg>
        </div>
        <div class="btn max" v-on:click="sendIpc('app-max')">
            <svg width="6" height="6" viewBox="0 0 25 24" xmlns="http://www.w3.org/2000/svg">
                <g id="Layer_2" data-name="Layer 2">
                    <g id="Layer_1-2" data-name="Layer 1">
                        <path class="cls-1" d="M2,18V2H18" />
                        <path class="cls-1" d="M23,6V22H7" />
                    </g>
                </g>
            </svg>
        </div> -->
    </div>
</template>

<script lang="ts" setup>
import { sendIpc } from '@/common/ipcRenderer';
</script>

<style lang="scss" scoped>
.left_title {
    display: flex;
    float: left;
    margin-top: 12px;
    margin-left: 14px;
    -webkit-app-region: none;

    .btn {
        width: 12px;
        height: 12px;
        border-radius: 100rem;
        margin-right: 6px;
        display: flex;
        align-items: center;
        justify-content: center;

        svg {
            fill: #555;
            stroke-width: 3.6px;
            stroke: #555;
            padding: 1px;
            visibility: hidden;
        }
    }

    .close {
        background-color: #fc605c;
    }

    .max {
        background-color: #34c648;
        margin-right: 0px !important;
    }

    .min {
        background-color: #fdbc40;

        svg {
            stroke-width: 28px !important;
        }
    }
}

.left_title:hover {
    svg {
        visibility: visible;
    }
}
</style>