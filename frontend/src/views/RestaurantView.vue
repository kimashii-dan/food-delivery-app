<template>
  <div class="page">
    <!-- Loading State -->
    <div v-if="loading" class="loading-container">
      <div class="spinner"></div>
      <p>Loading restaurant details...</p>
    </div>

    <!-- Error State -->
    <div v-else-if="apiError" class="error-container">
      <p class="error-message">{{ apiError }}</p>
      <button @click="fetchMenu" class="retry-btn">Try Again</button>
    </div>

    <!-- Restaurant Content -->
    <div v-else-if="restaurant" class="restaurant-container">
      <!-- Restaurant Header -->
      <div class="restaurant-hero">
        <div class="hero-content">
          <div class="hero-text">
            <h1 class="restaurant-title">{{ restaurant.name }}</h1>
            <p class="restaurant-address">{{ restaurant.address }}</p>
            <p class="restaurant-description">{{ restaurant.description }}</p>
            <div class="restaurant-info">
              <span class="info-badge">
                {{ restaurant.opening_time }} - {{ restaurant.closing_time }}
              </span>
            </div>
          </div>
          <div class="hero-image">
            <img :src="restaurant.logo_url" :alt="restaurant.name" />
          </div>
        </div>
      </div>

      <!-- Menu Section -->
      <div class="menu-section">
        <h2 class="section-title">Menu</h2>

        <div v-if="menu && menu.length > 0" class="menu-grid">
          <div v-for="item in menu" :key="item.id" class="menu-item-card">
            <div class="menu-item-image" v-if="item.image_url">
              <img :src="item.image_url" :alt="item.name" />
            </div>
            <div class="menu-item-content">
              <div class="menu-item-header">
                <h3 class="menu-item-name">{{ item.name }}</h3>
                <span class="menu-item-price">₸{{ item.price.toFixed(2) }}</span>
              </div>
              <p class="menu-item-description">{{ item.description }}</p>
              <button class="add-to-cart-btn">Add to Cart</button>
            </div>
          </div>
        </div>

        <div v-else class="empty-menu">
          <p>No menu items available yet.</p>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { restaurantService } from '@/modules/restaurant/restaurant-service'
import type { MenuItem, Restaurant } from '@/modules/restaurant/types'
import type { ApiError } from '@/shared/types'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

const restaurant = ref<Restaurant | null>(null)
const menu = ref<MenuItem[] | null>(null)
const loading = ref<boolean>(false)
const apiError = ref<string>('')
const route = useRoute()
const id = String(route.params.id)

const fetchMenu = async () => {
  try {
    loading.value = true
    const [restaurantResponse, menuResponse] = await Promise.all([
      restaurantService.getRestaurant(id),
      restaurantService.getRestaurantMenu(id),
    ])

    restaurant.value = restaurantResponse.restaurant
    menu.value = menuResponse.items
  } catch (error) {
    apiError.value = (error as ApiError)?.response?.data?.error
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchMenu()
  console.log(restaurant, menu)
})
</script>
<style scoped>
/* Loading State */
.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 20px;
  padding: 80px 20px;
  color: var(--text-muted);
}

.spinner {
  width: 50px;
  height: 50px;
  border: 4px solid var(--border);
  border-top-color: var(--text-primary);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

/* Error State */
.error-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
  padding: 40px;
  text-align: center;
}

.error-message {
  color: #ef4444;
  font-size: 16px;
}

.retry-btn {
  padding: 10px 24px;
  background-color: var(--text-primary);
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 16px;
  transition: opacity 0.2s;
}

.retry-btn:hover {
  opacity: 0.9;
}

/* Restaurant Container */
.restaurant-container {
  width: 100%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 48px;
}

/* Restaurant Hero Section */
.restaurant-hero {
  background-color: var(--bg-light);
  border: 1px solid var(--border);
  border-radius: 8px;
  padding: 24px;
}

.hero-content {
  display: grid;
  grid-template-columns: 1fr auto;
  gap: 40px;
  align-items: center;
}

.hero-text {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.restaurant-title {
  font-size: 28px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
}

.restaurant-address {
  font-size: 14px;
  color: var(--text-muted);
}

.restaurant-description {
  font-size: 14px;
  line-height: 1.5;
  color: var(--text-primary);
  max-width: 600px;
}

.restaurant-info {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  margin-top: 8px;
}

.info-badge {
  display: inline-block;
  padding: 6px 12px;
  background-color: transparent;
  border: 1px solid var(--border);
  border-radius: 4px;
  font-size: 13px;
  color: var(--text-muted);
}

.hero-image {
  width: 150px;
  height: 150px;
  border-radius: 8px;
  overflow: hidden;
}

.hero-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* Menu Section */
.menu-section {
  display: flex;
  flex-direction: column;
  gap: 28px;
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
  padding-bottom: 8px;
  border-bottom: 1px solid var(--border);
}

.menu-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 24px;
}

.menu-item-card {
  background-color: var(--bg-light);
  border: 1px solid var(--border);
  border-radius: 8px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.menu-item-image {
  width: 100%;
  height: 160px;
  overflow: hidden;
  background-color: var(--border);
}

.menu-item-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.menu-item-content {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.menu-item-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 16px;
}

.menu-item-name {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
  flex: 1;
}

.menu-item-price {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  white-space: nowrap;
}

.menu-item-description {
  font-size: 13px;
  line-height: 1.4;
  color: var(--text-muted);
  margin: 0;
}

.add-to-cart-btn {
  margin-top: 8px;
  padding: 8px 16px;
  background-color: var(--text-primary);
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 14px;
  cursor: pointer;
}

.empty-menu {
  text-align: center;
  padding: 60px 20px;
  color: var(--text-muted);
  font-size: 18px;
}

/* Responsive Design */
@media (max-width: 768px) {
  .hero-content {
    grid-template-columns: 1fr;
  }

  .hero-image {
    margin: 0 auto;
  }

  .restaurant-title {
    font-size: 24px;
  }

  .menu-grid {
    grid-template-columns: 1fr;
  }
}
</style>
