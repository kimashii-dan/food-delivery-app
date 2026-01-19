interface Address {
  id: string
  user_id: string
  street: string
  city: string
  postal_code: string
  longitude: number
  latitude: number
  is_default: boolean
  created_at: string
}

type GetAddressesResponse = {
  addresses: Address[]
}

type AddAddressResponse = {
  address_id: string
}

type AddAddressRequest = {
  city: string
  street: string
  postal_code: string
  longitude: number
  latitude: number
  is_default: boolean
}

export type { GetAddressesResponse, AddAddressRequest, AddAddressResponse, Address }
