import { authService } from '@/modules/auth/auth-service'
import axios, { AxiosError, type AxiosRequestConfig } from 'axios'

export const api = axios.create({
  baseURL: 'http://localhost:8000',
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true,
})

api.interceptors.response.use(
  (response) => response,
  async (error: AxiosError & { config?: AxiosRequestConfig & { _retry?: boolean } }) => {
    const originalRequest = error.config
    if (
      error.response?.status === 401 &&
      originalRequest &&
      !originalRequest._retry &&
      !originalRequest.url?.includes('/api/users/refresh') &&
      !originalRequest.url?.includes('/api/users/login')
    ) {
      originalRequest._retry = true
      try {
        await authService.refresh()
        return api(originalRequest)
      } catch (refreshError) {
        const { useAuthStore } = await import('@/modules/auth/auth-store')
        const authStore = useAuthStore()
        authStore.clearAuth()

        // Redirect to login
        if (typeof window !== 'undefined') {
          window.location.href = '/login'
        }
        return Promise.reject(refreshError)
      }
    }
    return Promise.reject(error)
  },
)
