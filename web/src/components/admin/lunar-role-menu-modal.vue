<script setup lang="ts">
import { menuTreeApi, type menuType } from '@/api/menu-api';
import { Message } from '@arco-design/web-vue';
import { ref } from 'vue';

interface Props {
  visible: boolean
}
const props = defineProps<Props>()

const emits = defineEmits(['update:visible', 'ok'])
const checkedKeys = ref([])

const data = ref<menuType[]>([])
async function getMenuTree() {
  const res = await menuTreeApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  data.value = res.data as any
}

async function updateRoleMenuHandler() {
  
}

function cancel() {
  emits('update:visible', false)
}
</script>

<template>
  <a-modal width="400px" @before-open="getMenuTree" @cancel="cancel" :on-before-ok="updateRoleMenuHandler" :visible="props.visible" title="设置菜单">
    <a-tree
      checkable
      check-strictly
      :data="data"
      v-model:checked-keys="checkedKeys"
    >
      <template #title="data">
        {{ data.meta.title }}
      </template>
    </a-tree>
  </a-modal>
</template>

<style lang="less">

</style>