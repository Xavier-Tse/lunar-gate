<script setup lang="ts">
import { apiGroupApi, type apiType } from '@/api/api-api';
import { type menuType } from '@/api/menu-api';
import { permissionRoleApiApi, type permissionRoleApiRequest } from '@/api/permission-api';
import { Message } from '@arco-design/web-vue';
import { reactive, ref } from 'vue';

interface Props {
  visible: boolean
  roleId: number
  apiIdList: number[]
}
const props = defineProps<Props>()

const emits = defineEmits(['update:visible', 'ok'])

const form = reactive<permissionRoleApiRequest>({
  roleID: 0,
  apiIDList: [],
})

const data = ref<Record<string, apiType[]>>({})
const groupList = ref<string[]>([])

async function getApiGroup() {
  const res = await apiGroupApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  data.value = res.data as any
  form.apiIDList = props.apiIdList
}

function groupChange(key: string, val: boolean) {
  const groupApiIDList = data.value[key].map(value => value.id)
  if (val) {
    form.apiIDList = [...new Set([...form.apiIDList, ...groupApiIDList])]
  } else {
    form.apiIDList = form.apiIDList.filter(value => !groupApiIDList.includes(value))
  }
}

function checked(val: string): boolean {
  return data.value[val].every(api => form.apiIDList.includes(api.id)) || false
}

async function updateRoleApiHandler() {
  form.roleID = props.roleId
  const res = await permissionRoleApiApi(form)
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
  <a-modal body-class="lunar-role-api-modal-body" width="400px" @before-open="getApiGroup" @cancel="cancel" :on-before-ok="updateRoleApiHandler" :visible="props.visible" title="设置API">
    <div class="api-tree">
        <div class="group" v-for="(group, key) in data">
          <div class="title">
            <a-checkbox @change="groupChange(key, $event)" :model-value="checked(key)">{{ key }}</a-checkbox>
          </div>
          <div class="children">
            <a-checkbox-group v-model="form.apiIDList">
              <div class="router" v-for="r in group">
                <a-checkbox :value="r.id">
                  <span class="name">{{ r.name }}</span>
                  <span class="path">{{ r.path }}</span>
                </a-checkbox>
              </div>
            </a-checkbox-group>
          </div>
        </div>
    </div>
  </a-modal>
</template>

<style lang="less">
.lunar-role-api-modal-body {
  .api-tree {
    .arco-checkbox-group {
      width: 100%;
    }

    .children {
      margin-left: 20px;

      .arco-checkbox {
        width: 100%;
      }

      .path {
        position: absolute;
        right: 0;
        color: var(--color-text-2);
      }
    }
  }
}
</style>