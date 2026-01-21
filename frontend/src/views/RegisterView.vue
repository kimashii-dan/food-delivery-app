<template>
  <div class="auth-page">
    <div class="form">
      <div class="">
        <h1 class="title">Register</h1>
        <p class="description">Create your account to get started.</p>
      </div>

      <div class="form-field">
        <label for="email">Email</label>
        <input id="email" v-model="email" type="email" placeholder="Enter your email" />
      </div>
      <div class="form-field">
        <label for="password">Name</label>
        <input id="name" v-model="name" placeholder="Enter your name" />
      </div>
      <div class="form-field">
        <label for="password">Phone</label>
        <input id="phone" v-model="phone" type="tel" placeholder="Enter your phone" />
      </div>
      <div class="form-field">
        <label for="password">Password</label>
        <input id="password" v-model="password" type="password" placeholder="Enter password" />
      </div>
      <div class="form-field">
        <label for="password">Confirm password</label>
        <input
          id="confirmPassword"
          v-model="confirmPassword"
          type="password"
          placeholder="Confirm your password"
        />
      </div>

      <p v-if="password && confirmPassword && password !== confirmPassword">
        Passwords don't match
      </p>

      <button type="button" @click="submit">Submit</button>

      <RouterLink class="account-exists" to="/login">Already have an account?</RouterLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/modules/auth/auth-store'
import type { RegisterRequest } from '@/modules/auth/types'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const email = ref('')
const name = ref('')
const phone = ref('')
const password = ref('')
const confirmPassword = ref('')
const { register } = useAuthStore()
const router = useRouter()

const submit = async () => {
  const registerRequest: RegisterRequest = {
    email: email.value,
    password: password.value,
    name: name.value,
    phone: phone.value,
    role: 'customer',
  }

  try {
    const response = await register(registerRequest)
    console.log(response)
    router.push('/login')
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

.account-exists {
  color: var(--info);
  text-decoration: underline;
}
</style>
