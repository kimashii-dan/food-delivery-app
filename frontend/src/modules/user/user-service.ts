import { api } from '@/lib/axios'
import type { User } from './types'

type GetMeResponse = {
  user: User
}

export const userService = {
  async getMe() {
    const { data } = await api.get<GetMeResponse>('/api/users/me')
    return data
  },
}
