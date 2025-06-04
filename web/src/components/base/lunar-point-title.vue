<script setup lang="ts">

</script>

<template>
  <div class="lunar-point-title">
    <slot></slot>
  </div>
</template>

<style lang="less">
.lunar-point-title {
  font-weight: 600;
  display: flex;
  align-items: center;
  font-size: 20px;

  &::before {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background-color: rgb(var(--arcoblue-6));
    content: '';
    display: block;
    margin-right: 5px;
  }
}
</style>