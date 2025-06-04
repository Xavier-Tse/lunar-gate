<script setup lang="ts">
import {reactive, ref} from "vue";
import type {TableColumn} from "@arco-design/web-vue";

interface actionGroupType {
  label: string
  value: string
  callback: (idList: any[]) => void
}

interface Props {
  url: string
  columns: TableColumn,
  noAdd?: boolean
  noDelete?: boolean
  noEdit?:boolean
  noAction?: boolean
  noPage?: boolean
  noSearch?: boolean
  actionGroup?: actionGroupType[]
}

const params = reactive({
  limit: 2,
  page: 1,
  key: "",
})

const props = defineProps<Props>()
const actionValue = ref('')
const selectedKeys = ref([])
const emits = defineEmits(['add', 'edit', 'remove', 'removeBatch'])

const data = reactive({
  list: [],
  count: 0,
})

const rowSelection = {
  type: 'checkbox',
  showCheckedAll: true
}

function getList() {
  Object.assign(data, {
    count: 10,
    list: [
      {id: 1, username: "fengfeng", createdAt: "2025-05-18 13:30:00"},
      {id: 2, username: "zhangsan", createdAt: "2025-05-18 13:30:00"},
      {id: 3, username: "lisi", createdAt: "2025-05-18 13:30:00"},
    ]
  })
}
getList()

function add() {
  emits("add")
}

function edit(record: any) {
  emits("edit", record)
}

function remove(record: any) {
  emits("remove", record)
}

function removeBatch(idList: any[]) {
  emits("removeBatch", idList)
}

function pageChange() {
  console.log(params)
}

const actions = ref([
  {label: "批量删除", value: "deleteBatch", callback: removeBatch}
])

function initActionGroup() {
  if (props.actionGroup) {
    for (const ac of props.actionGroup) {
      actions.value.push(ac)
    }
  }

}
initActionGroup()

function actionHandler() {
  const item = actions.value.find((v) => v.value === actionValue.value)
  if (item) {
    item.callback(selectedKeys.value)
    return;
  }
}
</script>

<template>
  <div class="lunar-list">
    <div class="head">
      <div class="left">
        <a-button @click="add" v-if="!props.noAdd" type="primary">添加</a-button>
        <a-select
          :options="actions"
          v-if="!props.noAction"
          v-model="actionValue"
          class="select"
          placeholder="操作"
        />
        <a-button
          type="primary"
          status="danger"
          v-if="selectedKeys.length && actionValue"
          @click="actionHandler"
        >执行</a-button>
        <a-input-search v-model="params.key" v-if="!props.noSearch" class="search" placeholder="搜索" />
        <slot name="search_other"></slot>
      </div>
      <div class="refresh">
        <a-button>
          <IconRefresh />
        </a-button>
      </div>
    </div>
    <div class="record">
      <a-table
        :pagination="false"
        :row-selection="rowSelection"
        v-model:selected-keys="selectedKeys"
        row-key="id"
        :columns="props.columns"
        :data="data.list"
      >
        <template #columns>
          <template v-for="item in props.columns">
            <a-table-column v-if="item.dataIndex" v-bind="item" />
            <a-table-column v-if="item.slotName" v-bind="item">
              <template #cell="data">
                <template v-if="item.slotName === 'action'">
                  <div class="actions">
                    <slot name="actions" :record="data.record">
                      <a-button
                        v-if="!props.noEdit"
                        type="primary"
                        @click="edit(data.record)"
                      >编辑</a-button>
                      <a-button
                        v-if="!props.noDelete"
                        type="primary"
                        @click="remove(data.record)"
                        status="danger"
                      >删除</a-button>
                    </slot>
                  </div>
                </template>
                <slot :name="item.slotName" v-bind="data"></slot>
              </template>
            </a-table-column>
          </template>
        </template>
      </a-table>
    </div>
    <div class="page" v-if="!props.noPage">
      <a-pagination
        @change="pageChange"
        v-model:current="params.page"
        :page-size="params.limit"
        :total="data.count"
        show-total
        show-jumper
      />
    </div>
  </div>
</template>

<style lang="less">
.lunar-list {
  padding: 0 !important;

  .head {
    border-bottom: 1px solid var(--color-neutral-2);
    display: flex;
    padding: 20px 20px 10px 20px;
    position: relative;

    .left{
      >*{
        margin-right: 10px;
      }
    }

    .select {
      width: 120px;
    }

    .search {
      width: 200px;
    }

    .refresh {
      position: absolute;
      right: 20px;
      
      .arco-btn{
        width: 32px;
        padding: 0;
      }
    }
  }

  .record {
    padding: 10px 20px 20px 20px;
  }

  .page {
    display: flex;
    justify-content: center;
  }
}
</style>