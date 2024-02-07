<template>
  <div id="app">
    <component :is="layout">
      <router-view />
    </component>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

const layout = computed(() => {
  return "layout-" + (route.meta.layout || "default").toLowerCase();
});
</script>

<style>
* {
  margin: 0;
  padding: 0;
  list-style: none;
  text-decoration: none;
  outline: none;
}

#app {
}
</style>
