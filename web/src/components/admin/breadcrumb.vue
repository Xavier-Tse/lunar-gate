<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()
const list = ref<string[]>([])

function initRoute() {
  list.value = []
  for (const r of route.matched) {
    if (!r.meta.title) {
      continue
    }
    list.value.push(r.meta.title as string)
  }
}
initRoute()

watch(() => route.path, () => {
  initRoute()
}, {immediate: true})
</script>

<template>
  <a-breadcrumb>
    <a-breadcrumb-item v-for="item in list">
      <span>{{item}}</span>
    </a-breadcrumb-item>
  </a-breadcrumb>
</template>

<style lang="less">

</style>