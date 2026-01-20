<template>
  <div class="page">
    <h1 class="page-title">Profile</h1>
    <div v-if="loading" class="skeleton">loading...</div>
    <div v-else-if="apiError" class="error">
      <p>Error: {{ apiError }}</p>
      <button @click="retry" type="button">Retry</button>
    </div>
    <div v-else class="content">
      <div class="profile-header">
        <div class="user-card">
          <div class="idk">
            <div class="user-details">
              <h2 class="user-name">{{ user?.name }}</h2>
              <p class="id">{{ user?.id }}</p>
            </div>
            <button @click="onLogout" class="danger logout-btn" type="button">Logout</button>
          </div>

          <div class="user-info">
            <div class="info-item">
              <span class="label">Email</span>
              <span class="value">{{ user?.email }}</span>
            </div>
            <div class="info-item">
              <span class="label">Phone</span>
              <span class="value">{{ user?.phone }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="addresses-section">
        <div class="titleAndButton">
          <h2 class="page-subtitle">Addresses</h2>
          <button @click="openAddressForm">Add address</button>
        </div>

        <div v-if="addresses.length > 0" class="addresses-grid">
          <div v-for="address in addresses" :key="address.id" class="address-card">
            <div class="address-header">
              <h3>{{ address.city }}</h3>
              <span v-if="address.is_default" class="default-badge">Default</span>
            </div>
            <p class="address-street">{{ address.street }}</p>
            <p class="address-postal">{{ address.postal_code }}</p>
            <div class="address-coords">
              <span class="id"
                >{{ address.latitude.toFixed(4) }}, {{ address.longitude.toFixed(4) }}</span
              >
            </div>
          </div>
        </div>
        <div v-else class="">
          <div class="no-addresses">
            <h2>You have no addresses</h2>
            <button @click="openAddressForm">Add address</button>
          </div>
        </div>

        <AddressForm
          v-if="isAddressFormOpened"
          v-model="isAddressFormOpened"
          @success="fetchUser"
        />
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { addressService } from '@/modules/address/address-service'
import type { Address } from '@/modules/address/types'
import { useAuthStore } from '@/modules/auth/auth-store'
import type { User } from '@/modules/user/types'
import { userService } from '@/modules/user/user-service'
import AddressForm from '@/shared/components/AddressForm.vue'
import type { ApiError } from '@/shared/types'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const { logout } = useAuthStore()
const user = ref<User | null>(null)
const addresses = ref<Address[]>([])
const loading = ref<boolean>(false)
const apiError = ref<string>('')
const isAddressFormOpened = ref<boolean>(false)

const fetchUser = async () => {
  try {
    loading.value = true
    const [userResponse, addressResponse] = await Promise.all([
      userService.getMe(),
      addressService.getAddresses(),
    ])

    user.value = userResponse.user
    addresses.value = addressResponse.addresses
  } catch (error) {
    apiError.value = (error as ApiError)?.response?.data?.error
  } finally {
    loading.value = false
  }
}

const openAddressForm = () => {
  isAddressFormOpened.value = true
}

onMounted(() => {
  fetchUser()
})

const retry = () => {
  fetchUser()
}

const onLogout = async () => {
  try {
    await logout()
    router.push('/')
  } catch (error) {
    console.log(error)
  }
}
</script>
<style scoped>
.error {
  color: var(--danger);
}

.profile-header {
  display: flex;
  gap: 16px;
  align-items: flex-start;
  justify-content: space-between;
}

.idk {
  display: flex;
  align-items: start;
  justify-content: space-between;
}

.user-card {
  flex: 1;
  padding: 32px;
  background-color: var(--bg-light);
  border: 1px solid var(--border);
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  gap: 36px;
}

.user-details {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.user-name {
  font-size: 28px;
  font-weight: 600;
}

.user-info {
  display: flex;
  justify-content: space-between;
  gap: 20px;
}

.info-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.label {
  font-size: 14px;
  color: var(--text-muted);
}

.value {
  font-size: 16px;
}

.id {
  color: var(--text-muted);
  font-size: 14px;
}

.logout-btn {
  height: fit-content;
}

.addresses-section {
  display: flex;
  flex-direction: column;
  gap: 32px;
}

.titleAndButton {
  display: flex;
  justify-content: space-between;
}

.addresses-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.address-card {
  padding: 24px;
  background-color: var(--bg-light);
  border: 1px solid var(--border);
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.address-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.address-header h3 {
  font-size: 20px;
  font-weight: 600;
}

.default-badge {
  padding: 4px 12px;
  background-color: var(--info);
  border-radius: 6px;
  font-size: 12px;
  font-weight: 500;
}

.address-street {
  font-size: 16px;
}

.address-postal {
  font-size: 14px;
  color: var(--text-muted);
}

.address-coords {
  margin-top: 8px;
  padding-top: 12px;
  border-top: 1px solid var(--border-muted);
}

.no-addresses {
  height: 250px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  gap: 16px;
  padding: 32px;
  background-color: var(--bg-light);
  border: 1px solid var(--border);
  border-radius: 12px;
}
</style>
