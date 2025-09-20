# 🐳 Docker Compose 文件说明

## 📋 文件概览

| 文件名 | 用途 | 使用场景 |
|--------|------|----------|
| `docker-compose.yml` | **基础配置** | 生产环境基础配置 |
| `docker-compose.dev.yml` | **开发环境** | 本地开发使用 |
| `docker-compose.prod.yml` | **生产环境** | 生产部署使用 |
| `docker-compose.override.yml.example` | **覆盖配置模板** | 本地个性化配置 |

## 🔧 各文件详细说明

### 1. `docker-compose.yml` - 基础配置
**用途**: 生产环境的基础配置模板
**特点**:
- 使用预构建的镜像 (`build: ./backend`)
- 基础的安全配置
- 适合生产环境部署

**包含服务**:
- `mysql` - MySQL 8.0 数据库
- `redis` - Redis 7 缓存
- `backend` - Go 后端服务
- `frontend` - Vue3 前端服务

### 2. `docker-compose.dev.yml` - 开发环境
**用途**: 本地开发环境配置
**特点**:
- 使用开发镜像 (`golang:1.21-alpine`, `node:20-alpine`)
- 支持热重载和实时调试
- 包含 Nginx 反向代理
- 详细的日志输出

**包含服务**:
- `mysql` - 开发数据库
- `redis` - 开发缓存
- `backend-dev` - Go 开发服务
- `frontend-dev` - Vue3 开发服务
- `nginx` - 开发环境反向代理

**开发特性**:
- 代码热重载
- 调试模式
- 详细的错误信息
- 开发工具集成

### 3. `docker-compose.prod.yml` - 生产环境
**用途**: 生产环境完整配置
**特点**:
- 严格的安全配置
- 资源限制
- 只读文件系统
- 非 root 用户运行
- SSL/TLS 支持

**安全特性**:
- 用户权限限制
- 只读文件系统
- 资源使用限制
- 安全选项配置

### 4. `docker-compose.override.yml.example` - 覆盖配置模板
**用途**: 本地个性化配置模板
**特点**:
- 覆盖默认配置
- 设置敏感信息
- 不被 Git 跟踪

## 🚀 使用方法

### 开发环境
```bash
# 启动开发环境
docker-compose -f docker-compose.dev.yml up -d

# 查看日志
docker-compose -f docker-compose.dev.yml logs -f

# 停止服务
docker-compose -f docker-compose.dev.yml down
```

### 生产环境
```bash
# 启动生产环境
docker-compose -f docker-compose.prod.yml up -d

# 查看状态
docker-compose -f docker-compose.prod.yml ps

# 停止服务
docker-compose -f docker-compose.prod.yml down
```

### 本地个性化配置
```bash
# 1. 复制覆盖配置模板
cp docker-compose.override.yml.example docker-compose.override.yml

# 2. 编辑配置文件
# 修改密码、密钥等敏感信息

# 3. 启动服务（会自动应用覆盖配置）
docker-compose up -d
```

## 🔍 配置对比

| 特性 | 开发环境 | 生产环境 |
|------|----------|----------|
| **镜像类型** | 开发镜像 | 生产镜像 |
| **热重载** | ✅ 支持 | ❌ 不支持 |
| **调试模式** | ✅ 开启 | ❌ 关闭 |
| **安全配置** | 基础 | 严格 |
| **资源限制** | 无 | 有 |
| **用户权限** | root | 非root |
| **文件系统** | 可写 | 只读 |
| **日志级别** | 详细 | 精简 |

## 🛠️ 服务说明

### MySQL 数据库
- **开发**: `sajuna-blog-mysql-dev`
- **生产**: `sajuna-blog-mysql-prod`
- **端口**: 3306
- **数据持久化**: 使用 Docker Volume

### Redis 缓存
- **开发**: `sajuna-blog-redis-dev`
- **生产**: `sajuna-blog-redis-prod`
- **端口**: 6379
- **密码**: 生产环境需要密码

### 后端服务
- **开发**: `sajuna-blog-backend-dev`
- **生产**: `sajuna-blog-backend-prod`
- **端口**: 8080
- **语言**: Go + Gin

### 前端服务
- **开发**: `sajuna-blog-frontend-dev`
- **生产**: `sajuna-blog-frontend-prod`
- **端口**: 3000
- **框架**: Vue3 + TypeScript

### Nginx 代理
- **开发**: `sajuna-blog-nginx-dev`
- **生产**: `sajuna-blog-nginx-prod`
- **端口**: 80 (HTTP), 443 (HTTPS)
- **功能**: 反向代理、负载均衡

## 🔧 环境变量

### 数据库配置
```bash
MYSQL_ROOT_PASSWORD=your-root-password
MYSQL_USER=your-database-user
MYSQL_PASSWORD=your-database-password
```

### Redis 配置
```bash
REDIS_PASSWORD=your-redis-password
```

### 应用配置
```bash
JWT_SECRET=your-jwt-secret-key
GIN_MODE=release  # 或 debug
```

## 📝 最佳实践

1. **开发时使用** `docker-compose.dev.yml`
2. **生产部署使用** `docker-compose.prod.yml`
3. **敏感信息使用** `docker-compose.override.yml`
4. **定期备份** 数据库数据
5. **监控资源使用** 情况
6. **定期更新** 基础镜像

## 🚨 注意事项

- 生产环境必须设置强密码
- 定期更新 Docker 镜像
- 监控容器资源使用
- 备份重要数据
- 使用 HTTPS 证书
