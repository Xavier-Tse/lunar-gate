<script setup lang="ts">
import { reactive, ref } from 'vue'
import LunarCaptcha from '@/components/input/lunar-captcha.vue'
import { useStore } from '@/stores'
import { Message, type FormInstance } from '@arco-design/web-vue'
import { emailSendApi } from '@/api/email-api'
import { userRegisterApi } from '@/api/user-api'
import router from '@/router'

const store = useStore()
const formRef = ref<FormInstance>()

const form = reactive({
  email: '',
  emailCode: '',
  emailID: '',
  password: '',
  rePassword: '',
  captcha: '',
  captchaID: '',
})

function pwdValidator(value: string | undefined, callback: (error?: string) => void) {
  if (value !== form.password) {
    callback('两次输入的密码不一致')
  }
}

function handleCaptchaChange(captchaID: string) {
  form.captchaID = captchaID
}

async function handleSendEmail() {
  const val = await formRef.value.validateField(['email', 'captcha'])
  if (val) {
    return
  }
  const res = await emailSendApi({
    email: form.email,
    captchaCode: form.captcha,
    captchaID: form.captchaID,
  })
  if (res.code) {
    Message.error(res.message)
    return
  }
  form.emailID = res.data!.emailID
}

async function handleRegister() {
  const res = await userRegisterApi(form)
  if (res.code) {
    Message.error(res.message)
    return
  }
  Message.success(res.message)
  router.push({ name: 'login' })
}
</script>

<template>
  <div>
    <div class="title">用户注册</div>
      <a-form :model="form" ref="formRef" :label-col-props="{span: 7}" :wrapper-col-props="{span: 17}">
        <a-form-item label="邮箱" field="email" validate-trigger="blur" :rules="[{ required: true, message: '请输入邮箱' }, { type: 'email', message: '请输入正确的邮箱格式' }]">
          <a-input v-model="form.email" placeholder="邮箱" />
        </a-form-item>
        <a-form-item v-if="store.siteInfo.login.captcha.enable" label="图形验证码" field="captcha" validate-trigger="blur" :rules="[{ required: true, message: '请输入图形验证码' }]">
          <a-input v-model="form.captcha" placeholder="图形验证码" />
          <lunar-captcha :on-captcha-change="handleCaptchaChange" />
        </a-form-item>
        <a-form-item label="邮箱验证码" field="emailCode" validate-trigger="blur" :rules="[{ required: true, message: '请输入邮箱验证码' }]">
          <a-input v-model="form.emailCode" placeholder="邮箱验证码" />
          <a-button type="primary" style="margin-left: 10px;" @click="handleSendEmail">校验邮箱</a-button>
        </a-form-item>
        <a-form-item label="密码" field="password" validate-trigger="blur" :rules="[{ required: true, message: '请输入密码' }]">
          <a-input-password v-model="form.password" placeholder="密码" />
        </a-form-item>
        <a-form-item label="确认密码" field="rePassword" validate-trigger="blur" :rules="[{ required: true, message: '请再次输入密码' }, { validator: pwdValidator }]">
          <a-input-password v-model="form.rePassword" placeholder="确认密码" />
        </a-form-item>
        <a-button long type="primary" @click="handleRegister">注册</a-button>
        <div class="other">
          <RouterLink :to="{ name: 'login' }">去登陆</RouterLink>
        </div>
      </a-form>
  </div>
</template>