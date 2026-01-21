<template>
  <div class="page">
    <div class="restaurants-section">
      <div class="header-section">
        <h1 class="page-title">Discover Restaurants</h1>
        <p class="page-subtitle">Find your favorite food from local restaurants</p>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="loading-container">
        <div class="spinner"></div>
        <p>Loading restaurants...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="apiError" class="error-container">
        <p class="error-message">{{ apiError }}</p>
        <button @click="fetchRestaurants" class="retry-btn">Try Again</button>
      </div>

      <!-- Restaurants Grid -->
      <div v-else-if="restaurants && restaurants.length > 0" class="restaurants-grid">
        <RouterLink
          :to="`/restaurants/${restaurant.id}`"
          v-for="restaurant in restaurants"
          :key="restaurant.id"
          class="restaurant-card"
        >
          <div class="card-image-container">
            <img class="restaurant-logo" :src="restaurant.logo_url" :alt="restaurant.name" />
          </div>
          <div class="card-content">
            <h3 class="restaurant-name">{{ restaurant.name }}</h3>
            <p class="restaurant-address">{{ restaurant.address }}</p>
            <p class="restaurant-description">{{ restaurant.description }}</p>
            <span class="restaurant-hours">
              🕒 {{ restaurant.opening_time }} - {{ restaurant.closing_time }}
            </span>
          </div>
        </RouterLink>
      </div>

      <!-- Empty State -->
      <div v-else class="empty-container">
        <p>No restaurants available at the moment.</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.restaurants-section {
  display: flex;
  flex-direction: column;
  gap: 40px;
  width: 100%;
  margin: 0 auto;
}

.header-section {
  text-align: left;
  margin-bottom: 20px;
}

.page-title {
  font-size: 32px;
  font-weight: 600;
  margin-bottom: 8px;
  color: var(--text-primary);
}

.page-subtitle {
  font-size: 16px;
  color: var(--text-muted);
}

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
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
}

/* Empty State */
.empty-container {
  text-align: center;
  padding: 60px 20px;
  color: var(--text-muted);
  font-size: 18px;
}

/* Restaurants Grid */
.restaurants-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
}

.restaurant-card {
  background-color: var(--bg-light);
  border: 1px solid var(--border);
  border-radius: 8px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  text-decoration: none;
  color: inherit;
}

.card-image-container {
  width: 100%;
  height: 200px;
  overflow: hidden;
  background-color: var(--border);
}

.restaurant-logo {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.card-content {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.restaurant-name {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
}

.restaurant-address {
  font-size: 13px;
  color: var(--text-muted);
}

.restaurant-description {
  font-size: 14px;
  line-height: 1.4;
  color: var(--text-primary);
  margin: 4px 0;
}

.restaurant-hours {
  font-size: 13px;
  color: var(--text-muted);
  margin-top: 4px;
}
</style>

<script setup lang="ts">
import { restaurantService } from '@/modules/restaurant/restaurant-service'
import type { Restaurant } from '@/modules/restaurant/types'
import type { ApiError } from '@/shared/types'
import { onMounted, ref } from 'vue'

const restaurants = ref<Restaurant[] | null>(null)
const loading = ref<boolean>(false)
const apiError = ref<string>('')
const fetchRestaurants = async () => {
  try {
    loading.value = true
    const response = await restaurantService.getRestaurants()
    restaurants.value = response.restaurants
  } catch (error) {
    apiError.value = (error as ApiError)?.response?.data?.error
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchRestaurants()
  console.log(restaurants)
})
</script>
