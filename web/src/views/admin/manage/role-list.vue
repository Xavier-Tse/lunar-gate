<script setup lang="ts">
import { roleCreateApi, roleListApi, type roleCreateRequest, type roleType } from '@/api/role-api';
import LunarList from '@/components/admin/lunar-list.vue';
import LunarRoleMenuModal from '@/components/admin/lunar-role-menu-modal.vue';
import { Message } from '@arco-design/web-vue';
import { reactive, ref } from 'vue';

const lunarListRef = ref()
const formRef = ref()
const visible = ref(false)
const menuVisible = ref(false)

const form = reactive<roleCreateRequest>({
  id: 0,
  title: '',
})

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '角色名称', dataIndex: 'title' },
  { title: '角色API数', dataIndex: 'roleApiCount' },
  { title: '角色用户数', dataIndex: 'roleUserCount' },
  { title: '创建时间', slotName: 'createdAt' },
  { title: '更新时间', dataIndex: 'updatedAt', type: 'datetime' },
  { title: '操作', slotName: 'action' },
]

function edit(record: roleType) {
  form.title = record.title
  form.id = record.id
  visible.value = true
}

function add() {
  form.id = 0
  visible.value = true
}

function cleanForm() {
  form.id = 0
  form.title = ''
}

async function handler() {
  const val = await formRef.value.validate()
  if (val) {
    return
  }
  const res = await roleCreateApi(form)
  if (res.code) {
    Message.error(res.message)
    return
  }
  Message.success(res.message)
  lunarListRef.value.getList()
}

function showModal(record: roleType) {
  menuVisible.value = true
}

</script>

<template>
  <div class="user-list-view no-padding">
    <a-modal @cancle="cleanForm" :on-before-ok="handler" :title="form.id === 0 ? '创建角色' : '编辑角色'" v-model:visible="visible">
      <a-form :model="form" ref="formRef">
        <a-form-item label="角色名称" validate-trigger="blur" field="title" :rules="[{ required: true, message: '请输入角色名称' }]">
          <a-input v-model="form.title" placeholder="角色名称" allow-clear />
        </a-form-item>
      </a-form>
    </a-modal>
    <LunarRoleMenuModal v-model:visible="menuVisible"></LunarRoleMenuModal>
    <LunarList ref="lunarListRef" :columns="columns" :url="roleListApi" @edit="edit" @add="add">
      <template #action-left="{record}: {record: roleType}">
        <a-button @click="showModal(record)" type="outline">设置菜单</a-button>
        <a-button type="outline">设置API</a-button>
      </template>
    </LunarList>
  </div>
</template>

<style lang="less">
.user-list-view {
  padding: 0!important;

  .role-list {
    display: grid;
    grid-template-columns: 1fr;
    grid-row-gap: 10px;

    >span {
      width: fit-content;
    }
  }
}
</style>