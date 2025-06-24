<script setup lang="ts">
import { generateOptionsCache } from '@/api';
import { buttonCreateApi, buttonGroupOptionsApi, buttonListApi, type buttonCreateRequest, type buttonType } from '@/api/button-api';
import LunarList from '@/components/admin/lunar-list.vue';
import type { filterGroupType } from '@/types';
import { Message } from '@arco-design/web-vue';
import { reactive, ref } from 'vue';

const lunarListRef = ref()
const visible = ref(false)

const groupOptions = generateOptionsCache(buttonGroupOptionsApi)

const form = reactive<buttonCreateRequest>({
  id: 0,
  name: '',
  title: '',
  group: '',
})

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '名称', dataIndex: 'title' },
  { title: '代号', dataIndex: 'name' },
  { title: '分组', dataIndex: 'group' },
  { title: '创建时间', slotName: 'createdAt' },
  { title: '操作', slotName: 'action' },
]

async function getList(params?: any) {
  lunarListRef.value?.getList(params)
}

function edit(record: buttonType) {
  form.id = record.id
  form.name = record.name
  form.title = record.title
  form.group = record.group
  visible.value = true
}

function add() {
  form.id = 0
  form.name = ''
  form.title = ''
  form.group = ''
  visible.value = true
}

const filterGroup: filterGroupType[] = [
  {
    label: '过滤分组',
    column: 'group',
    source: buttonGroupOptionsApi,
  }
]

async function buttonHandler() {
  const res = await buttonCreateApi(form)
  if (res.code) {
    Message.error(res.message)
    return
  }
  Message.success(res.message)
  lunarListRef.value.getList()
  getList()
}
</script>

<template>
  <div class="button-list-view no-padding">
    <a-modal :title="form.id === 0 ? '创建按钮' : '编辑按钮'" v-model:visible="visible" :on-before-ok="buttonHandler">
      <a-form :model="form">
        <a-form-item label="分组">
          <a-select placeholder="按钮分组" :options="groupOptions" v-model="form.group" allow-create allow-clear />
        </a-form-item>
        <a-form-item label="名称" :rules="[{ required: true, message: '请输入按钮名称' }]" field="title" validate-trigger="blur">
          <a-input placeholder="按钮名称" v-model="form.title" />
        </a-form-item>
        <a-form-item label="代号" :rules="[{ required: true, message: '请输入按钮代号' }]" field="name" validate-trigger="blur">
          <a-input placeholder="按钮代号" v-model="form.name" />
        </a-form-item>
      </a-form>
    </a-modal>
    <LunarList ref="lunarListRef" @add="add" @edit="edit" :columns="columns" :url="buttonListApi" :filter-group="filterGroup"></LunarList>
  </div>
</template>