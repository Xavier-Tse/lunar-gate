<script setup lang="ts">
import { useStore } from "@/stores";
import {Message} from "@arco-design/web-vue";
import {useRouter} from "vue-router";

const router = useRouter()
const store = useStore()

function handleSelect(val: string) {
  if (val === "logout"){
    store.logout()
    return
  }
  router.push({name: val})
}
</script>

<template>
  <a-dropdown @select="handleSelect">
    <span class="user-dropdown" >
      <a-avatar :size="35" :image-url="store.userInfo.avatar" />
      <span>{{ store.userInfo.nickname }}</span>
      <icon-down />
    </span>
    <template #content>
      <a-doption value="user-list">用户列表</a-doption>
      <a-doption value="role-list">角色列表</a-doption>
      <a-doption value="logout" style="color: red">退出登录</a-doption>
    </template>
  </a-dropdown>
</template>

<style lang="less">
.user-dropdown{
  >span{
    margin-left: 10px;
  }
}
</style>