import { api } from '@/shared/lib/axios'
import type { GetMeResponse, User } from './types'
export const userService = {
  async getMe() {
    const { data } = await api.get<GetMeResponse>('/api/users/me')
    return data
  },
}
