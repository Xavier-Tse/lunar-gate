<script setup lang="ts">
import { useStore } from '@/stores'

const store = useStore()
const wsUrl = location.origin.replace('http', 'ws')
const ws = new WebSocket(`${wsUrl}/api/ws?Authorization=Bearer%20${store.userInfo.token}`)

ws.onopen = () => {
  console.log('ws连接成功')
}
ws.onclose = () => {
  console.log('ws连接断开')
}
ws.onerror = (val: any) => {
  console.log('ws连接失败 ', val)
}
ws.onmessage = (val: any) => {
  console.log('ws接收消息 ', val)
}
</script>

<template>
  <div>系统日志</div>
</template>

<style lang="less">

</style>