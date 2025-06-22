<script setup lang="ts">
import ImgCutter from 'vue-img-cutter'
import {Message} from "@arco-design/web-vue";
import { imageUploadApi } from '@/api/image-api';

interface Props {
  modelValue: string
  rate?: string
  size?: number
}

const props = defineProps<Props>()
const {rate = "1:1", size = 90} = props
const emits = defineEmits(["update:modelValue", 'ok'])

async function cutDown(e: any) {
  const file = e.file as File
  const res = await imageUploadApi(file)
  if (res.code) {
    Message.error(res.message)
    return
  }
  emits("update:modelValue", res.data)
  emits('ok')
}


</script>

<template>
  <div>
    <ImgCutter :rate="rate" :toolBoxOverflow="false" @cutDown="cutDown">
      <template #open>
        <a-avatar :size="size" :image-url="props.modelValue" />
      </template>
    </ImgCutter>
  </div>
</template>