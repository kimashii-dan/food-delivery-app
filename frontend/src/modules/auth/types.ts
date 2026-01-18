import type { User } from '../user/types'

type RegisterRequest = {
  email: string
  name: string
  password: string
  phone: string
  role: string
}

type RegisterResponse = {
  user_id: string
}

type LoginResponse = {
  user: User
}

type LoginRequest = {
  email: string
  password: string
}

export type { RegisterRequest, RegisterResponse, LoginRequest, LoginResponse }
