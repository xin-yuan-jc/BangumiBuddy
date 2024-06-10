import router from '@/router/index'
import { useUserStore } from '@/stores/token'

router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  if (to.path == "/login") {
    if (!userStore.accessTokenExpired() || !userStore.refreshTokenExpired()) {
      next("/home")
    }else {
      next()
    }
  }else {
    if (userStore.accessTokenExpired() && userStore.refreshTokenExpired()) {
      next(`/login?redirect=${to.path}`)
    }else {
      next()
    }
  }
})