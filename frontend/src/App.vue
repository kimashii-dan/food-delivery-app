<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useAuthStore } from './modules/auth/auth-store'
import ToggleTheme from './shared/components/ToggleTheme.vue'

const authStore = useAuthStore()

const route = useRoute()

console.log(route.path)
</script>

<template>
  <header v-if="route.path !== '/login' && route.path !== '/register'">
    <nav>
      <div class="link-group">
        <li><RouterLink to="/">Home</RouterLink></li>
        <li><RouterLink to="/about">About</RouterLink></li>
      </div>
      <div class="link-group">
        <li><ToggleTheme /></li>
        <li v-if="authStore.isAuthenticated"><RouterLink to="/profile">Profile</RouterLink></li>
        <li v-else><RouterLink to="/login">Login</RouterLink></li>
      </div>
    </nav>
  </header>
  <div class="utilities" v-else>
    <ToggleTheme />
  </div>

  <main><RouterView /></main>
</template>

<style scoped>
.link-group {
  display: flex;
  gap: 32px;
  align-items: center;
}

.utilities {
  position: absolute;
  right: 0;
  top: 0;
  padding: 16px;
}
</style>
