<template>
  <div class="responsive-test">
    <ResponsiveLayout
      title="响应式测试"
      mobile-title="测试"
      :menu-items="menuItems"
      :active-index="activeIndex"
      @navigate="handleNavigate"
    >
      <!-- 响应式网格测试 -->
      <div class="test-section">
        <h2>响应式网格测试</h2>
        <div class="responsive-grid">
          <el-card v-for="i in 6" :key="i" class="test-card">
            <h3>卡片 {{ i }}</h3>
            <p>这是一个测试卡片，用于验证响应式网格布局。</p>
          </el-card>
        </div>
      </div>

      <!-- 响应式按钮组测试 -->
      <div class="test-section">
        <h2>响应式按钮组测试</h2>
        <div class="button-group">
          <el-button type="primary" class="touch-target">主要按钮</el-button>
          <el-button type="success" class="touch-target">成功按钮</el-button>
          <el-button type="warning" class="touch-target">警告按钮</el-button>
          <el-button type="danger" class="touch-target">危险按钮</el-button>
        </div>
      </div>

      <!-- 响应式文本测试 -->
      <div class="test-section">
        <h2>响应式文本测试</h2>
        <p class="responsive-text">
          这是一段响应式文本，在移动端会显示较小的字体，在桌面端会显示较大的字体。
          这段文本用于测试响应式字体大小的效果。
        </p>
      </div>

      <!-- 响应式图片测试 -->
      <div class="test-section">
        <h2>响应式图片测试</h2>
        <div class="image-container">
          <img 
            src="https://via.placeholder.com/800x400/409eff/ffffff?text=响应式图片" 
            alt="响应式图片"
            class="responsive-image"
          />
        </div>
      </div>

      <!-- 断点测试 -->
      <div class="test-section">
        <h2>断点测试</h2>
        <div class="breakpoint-info">
          <p>当前屏幕宽度：<span class="screen-width">{{ screenWidth }}px</span></p>
          <p>当前断点：<span class="breakpoint">{{ currentBreakpoint }}</span></p>
        </div>
      </div>
    </ResponsiveLayout>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import ResponsiveLayout from '../components/ResponsiveLayout.vue'
import { Home, Document, User, Setting } from '@element-plus/icons-vue'

const router = useRouter()
const activeIndex = ref('test')
const screenWidth = ref(0)
const currentBreakpoint = ref('')

const menuItems = [
  { index: 'home', label: '首页', icon: Home, action: () => router.push('/') },
  { index: 'articles', label: '文章', icon: Document },
  { index: 'about', label: '关于', action: () => router.push('/about') },
  { index: 'admin', label: '管理', action: () => router.push('/admin') },
  { index: 'test', label: '测试', icon: Setting }
]

const updateScreenInfo = () => {
  screenWidth.value = window.innerWidth
  if (screenWidth.value <= 768) {
    currentBreakpoint.value = '移动端 (≤768px)'
  } else if (screenWidth.value <= 1024) {
    currentBreakpoint.value = '平板端 (769px-1024px)'
  } else {
    currentBreakpoint.value = '桌面端 (≥1025px)'
  }
}

const handleNavigate = (index: string) => {
  activeIndex.value = index
}

onMounted(() => {
  updateScreenInfo()
  window.addEventListener('resize', updateScreenInfo)
})

onUnmounted(() => {
  window.removeEventListener('resize', updateScreenInfo)
})
</script>

<style scoped>
.responsive-test {
  min-height: 100vh;
}

.test-section {
  margin-bottom: 32px;
}

.test-section h2 {
  color: #409eff;
  margin-bottom: 16px;
  border-bottom: 2px solid #409eff;
  padding-bottom: 8px;
}

.test-card {
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.test-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.test-card h3 {
  color: #409eff;
  margin-bottom: 8px;
}

.image-container {
  text-align: center;
  margin: 16px 0;
}

.breakpoint-info {
  background: #f8f9fa;
  padding: 16px;
  border-radius: 8px;
  border-left: 4px solid #409eff;
}

.breakpoint-info p {
  margin: 8px 0;
  font-size: 14px;
}

.screen-width {
  font-weight: bold;
  color: #409eff;
}

.breakpoint {
  font-weight: bold;
  color: #67c23a;
}

/* 移动端样式调整 */
@media (max-width: 768px) {
  .test-section {
    margin-bottom: 20px;
  }
  
  .test-section h2 {
    font-size: 18px;
  }
  
  .breakpoint-info {
    padding: 12px;
  }
}
</style>
