<script setup lang="ts">
import { reactive } from 'vue'
import LunarCaptcha from '@/components/input/lunar-captcha.vue'
import { useStore } from '@/stores'

const store = useStore()

const form = reactive({
  email: '',
  captchaCode: '',
  captchaID: '',
  emailCode: '',
  password: '',
  confirmPassword: ''
})

function handleCaptchaChange(captchaID: string) {
  form.captchaID = captchaID
}
</script>

<template>
  <div>
    <div class="title">用户注册</div>
      <a-form :model="form" :label-col-props="{span: 7}" :wrapper-col-props="{span: 17}">
        <a-form-item label="邮箱" field="email" :rules="[{ required: true, message: '请输入邮箱' }]">
          <a-input v-model="form.email" placeholder="邮箱" />
        </a-form-item>
        <a-form-item v-if="store.siteInfo.login.captcha.enable" label="图形验证码" field="captchaCode" :rules="[{ required: true, message: '请输入图形验证码' }]">
          <a-input v-model="form.captchaCode" placeholder="图形验证码" />
          <lunar-captcha :on-captcha-change="handleCaptchaChange" />
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