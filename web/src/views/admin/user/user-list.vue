<script setup lang="ts">
import { userListApi, userRemoveApi, type userListResponse } from '@/api/user-api';
import LunarList from '@/components/admin/lunar-list.vue';
import { Message } from '@arco-design/web-vue';
import { ref } from 'vue';

const lunarListRef = ref()

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
];

async function remove (record: userListResponse) {
  const res = await userRemoveApi([record.id])
  if (res.code) {
    Message.error(res.message)
    return
  }
  Message.success('删除成功')
  lunarListRef.value.getList()
}

async function removeBatch (idList: number[]) {

}

</script>

<template>
  <LunarList ref="lunarListRef" no-add class="user-list-view" :columns="columns" :url="userListApi" @remove="remove" @remove-batch="removeBatch">
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
</template>

<style lang="less">
.user-list-view {
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