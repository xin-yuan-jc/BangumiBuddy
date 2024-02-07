<template>
  <a-layout-sider
    class="sidebar"
    v-model:collapsed="collapsed"
    :trigger="null"
    collapsible
  >
    <!-- Sidebar Menu -->
    <div class="logo"></div>
    <!-- Sidebar Menu -->
    <a-menu v-model:selectedKeys="selectedKeys" mode="inline" class="menu">
      <a-menu-item
        class="fold"
        key="fold"
        @click="() => (collapsed = !collapsed)"
      >
        <MenuUnfoldOutlined v-if="collapsed" class="trigger" />
        <MenuFoldOutlined v-else class="trigger" />
        <span> 菜单 </span>
      </a-menu-item>
      <a-menu-item key="home">
        <router-link to="/">
          <HomeOutlined />
          <span>主页</span>
        </router-link>
      </a-menu-item>
      <a-menu-item key="setting">
        <router-link to="/setting">
          <SettingOutlined />
          <span>设置</span>
        </router-link>
      </a-menu-item>
      <a-menu-item key="logout" class="logout">
        <LogoutOutlined />
        <span>退出</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>
<script setup>
import { ref } from "vue";
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  HomeOutlined,
  SettingOutlined,
  LogoutOutlined,
} from "@ant-design/icons-vue";
const collapsed = ref(false);
const selectedKeys = ref(["home"]);
</script>

<style lang="scss" scoped>
.sidebar {
  height: 100vh;
  width: 256px;
  background-image: linear-gradient(to bottom, #fbc2eb, #a6c1ee);
  .logo {
    height: 32px;
    margin: 16px;
    background: rgba(255, 255, 255, 0.3);
  }
  .menu {
    display: flex;
    flex-direction: column;
    height: 100%;
    background-image: linear-gradient(to bottom, #fbc2eb, #a6c1ee);
  }
}
</style>
