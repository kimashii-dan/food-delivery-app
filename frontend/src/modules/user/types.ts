interface User {
  id: string
  email: string
  name: string
  phone: string
  role: string
  created_at: string
}

type GetMeResponse = {
  user: User
}

export type { User, GetMeResponse }
