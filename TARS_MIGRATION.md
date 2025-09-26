# Tars 微服务架构迁移说明

## 概述

本项目已成功从基于 Gin 框架的单一服务架构迁移到基于 Tars 微服务框架的分布式架构。通过使用 TarsGateway 作为 API 网关，实现了服务的统一管理和路由分发。

## 主要变更

### 1. 架构变更
- **之前**：Gin + GORM 单一服务
- **现在**：Tars 微服务框架 + TarsGateway

### 2. 服务组件
- **Tars Registry**：服务注册中心 (端口: 17890)
- **TarsGateway**：API 网关 (端口: 8080)
- **BlogServer**：博客主服务 (端口: 10000)
  - UserService：用户管理
  - ArticleService：文章管理
  - CategoryService：分类管理
  - TagService：标签管理
  - CommentService：评论管理

### 3. 文件结构变更

#### 新增文件
```
backend/
├── api/
│   └── blog.tars                    # Tars 接口定义
├── cmd/
│   └── blog_server/
│       ├── main.go                  # Tars 服务主程序
│       └── BlogServer.go           # Tars 服务接口
├── internal/
│   └── service/                     # 业务服务层
│       ├── user_service.go
│       ├── article_service.go
│       ├── category_service.go
│       ├── tag_service.go
│       └── comment_service.go
├── configs/
│   ├── sajuna-blog-server.conf     # Tars 服务配置
│   └── sajuna-blog-client.conf     # Tars 客户端配置
└── internal/utils/
    └── jwt.go                      # JWT 工具类
```

#### 修改文件
- `go.mod`：添加 Tars 框架依赖
- `docker-compose.yml`：添加 Tars 相关服务
- `Dockerfile`：修改构建目标为 Tars 服务
- `docs/sajuna_blog开发文档.md`：更新架构说明

## 启动方式

### 使用脚本启动（推荐）
```bash
# Windows
start-tars.bat

# Linux/Mac
./start-tars.sh
```

### 手动启动
```bash
# 1. 启动基础服务
docker-compose up -d mysql redis

# 2. 等待数据库启动后，启动 Tars Registry
docker-compose up -d tars-registry

# 3. 启动 Tars Gateway
docker-compose up -d tars-gateway

# 4. 启动博客服务
docker-compose up -d backend

# 5. 启动前端服务
docker-compose up -d frontend
```

## 服务访问

- **前端界面**：http://localhost:3000
- **API 网关**：http://localhost:8080
- **Tars Registry**：http://localhost:17890

## API 接口

所有 API 接口保持不变，通过 TarsGateway 进行路由：

- `GET /api/v1/articles` - 获取文章列表
- `GET /api/v1/articles/:id` - 获取文章详情
- `POST /api/v1/articles` - 创建文章
- `PUT /api/v1/articles/:id` - 更新文章
- `DELETE /api/v1/articles/:id` - 删除文章
- `GET /api/v1/categories` - 获取分类列表
- `GET /api/v1/tags` - 获取标签列表
- `POST /api/v1/users/login` - 用户登录
- `POST /api/v1/users/register` - 用户注册

## 配置说明

### Tars 服务配置
- 服务名：`sajuna_blog.blog_server`
- 端口：10000
- 注册中心：`tars-registry:17890`

### 网关路由配置
网关根据路径前缀路由到相应的服务：
- `/api/v1/users/*` → UserService
- `/api/v1/articles/*` → ArticleService
- `/api/v1/categories/*` → CategoryService
- `/api/v1/tags/*` → TagService
- `/api/v1/comments/*` → CommentService

## 监控和管理

通过 Tars 平台可以：
- 查看服务状态和健康检查
- 监控服务性能指标
- 管理服务配置
- 查看服务日志

## 优势

1. **服务解耦**：各业务模块独立部署和扩展
2. **统一管理**：通过 Tars 平台统一管理所有服务
3. **高可用**：支持服务注册发现和负载均衡
4. **监控完善**：内置监控和日志收集
5. **扩展性强**：易于添加新的微服务

## 注意事项

1. 确保 Tars Registry 服务正常运行
2. 服务注册需要正确的服务名称和端口
3. 网关路由配置需要与实际服务接口匹配
4. 数据库连接配置需要正确
5. 日志路径需要确保有写入权限

## 故障排除

### 服务启动失败
```bash
# 查看服务日志
docker-compose logs [服务名]

# 查看服务状态
docker-compose ps
```

### 服务注册失败
1. 检查 Tars Registry 是否正常运行
2. 检查服务配置中的注册中心地址
3. 检查网络连接

### API 请求失败
1. 检查 TarsGateway 是否正常运行
2. 检查路由配置是否正确
3. 检查后端服务是否正常注册
