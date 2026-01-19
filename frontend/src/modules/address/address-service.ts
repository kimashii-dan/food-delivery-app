import { api } from '@/shared/lib/axios'
import type { AddAddressRequest, AddAddressResponse, GetAddressesResponse } from './types'

export const addressService = {
  async getAddresses() {
    const { data } = await api.get<GetAddressesResponse>('/api/users/addresses')
    return data
  },

  async addAddress(address: AddAddressRequest) {
    const { data } = await api.post<AddAddressResponse>('/api/users/addresses', address)
    return data
  },
}

export type {}
