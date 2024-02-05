import { svelte } from '@sveltejs/vite-plugin-svelte';
import path from 'path';
import { defineConfig } from 'vite';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  resolve: {
    alias: {
      types: path.resolve('src/types'),
      '@components': path.resolve('src/components'),
      '@assets': path.resolve('src/assets'),
      '@services': path.resolve('src/services'),
      '@stores': path.resolve('src/stores'),
    },
  },
  server: {
    host: true,
    port: 5173,
  },
});
