<script setup lang="ts">
import { menuListApi, type menuType } from '@/api/menu-api';
import LunarList from '@/components/admin/lunar-list.vue';
import LunarIcon from '@/components/base/lunar-icon.vue';
import { ref } from 'vue';

const lunarListRef = ref()
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

function edit(record: menuType) {
  visible.value = true
}

function add() {
  visible.value = true
}

</script>

<template>
  <div class="menu-list-view no-padding">
    <LunarList ref="lunarListRef" @add="add" @edit="edit" :columns="columns" :url="menuListApi">
      <template #icon="{record}: {record: menuType}">
        <div v-if="record.meta.icon" style="display: flex;">
          <LunarIcon :icon="record.meta.icon" />
          <code style="padding-left: 5px;">{{ record.meta.icon }}</code>
        </div>
        <span v-else>未设置图标</span>
      </template>
    </LunarList>
  </div>
</template>