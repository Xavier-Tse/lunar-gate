<script setup lang="ts">
import { RouterView } from 'vue-router'
import { useStore } from './stores';

const store = useStore()
store.loadUser()
</script>

<template>
  <RouterView />
</template>