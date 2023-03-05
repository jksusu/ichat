<template>
  <div class="login_method">
    <div class="login_method_tips"><img style="float: right;" :src="props.tripUrl" /></div>
    <img v-on:click="jumpTo(props.route)" class="login_method_icon" :src="props.iconUrl" />
  </div>
  <div class="login_logo">
    <img class="logo" :src="LogoUrl" />
    <div class="cn">IM即时通讯</div>
    <div class="en">im instant messaging</div>
  </div>
</template>
<script lang="ts" setup>
import { jump } from '@/renderer/tool/tool';
import LogoUrl from '@/renderer/assets/images/logo.png'
const props = defineProps<{ route: string, iconUrl: string, tripUrl: string }>()
const jumpTo = (route: string) => {
  jump(route)
}
</script>

<style lang="scss">
:root {
  --loginMethodTipsShow: hidden;
}

.login_method {
  display: flex;
  width: 107px;
  height: 30px;
  margin: 17px 30px 0 223px;
  -webkit-user-drag: none;

  .login_method_tips {
    width: 77px;
    height: 19px;
    visibility: var(--loginMethodTipsShow, hidden);
  }

  .login_method_icon {
    width: 30px;
    height: 30px;
  }

  &:hover .login_method_tips {
    --loginMethodTipsShow: visible
  }
}

.login_logo {
  height: 123px;
  width: 189px;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 0 83px 0 88px;
  user-select: none;
  -webkit-user-drag: none;
  
  .logo {
    width: 55px;
    height: 55px;
  }

  .cn {
    margin-top: 20px;
    width: 158px;
    height: 31px;
    font-size: 22px;
    font-weight: 500;
    color: #333333;
    line-height: 26px;
    text-align: center;
  }

  .en {
    height: 17px;
    width: 189px;
    font-size: 12px;
    font-weight: 400;
    color: #999999;
    line-height: 14px;
    letter-spacing: 1px;
    text-align: center;
  }
}
</style>