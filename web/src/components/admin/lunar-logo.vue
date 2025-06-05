<script setup lang="ts">
import { useStore } from '@/stores';

const store = useStore()
</script>

<template>
  <div class="lunar-logo">
    <div class="left">
      <img :src="store.siteInfo.site.logo" alt="logo" />
    </div>
    <div class="right">
      <div>{{ store.siteInfo.site.title }}</div>
      <div>{{ store.siteInfo.site.enTitle }}</div>
    </div>
  </div>
</template>

<style lang="less">
.lunar-logo {
  .left{
    width: 80px;
    display: flex;
    justify-content: center;
    transition: all 0.3s ease;
    align-items: center;

    img {
      width: 50px;
      height: 50px;
      transition: all 0.3s ease;
    }
  }

  .right {
    display: flex;
    flex-direction: column;
    justify-content: center;

    div:nth-child(1) {
      font-size: 17px;
      font-weight: 600;
    }

    div:nth-child(2) {
      font-size: 12px;
    }
  }
}

.collapsed {
  .lunar-logo {
    padding: 5px!important;

    .left {
      align-items: center;

      img {
        width: 30px;
        height: 30px;
      }
    }

    .right {
      display: none;
    }
  }
}
</style>