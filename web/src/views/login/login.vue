<template>
  <div>
    <div class="title">用户登陆</div>
      <a-form :label-col-props="{span: 5}" :wrapper-col-props="{span: 19}">
        <a-form-item label="用户名">
          <a-input placeholder="用户名" />
        </a-form-item>
        <a-form-item label="密码">
          <a-input-password placeholder="密码" />
        </a-form-item>
        <a-form-item label="验证码">
          <a-input placeholder="验证码" />
          <img src="https://example.com/xxx.png" />
        </a-form-item>
        <a-button long type="primary">登陆</a-button>
        <div class="other">
          <RouterLink :to="{ name: 'register' }">注册</RouterLink>
        </div>
      </a-form>
  </div>
</template>