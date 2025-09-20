# Sajuna Blog

一个基于 Vue3 + Go + MySQL 的现代化博客系统，支持 Docker 容器化部署。

## 🚀 快速开始

### 一键启动
```bash
# 启动开发环境
start.bat

# 停止开发环境  
stop.bat
```

### 访问地址
- 🌐 **完整应用**: http://localhost:80
- 🎨 **前端界面**: http://localhost:3000  
- ⚙️ **API接口**: http://localhost:8080

## 📁 项目结构

```
sajunaBlog/
├── 🎨 frontend/          # Vue3 前端
├── ⚙️ backend/           # Go 后端
├── 📋 configs/           # 配置文件
│   ├── init.sql         # 数据库初始化
│   ├── nginx-dev.conf   # 开发环境Nginx
│   └── nginx-prod.conf  # 生产环境Nginx
├── 🛠️ scripts/          # 管理脚本
│   ├── sajuna-blog.bat  # 完整管理脚本
│   ├── db-connect.bat   # 数据库连接
│   └── security-check.bat # 安全检查
├── 📚 docs/             # 文档
│   ├── Docker开发说明.md
│   ├── SECURITY.md
│   └── 项目结构说明.md
├── 🐳 docker-compose.yml      # 生产环境
├── 🐳 docker-compose.dev.yml  # 开发环境
├── 🚀 start.bat              # 快速启动
├── 🛑 stop.bat               # 快速停止
└── 📖 README.md              # 项目说明
```

## 🛠️ 管理工具

### 主要脚本
- `start.bat` - **一键启动**开发环境
- `stop.bat` - **一键停止**开发环境
- `scripts/sajuna-blog.bat` - **完整管理**脚本（启动/停止/日志/数据库等）

### 开发工具
- `scripts/db-connect.bat` - 安全连接数据库
- `scripts/security-check.bat` - 安全检查工具

## 🔧 技术栈

### 前端
- Vue 3 + TypeScript + Vite
- Pinia + Vue Router
- Element Plus + Markdown-it

### 后端  
- Go 1.21+ + Gin + GORM
- JWT + bcrypt + Swagger

### 数据库
- MySQL 8.0 + Redis

### 部署
- Docker + Docker Compose + Nginx

## 📖 详细文档

- [项目结构说明](docs/项目结构说明.md) - 详细的项目结构说明
- [Docker开发说明](docs/Docker开发说明.md) - Docker使用指南
- [安全配置说明](docs/SECURITY.md) - 安全配置指南

## 🔍 故障排除

### 常见问题
1. **端口冲突** - 确保 3000、8080、3306、6379 端口未被占用
2. **数据库连接失败** - 检查 MySQL 容器状态
3. **前端无法访问后端** - 检查 Docker 网络配置

### 查看日志
```bash
# 查看所有服务日志
docker-compose -f docker-compose.dev.yml logs -f

# 查看特定服务日志
docker-compose -f docker-compose.dev.yml logs -f backend-dev
```

## 📝 开发指南

### 环境变量
复制 `env.example` 为 `env.dev` 并修改配置：
```bash
cp env.example env.dev
# 编辑 env.dev 文件
```

### 数据库管理
```bash
# 使用安全连接脚本
scripts/db-connect.bat

# 或直接连接
docker exec -it sajuna-blog-mysql-dev mysql -u blog_user -p sajuna_blog
```

## 📄 许可证

MIT License