<script setup lang="ts">
import type { MenuType } from './lunar-menu-variables';
import LunarIcon from '../base/lunar-icon.vue';

interface Props {
  list: MenuType[]
}
const props = defineProps<Props>()
</script>

<template>
  <template v-for="item in props.list" :key="item.name">
    <a-menu-item :key="item.name" v-if="!item.children">
      <template #icon>
        <LunarIcon :icon="item.icon" />
      </template>
      {{ item.title }}
    </a-menu-item>
    <a-sub-menu :key="item.name" v-if="item.children">
      <template #icon>
        <LunarIcon :icon="item.icon" />
      </template>
      <template #title>
        {{ item.title }}
      </template>
      <LunarMenuList :list="item.children" />
    </a-sub-menu>
  </template>
</template>

<style lang="less">

</style>