<template>
  <div>
    <h1>Login</h1>
    <form>
      <div class="form-field">
        <label for="email">Email</label>
        <input id="email" v-model="email" type="email" placeholder="Enter your email" />
      </div>
      <div class="form-field">
        <label for="password">Password</label>
        <input id="password" v-model="password" type="password" placeholder="Enter password" />
      </div>

      <button type="button" @click="submit">Submit</button>
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
  color: #42b983;
}
form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.form-field {
  display: flex;
  gap: 8px;
  flex-direction: column;
}
</style>
