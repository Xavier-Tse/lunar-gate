<script setup lang="ts">
import { roleCreateApi, roleListApi, type roleCreateRequest, type roleType } from '@/api/role-api';
import LunarList from '@/components/admin/lunar-list.vue';
import LunarRoleApiModal from '@/components/admin/lunar-role-api-modal.vue';
import LunarRoleMenuModal from '@/components/admin/lunar-role-menu-modal.vue';
import { Message } from '@arco-design/web-vue';
import { reactive, ref } from 'vue';

const lunarListRef = ref()
const formRef = ref()
const visible = ref(false)
const menuVisible = ref(false)
const apiVisible = ref(false)

const form = reactive<roleCreateRequest>({
  id: 0,
  title: '',
})

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '角色名称', dataIndex: 'title' },
  { title: '角色API数', dataIndex: 'roleApiCount' },
  { title: '角色用户数', dataIndex: 'roleUserCount' },
  { title: '角色菜单数', dataIndex: 'roleMenuCount' },
  { title: '创建时间', slotName: 'createdAt' },
  { title: '更新时间', dataIndex: 'updatedAt', type: 'datetime' },
  { title: '操作', slotName: 'action' },
]

interface roleDataType {
  roleID: number
  menuIDList: number[]
  apiIDList: number[]
}

const checkRoleData = reactive<roleDataType>({
  roleID: 0,
  menuIDList: [],
  apiIDList: [],
})

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

function showMenuModal(record: roleType) {
  checkRoleData.roleID = record.id
  checkRoleData.menuIDList = record.menuIDList as any
  menuVisible.value = true
}

function updateMenuOk() {
  lunarListRef.value.getList()
}

function showApiModal(record: roleType) {
  checkRoleData.roleID = record.id
  checkRoleData.apiIDList = record.apiIDList as any
  apiVisible.value = true
}

function updateApiOk() {
  lunarListRef.value.getList()
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
    <LunarRoleMenuModal @ok="updateMenuOk" :menu-id-list="checkRoleData.menuIDList" :role-id="checkRoleData.roleID" v-model:visible="menuVisible"></LunarRoleMenuModal>
    <LunarRoleApiModal @ok="updateApiOk" :api-id-list="checkRoleData.apiIDList" :role-id="checkRoleData.roleID" v-model:visible="apiVisible"></LunarRoleApiModal>
    <LunarList ref="lunarListRef" :columns="columns" :url="roleListApi" @edit="edit" @add="add">
      <template #action-left="{record}: {record: roleType}">
        <a-button @click="showMenuModal(record)" type="outline">设置菜单</a-button>
        <a-button @click="showApiModal(record)" type="outline">设置API</a-button>
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