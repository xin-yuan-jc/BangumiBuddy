import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers'


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [
        AntDesignVueResolver({
            importStyle: false,
          resolveIcons: true,
        }),
      ],
      dts: true,
      types: [{
        from: 'vue-router',
        names: ['RouterLink', 'RouterView'],
      }],
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
      '@v':fileURLToPath(new URL('./src/views', import.meta.url)),
    }
  },
  build:{
    outDir: "../backend/web"
  },
  server:{
    proxy: {
      '/apis': {
        target: "http://localhost:6937",
        changeOrigin: true,
      }
    }
  }
})
