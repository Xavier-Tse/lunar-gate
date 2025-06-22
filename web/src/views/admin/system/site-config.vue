<script setup lang="ts">
import {reactive} from "vue";
import {siteInfoApi, siteInfoUpdateApi, type siteInfoResponse} from "@/api/site-api";
import {Message} from "@arco-design/web-vue";
import LunarPointTitle from "@/components/base/lunar-point-title.vue";
import LunarCutterInput from "@/components/input/lunar-cutter-input.vue";

const data = reactive<siteInfoResponse>({
  site: {
    title: '',
    enTitle: '',
    slogan: '',
    logo: '',
    icp: '',
  },
  project: {
    title: '',
    icon: '',
    path: '',
  },
  login: {
    captcha: {
      enable: false,
      type: '',
    },
  },
})

async function getData() {
  const res = await siteInfoApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  Object.assign(data, res.data)
}

async function siteUpdate() {
  const res = await siteInfoUpdateApi(data)
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
  <div class="site-config-view">
    <div class="site_config">
      <LunarPointTitle>网站设置</LunarPointTitle>
      <div class="body">
        <a-form :model="data.site" :label-col-props="{span: 4}" :wrapper-col-props="{span:20}">
          <a-form-item label="网站标题">
            <a-input v-model="data.site.title" placeholder="请输入网站标题" />
          </a-form-item>
          <a-form-item label="英文标题">
            <a-input v-model="data.site.enTitle" placeholder="请输入英文标题" />
          </a-form-item>
          <a-form-item label="logo">
            <LunarCutterInput placeholder="请输入网站logo" v-model="data.site.logo" :size="50" />
          </a-form-item>
          <a-form-item label="备案号">
            <a-input v-model="data.site.icp" placeholder="请输入网站备案号" />
          </a-form-item>
        </a-form>
      </div>
    </div>
    <div class="project-config">
      <LunarPointTitle>项目设置</LunarPointTitle>
      <div class="body">
        <a-form :model="data.project" :label-col-props="{span: 4}" :wrapper-col-props="{span:20}">
          <a-form-item label="网站title">
            <a-input v-model="data.project.title" placeholder="请输入网站title" />
          </a-form-item>
          <a-form-item label="前端地址">
            <a-input v-model="data.project.path" placeholder="请输入前端地址" />
          </a-form-item>
        </a-form>
      </div>
    </div>
    <div class="login-config">
      <LunarPointTitle>登录设置</LunarPointTitle>
      <div class="body">
        <a-form :model="data.login.captcha" :label-col-props="{span: 6}" :wrapper-col-props="{span:18}">
          <a-form-item label="启用图片验证码">
            <a-switch v-model="data.login.captcha.enable" />
          </a-form-item>
          <a-form-item label="图片验证码类型">
            <a-select v-model="data.login.captcha.type">
              <a-option>string</a-option>
              <a-option>math</a-option>
            </a-select>
          </a-form-item>
        </a-form>
      </div>
    </div>
    <a-button @click="siteUpdate" type="primary">更新</a-button>
  </div>
</template>

<style lang="less">
.site-config-view {
  > div {
    width: 500px;
  }

  .body {
    margin-top: 10px;

    .arco-form-item {
      margin-bottom: 10px;
    }
  }

  .project-config {
    margin-top: 20px;
  }

  .login-config {
    margin-top: 20px;
  }
}
</style>