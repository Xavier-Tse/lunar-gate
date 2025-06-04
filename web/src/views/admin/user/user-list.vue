<script setup lang="ts">
import LunarList from '@/components/admin/lunar-list.vue';

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
  },
  {
    title: '用户名',
    dataIndex: 'username',
  },
  {
    title: '时间',
    dataIndex: 'createdAt',
  }
];

</script>

<template>
  <LunarList :columns="columns" url="/pai/user/list"></LunarList>
</template>

<style lang="less">

</style>