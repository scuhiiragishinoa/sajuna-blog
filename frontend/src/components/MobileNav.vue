<template>
  <div class="mobile-nav">
    <div class="mobile-nav-header">
      <h1 class="mobile-logo">{{ title }}</h1>
      <el-button 
        type="text" 
        @click="toggleMobileMenu"
        class="mobile-menu-toggle"
      >
        <el-icon><Menu /></el-icon>
      </el-button>
    </div>
    <div class="mobile-nav-menu" :class="{ open: mobileMenuOpen }">
      <a 
        v-for="item in menuItems" 
        :key="item.index"
        href="#" 
        class="mobile-nav-item" 
        @click="handleNavClick(item)"
        :class="{ active: activeIndex === item.index }"
      >
        <el-icon v-if="item.icon" class="nav-icon">
          <component :is="item.icon" />
        </el-icon>
        {{ item.label }}
      </a>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Menu } from '@element-plus/icons-vue'

interface MenuItem {
  index: string
  label: string
  icon?: any
  action?: () => void
}

interface Props {
  title: string
  menuItems: MenuItem[]
  activeIndex: string
}

const props = defineProps<Props>()
const emit = defineEmits<{
  navigate: [index: string]
}>()

const mobileMenuOpen = ref(false)

const toggleMobileMenu = () => {
  mobileMenuOpen.value = !mobileMenuOpen.value
}

const handleNavClick = (item: MenuItem) => {
  mobileMenuOpen.value = false
  emit('navigate', item.index)
  if (item.action) {
    item.action()
  }
}
</script>

<style scoped>
.mobile-nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.mobile-nav-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid #e4e7ed;
  height: 60px;
}

.mobile-logo {
  font-size: 18px;
  font-weight: bold;
  color: #409eff;
  margin: 0;
}

.mobile-menu-toggle {
  font-size: 20px;
  color: #409eff;
}

.mobile-nav-menu {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background: #fff;
  border-top: 1px solid #e4e7ed;
  transform: translateY(-100%);
  transition: transform 0.3s ease;
  z-index: 999;
}

.mobile-nav-menu.open {
  transform: translateY(0);
}

.mobile-nav-item {
  display: flex;
  align-items: center;
  padding: 16px;
  border-bottom: 1px solid #f0f0f0;
  text-decoration: none;
  color: #333;
  transition: background-color 0.3s ease;
  cursor: pointer;
}

.mobile-nav-item:hover {
  background-color: #f5f7fa;
}

.mobile-nav-item.active {
  background-color: #e6f7ff;
  color: #409eff;
}

.nav-icon {
  margin-right: 8px;
  font-size: 16px;
}

/* 桌面端隐藏 */
@media (min-width: 769px) {
  .mobile-nav {
    display: none;
  }
}
</style>
