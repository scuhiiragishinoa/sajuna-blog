<template>
  <div class="admin-dashboard">
    <el-container>
      <!-- 桌面端导航 -->
      <el-header class="desktop-nav">
        <div class="header-content">
          <h1 class="logo">Sajuna Blog - 管理后台</h1>
          <el-menu mode="horizontal" :default-active="activeIndex" class="desktop-menu">
            <el-menu-item index="1" @click="goHome">返回首页</el-menu-item>
            <el-menu-item index="2">仪表板</el-menu-item>
            <el-menu-item index="3">文章管理</el-menu-item>
            <el-menu-item index="4">用户管理</el-menu-item>
          </el-menu>
        </div>
      </el-header>

      <!-- 移动端导航 -->
      <div class="mobile-nav">
        <div class="mobile-nav-header">
          <h1 class="mobile-logo">管理后台</h1>
          <el-button 
            type="text" 
            @click="toggleMobileMenu"
            class="mobile-menu-toggle"
          >
            <el-icon><Menu /></el-icon>
          </el-button>
        </div>
        <div class="mobile-nav-menu" :class="{ open: mobileMenuOpen }">
          <a href="#" class="mobile-nav-item" @click="navigateTo('1')">返回首页</a>
          <a href="#" class="mobile-nav-item" @click="navigateTo('2')">仪表板</a>
          <a href="#" class="mobile-nav-item" @click="navigateTo('3')">文章管理</a>
          <a href="#" class="mobile-nav-item" @click="navigateTo('4')">用户管理</a>
        </div>
      </div>

      <el-main class="responsive-padding">
        <div class="container">
          <!-- 统计卡片 -->
          <div class="stats-grid">
            <el-card class="stat-card">
              <div class="stat-content">
                <el-icon class="stat-icon"><Document /></el-icon>
                <div class="stat-info">
                  <h3>文章总数</h3>
                  <p class="stat-number">{{ stats.articles }}</p>
                </div>
              </div>
            </el-card>
            <el-card class="stat-card">
              <div class="stat-content">
                <el-icon class="stat-icon"><User /></el-icon>
                <div class="stat-info">
                  <h3>用户总数</h3>
                  <p class="stat-number">{{ stats.users }}</p>
                </div>
              </div>
            </el-card>
            <el-card class="stat-card">
              <div class="stat-content">
                <el-icon class="stat-icon"><ChatDotRound /></el-icon>
                <div class="stat-info">
                  <h3>评论总数</h3>
                  <p class="stat-number">{{ stats.comments }}</p>
                </div>
              </div>
            </el-card>
            <el-card class="stat-card">
              <div class="stat-content">
                <el-icon class="stat-icon"><View /></el-icon>
                <div class="stat-info">
                  <h3>访问量</h3>
                  <p class="stat-number">{{ stats.views }}</p>
                </div>
              </div>
            </el-card>
          </div>
          
          <!-- 快速操作 -->
          <el-card class="action-card">
            <h3>快速操作</h3>
            <div class="action-buttons">
              <el-button type="primary" class="touch-target">
                <el-icon><Edit /></el-icon>
                写新文章
              </el-button>
              <el-button type="success" class="touch-target">
                <el-icon><FolderOpened /></el-icon>
                管理分类
              </el-button>
              <el-button type="warning" class="touch-target">
                <el-icon><Setting /></el-icon>
                系统设置
              </el-button>
            </div>
          </el-card>

          <!-- 最近活动 -->
          <el-card class="activity-card">
            <h3>最近活动</h3>
            <div class="activity-list">
              <div class="activity-item">
                <el-icon class="activity-icon"><Document /></el-icon>
                <div class="activity-content">
                  <p>发布了新文章《Vue3响应式设计实践》</p>
                  <span class="activity-time">2小时前</span>
                </div>
              </div>
              <div class="activity-item">
                <el-icon class="activity-icon"><User /></el-icon>
                <div class="activity-content">
                  <p>新用户注册：张三</p>
                  <span class="activity-time">5小时前</span>
                </div>
              </div>
              <div class="activity-item">
                <el-icon class="activity-icon"><ChatDotRound /></el-icon>
                <div class="activity-content">
                  <p>收到新评论：李四在《Go语言入门》中评论</p>
                  <span class="activity-time">1天前</span>
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
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { 
  Menu, 
  Document, 
  User, 
  ChatDotRound, 
  View, 
  Edit, 
  FolderOpened, 
  Setting 
} from '@element-plus/icons-vue'

const router = useRouter()
const activeIndex = ref('2')
const mobileMenuOpen = ref(false)

const stats = reactive({
  articles: 12,
  users: 8,
  comments: 45,
  views: 1234
})

const goHome = () => {
  router.push('/')
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
      // 已经在仪表板
      break
    case '3':
      // 文章管理
      console.log('跳转到文章管理')
      break
    case '4':
      // 用户管理
      console.log('跳转到用户管理')
      break
  }
}
</script>

<style scoped>
.admin-dashboard {
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
  font-size: 20px;
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
  font-size: 16px;
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

/* 统计卡片网格 */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-bottom: 24px;
}

.stat-card {
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.stat-content {
  display: flex;
  align-items: center;
  gap: 16px;
}

.stat-icon {
  font-size: 32px;
  color: #409eff;
}

.stat-info h3 {
  color: #666;
  font-size: 14px;
  margin: 0 0 8px 0;
  font-weight: normal;
}

.stat-number {
  font-size: 24px;
  font-weight: bold;
  color: #409eff;
  margin: 0;
}

/* 操作卡片 */
.action-card {
  margin-bottom: 24px;
}

.action-card h3 {
  color: #409eff;
  margin-bottom: 16px;
}

.action-buttons {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
}

.action-buttons .el-button {
  flex: 1;
  min-width: 120px;
}

/* 活动列表 */
.activity-card h3 {
  color: #409eff;
  margin-bottom: 16px;
}

.activity-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.activity-item {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  padding: 12px;
  background: #f8f9fa;
  border-radius: 8px;
  transition: background-color 0.3s ease;
}

.activity-item:hover {
  background: #f0f2f5;
}

.activity-icon {
  font-size: 20px;
  color: #409eff;
  margin-top: 2px;
}

.activity-content {
  flex: 1;
}

.activity-content p {
  margin: 0 0 4px 0;
  color: #333;
  font-size: 14px;
}

.activity-time {
  color: #999;
  font-size: 12px;
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
  
  .stats-grid {
    grid-template-columns: 1fr;
    gap: 12px;
  }
  
  .action-buttons {
    flex-direction: column;
  }
  
  .action-buttons .el-button {
    width: 100%;
  }
  
  .stat-content {
    gap: 12px;
  }
  
  .stat-icon {
    font-size: 24px;
  }
  
  .stat-number {
    font-size: 20px;
  }
}

/* 平板端样式 */
@media (min-width: 769px) and (max-width: 1024px) {
  .stats-grid {
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



