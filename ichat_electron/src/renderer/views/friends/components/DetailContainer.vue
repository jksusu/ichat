<!-- <template>
    <div class="detail_container">
        <PersonalDataCard v-if="{ props.nickname != '' ? true : false }" :nickname="props.nickname"
            :account="props.uid" :message="props.message" :headPortraitUrl='props.headPortraitUrl'>
        </PersonalDataCard>

        <DefaultNullValue v-else></DefaultNullValue>
    </div>
</template>
<script lang="ts" setup>
import { reactive } from 'vue'
import DefaultNullValue from "@/renderer/components/DefaultNullValue.vue";
import PersonalDataCard from "@/renderer/components/personal/PersonalDataCard.vue";

const props = withDefaults(defineProps<{
    uid: string,
    nickname: string
    message: string
    headPortraitUrl: string
    status: number
}>(), {
    uid: "",
    nickname: "",
    message: "",
    headPortraitUrl: "",
    status: 0,
})
if (props.nickname != '') {
    console.log(props.nickname)
}
</script>
<style lang="scss" scoped>
.detail_container {
    height: 100%;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
}
</style> -->