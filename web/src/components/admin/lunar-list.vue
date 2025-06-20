<script setup lang="ts">
import {reactive, ref} from "vue";
import {Message, type TableColumn} from "@arco-design/web-vue";
import { defaultDeleteApi, generateOptionsCache, type baseParams, type baseResponse, type listResponse, type optionsResponse } from "@/api";
import { dateTimeFormat } from "@/utils/date";
import LunarOptions from "../base/lunar-options.vue";
import type { filterGroupType } from "@/types";

interface actionGroupType {
  label: string
  value: string
  callback: (idList: any[]) => void
}

interface Props {
  url: (params?: baseParams) => Promise<baseResponse<listResponse<any>>>
  columns: TableColumn
  noAdd?: boolean
  noDelete?: boolean
  noEdit?:boolean
  noAction?: boolean
  noPage?: boolean
  noSearch?: boolean
  addLabel?: string
  actionGroup?: actionGroupType[]
  filterGroup?: filterGroupType[]
  defaultParams?: Object
  noDefaultDelete?: boolean
}

const params = reactive<baseParams>({
  limit: 10,
  page: 1,
  key: "",
})

const props = defineProps<Props>()
const { addLabel = '添加' } = props
const actionValue = ref('')
const selectedKeys = ref([])
const emits = defineEmits(['add', 'edit', 'remove', 'removeBatch'])

const data = reactive<listResponse<any>>({
  list: [],
  count: 0,
})

const rowSelection = {
  type: 'checkbox',
  showCheckedAll: true
}

const filterList = ref<filterGroupType[]>([])

async function initFilterGroup() {
  if (props.filterGroup) {
    filterList.value = []
    for (const f of props.filterGroup) {
      if (typeof f.source === 'function') {
        const res = generateOptionsCache(f.source() as any)
        f.options = res as any
      } else {
        f.options = f.source
      }
      filterList.value.push(f)
    }
  }
}
initFilterGroup()

async function getList(p?: baseParams) {
  if (props.defaultParams) {
    Object.assign(params, props.defaultParams)
  }
  if (p) {
    Object.assign(params, p)
  }
  const res = await props.url(params)
  if (res.code) {
    Message.error(res.message)
    return
  }
  data.list = res.data!.list
  data.count = res.data!.count
}
getList(props.defaultParams as any)

function add() {
  emits("add")
}

function edit(record: any) {
  emits("edit", record)
}

async function defaultRemove(idList: number[]) {
  const url =  /\"(.*)\"/.exec(props.url.toString())?.[1] ?? ''
  const res = await defaultDeleteApi(url, idList)
  if (res.code){
    Message.error(res.message)
    return
  }
  Message.success(res.message)
  getList()
}

function remove(record: any) {
  if (!props.noDefaultDelete) {
    defaultRemove([record.id])
    return
  }
  emits("remove", record)
}

function removeBatch(idList: any[]) {
  if (!props.noDefaultDelete) {
    defaultRemove(idList)
    return
  }
  emits("removeBatch", idList)
}

function pageChange() {
  getList()
}

async function refresh() {
  await getList()
  Message.success("刷新成功")
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
    selectedKeys.value = []
    return
  }
}

async function filterChange(index: number, column: string, val: string) {
  const item = props.filterGroup?.find((value, index, obj) => index === index)
  if (item) {
    if (item.callback) {
      item.callback(column, val)
    } else {
      const obj: { [key: string]: string } = {}
      obj[column] = val
      getList(obj)
    }
  }
}

defineExpose({
  getList,
})
</script>

<template>
  <div class="lunar-list">
    <div class="head">
      <div class="left">
        <a-button @click="add" v-if="!props.noAdd" type="primary">{{ addLabel }}</a-button>
        <slot name="action1"></slot>
        <a-select
          allow-clear
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
        <a-input-search
          class="search"
          allow-clear v-model="params.key"
          v-if="!props.noSearch"
          @search="getList"
          @keydown.enter="getList"
          placeholder="搜索"
        />
        <div class="filter-group">
          <a-select
            allow-clear
            v-for="(item, index) in filterList"
            :placeholder="item.label"
            :options="item.options"
            style="width: 160px;"
            @change="filterChange(index, item.column, $event)"
          />
        </div>
        <slot name="search_other"></slot>
      </div>
      <div class="refresh">
        <a-button @click="refresh">
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
            <a-table-column v-if="item.type || item.slotName" v-bind="item">
              <template #cell="data">
                <template v-if="item.slotName === 'action'">
                  <div class="actions">
                    <slot name="actions" :record="data.record">
                      <slot name="action-left" :record="data.record"></slot>
                      <a-button v-if="!props.noEdit" type="primary" @click="edit(data.record)">编辑</a-button>
                      <slot name="action-middle" :record="data.record"></slot>
                      <a-popconfirm v-if="!props.noDelete" content="确认删除？" @ok="remove(data.record)">
                        <a-button type="primary" status="danger">删除</a-button>
                      </a-popconfirm>
                      <slot name="action-right" :record="data.record"></slot>
                    </slot>
                  </div>
                </template>
                <template v-else-if="item.slotName === 'createdAt'">
                  <span>{{ dateTimeFormat(data.record.createdAt) }}</span>
                </template>
                <template v-else-if="item.type === 'date' || item.type === 'datetime' || item.type === 'relative'">
                  <span>{{ dateTimeFormat(data.record[item.dataIndex]) }}</span>
                </template>
                <template v-else-if="item.type === 'switch'">
                  <a-switch :model-value="data.record[item.dataIndex]" style="cursor: not-allowed;" />
                </template>
                <template v-else-if="item.type === 'avatar'">
                  <a-avatar :image-url="data.record[item.dataIndex]" />
                </template>
                <template v-else-if="item.type === 'options'">
                  <LunarOptions :options="item.options" :value="data.record[item.dataIndex]" />
                </template>
                <slot :name="item.slotName" v-bind="data"></slot>
              </template>
            </a-table-column>
            <a-table-column v-else-if="item.dataIndex" v-bind="item" />
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
      display: flex;

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

    .actions {
      .arco-btn {
        margin-right: 10px;
      }
    }
  }

  .page {
    display: flex;
    justify-content: center;
  }
}
</style>