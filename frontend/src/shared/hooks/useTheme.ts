import { ref } from 'vue'

type Theme = 'light' | 'dark'

const theme = ref<Theme>('dark')

export function useTheme() {
  const setTheme = (newTheme: Theme) => {
    theme.value = newTheme
    document.documentElement.dataset.theme = newTheme
    localStorage.setItem('theme', newTheme)
  }

  const toggleTheme = () => {
    setTheme(theme.value === 'light' ? 'dark' : 'light')
    console.log(theme.value)
  }

  const initTheme = () => {
    const saved = localStorage.getItem('theme') as Theme | null
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
    setTheme(saved ?? (prefersDark ? 'dark' : 'light'))
  }

  return {
    theme,
    setTheme,
    toggleTheme,
    initTheme,
  }
}
