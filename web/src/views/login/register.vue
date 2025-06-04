<script setup lang="ts">

</script>

<template>
  <div>
    <div class="title">用户注册</div>
      <a-form :label-col-props="{span: 7}" :wrapper-col-props="{span: 17}">
        <a-form-item label="邮箱">
          <a-input placeholder="邮箱" />
        </a-form-item>
        <a-form-item label="图形验证码">
          <a-input placeholder="图形验证码" />
          <img src="https://example.com/xxx.png" />
        </a-form-item>
        <a-form-item label="邮箱验证码">
          <a-input placeholder="邮箱验证码" />
        </a-form-item>
        <a-form-item label="密码">
          <a-input-password placeholder="密码" />
        </a-form-item>
        <a-form-item label="确认密码">
          <a-input-password placeholder="确认密码" />
        </a-form-item>
        <a-button long type="primary">注册</a-button>
        <div class="other">
          <RouterLink :to="{ name: 'login' }">去登陆</RouterLink>
        </div>
      </a-form>
  </div>
</template>