<script setup lang="ts">
import LunarImgCutter from '../base/lunar-img-cutter.vue';

interface Props {
  modelValue: string
  size?: number
  placeholder: string
}
const props = defineProps<Props>()
const emits = defineEmits(['update:modelValue'])

function cutterOk(val: string) {
  emits('update:modelValue', val)
}

function inputChange(val: string) {
  emits('update:modelValue', val)
}
</script>

<template>
  <div class="lunar-cutter-input">
    <a-input @input="inputChange" @change="inputChange" :model-value="props.modelValue" :placeholder="props.placeholder" />
    <LunarImgCutter class="lunar-img-cutter" @ok="cutterOk" :model-value="props.modelValue" :size="props.size" />
  </div>
</template>

<style lang="less">
.lunar-cutter-input {
  width: 100%;

  .lunar-img-cutter {
    margin-top: 10px;
  }
}
</style>