<template>
  <div class="home">
    <el-container>
      <el-header>
        <h1>Sajuna Blog</h1>
        <el-menu mode="horizontal" :default-active="activeIndex">
          <el-menu-item index="1">首页</el-menu-item>
          <el-menu-item index="2">文章</el-menu-item>
          <el-menu-item index="3">关于</el-menu-item>
          <el-menu-item index="4">管理</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <el-card>
          <h2>欢迎来到Sajuna Blog</h2>
          <p>这是一个使用Vue3 + Go构建的私人博客系统</p>
          <el-button type="primary" @click="goToAdmin">进入管理后台</el-button>
        </el-card>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const activeIndex = ref('1')

const goToAdmin = () => {
  router.push('/admin')
}
</script>

<style scoped>
.home {
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
</style>



