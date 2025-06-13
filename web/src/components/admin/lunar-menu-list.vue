<script setup lang="ts">
import LunarIcon from '../base/lunar-icon.vue';
import type { menuType } from '@/api/menu-api';

interface Props {
  list: menuType[]
}
const props = defineProps<Props>()
</script>

<template>
  <template v-for="item in props.list" :key="item.name">
    <a-menu-item :key="item.name" v-if="!item.children">
      <template #icon>
        <LunarIcon :icon="item.meta.icon" />
      </template>
      {{ item.meta.title }}
    </a-menu-item>
    <a-sub-menu :key="item.name" v-if="item.children">
      <template #icon>
        <LunarIcon :icon="item.meta.icon" />
      </template>
      <template #title>
        {{ item.meta.title }}
      </template>
      <LunarMenuList :list="item.children" />
    </a-sub-menu>
  </template>
</template>

<style lang="less">

</style>