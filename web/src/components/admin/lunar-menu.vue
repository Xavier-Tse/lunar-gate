<script setup lang="ts">
import { ref, watch, type Ref } from 'vue';
import { collapsed, type MenuType } from './lunar-menu-variables';
import router from '@/router';
import LunarMenuList from './lunar-menu-list.vue';
import { useRoute } from 'vue-router';

const route = useRoute()

const menuList: Ref<MenuType[]> = ref([
  {title: "数据统计", name: "data", icon: "fluent:data-bar-vertical-20-regular"},
  {
    title: "管理页", name: "manage", icon: "radix-icons:component-1", children: [
      {title: "用户列表", name: "user-list", icon: "fluent:layer-diagonal-person-16-regular"},
      {title: "角色列表", name: "role-list", icon: "fluent:person-key-20-regular"},
      {title: "API列表", name: "api-list", icon: "tabler:api-app"},
      {title: "菜单列表", name: "menu-list", icon: "tabler:menu-order"},
    ]
  },
  {
    title: "系统管理", name: "settings", icon: "radix-icons:gear", children: [
      {
        title: "网站配置", name: "site", icon: "radix-icons:globe", children: [
          {title: "登录配置", name: "login", icon: "radix-icons:enter"},
        ]
      },
      {title: "日志", name: "logs", icon: "radix-icons:reader"},
    ]
  },
])

function select(val: string) {
  router.push({name: val})
}

const selectedKeys = ref<string[]>([])
const openKeys = ref<string[]>([])
function initRoute() {
  selectedKeys.value = []
  openKeys.value = []
  if (route.matched.length >= 1) {
    for (const r of route.matched.slice(1, -1)) {
      openKeys.value.push(r.name as string)
    }
    selectedKeys.value = [route.matched[route.matched.length - 1].name as string]
  }
}
initRoute()

watch(() => route.path, () => {
  initRoute()
}, { immediate: true })
</script>

<template>
  <div class="lunar-menu">
    <a-menu
      show-collapse-button
      v-model:selected-keys="selectedKeys"
      v-model:open-keys="openKeys"
      v-model:collapsed="collapsed"
      @menu-item-click="select"
    >
      <LunarMenuList :list="menuList" />
    </a-menu>
  </div>
</template>

<style lang="less">
.lunar-menu {
  height: calc(100vh - 90px);

  .arco-menu {
    height: 100%;
  }
}
</style>