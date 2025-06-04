<script setup lang="ts">
import { Icon } from "@iconify/vue";

interface Props {
  icon: string
}
const props = defineProps<Props>()

</script>

<template>
  <div class="lunar-icon">
    <Icon :icon="props.icon" width="16" height="16" />
  </div>
</template>

<style lang="less">
.lunar-icon {
  padding-top: 3px;
}
</style>