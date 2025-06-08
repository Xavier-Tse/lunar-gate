<script setup lang="ts">
import { menuTreeApi, type menuType } from '@/api/menu-api';
import { permissionRoleMenuApi, type permissionRoleMenuRequest } from '@/api/permission-api';
import { Message } from '@arco-design/web-vue';
import { reactive, ref } from 'vue';

interface Props {
  visible: boolean
  roleId: number
  menuIdList: number[]
}
const props = defineProps<Props>()

const emits = defineEmits(['update:visible', 'ok'])

const form = reactive<permissionRoleMenuRequest>({
  roleID: 0,
  menuIDList: [],
})

const data = ref<menuType[]>([])

async function getMenuTree() {
  const res = await menuTreeApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  data.value = res.data as any
  form.menuIDList = props.menuIdList
}

async function updateRoleMenuHandler() {
  form.roleID = props.roleId
  const res = await permissionRoleMenuApi(form)
  if (res.code) {
    Message.error(res.message)
    return
  }
  Message.success(res.message)
  emits('update:visible', false)
  emits('ok')
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
      v-model:checked-keys="form.menuIDList"
    >
      <template #title="data">
        {{ data.meta.title }}
      </template>
    </a-tree>
  </a-modal>
</template>

<style lang="less">

</style>