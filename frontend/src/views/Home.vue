<template>
  <div class="home">
    <el-container>
      <!-- 桌面端导航 -->
      <el-header class="desktop-nav">
        <div class="header-content">
          <h1 class="logo">Sajuna Blog</h1>
          <el-menu mode="horizontal" :default-active="activeIndex" class="desktop-menu">
            <el-menu-item index="1">首页</el-menu-item>
            <el-menu-item index="2">文章</el-menu-item>
            <el-menu-item index="3">关于</el-menu-item>
            <el-menu-item index="4">管理</el-menu-item>
          </el-menu>
        </div>
      </el-header>

      <!-- 移动端导航 -->
      <div class="mobile-nav">
        <div class="mobile-nav-header">
          <h1 class="mobile-logo">Sajuna Blog</h1>
          <el-button 
            type="text" 
            @click="toggleMobileMenu"
            class="mobile-menu-toggle"
          >
            <el-icon><Menu /></el-icon>
          </el-button>
        </div>
        <div class="mobile-nav-menu" :class="{ open: mobileMenuOpen }">
          <a href="#" class="mobile-nav-item" @click="navigateTo('1')">首页</a>
          <a href="#" class="mobile-nav-item" @click="navigateTo('2')">文章</a>
          <a href="#" class="mobile-nav-item" @click="navigateTo('3')">关于</a>
          <a href="#" class="mobile-nav-item" @click="navigateTo('4')">管理</a>
        </div>
      </div>

      <el-main class="responsive-padding">
        <div class="container">
          <el-card class="welcome-card">
            <h2>欢迎来到Sajuna Blog</h2>
            <p class="responsive-text">这是一个使用Vue3 + Go构建的私人博客系统</p>
            <div class="button-group">
              <el-button type="primary" @click="goToAdmin" class="touch-target">
                进入管理后台
              </el-button>
              <el-button type="success" @click="goToArticles" class="touch-target">
                浏览文章
              </el-button>
              <el-button type="info" @click="goToTest" class="touch-target">
                响应式测试
              </el-button>
            </div>
          </el-card>

          <!-- 响应式功能展示 -->
          <div class="responsive-grid" style="margin-top: 20px;">
            <el-card>
              <h3>最新文章</h3>
              <p>这里将显示最新的博客文章</p>
            </el-card>
            <el-card>
              <h3>热门标签</h3>
              <p>这里将显示热门标签</p>
            </el-card>
            <el-card>
              <h3>关于作者</h3>
              <p>这里将显示作者信息</p>
            </el-card>
          </div>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Menu } from '@element-plus/icons-vue'

const router = useRouter()
const activeIndex = ref('1')
const mobileMenuOpen = ref(false)

const goToAdmin = () => {
  router.push('/admin')
}

const goToArticles = () => {
  // 这里可以添加文章列表页面的路由
  console.log('跳转到文章列表')
}

const goToTest = () => {
  router.push('/test')
}

const toggleMobileMenu = () => {
  mobileMenuOpen.value = !mobileMenuOpen.value
}

const navigateTo = (index: string) => {
  activeIndex.value = index
  mobileMenuOpen.value = false
  
  switch (index) {
    case '1':
      // 已经在首页
      break
    case '2':
      goToArticles()
      break
    case '3':
      router.push('/about')
      break
    case '4':
      goToAdmin()
      break
  }
}
</script>

<style scoped>
.home {
  min-height: 100vh;
}

/* 桌面端头部样式 */
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

/* 移动端导航样式 */
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
  display: block;
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

/* 主内容区域 */
.el-main {
  padding: 0;
  margin-top: 60px; /* 为固定导航留出空间 */
}

/* 欢迎卡片样式 */
.welcome-card {
  margin-bottom: 20px;
}

.welcome-card h2 {
  color: #409eff;
  margin-bottom: 16px;
}

.button-group {
  display: flex;
  gap: 12px;
  margin-top: 20px;
}

/* 响应式网格卡片 */
.responsive-grid .el-card {
  text-align: center;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.responsive-grid .el-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.responsive-grid .el-card h3 {
  color: #409eff;
  margin-bottom: 12px;
}

/* 移动端样式调整 */
@media (max-width: 768px) {
  .desktop-nav {
    display: none;
  }
  
  .mobile-nav {
    display: block;
  }
  
  .el-main {
    margin-top: 60px;
    padding: 12px;
  }
  
  .button-group {
    flex-direction: column;
  }
  
  .button-group .el-button {
    width: 100%;
  }
  
  .welcome-card {
    margin-bottom: 12px;
  }
}

/* 桌面端样式 */
@media (min-width: 769px) {
  .desktop-nav {
    display: block;
  }
  
  .mobile-nav {
    display: none;
  }
  
  .el-main {
    margin-top: 0;
    padding: 20px;
  }
}
</style>



