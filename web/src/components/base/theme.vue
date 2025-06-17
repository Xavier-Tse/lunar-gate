<script setup lang="ts">
import { IconSunFill } from '@arco-design/web-vue/es/icon'
import { theme } from './theme'

function setTheme(t: '' | 'dark') {
  theme.value = t
  document.body.setAttribute('arco-theme', theme.value)
  localStorage.setItem('theme', theme.value)
}

function initTheme() {
  const val = localStorage.getItem('theme')
  if (val === null) {
    return
  }
  if (val === 'dark') {
    setTheme(val)
  }
}
initTheme()
</script>

<template>
  <span>
    <IconSunFill @click="setTheme('dark')" v-if="theme === ''" />
    <IconMoonFill @click="setTheme('')" v-else />
  </span>
</template>

<style lang="less">

</style>