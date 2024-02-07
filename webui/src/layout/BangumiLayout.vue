<template>
  <a-layout>
    <BangumiSidebar />
    <a-layout-content style="padding: 24px">
      <router-view />
    </a-layout-content>
  </a-layout>
</template>
<script setup>
import BangumiSidebar from "@/components/BangumiSidebar.vue";
</script>
<style scoped></style>
