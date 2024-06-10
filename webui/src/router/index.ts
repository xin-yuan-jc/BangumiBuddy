import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: "/login",
    name: "login",
    component: () => import("@v/LoginView.vue"),
  },
  {
    path: "/",
    redirect: "/home",
  },
  {
    path: "/home",
    name: "home",
    component: () => import("@v/HomeView.vue"),
  }
]

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: routes,
})


export default router
