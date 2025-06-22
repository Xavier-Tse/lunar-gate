<script setup lang="ts">
import {reactive} from "vue";
import {Message} from "@arco-design/web-vue";
import LunarPointTitle from "@/components/base/lunar-point-title.vue";
import { emailInfoApi, emailInfoUpdateApi, type emailInfoType } from "@/api/site-api";

const data = reactive<emailInfoType>({
  enable: false,
  domain: '',
  port: 0,
  sendEmail: '',
  authCode: '',
  sendNickname: '',
})

async function getData() {
  const res = await emailInfoApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  Object.assign(data, res.data)
}

async function emailUpdate() {
  const res = await emailInfoUpdateApi(data)
  if (res.code) {
    Message.error(res.message)
    return
  }
  Message.success(res.message)
  getData()
}
getData()
</script>

<template>
  <div class="email-config-view">
    <div class="email-config">
      <LunarPointTitle>邮箱设置</LunarPointTitle>
      <div class="body">
        <a-form :model="data" :label-col-props="{span: 4}" :wrapper-col-props="{span:20}">
          <a-form-item label="是否启用">
            <a-switch v-model="data.enable" />
          </a-form-item>
          <a-form-item label="host">
            <a-input v-model="data.domain" placeholder="请输入host" />
          </a-form-item>
          <a-form-item label="port">
            <a-input-number v-model="data.port" placeholder="请输入port" />
          </a-form-item>
          <a-form-item label="发件邮箱">
            <a-input placeholder="请输入发件邮箱" v-model="data.sendEmail" :size="50" />
          </a-form-item>
          <a-form-item label="发件昵称">
            <a-input v-model="data.sendNickname" placeholder="请输入发件昵称" />
          </a-form-item>
          <a-form-item label="授权码">
            <a-input v-model="data.authCode" placeholder="请输入授权码" />
          </a-form-item>
        </a-form>
      </div>
      <a-button @click="emailUpdate" type="primary">更新</a-button>
    </div>
  </div>
</template>

<style lang="less">
.email-config-view {
  .email-config {
    width: 500px;
    .body{
      margin-top: 10px;
    }
  }
}
</style>