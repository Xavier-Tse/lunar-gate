<script setup lang="ts">

</script>

<template>
  <div class="login-view">
    <div class="login-container">
      <RouterView />
    </div>
  </div>
</template>

<style lang="less">
.login-view {
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, oklch(68.5% 0.169 237.323), oklch(84.5% 0.143 164.978));
  font-family: "LXGW WenKai", sans-serif;

  .login-container {
    width: 420px;
    padding: 40px;
    background: white;
    border-radius: 12px;
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);

    .title {
      font-size: 24px;
      font-weight: 600;
      color: #1a1a1a;
      text-align: center;
      margin-bottom: 32px;
    }

    .ant-form {
      .ant-form-item {
        margin-bottom: 24px;

        .ant-input,
        .ant-input-password {
          height: 36px;
          border-radius: 6px;
        }
      }

      .ant-btn {
        height: 44px;
        width: 100%;
        font-size: 16px;
        margin-top: 8px;
        border-radius: 6px;
      }

      .other {
        margin-top: 16px;
        text-align: center;

        a {
          color: #1890ff;
          text-decoration: none;
          transition: color 0.3s;

          &:hover {
            color: #40a9ff;
          }
        }
      }
    }
  }
}
</style>
