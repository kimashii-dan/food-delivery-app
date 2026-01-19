import { defineStore } from 'pinia'
import type { User } from '../user/types'
import type { LoginRequest, RegisterRequest } from './types'
import { authService } from './auth-service'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isAuthenticated: false as boolean,
    user: null as null | User,
  }),

  getters: {
    userRole: (state) => state.user?.role,
  },

  actions: {
    initialize() {
      const storedUser = localStorage.getItem('user')
      if (storedUser) {
        try {
          this.user = JSON.parse(storedUser)
          this.isAuthenticated = true
        } catch (error) {
          localStorage.removeItem('user')
        }
      }
    },

    async register(credentials: RegisterRequest) {
      try {
        const response = await authService.register(credentials)
        return response
      } catch (error) {
        console.error('Register error:', error)
      }
    },

    async login(credentials: LoginRequest) {
      try {
        const response = await authService.login(credentials)
        this.setAuthenticated(response.user)
        return response
      } catch (error) {
        console.error('Login error:', error)
      }
    },

    async logout() {
      try {
        await authService.logout()
        this.clearAuth()
      } catch (error) {
        console.error('Logout error:', error)
      }
    },

    setAuthenticated(user: User | null) {
      this.user = user
      this.isAuthenticated = true
      localStorage.setItem('user', JSON.stringify(user))
    },

    updateUser(user: User) {
      this.user = user
      localStorage.setItem('user', JSON.stringify(user))
    },

    clearAuth() {
      this.user = null
      this.isAuthenticated = false
      localStorage.removeItem('user')
    },
  },
})
