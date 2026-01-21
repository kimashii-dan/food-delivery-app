interface Restaurant {
  id: string
  name: string
  description: string
  address: string
  phone: string
  latitude: number
  longitude: number
  logo_url: string
  opening_time: string
  closing_time: string
  created_at: string
  updated_at: string
}

interface MenuItem {
  id: string
  restaurant_id: string
  name: string
  description: string
  price: number
  image_url: string
  is_available: boolean
  category: string
  created_at: string
  updated_at: string
}

type getRestaurantsResponse = {
  restaurants: Restaurant[]
  total: number
}

type getRestaurantResponse = {
  restaurant: Restaurant
}

type getMenuResponse = {
  items: MenuItem[]
  total: number
}

type getMenuItemResponse = {
  item: MenuItem
}

export type {
  Restaurant,
  MenuItem,
  getRestaurantsResponse,
  getRestaurantResponse,
  getMenuItemResponse,
  getMenuResponse,
}
