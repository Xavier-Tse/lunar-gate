<script setup lang="ts">
import LunarCard from '@/components/base/lunar-card.vue';
import LunarPointTitle from '@/components/base/lunar-point-title.vue';


</script>

<template>
  <div class="data-view">
    <div class="left">
      <div class="sector greeting">
        <div>早安，${username}，请开始一天的工作吧。</div>
        <div>lunar admin vue后台管理系统</div>
        <div class="data">
          <span>用户总数：114514</span>
          <span>今日登陆：1919</span>
          <span>今日注册：810</span>
        </div>
      </div>
      <div class="sector charts">
        <div class="title">
          <LunarPointTitle>最近7日数据</LunarPointTitle>
        </div>
        <div class="body"></div>
      </div>
    </div>
    <div class="right">
      <div class="sector quick">
        <LunarCard title="快捷操作"></LunarCard>
      </div>
      <div class="sector version">
        <LunarCard title="更新日志"></LunarCard>
      </div>
      <div class="sector system">
        <div class="title">
          <LunarPointTitle>系统信息</LunarPointTitle>
        </div>
        <div class="body"></div>
      </div>
    </div>
  </div>
</template>

<style lang="less">
.data-view {
  display: flex;
  justify-content: space-between;
  background-color: transparent!important;
  padding: 0!important;

  .left {
    width: calc(100% - 320px);

    .greeting {
      padding: 20px;

      div:nth-child(1) {
        font-size: 18px;
        font-weight: 600;
      }

      div:nth-child(2) {
        font-size: 16px;
        margin-top: 10px;
      }

      .data {
        margin-top: 20px;
        
        span {
          margin-right: 20px;
          color: var(--color-text-2);
        }
      }
    }

    .charts {
      margin-top: 20px;
      padding: 20px;

      .body {
        min-height: calc(100vh - 370px);
      }
    }
  }

  .right {
    width: 300px;

    .quick {
      .body {
        min-height: 120px;
      }
    }

    .version {
      margin-top: 20px;

      .body {
        min-height: 200px;
      }
    }

    .system {
      margin-top: 20px;

      .title {
        padding: 20px 20px 10px 20px;
      }

      .body {
        min-height: 300px;
      }
    }
  }

  .sector {
    background-color: var(--color-bg-1);
    border-radius: 5px;
  }
}
</style>