<script setup lang="ts">
import { RouterView } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useAuthStore } from './modules/auth/auth-store'

const authStore = useAuthStore()
const { isLoggedIn } = storeToRefs(authStore)
</script>

<template>
  <header>
    <nav>
      <li><RouterLink to="/">Home</RouterLink></li>
      <li><RouterLink to="/about">About</RouterLink></li>
      <li v-if="isLoggedIn"><RouterLink to="/profile">Profile</RouterLink></li>
      <li v-else><RouterLink to="/login">Login</RouterLink></li>
    </nav>
  </header>

  <main><RouterView /></main>
</template>

<style scoped></style>
