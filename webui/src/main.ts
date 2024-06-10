import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import Antd from 'ant-design-vue'
import '@/router/auth'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'


createApp(App)
  .use(createPinia().use(piniaPluginPersistedstate))
  .use(router)
  .use(Antd)
  .mount('#app')
