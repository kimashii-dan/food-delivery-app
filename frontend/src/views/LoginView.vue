<template>
  <div class="auth-page">
    <form class="form">
      <div class="">
        <h1 class="title">Login</h1>
        <p class="description">Welcome back! Please sign in to continue.</p>
      </div>

      <div class="form-field">
        <label for="email">Email</label>
        <input id="email" v-model="email" type="email" placeholder="Enter your email" />
      </div>
      <div class="form-field">
        <label for="password">Password</label>
        <input id="password" v-model="password" type="password" placeholder="Enter password" />
      </div>

      <button type="button" @click="submit">Submit</button>

      <RouterLink class="no-account" to="/register">Don't have an account yet?</RouterLink>
    </form>
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/modules/auth/auth-store'
import type { LoginRequest } from '@/modules/auth/types'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const email = ref('')
const password = ref('')
const { login } = useAuthStore()
const router = useRouter()

const submit = async () => {
  const loginRequest: LoginRequest = {
    email: email.value,
    password: password.value,
  }

  try {
    const response = await login(loginRequest)
    console.log(response?.user)
    router.push('/')
  } catch (error) {
    console.log(error)
  }
}
</script>

<style scoped>
p {
  color: var(--primary);
}

.form-field {
  display: flex;
  gap: 8px;
  flex-direction: column;
}

.form {
  display: flex;
  flex-direction: column;
  gap: 32px;
  padding: 24px;
  border-radius: 12px;
  background-color: var(--bg-light);
  border: 1px solid var(--border);
  width: 384px;
}

.title {
  margin: 8px 0;
  font-weight: 500;
}

.description {
  margin: 0;
}

.no-account {
  color: var(--info);
  text-decoration: underline;
}
</style>
