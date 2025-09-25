<template>
  <el-header class="desktop-nav">
    <div class="header-content">
      <h1 class="logo">{{ title }}</h1>
      <el-menu 
        mode="horizontal" 
        :default-active="activeIndex" 
        class="desktop-menu"
        @select="handleMenuSelect"
      >
        <el-menu-item 
          v-for="item in menuItems" 
          :key="item.index"
          :index="item.index"
        >
          <el-icon v-if="item.icon" class="nav-icon">
            <component :is="item.icon" />
          </el-icon>
          {{ item.label }}
        </el-menu-item>
      </el-menu>
    </div>
  </el-header>
</template>

<script setup lang="ts">
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

const handleMenuSelect = (index: string) => {
  emit('navigate', index)
  const item = props.menuItems.find(item => item.index === index)
  if (item && item.action) {
    item.action()
  }
}
</script>

<style scoped>
.desktop-nav .el-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #f5f5f5;
  padding: 0 20px;
  height: 60px;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.logo {
  font-size: 24px;
  font-weight: bold;
  color: #409eff;
  margin: 0;
}

.desktop-menu {
  border-bottom: none;
}

.nav-icon {
  margin-right: 4px;
}

/* 移动端隐藏 */
@media (max-width: 768px) {
  .desktop-nav {
    display: none;
  }
}
</style>
