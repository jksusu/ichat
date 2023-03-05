<template>
    <div>
        <el-button></el-button>
    </div>
</template>

<script lang="ts" setup>
import { ElMessage } from 'element-plus'
import { type } from 'os';

const open2 = () => {
    ElMessage({
        message: 'Congrats, this is a success message.',
        type: 'success',
    })
}
</script>

<style lang="scss">

</style>