<script setup lang="ts">
import {IconFullscreen, IconFullscreenExit} from '@arco-design/web-vue/es/icon'
import { ref } from 'vue'

const isFull = ref(false)

function enterFullscreen() {
  const elem = document.documentElement
  elem.requestFullscreen()
    .catch((err) => {
      console.error('进入全屏失败 ', err)
  })
  isFull.value = true
}

function exitFullscreen() {
  if (document.fullscreenElement) {
    document.exitFullscreen()
      .catch((err) => {
        console.error('退出全屏失败 ', err)
    })
  }
  isFull.value = false
}

window.onresize = function() {
  isFull.value = Math.abs(window.screen.height-window.document.documentElement.clientHeight) <= 17
}

</script>

<template>
  <span>
    <span @click="enterFullscreen" v-if="!isFull">
      <icon-fullscreen />
    </span>
    <span @click="exitFullscreen" v-else>
      <icon-fullscreen-exit />
    </span>
  </span>
</template>