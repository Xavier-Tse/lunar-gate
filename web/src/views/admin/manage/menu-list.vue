<script setup lang="ts">
import { menuCreateApi, menuListApi, type menuCreateRequest, type menuType } from '@/api/menu-api';
import LunarList from '@/components/admin/lunar-list.vue';
import LunarIcon from '@/components/base/lunar-icon.vue';
import { Message } from '@arco-design/web-vue';
import { reactive, ref } from 'vue';

const lunarListRef = ref()
const formRef = ref()
const visible = ref(false)

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '名称', dataIndex: 'name' },
  { title: '菜单名称', dataIndex: 'meta.title' },
  { title: '图标', slotName: 'icon' },
  { title: '路由', dataIndex: 'path' },
  { title: '是否显示', dataIndex: 'enable', type: 'switch' },
  { title: '排序', dataIndex: 'sort' },
  { title: '创建时间', slotName: 'createdAt' },
  { title: '操作', slotName: 'action' },
]

const form = reactive<menuCreateRequest>({
  id: 0,
  icon: "",
  title: "",
  enable: false,
  name: "",
  path: "",
  component: "",
  parentMenuID: undefined,
  sort: 0
})

function edit(record: menuType) {
  visible.value = true
}

function add() {
  form.id = 0
  form.parentMenuID = undefined
  visible.value = true
}

function addSubMenu(record: menuType) {
  form.id = 0
  form.parentMenuID = record.id
  visible.value = true
}

function cleanForm() {
  form.icon = ''
  form.title = ''
  form.name = ''
  form.path = ''
  form.component = ''
}

async function createMenuHandler() {
  const val = await formRef.value.validate()
  if (val) {
    return
  }
  const res = await menuCreateApi(form)
  if (res.code) {
    Message.error(res.message)
    return
  }
  Message.success(res.message)
  lunarListRef.value.getList()
  cleanForm()
}

</script>

<template>
  <div class="menu-list-view no-padding">
    <a-modal v-model:visible="visible" :title="form.id === 0 ? !form.parentMenuID ? '创建根菜单' : '创建子菜单' : '更新菜单'" :on-before-ok="createMenuHandler">
      <a-form :model="form" ref="formRef">
        <a-form-item label="菜单名称" field="title" :rules="[{ required: true, message: '请输入菜单名称' }]" validate-trigger="blur">
          <a-input placeholder="请输入菜单名称" v-model="form.title" />
        </a-form-item>
        <a-form-item label="路由名称" field="name" :rules="[{ required: true, message: '请输入路由名称' }]" validate-trigger="blur">
          <a-input placeholder="请输入路由名称" v-model="form.name" />
        </a-form-item>
        <a-form-item label="图标" field="icon">
          <a-input placeholder="请输入图标" v-model="form.icon" />
          <template #help>
            仅支持<a href="https://icon-sets.iconify.design/" target="_blank">iconify</a>图标，例如"radix-icons:gear"
          </template>
        </a-form-item>
        <a-form-item label="路由路径" field="path" :rules="[{ required: true, message: '请输入路由路径' }]" validate-trigger="blur">
          <a-input placeholder="请输入路由路径" v-model="form.path" />
        </a-form-item>
        <a-form-item label="组件" field="component">
          <a-input placeholder="请输入组件" v-model="form.component" />
        </a-form-item>
        <a-form-item label="是否显示" field="enable">
          <a-switch v-model="form.enable" />
        </a-form-item>
        <a-form-item label="排序" field="sort">
          <a-input-number placeholder="排序" v-model="form.sort" />
        </a-form-item>
      </a-form>
    </a-modal>
    <LunarList ref="lunarListRef" add-label="添加根菜单" @add="add" @edit="edit" :columns="columns" :url="menuListApi">
      <template #icon="{record}: {record: menuType}">
        <div v-if="record.meta.icon" style="display: flex;">
          <LunarIcon :icon="record.meta.icon" />
          <code style="padding-left: 5px;">{{ record.meta.icon }}</code>
        </div>
        <span v-else>未设置图标</span>
      </template>
      <template #action-left="{record}: {record: menuType}">
        <a-button type="outline" @click="addSubMenu(record)">创建子菜单</a-button>
      </template>
    </LunarList>
  </div>
</template>