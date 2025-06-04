<script setup lang="ts">
import { RouterView } from 'vue-router'
import { userLogin } from './api/user-api';
async function login() {
  const res = await userLogin({
    username: 'admin',
    password: 'root'})
  console.log(res)
}
login()
</script>

<template>
  <RouterView />
</template>