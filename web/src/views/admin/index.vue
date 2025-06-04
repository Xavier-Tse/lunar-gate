<script setup lang="ts">
import Breadcrumb from "@/components/admin/breadcrumb.vue";
import LunarLogo from "@/components/admin/lunar-logo.vue";
import { collapsed } from "@/components/admin/lunar-menu-variables";
import LunarMenu from "@/components/admin/lunar-menu.vue";
import LunarTabs from "@/components/admin/lunar-tabs.vue";
import FullScreen from "@/components/base/full-screen.vue";
import Theme from "@/components/base/theme.vue";
import UserDropdown from "@/components/base/user-dropdown.vue";
import router from "@/router";
import { IconHome } from "@arco-design/web-vue/es/icon";

function goHome() {
  router.push({ name: 'web' })
}
</script>

<template>
  <div class="admin-view">
    <div class="aside" :class="{ collapsed: collapsed }">
      <LunarLogo />
      <LunarMenu />
    </div>
    <div class="container">
      <div class="head">
        <div class="crumbs">
          <Breadcrumb />
        </div>
        <div class="actions">
          <span>
            <IconHome @click="goHome" />
          </span>
          <FullScreen />
          <Theme />
          <UserDropdown />
        </div>
      </div>
      <LunarTabs />
      <div class="main">
        <RouterView class="base-view" v-slot="{ Component: component }">
          <Transition name="fade" mode="out-in">
            <component :is="component" />
          </Transition>
        </RouterView>
      </div>
    </div>
  </div>
</template>

<style lang="less">
.admin-view {
  font-family: "LXGW WenKai", sans-serif;
  display: flex;
  width: 100%;
  background-color: var(--color-bg-1);
  color: var(--color-text-1);
  overflow: hidden;

  .aside {
    width: 240px;
    border-right: 1px solid var(--color-neutral-2);
    height: 100vh;
    transition: all 0.3s ease;

    .lunar-logo {
      border-bottom: 1px solid var(--color-neutral-2);
      height: 90px;
      display: flex;
      padding: 20px;
    }

    .lunar-menu {
      padding: 10px 0;
    }
  }

  .aside.collapsed {
    width: 48px;

    &~.container {
      width: calc(100% - 48px);
    }
  }

  >.container {
    width: calc(100% - 240px);
    transition: all 0.3s ease;

    >.head {
      border-bottom: 1px solid var(--color-neutral-2);
      display: flex;
      justify-content: space-between;
      height: 60px;
      padding: 0 20px;
      align-items: center;

      .actions {
        >span {
          margin: 0 10px;
          cursor: pointer;

          &:last-child {
            margin-right: 0;
          }

          &:hover {
            color: rgb(var(--arcoblue-6));
          }
        }
      }
    }

    .lunar-tabs{
      height: 30px;
      border-bottom: 1px solid var(--color-neutral-2);
      display: flex;
      align-items: center;
      padding: 0 20px;
    }

    .main{
      padding: 20px;
      background-color: var(--color-neutral-2);
      height: calc(100vh - 90px);
    }
  }
}

.base-view {
  background-color: var(--color-bg-1);
  border-radius: 5px;
  padding: 20px;
  height: calc(100vh - 130px);
}


// 组件刚开始离开
// .fade-leave-active {
// }

// 组件离开结束
.fade-leave-to {
  transform: translateX(30px);
  opacity: 0;
}

// 组件刚开始进入
.fade-enter-active {
  transform: translateX(-30px);
  opacity: 0;
}

// 组件进入完成
.fade-enter-to {
  transform: translateX(0);
  opacity: 1;
}

// 正在进入和离开
.fade-leave-active, .fade-enter-active {
  transition: all 0.3s ease-out;
}
</style>