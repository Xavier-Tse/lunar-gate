<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const goBack = () => {
  router.back()
}

const goHome = () => {
  router.push({ name: 'data' })
}
</script>

<template>
  <div class="not-found">
    <div class="container">
      <div class="image">404</div>
      <div class="title">叽里咕噜找什么呢</div>
      <div class="actions">
        <a-button type="primary" @click="goBack">
          <template #icon><icon-undo /></template>
          返回上页
        </a-button>
        <a-button @click="goHome">
          <template #icon><icon-home /></template>
          回到首页
        </a-button>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
.not-found {
  font-family: "LXGW WenKai", sans-serif;
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--color-bg-1);

  .container {
    text-align: center;
    
    .image {
      font-size: 120px;
      font-weight: bold;
      color: var(--color-text-1);
      margin-bottom: 24px;
      animation: bounce 2s ease infinite;
    }

    .title {
      font-size: 24px;
      color: var(--color-text-2);
      margin-bottom: 32px;
    }

    .actions {
      .arco-btn {
        margin: 0 8px;
      }
    }
  }
}

@keyframes bounce {
  0%, 20%, 50%, 80%, 100% {
    transform: translateY(0);
  }
  40% {
    transform: translateY(-30px);
  }
  60% {
    transform: translateY(-15px);
  }
}
</style>