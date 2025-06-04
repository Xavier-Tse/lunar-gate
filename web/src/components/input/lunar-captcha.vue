<script setup lang="ts">
import { captchaGenerateApi } from '@/api/captcha-api'
import { Message } from '@arco-design/web-vue'
import { ref } from 'vue'

const props = defineProps<{
  onCaptchaChange?: (captchaID: string) => void
}>()

const captcha = ref<string>('')
const captchaID = ref<string>('')

async function getImage() {
  const res = await captchaGenerateApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  captcha.value = res.data!.captcha
  captchaID.value = res.data!.captchaID
  props.onCaptchaChange?.(res.data!.captchaID)
}

defineExpose({
  refresh: getImage
})
getImage()
</script>

<template>
  <img class="lunar-captcha" :src="captcha" alt="图片验证码" @click="getImage" />
</template>

<style lang="less">
.lunar-captcha {
  height: 36px;
  width: 120;
  object-fit: contain;
  margin-left: 8px;
  border-radius: 3px;
  cursor: pointer;
  transition: opacity 0.3s;

  &:hover {
    opacity: 0.8;
  }
}
</style>