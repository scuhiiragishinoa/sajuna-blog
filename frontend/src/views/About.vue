<template>
  <div class="about">
    <el-container>
      <el-header>
        <h1>Sajuna Blog</h1>
        <el-menu mode="horizontal" :default-active="activeIndex">
          <el-menu-item index="1" @click="goHome">首页</el-menu-item>
          <el-menu-item index="2">文章</el-menu-item>
          <el-menu-item index="3">关于</el-menu-item>
          <el-menu-item index="4" @click="goAdmin">管理</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <el-card>
          <h2>关于我们</h2>
          <p>这是一个基于Vue3 + Go技术栈构建的现代化博客系统</p>
          <h3>技术特点：</h3>
          <ul>
            <li>前端：Vue3 + TypeScript + Element Plus</li>
            <li>后端：Go + Gin + GORM</li>
            <li>数据库：MySQL</li>
            <li>部署：Docker + Nginx</li>
          </ul>
        </el-card>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const activeIndex = ref('3')

const goHome = () => {
  router.push('/')
}

const goAdmin = () => {
  router.push('/admin')
}
</script>

<style scoped>
.about {
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

ul {
  margin: 20px 0;
  padding-left: 20px;
}

li {
  margin: 10px 0;
}
</style>



