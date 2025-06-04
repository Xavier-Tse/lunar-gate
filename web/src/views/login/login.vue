<script setup lang="ts">
import { userLogin, type userLoginRequest } from '@/api/user-api';
import { useStore } from '@/stores';
import { Message, type FormInstance } from '@arco-design/web-vue';
import { reactive, ref } from 'vue';

const form = reactive<userLoginRequest>({
  username: '',
  password: '',
  captchaID: '',
  captchaCode: '',
})

const store = useStore()
const formRef = ref<FormInstance>()

async function handleLogin() {
  const val = formRef.value.validate()
  if (!val) {
    return
  }
  const res = await userLogin(form)
  if (res.code) {
    Message.error(res.message)
    return
  }
  Message.success(res.message)
  store.saveUser(res.data!.token)
}
</script>

<template>
  <div>
    <div class="title">用户登陆</div>
      <a-form :model="form" ref="formRef" :label-col-props="{span: 5}" :wrapper-col-props="{span: 19}">
        <a-form-item label="用户名" field="username" :rules="[{ required: true, message: '请输入用户名' }]" validate-trigger="blur">
          <a-input v-model="form.username" placeholder="用户名" />
        </a-form-item>
        <a-form-item label="密码" field="password" :rules="[{ required: true, message: '请输入密码' }]" validate-trigger="blur">
          <a-input-password v-model="form.password" placeholder="密码" />
        </a-form-item>
        <a-form-item label="验证码" field="captchaCode" :rules="[{ required: true, message: '请输入验证码' }]" validate-trigger="blur">
          <a-input v-model="form.captchaCode" placeholder="验证码" />
          <img src="https://example.com/xxx.png" />
        </a-form-item>
        <a-button long type="primary" @click="handleLogin">登陆</a-button>
        <div class="other">
          <RouterLink :to="{ name: 'register' }">注册</RouterLink>
        </div>
      </a-form>
  </div>
</template>