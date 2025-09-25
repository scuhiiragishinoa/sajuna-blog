<template>
  <div class="about">
    <el-container>
      <!-- 桌面端导航 -->
      <el-header class="desktop-nav">
        <div class="header-content">
          <h1 class="logo">Sajuna Blog</h1>
          <el-menu mode="horizontal" :default-active="activeIndex" class="desktop-menu">
            <el-menu-item index="1" @click="goHome">首页</el-menu-item>
            <el-menu-item index="2">文章</el-menu-item>
            <el-menu-item index="3">关于</el-menu-item>
            <el-menu-item index="4" @click="goAdmin">管理</el-menu-item>
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
          <el-card class="about-card">
            <h2>关于我们</h2>
            <p class="responsive-text">这是一个基于Vue3 + Go技术栈构建的现代化博客系统</p>
            
            <div class="tech-stack">
              <h3>技术特点：</h3>
              <div class="tech-grid">
                <div class="tech-item">
                  <h4>前端技术</h4>
                  <ul>
                    <li>Vue3 + TypeScript</li>
                    <li>Element Plus UI</li>
                    <li>响应式设计</li>
                  </ul>
                </div>
                <div class="tech-item">
                  <h4>后端技术</h4>
                  <ul>
                    <li>Go + Gin框架</li>
                    <li>GORM ORM</li>
                    <li>MySQL数据库</li>
                  </ul>
                </div>
                <div class="tech-item">
                  <h4>部署运维</h4>
                  <ul>
                    <li>Docker容器化</li>
                    <li>Nginx反向代理</li>
                    <li>自动化部署</li>
                  </ul>
                </div>
              </div>
            </div>

            <div class="features">
              <h3>功能特色</h3>
              <div class="feature-list">
                <div class="feature-item">
                  <el-icon class="feature-icon"><Document /></el-icon>
                  <div>
                    <h4>文章管理</h4>
                    <p>支持Markdown编辑，实时预览</p>
                  </div>
                </div>
                <div class="feature-item">
                  <el-icon class="feature-icon"><User /></el-icon>
                  <div>
                    <h4>用户系统</h4>
                    <p>完整的用户注册登录体系</p>
                  </div>
                </div>
                <div class="feature-item">
                  <el-icon class="feature-icon"><ChatDotRound /></el-icon>
                  <div>
                    <h4>评论系统</h4>
                    <p>支持文章评论和互动</p>
                  </div>
                </div>
                <div class="feature-item">
                  <el-icon class="feature-icon"><Monitor /></el-icon>
                  <div>
                    <h4>管理后台</h4>
                    <p>功能完善的后台管理系统</p>
                  </div>
                </div>
              </div>
            </div>
          </el-card>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Menu, Document, User, ChatDotRound, Monitor } from '@element-plus/icons-vue'

const router = useRouter()
const activeIndex = ref('3')
const mobileMenuOpen = ref(false)

const goHome = () => {
  router.push('/')
}

const goAdmin = () => {
  router.push('/admin')
}

const toggleMobileMenu = () => {
  mobileMenuOpen.value = !mobileMenuOpen.value
}

const navigateTo = (index: string) => {
  activeIndex.value = index
  mobileMenuOpen.value = false
  
  switch (index) {
    case '1':
      goHome()
      break
    case '2':
      // 文章页面
      console.log('跳转到文章列表')
      break
    case '3':
      // 已经在关于页面
      break
    case '4':
      goAdmin()
      break
  }
}
</script>

<style scoped>
.about {
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
  margin-top: 60px;
}

/* 关于卡片样式 */
.about-card h2 {
  color: #409eff;
  margin-bottom: 16px;
}

.tech-stack {
  margin: 24px 0;
}

.tech-stack h3 {
  color: #409eff;
  margin-bottom: 16px;
}

.tech-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-top: 16px;
}

.tech-item {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  border-left: 4px solid #409eff;
}

.tech-item h4 {
  color: #409eff;
  margin-bottom: 12px;
}

.tech-item ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.tech-item li {
  padding: 4px 0;
  color: #666;
}

.features {
  margin: 24px 0;
}

.features h3 {
  color: #409eff;
  margin-bottom: 16px;
}

.feature-list {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 16px;
}

.feature-item {
  display: flex;
  align-items: flex-start;
  padding: 16px;
  background: #f8f9fa;
  border-radius: 8px;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.feature-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.feature-icon {
  font-size: 24px;
  color: #409eff;
  margin-right: 12px;
  margin-top: 4px;
}

.feature-item h4 {
  color: #409eff;
  margin: 0 0 8px 0;
  font-size: 16px;
}

.feature-item p {
  color: #666;
  margin: 0;
  font-size: 14px;
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
  
  .tech-grid {
    grid-template-columns: 1fr;
    gap: 12px;
  }
  
  .feature-list {
    grid-template-columns: 1fr;
    gap: 12px;
  }
  
  .feature-item {
    padding: 12px;
  }
  
  .tech-item {
    padding: 16px;
  }
}

/* 平板端样式 */
@media (min-width: 769px) and (max-width: 1024px) {
  .tech-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .feature-list {
    grid-template-columns: repeat(2, 1fr);
  }
}

/* 桌面端样式 */
@media (min-width: 1025px) {
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



