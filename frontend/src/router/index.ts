import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/Home.vue')
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/views/About.vue')
  },
  {
    path: '/admin',
    name: 'Admin',
    component: () => import('@/views/admin/Dashboard.vue')
  },
  {
    path: '/test',
    name: 'ResponsiveTest',
    component: () => import('@/views/ResponsiveTest.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router



