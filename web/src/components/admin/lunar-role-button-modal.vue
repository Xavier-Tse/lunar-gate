<script setup lang="ts">
import { buttonGroupListApi, type buttonGroupListResponse, type buttonType } from '@/api/button-api';
import { permissionButtonUpdateApi, type permissionRoleButtonRequest } from '@/api/permission-api';
import { Message } from '@arco-design/web-vue';
import { reactive, ref } from 'vue';

interface Props {
  visible: boolean
  roleId: number
  buttonIdList: number[]
}
const props = defineProps<Props>()

const emits = defineEmits(['update:visible', 'ok'])

const form = reactive<permissionRoleButtonRequest>({
  roleID: 0,
  buttonIDList: [],
})

const data = ref<buttonGroupListResponse[]>([])

async function getButtonGroup() {
  const res = await buttonGroupListApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  data.value = res.data as any
  form.buttonIDList = props.buttonIdList
}

function groupChange(val: boolean, list: buttonType[]) {
  const groupButtonIDList = list.map(value => value.id)
  if (val) {
    form.buttonIDList = [...new Set([...form.buttonIDList, ...groupButtonIDList])]
  } else {
    form.buttonIDList = form.buttonIDList.filter(value => !groupButtonIDList.includes(value))
  }
}

function checked(list: buttonType[]): boolean {
  return list.every(btn => form.buttonIDList.includes(btn.id)) || false
}

async function updateRoleButtonHandler() {
  form.roleID = props.roleId
  const res = await permissionButtonUpdateApi(form)
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
  <a-modal
    body-class="lunar-role-button-modal-body"
    width="400px"
    @before-open="getButtonGroup"
    @cancel="cancel"
    :on-before-ok="updateRoleButtonHandler"
    :visible="props.visible"
    title="设置按钮">
    <div class="button-tree">
        <div class="group" v-for="item in data">
          <div class="title">
            <a-checkbox @change="groupChange($event, item.list)" :model-value="checked(item.list)">{{ item.groupTitle }}</a-checkbox>
          </div>
          <div class="children">
            <a-checkbox-group v-model="form.buttonIDList">
              <div class="router" v-for="r in item.list">
                <a-checkbox :value="r.id">
                  <span class="title">{{ r.title }}</span>
                  <span class="name">{{ r.name }}</span>
                </a-checkbox>
              </div>
            </a-checkbox-group>
          </div>
        </div>
    </div>
  </a-modal>
</template>

<style lang="less">
.lunar-role-button-modal-body {
  max-height: 70vh;
  overflow-y: auto;

  .button-tree {
    .arco-checkbox-group {
      width: 100%;
    }

    .children {
      margin-left: 20px;

      .arco-checkbox {
        width: 100%;
      }

      .name {
        position: absolute;
        right: 0;
        color: var(--color-text-2);
      }
    }
  }
}
</style>