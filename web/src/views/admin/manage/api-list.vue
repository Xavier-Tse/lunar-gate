<script setup lang="ts">
import { apiCreateApi, apiListApi, type apiCreateRequest } from '@/api/api-api';
import { type userListResponse } from '@/api/user-api';
import LunarList from '@/components/admin/lunar-list.vue';
import { methodOptions } from '@/options';
import { Message } from '@arco-design/web-vue';
import { reactive, ref } from 'vue';

const lunarListRef = ref()
const visible = ref(false)

const form = reactive<apiCreateRequest>({
  name: '',
  path: '',
  group: '',
  method: '',
})

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '路径', dataIndex: 'path' },
  { title: '分组', dataIndex: 'group' },
  { title: '方法', dataIndex: 'method', type: 'options', options: methodOptions },
  { title: '创建时间', slotName: 'createdAt' },
  { title: '操作', slotName: 'action' },
]

function edit(record: userListResponse) {
  visible.value = true
}

async function apiHandler() {
  const res = await apiCreateApi(form)
  if (res.code) {
    Message.error(res.message)
    return
  }
  Message.success(res.message)
  lunarListRef.value.getList()
}

</script>

<template>
  <div class="api-list-view no-padding">
    <a-modal title="创建API" v-model:visible="visible" :on-before-ok="apiHandler">
      <a-form :model="form">
        <a-form-item label="分组">
          <a-select placeholder="API分组" v-model="form.group" allow-create allow-clear />
        </a-form-item>
        <a-form-item label="名称" :rules="[{ required: true, message: '请输入API名称' }]" field="name" validate-trigger="blur">
          <a-input placeholder="API名称" v-model="form.name" />
        </a-form-item>
        <a-form-item label="路径" :rules="[{ required: true, message: '请输入API路径' }]" field="path" validate-trigger="blur">
          <a-input placeholder="API路径" v-model="form.path" />
        </a-form-item>
        <a-form-item label="方法" :rules="[{ required: true, message: '请选择API方法' }]" field="method" validate-trigger="blur">
          <a-select placeholder="API方法" :options="methodOptions" v-model="form.method" />
        </a-form-item>
      </a-form>
    </a-modal>
    <LunarList ref="lunarListRef" @add="visible=true" :columns="columns" :url="apiListApi" @edit="edit"></LunarList>
  </div>
</template>

<style lang="less">
.api-list-view {
}
</style>