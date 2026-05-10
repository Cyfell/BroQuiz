import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import unocss from 'unocss/vite';
import { fileURLToPath, URL } from 'node:url';

const apiTarget = process.env.BROQUIZ_API_URL ?? 'http://localhost:8080';

export default defineConfig({
  plugins: [vue(), unocss()],
  resolve: {
    alias: {
      '@shared': fileURLToPath(new URL('./shared', import.meta.url)),
      '@client': fileURLToPath(new URL('./client', import.meta.url)),
    },
  },
  server: {
    port: 5173,
    proxy: {
      '/buzzer': apiTarget,
      '/teams': apiTarget,
      '/round': apiTarget,
      '/score': apiTarget,
      '/state': apiTarget,
      '/reset': apiTarget,
      '/lock': apiTarget,
      '/events': { target: apiTarget.replace(/^http/, 'ws'), ws: true },
    },
  },
  build: {
    outDir: 'dist',
    emptyOutDir: true,
  },
});
