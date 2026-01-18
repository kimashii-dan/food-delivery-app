<template>
  <div>
    <p>ID: {{ user?.id }}</p>
    <p>Name: {{ user?.name }}</p>
    <p>Email: {{ user?.email }}</p>
    <p>Phone: {{ user?.phone }}</p>
    <p>Role: {{ user?.role }}</p>
    <button @click="onLogout" type="button">Logout</button>
  </div>
</template>
<script setup lang="ts">
import { useAuthStore } from '@/modules/auth/auth-store'
import type { User } from '@/modules/user/types'
import { userService } from '@/modules/user/user-service'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const { logout } = useAuthStore()
const user = ref<User | null>(null)

onMounted(async () => {
  console.log('Component mounted!')
  const response = await userService.getMe()
  user.value = response.user
})

const onLogout = async () => {
  try {
    await logout()
    router.push('/')
  } catch (error) {
    console.log(error)
  }
}
</script>
<style scoped></style>
