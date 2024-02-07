import { createRouter, createWebHistory } from "vue-router";

let routes = [
  {
    path: "/login",
    name: "login",
    layout: "default",
    component: () => import("../views/LoginView.vue"),
  },
  {
    path: "/",
    name: "home",
    layout: "bangumi",
    component: () => import("../views/HomeView.vue"),
  },
  {
    path: "/setting",
    name: "setting",
    layout: "bangumi",
    component: () => import("../views/SettingView.vue"),
  },
];

function addLayoutToRoute(route, parentLayout = "default") {
  route.meta = route.meta || {};
  route.meta.layout = route.layout || parentLayout;
  if (route.children) {
    route.children = route.children.map((childRoute) =>
      addLayoutToRoute(childRoute, route.meta.layout)
    );
  }
  return route;
}

routes = routes.map((route) => addLayoutToRoute(route));

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
