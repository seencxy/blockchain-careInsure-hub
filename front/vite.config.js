import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': '/src', // 确保别名设置正确
    },
  },
  plugins: [vue()],
})
