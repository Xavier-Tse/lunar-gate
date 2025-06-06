<script setup lang="ts">
import type { optionsResponse } from '@/api';
import { roleOptionsApi } from '@/api/role-api';
import { userListApi, userRoleUpdateApi, type userListResponse, type userRoleUpdateRequest } from '@/api/user-api';
import LunarList from '@/components/admin/lunar-list.vue';
import { Message } from '@arco-design/web-vue';
import { reactive, ref } from 'vue';

const lunarListRef = ref()
const visible = ref(false)
const roleOptions = ref<optionsResponse[]>([])

const form = reactive<userRoleUpdateRequest>({
  userID: 0,
  roleIDList: [],
})

async function initRoleOptions() {
  const res = await roleOptionsApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  roleOptions.value = res.data as any
}
initRoleOptions()

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '用户名', dataIndex: 'username' },
  { title: '昵称', dataIndex: 'nickname' },
  { title: '头像', dataIndex: 'avatar', type: 'avatar' },
  { title: '邮箱', dataIndex: 'email' },
  { title: '角色', slotName: 'roleList' },
  { title: '管理员', dataIndex: 'isAdmin', type: 'switch' },
  { title: '注册时间', dataIndex: 'createdAt', type: 'date' },
  { title: '操作', slotName: 'action' },
]

function edit(record: userListResponse) {
  form.userID = record.id
  let roleIDList = []
  for (const item of record.roleList) {
    roleIDList.push(item.id)
  }
  form.roleIDList = roleIDList
  visible.value = true
}

async function updateUserRole() {
  const res = await userRoleUpdateApi(form)
  if (res.code) {
    Message.error(res.message)
    return
  }
  Message.success(res.message)
  lunarListRef.value.getList()
}

</script>

<template>
  <div class="user-list-view no-padding">
    <a-modal title="编辑用户" v-model:visible="visible" :on-before-ok="updateUserRole">
      <a-form :model="form">
        <a-form-item label="角色">
          <a-select v-model="form.roleIDList" placeholder="选择角色" :options="roleOptions" multiple allow-clear></a-select>
        </a-form-item>
      </a-form>
    </a-modal>
    <LunarList ref="lunarListRef" no-add :columns="columns" :url="userListApi" @edit="edit">
      <template #avatar="{record}:{record: userListResponse}">
        <a-avatar :image-url="record.avatar" />
      </template>
      <template #roleList="{record}:{record: userListResponse}">
        <div class="role-list">
          <a-tag v-for="item in record.roleList">
            {{ item.title }}
          </a-tag>
        </div>
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