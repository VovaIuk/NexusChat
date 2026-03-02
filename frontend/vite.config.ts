import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: true, // слушать на 0.0.0.0 — нужно для доступа из Docker
    allowedHosts: [
      'localhost',
      'localhost:5173',
      'xn----7sbbozvcgr0a7c4b.xn--p1ai',
      'xn----7sbbozvcgr0a7c4b.xn--p1ai:5173',
    ],
  },
})
