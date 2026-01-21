import { api } from '@/shared/lib/axios'
import type {
  getMenuItemResponse,
  getMenuResponse,
  getRestaurantResponse,
  getRestaurantsResponse,
} from './types'
export const restaurantService = {
  async getRestaurants() {
    const { data } = await api.get<getRestaurantsResponse>('/api/restaurants')
    return data
  },

  async getRestaurant(restaurant_id: string) {
    const { data } = await api.get<getRestaurantResponse>(`/api/restaurants/${restaurant_id}`)
    return data
  },

  async getRestaurantMenu(restaurant_id: string) {
    const { data } = await api.get<getMenuResponse>(`/api/restaurants/${restaurant_id}/menu`)
    return data
  },

  async getMenuItem(item_id: string) {
    const { data } = await api.get<getMenuItemResponse>(`api/restaurants/menu-items/${item_id}`)
    return data
  },
}
