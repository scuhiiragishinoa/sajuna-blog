<template>
  <div class="responsive-layout">
    <el-container>
      <!-- 桌面端导航 -->
      <DesktopNav
        :title="title"
        :menu-items="menuItems"
        :active-index="activeIndex"
        @navigate="handleNavigate"
      />

      <!-- 移动端导航 -->
      <MobileNav
        :title="mobileTitle"
        :menu-items="menuItems"
        :active-index="activeIndex"
        @navigate="handleNavigate"
      />

      <!-- 主内容区域 -->
      <el-main class="responsive-main">
        <div class="container">
          <slot />
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import DesktopNav from './DesktopNav.vue'
import MobileNav from './MobileNav.vue'

interface MenuItem {
  index: string
  label: string
  icon?: any
  action?: () => void
}

interface Props {
  title: string
  mobileTitle?: string
  menuItems: MenuItem[]
  activeIndex: string
}

const props = withDefaults(defineProps<Props>(), {
  mobileTitle: ''
})

const emit = defineEmits<{
  navigate: [index: string]
}>()

const handleNavigate = (index: string) => {
  emit('navigate', index)
}
</script>

<style scoped>
.responsive-layout {
  min-height: 100vh;
}

.responsive-main {
  padding: 0;
  margin-top: 60px; /* 为固定导航留出空间 */
}

.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 16px;
}

/* 移动端样式调整 */
@media (max-width: 768px) {
  .responsive-main {
    margin-top: 60px;
    padding: 12px;
  }
  
  .container {
    padding: 0 12px;
  }
}

/* 桌面端样式 */
@media (min-width: 769px) {
  .responsive-main {
    margin-top: 0;
    padding: 20px;
  }
  
  .container {
    padding: 0 20px;
  }
}
</style>
