import { post } from "@/utils/request";
import $router from "@/router";

export const login = async (params) => {
  await post("/login", params);
  await $router.push("/home");
};
