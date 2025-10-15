import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import App from './App.vue'
import router from './router'

// 样式导入
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import '@/styles/index.scss'

// 创建 Vue 应用实例
const app = createApp(App)

// 注册 Element Plus 图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// 使用插件
app.use(createPinia())
app.use(router)
app.use(ElementPlus)

// 全局错误处理
app.config.errorHandler = (err, instance, info) => {
  console.error('Vue 错误:', err)
  console.error('组件实例:', instance)
  console.error('错误信息:', info)
  
  // 在生产环境中可以上报错误到服务器
  if (import.meta.env.PROD) {
    // 错误上报逻辑
    reportError(err, info)
  }
}

// 全局属性扩展
declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $filters: {
      formatDate: (date: string | Date) => string
      truncate: (text: string, length: number) => string
    }
  }
}

// 全局过滤器
app.config.globalProperties.$filters = {
  formatDate(date: string | Date): string {
    if (!date) return ''
    const d = new Date(date)
    return d.toLocaleDateString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit'
    })
  },
  
  truncate(text: string, length: number = 100): string {
    if (!text || text.length <= length) return text
    return text.substring(0, length) + '...'
  }
}

// 挂载应用
app.mount('#app')

// 错误上报函数
function reportError(error: unknown, info: string) {
  // 这里可以集成 Sentry、Baidu Tongji 等错误监控服务
  if (import.meta.env.PROD) {
    console.log('上报错误到监控服务:', { error, info })
    
    // 示例：发送错误到后端
    fetch('/api/error-log', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        error: String(error),
        info,
        url: window.location.href,
        userAgent: navigator.userAgent,
        timestamp: new Date().toISOString()
      })
    }).catch(console.error)
  }
}

// PWA 支持（可选）
if ('serviceWorker' in navigator && import.meta.env.PROD) {
  window.addEventListener('load', () => {
    navigator.serviceWorker.register('/sw.js')
      .then(registration => {
        console.log('SW registered: ', registration)
      })
      .catch(registrationError => {
        console.log('SW registration failed: ', registrationError)
      })
  })
}

// 性能监控（可选）
if (import.meta.env.PROD) {
  // 监听页面性能
  window.addEventListener('load', () => {
    setTimeout(() => {
      const perfData = performance.timing
      const loadTime = perfData.loadEventEnd - perfData.navigationStart
      
      console.log(`页面加载耗时: ${loadTime}ms`)
      
      // 上报性能数据
      fetch('/api/performance', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          loadTime,
          url: window.location.href,
          timestamp: new Date().toISOString()
        })
      }).catch(console.error)
    }, 0)
  })
}