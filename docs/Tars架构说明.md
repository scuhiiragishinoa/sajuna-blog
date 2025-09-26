# Tars 微服务架构说明

## 概述

本项目已从基于 Gin 框架的单一服务架构迁移到基于 Tars 微服务框架的分布式架构。通过使用 TarsGateway 作为 API 网关，实现了服务的统一管理和路由分发。

## 架构组件

### 1. TarsGateway
- **作用**：API 网关，负责接收 HTTP 请求并路由到相应的 Tars 服务
- **端口**：8080
- **功能**：
  - 请求路由和负载均衡
  - 限流和熔断保护
  - 协议转换（HTTP 到 Tars RPC）
  - 统一鉴权和参数校验

### 2. Tars Registry
- **作用**：服务注册中心，管理所有 Tars 服务的注册和发现
- **端口**：17890
- **功能**：
  - 服务注册和发现
  - 健康检查
  - 配置管理
  - 监控数据收集

### 3. BlogServer
- **作用**：博客主服务，包含所有业务逻辑
- **端口**：10000
- **包含服务**：
  - UserService：用户管理
  - ArticleService：文章管理
  - CategoryService：分类管理
  - TagService：标签管理
  - CommentService：评论管理

## 服务接口定义

### 用户服务 (UserService)
```go
// 用户注册
int register(string username, string password, string email, out string message);

// 用户登录
int login(string username, string password, out string token, out string message);

// 获取用户信息
int getProfile(string token, out string userInfo, out string message);
```

### 文章服务 (ArticleService)
```go
// 获取文章列表
int getArticles(int page, int pageSize, string category, string tag, out string articles, out string message);

// 获取文章详情
int getArticle(int articleId, out string article, out string message);

// 创建文章
int createArticle(string token, string title, string content, string summary, int categoryId, string tags, out string message);

// 更新文章
int updateArticle(string token, int articleId, string title, string content, string summary, int categoryId, string tags, out string message);

// 删除文章
int deleteArticle(string token, int articleId, out string message);
```

### 分类服务 (CategoryService)
```go
// 获取分类列表
int getCategories(out string categories, out string message);

// 创建分类
int createCategory(string token, string name, string slug, out string message);
```

### 标签服务 (TagService)
```go
// 获取标签列表
int getTags(out string tags, out string message);

// 创建标签
int createTag(string token, string name, string slug, out string message);
```

### 评论服务 (CommentService)
```go
// 获取文章评论
int getComments(int articleId, int page, int pageSize, out string comments, out string message);

// 创建评论
int createComment(int articleId, string authorName, string authorEmail, string content, int parentId, out string message);

// 删除评论
int deleteComment(string token, int commentId, out string message);
```

## 配置文件

### 1. 服务配置 (sajuna-blog-server.conf)
```xml
<application>
    <server>
        app=sajuna-blog
        server=blog-server
        local=tcp -h 127.0.0.1 -p 10000 -t 60000
        logpath=/app/logs/
        logsize=10M
        lognum=10
        log=tars.tarslog.LogObj
        config=tars.tarsconfig.ConfigObj
        notify=tars.tarsnotify.NotifyObj
        deactivating-timeout=2000
        logLevel=DEBUG
    </server>
</application>
```

### 2. 客户端配置 (sajuna-blog-client.conf)
```xml
<application>
    <client>
        locator=tars.tarsregistry.QueryObj@tcp -h 127.0.0.1 -p 17890
        sync-invoke-timeout=3000
        async-invoke-timeout=5000
        refresh-endpoint-interval=60000
        report-interval=60000
        sample-rate=100000
        max-sample-count=100
        stat=tars.tarsstat.StatObj
        property=tars.tarsproperty.PropertyObj
        log=tars.tarslog.LogObj
        logpath=/app/logs/
        logsize=10M
        lognum=10
        logLevel=DEBUG
    </client>
</application>
```

### 3. 网关配置 (tars-gateway.conf)
```json
{
    "gateway": {
        "listen": "0.0.0.0:8080",
        "registry": "tars.tarsregistry.QueryObj@tcp -h tars-registry -p 17890",
        "timeout": 3000,
        "retry": 3,
        "rate_limit": {
            "enabled": true,
            "qps": 1000
        },
        "circuit_breaker": {
            "enabled": true,
            "failure_threshold": 5,
            "timeout": 30000
        }
    },
    "routes": [
        {
            "path": "/api/v1/users/*",
            "service": "sajuna_blog.blog_server.UserService",
            "method": "*"
        },
        {
            "path": "/api/v1/articles/*",
            "service": "sajuna_blog.blog_server.ArticleService",
            "method": "*"
        }
    ]
}
```

## 部署说明

### Docker Compose 服务
1. **mysql**：MySQL 数据库
2. **redis**：Redis 缓存
3. **tars-registry**：Tars 注册中心
4. **tars-gateway**：Tars 网关
5. **backend**：博客后端服务
6. **frontend**：前端服务

### 启动顺序
1. MySQL 和 Redis 先启动
2. Tars Registry 启动
3. Tars Gateway 启动
4. 博客后端服务启动
5. 前端服务启动

## 优势

1. **服务解耦**：各个业务模块独立部署和扩展
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
