<script setup lang="ts">
import { userListApi, type userListResponse } from '@/api/user-api';
import LunarList from '@/components/admin/lunar-list.vue';
import { dateTimeFormat } from '@/utils/date';

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

</script>

<template>
  <LunarList no-add class="user-list-view" :columns="columns" :url="userListApi">
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