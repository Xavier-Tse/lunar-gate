<script setup lang="ts">
import { generateOptions, generateOptionsCache } from '@/api';
import { apiBatchCreateApi, apiGroupOptionsApi, systemRouterApi, type systemRouterType } from '@/api/api-api';
import { Message } from '@arco-design/web-vue';
import { ref } from 'vue';

interface Props {
  visible: boolean
}
const props = defineProps<Props>()
const emits = defineEmits(['update:visible', 'ok'])

const data = ref<systemRouterType[]>([])
const checkboxIndexList = ref<number[]>([])
const groupOptions = generateOptionsCache(apiGroupOptionsApi)

function cancel() {
  emits('update:visible', false)
}

async function beforeOpen() {
  const res = await systemRouterApi()
  if (res.code) {
    Message.error(res.message)
    return
  }
  data.value = res.data as any
}

async function okHandler() {
  const list: systemRouterType[] = []
  for (const index of checkboxIndexList.value) {
    const item = data.value[index]
    if (item.name === '' || item.group === '') {
      Message.warning('请输入API名称和分组')
      return
    }
    list.push(item)
  }
  const res = await apiBatchCreateApi(list)
  if (res.code) {
    Message.error(res.message)
    return
  }
  emits('ok')
  Message.success(res.message)
  emits('update:visible', false)
  await beforeOpen()
  checkboxIndexList.value = []
}
</script>

<template>
  <a-modal width="auto" title="同步系统API" :on-before-ok="okHandler" :visible="props.visible" @before-open="beforeOpen" @cancel="cancel" modal-class="system-router-modal">
    <a-checkbox-group v-model="checkboxIndexList">
      <template v-for="(item, index) in data">
        <div class="item" v-if="!item.enable">
          <a-checkbox :value="index" />
          <span :class="'method ' + item.method">{{ item.method }}</span>
          <span class="path">{{ item.path }}</span>
          <a-input placeholder="API名称" v-model="item.name" />
          <a-select :options="groupOptions" allow-create allow-clear placeholder="API分组" v-model="item.group" />
        </div>
      </template>
    </a-checkbox-group>
  </a-modal>
</template>

<style lang="less">
.system-router-modal {
  .arco-modal-body {
    max-height: 80vh;
    overflow-y: auto;

    .item {
      display: flex;
      height: 40px;
      align-items: center;

      .arco-input-wrapper {
        margin: 0 10px;
      }

      .method {
        margin-right: 10px;

        &.POST {
          color: orange;
        }

        &.GET {
          color: rgb(0, 110, 255);
        }

        &.DELETE {
          color: red;
        }

        &.PUT {
          color: purple;
        }

        &.PATCH {
          color: rgba(var(--pinkpurple-6),.2);
        }
      }

      .path {
        margin-right: 10px;
        min-width: 150px;
      }
    }
  }
}
</style>