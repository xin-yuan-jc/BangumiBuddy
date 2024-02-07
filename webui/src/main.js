import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Antd from "ant-design-vue";
import "ant-design-vue/dist/reset.css";
import DefaultLayout from "./layout/DefaultLayout.vue";
import BangumiLayout from "./layout/BangumiLayout.vue";

createApp(App)
  .use(store)
  .use(router)
  .use(Antd)
  .component("layout-default", DefaultLayout)
  .component("layout-bangumi", BangumiLayout)
  .mount("#app");
