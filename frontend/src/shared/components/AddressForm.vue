<template>
  <div class="modal">
    <form class="address-form">
      <h1 class="title">Add address</h1>

      <div class="form-field">
        <label for="city">City</label>
        <input id="city" v-model="city" type="text" placeholder="Enter your city" />
      </div>
      <div class="form-field">
        <label for="street">Street</label>
        <input id="street" v-model="street" type="text" placeholder="Enter your street" />
      </div>
      <div class="form-field">
        <label for="postalCode">Postal code</label>
        <input
          id="postalCode"
          v-model="postalCode"
          type="text"
          placeholder="Enter your postal code"
        />
      </div>
      <div class="geolocation">
        <button type="button" class="getLocationButton" @click="getLocation">
          Get current location
        </button>
        <p class="explanation">We need your exact information to then sell it. Thanks.</p>
      </div>

      <div class="form-field">
        <label for="longitude">Longitude</label>
        <input id="longitude" v-model="longitude" type="text" placeholder="Enter your longitude" />
      </div>

      <div class="form-field">
        <label for="latitude">Latitude</label>
        <input id="latitude" v-model="latitude" type="text" placeholder="Enter your latitude" />
      </div>

      <div class="form-field">
        <label for="isDefault">Make this address default?</label>
        <input id="isDefault" v-model="isDefault" type="checkbox" class="isDefaultCheckbox" />
      </div>

      <button type="button" @click="submit">Submit</button>
      <button class="closeAddressForm" @click="closeAddressForm">×</button>
    </form>
  </div>
</template>
<script setup lang="ts">
import { addressService } from '@/modules/address/address-service'
import type { AddAddressRequest } from '@/modules/address/types'
import { ref, watch } from 'vue'

const isAddressFormOpened = defineModel()

const emit = defineEmits<{
  success: []
}>()

const closeAddressForm = () => {
  isAddressFormOpened.value = false
}

const city = ref<string>('')
const street = ref<string>('')
const postalCode = ref<string>('')
const longitude = ref<number>(0)
const latitude = ref<number>(0)
const isDefault = ref<boolean>(false)

const getLocation = () => {
  navigator.geolocation.getCurrentPosition((position) => {
    latitude.value = position.coords.latitude
    longitude.value = position.coords.longitude
    console.log('Location:', latitude.value, longitude.value)
  })
}

watch(longitude, () => [console.log(latitude.value, longitude.value)])

const submit = async () => {
  const addAddressRequest: AddAddressRequest = {
    city: city.value,
    street: street.value,
    postal_code: postalCode.value,
    longitude: longitude.value,
    latitude: latitude.value,
    is_default: isDefault.value,
  }

  console.log(addAddressRequest)

  try {
    const response = await addressService.addAddress(addAddressRequest)
    console.log(response)
    emit('success')
    closeAddressForm()
  } catch (error) {
    console.log(error)
  }
}
</script>
<style>
.address-form {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  gap: 24px;
  padding: 24px;
  border-radius: 12px;
  background-color: var(--bg-light);
  border: 1px solid var(--border);
  width: 384px;
  position: relative;
}

.form-field {
  display: flex;
  gap: 8px;
  flex-direction: column;
}

button.closeAddressForm {
  background-color: transparent;
  color: var(--danger);
  position: absolute;
  top: 0;
  right: 0;
  font-size: 24px;
}

.geolocation {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.explanation {
  color: var(--text-muted);
  font-size: 13px;
}

.isDefaultCheckbox {
  align-self: flex-start;
}

.getLocationButton {
  background-color: var(--info);
}
</style>
