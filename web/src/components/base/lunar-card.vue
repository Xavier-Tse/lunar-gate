<script setup lang="ts">
interface Props {
  title: string;
}
const props = defineProps<Props>()
</script>

<template>
  <div class="card">
    <div class="title">
      <span>{{ props.title }}</span>
    </div>
    <div class="body">
      <slot></slot>
    </div>
  </div>
</template>

<style lang="less">
.card {
  .title {
    font-size: 20px;
    font-weight: 600;
    border-bottom: 1px solid var(--color-neutral-2);
    padding: 20px 20px 10px 20px;
  }

  .body {
    padding: 10px 20px 20px 20px;
  }
}
</style>