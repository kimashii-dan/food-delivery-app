<template>
  <button @click="toggleTheme" class="theme-toggle">
    <Icon width="20" v-if="theme === 'light'" icon="majesticons:sun" color="var(--primary)" />
    <Icon width="20" v-else icon="majesticons:moon" color="var(--primary)" />
  </button>
</template>

<script setup lang="ts">
import { useTheme } from '../hooks/useTheme'
import { Icon } from '@iconify/vue'
const { theme, toggleTheme } = useTheme()
</script>

<style scoped>
.theme-toggle {
  background: var(--bg-light);
  border: 1px solid var(--border);
  display: flex;
  align-items: center;
  aspect-ratio: 1 / 1;
  padding: 8px;
}

.theme-toggle:hover {
  background: var(--highlight);
}
</style>
