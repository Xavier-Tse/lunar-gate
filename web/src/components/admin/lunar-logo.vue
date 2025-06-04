<script setup lang="ts">

</script>

<template>
  <div class="lunar-logo">
    <div class="left">
      <img src="https://avatars.githubusercontent.com/u/201599432?s=200&v=4" alt="" />
    </div>
    <div class="right">
      <div>后台管理系统</div>
      <div>developed by XavierTse</div>
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