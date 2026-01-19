import { api } from '@/shared/lib/axios'
import type { LoginRequest, LoginResponse, RegisterRequest, RegisterResponse } from './types'

export const authService = {
  async register(credentials: RegisterRequest) {
    const { data } = await api.post<RegisterResponse>('/api/users/register', credentials)
    return data
  },

  async login(credentials: LoginRequest) {
    const { data } = await api.post<LoginResponse>('/api/users/login', credentials)
    return data
  },

  async logout() {
    await api.post('/api/users/logout')
  },

  async refresh() {
    await api.post('/api/users/refresh')
  },
}
