<script setup lang="ts">
import { logListApi } from '@/api/log-api'
import { useStore } from '@/stores'
import { Message } from '@arco-design/web-vue'
import { AnsiUp } from 'ansi_up'
import { nextTick, ref } from 'vue'

interface wsMessageType {
  type: string
  content: string
}

const store = useStore()
const wsUrl = location.origin.replace('http', 'ws')
const ws = new WebSocket(`${wsUrl}/api/ws?Authorization=Bearer%20${store.userInfo.token}`)
const logsRef = ref()
const au = new AnsiUp()
const logList = ref<string[]>([])

async function getLogs() {
  const res = await logListApi({ limit: 30 })
  if (res.code) {
    Message.error(res.message)
    return
  }
  for (const u of res.data as any) {
    logList.value.push(au.ansi_to_html(u))
  }
}
getLogs()

ws.onmessage = (val: MessageEvent) => {
  const data = JSON.parse(val.data) as wsMessageType
  if (data.type === 'logs') {
    logList.value.push(au.ansi_to_html(data.content))

    nextTick(() => {
      const dom = logsRef.value as HTMLDivElement
      dom?.scrollTo({ top: dom.scrollHeight, behavior: 'smooth' })
    })
  }
}
ws.onclose = () => {
  console.log('ws连接断开')
}
</script>

<template>
  <div class="logs-view">
    <div class="logs-container" ref="logsRef">
      <div class="message" v-for="item in logList" v-html="item" />
    </div>
  </div>
</template>

<style lang="less">
.logs-view{
  .logs-container{
    height: 100%;
    overflow-y: auto;
    color: var(--color-text-1);
    line-height: 1.5rem;

    .message {
      word-break: break-word;
    }
  }
}
</style>