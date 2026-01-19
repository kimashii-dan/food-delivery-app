import './shared/styles/base.css'
import './shared/styles/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { useAuthStore } from './modules/auth/auth-store'
import { useTheme } from './shared/hooks/useTheme'
const app = createApp(App)

const pinia = createPinia()
app.use(pinia)
app.use(router)

const authStore = useAuthStore()
authStore.initialize()

const { initTheme } = useTheme()
initTheme()

app.mount('#app')
