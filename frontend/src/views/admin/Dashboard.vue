<template>
  <div class="admin-dashboard">
    <el-container>
      <el-header>
        <h1>Sajuna Blog - 管理后台</h1>
        <el-menu mode="horizontal" :default-active="activeIndex">
          <el-menu-item index="1" @click="goHome">返回首页</el-menu-item>
          <el-menu-item index="2">仪表板</el-menu-item>
          <el-menu-item index="3">文章管理</el-menu-item>
          <el-menu-item index="4">用户管理</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <el-row :gutter="20">
          <el-col :span="6">
            <el-card>
              <h3>文章总数</h3>
              <p class="stat-number">{{ stats.articles }}</p>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card>
              <h3>用户总数</h3>
              <p class="stat-number">{{ stats.users }}</p>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card>
              <h3>评论总数</h3>
              <p class="stat-number">{{ stats.comments }}</p>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card>
              <h3>访问量</h3>
              <p class="stat-number">{{ stats.views }}</p>
            </el-card>
          </el-col>
        </el-row>
        
        <el-card style="margin-top: 20px;">
          <h3>快速操作</h3>
          <el-button type="primary">写新文章</el-button>
          <el-button type="success">管理分类</el-button>
          <el-button type="warning">系统设置</el-button>
        </el-card>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const activeIndex = ref('2')

const stats = reactive({
  articles: 0,
  users: 0,
  comments: 0,
  views: 0
})

const goHome = () => {
  router.push('/')
}
</script>

<style scoped>
.admin-dashboard {
  min-height: 100vh;
}

.el-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #f5f5f5;
  padding: 0 20px;
}

.el-main {
  padding: 20px;
}

.stat-number {
  font-size: 2em;
  font-weight: bold;
  color: #409eff;
  margin: 10px 0;
}
</style>



