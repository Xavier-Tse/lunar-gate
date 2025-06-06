<script setup lang="ts">
import type { optionColorType } from '@/options';
import { computed } from 'vue';

interface Props {
  options: optionColorType[]
  value: string
}
const props = defineProps<Props>()

const item = computed(() => {
  const item = props.options.find((value) => value.value === props.value)
  return item
})
</script>

<template>
  <a-tag :color="item?.color">{{ item?.label }}</a-tag>
</template>

<style lang="less">

</style>