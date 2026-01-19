<template>
  <div class="page">
    <h1 class="page-title">Profile</h1>
    <div v-if="loading" class="skeleton">loading...</div>
    <div v-else-if="apiError" class="error">
      <p>Error: {{ apiError }}</p>
      <button @click="retry" type="button">Retry</button>
    </div>
    <div v-else class="content">
      <p class="user-id">ID: {{ user?.id }}</p>
      <div class="user-info">
        <p class="user-name">Name: {{ user?.name }}</p>
        <p class="user-email">Email: {{ user?.email }}</p>
        <p class="user-phone">Phone: {{ user?.phone }}</p>
      </div>
      <button @click="onLogout" type="button">Logout</button>
    </div>
  </div>
</template>
<script setup lang="ts">
import { useAuthStore } from '@/modules/auth/auth-store'
import type { User } from '@/modules/user/types'
import { userService } from '@/modules/user/user-service'
import type { ApiError } from '@/shared/types'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const { logout } = useAuthStore()
const user = ref<User | null>(null)
const loading = ref<boolean>(false)
const apiError = ref<string>('')

const fetchUser = async () => {
  try {
    loading.value = true
    await new Promise((resolve) => setTimeout(resolve, 2000))
    const response = await userService.getMe()
    user.value = response.user
  } catch (error) {
    apiError.value = (error as ApiError)?.response?.data?.error
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchUser()
})

const retry = () => {
  fetchUser()
}

const onLogout = async () => {
  try {
    await logout()
    router.push('/')
  } catch (error) {
    console.log(error)
  }
}
</script>
<style scoped>
.error {
  color: var(--danger);
}

.user-info {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 32px;
}

.user-info > * {
  padding: 20px 16px;
  background-color: var(--bg-light);
  font-size: 16px;
  border: 1px solid var(--border);
  border-radius: 12px;
}

.user-id {
  color: var(--text-muted);
}
</style>
