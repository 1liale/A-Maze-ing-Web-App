import { svelte } from '@sveltejs/vite-plugin-svelte';
import { defineConfig } from 'vite';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  resolve: {
    alias: {
      types: 'src/types',
      '@components': 'src/components',
      '@assets': 'src/assets',
      '@services': 'src/services',
    },
  },
  server: {
    host: true,
    port: 5173,
  },
});
