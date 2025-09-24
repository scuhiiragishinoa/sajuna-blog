# Sajuna Blog API 文档

## 概述

Sajuna Blog 是一个基于 Go + Vue3 构建的私人博客系统，提供完整的博客管理功能。本文档详细描述了系统的 RESTful API 接口。

**基础信息：**
- 基础URL: `http://localhost:8080/api/v1`
- 认证方式: JWT Bearer Token
- 数据格式: JSON
- 字符编码: UTF-8

## 认证说明

大部分API需要JWT认证，请在请求头中包含：
```
Authorization: Bearer <your_jwt_token>
```

## 通用响应格式

### 成功响应
```json
{
  "message": "操作成功",
  "data": { ... }
}
```

### 错误响应
```json
{
  "error": "错误信息"
}
```

### 分页响应
```json
{
  "data": [...],
  "total": 100,
  "page": 1,
  "page_size": 10
}
```

## API 接口

### 1. 健康检查

#### 检查服务状态
- **URL:** `GET /health`
- **描述:** 检查服务是否正常运行
- **认证:** 无需认证
- **响应示例:**
```json
{
  "status": "ok",
  "timestamp": "2024-01-01T00:00:00Z"
}
```

---

### 2. 用户管理

#### 2.1 用户注册
- **URL:** `POST /users/register`
- **描述:** 注册新用户
- **认证:** 无需认证
- **请求体:**
```json
{
  "username": "string (必填, 唯一, 英文+数字混合字符串)",
  "email": "string (必填, 邮箱格式, 唯一)",
  "password": "string (必填, 最少6位)",
  "nickname": "string (必填, 不能为空)"
}
```
- **响应Payload:**
```json
{
  "message": "success",
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "nickname": "测试用户"
  }
}
```

#### 2.2 用户登录
- **URL:** `POST /users/login`
- **描述:** 用户登录获取JWT token
- **认证:** 无需认证
- **请求体:**
```json
{
  "username": "string (必填)",
  "password": "string (必填)"
}
```
- **响应示例:**
```json
{
  "message": "success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "nickname": "测试用户",
    "role": "user"
  }
}
```

#### 2.3 获取用户信息
- **URL:** `GET /users/profile`
- **描述:** 获取当前登录用户的信息
- **认证:** 需要JWT认证
- **响应示例:**
```json
{
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "nickname": "测试用户",
    "avatar": "",
    "bio": "",
    "role": "user",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 2.4 忘记密码验证
- **URL:** `POST /users/forget_password_identify`
- **描述:** 验证用户名和邮箱是否匹配，确认用户身份
- **认证:** 无需认证
- **请求体:**
```json
{
  "username": "string (必填)",
  "email": "string (必填, 邮箱格式)"
}
```
- **响应示例:**
```json
{
  "message": "success",
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "nickname": "测试用户"
  }
}
```
- **错误响应:**
```json
{
  "error": "用户名或邮箱不匹配"
}
```

#### 2.5 发送忘记密码验证码
- **URL:** `POST /users/send_reset_code`
- **描述:** 向指定邮箱发送密码重置验证码
- **认证:** 无需认证
- **请求体:**
```json
{
  "email": "string (必填, 邮箱格式)"
}
```
- **响应示例:**
```json
{
  "message": "success",
  "expires_in": 300
}
```
- **错误响应:**
```json
{
  "error": "邮箱不存在或发送失败"
}
```

#### 2.6 验证重置码并修改密码
- **URL:** `POST /users/reset_password`
- **描述:** 验证重置码并设置新密码
- **认证:** 无需认证
- **请求体:**
```json
{
  "email": "string (必填, 邮箱格式)",
  "code": "string (必填, 6位数字验证码)",
  "new_password": "string (必填, 最少6位)"
}
```
- **响应示例:**
```json
{
  "message": "success",
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com"
  }
}
```
- **错误响应:**
```json
{
  "error": "验证码错误或已过期"
}
```

#### 2.7 忘记密码后登录
- **URL:** `POST /users/login`
- **描述:** 使用新密码登录（与普通登录接口相同）
- **认证:** 无需认证
- **请求体:**
```json
{
  "username": "string (必填)",
  "password": "string (必填, 新密码)"
}
```
- **响应示例:**
```json
{
  "message": "登录成功",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "nickname": "测试用户",
    "role": "user"
  }
}
```

---

### 3. 文章管理

#### 3.1 获取文章列表
- **URL:** `GET /articles`
- **描述:** 分页获取文章列表
- **认证:** 无需认证
- **查询参数:**
  - `page` (int, 可选): 页码，默认1
  - `page_size` (int, 可选): 每页数量，默认10
  - `status` (string, 可选): 文章状态，默认"published"
  - `category_id` (int, 可选): 分类ID筛选
- **响应示例:**
```json
{
  "articles": [
    {
      "id": 1,
      "title": "文章标题",
      "slug": "article-slug",
      "content": "文章内容...",
      "excerpt": "文章摘要",
      "cover_image": "封面图片URL",
      "status": "published",
      "view_count": 100,
      "like_count": 10,
      "is_top": false,
      "published_at": "2024-01-01T00:00:00Z",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z",
      "user": {
        "id": 1,
        "username": "author",
        "nickname": "作者"
      },
      "category": {
        "id": 1,
        "name": "技术",
        "slug": "tech"
      },
      "tags": [
        {
          "id": 1,
          "name": "Vue",
          "slug": "vue"
        }
      ]
    }
  ],
  "total": 100,
  "page": 1,
  "page_size": 10
}
```

#### 3.2 获取单篇文章
- **URL:** `GET /articles/{id}`
- **描述:** 根据ID获取文章详情
- **认证:** 无需认证
- **路径参数:**
  - `id` (int): 文章ID
- **响应示例:**
```json
{
  "article": {
    "id": 1,
    "title": "文章标题",
    "slug": "article-slug",
    "content": "文章内容...",
    "excerpt": "文章摘要",
    "cover_image": "封面图片URL",
    "status": "published",
    "view_count": 101,
    "like_count": 10,
    "is_top": false,
    "published_at": "2024-01-01T00:00:00Z",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z",
    "user": {
      "id": 1,
      "username": "author",
      "nickname": "作者"
    },
    "category": {
      "id": 1,
      "name": "技术",
      "slug": "tech"
    },
    "tags": [
      {
        "id": 1,
        "name": "Vue",
        "slug": "vue"
      }
    ],
    "comments": [
      {
        "id": 1,
        "content": "评论内容",
        "status": "approved",
        "created_at": "2024-01-01T00:00:00Z",
        "user": {
          "id": 2,
          "username": "commenter",
          "nickname": "评论者"
        }
      }
    ]
  }
}
```

#### 3.3 创建文章
- **URL:** `POST /articles`
- **描述:** 创建新文章
- **认证:** 需要JWT认证
- **请求体:**
```json
{
  "title": "string (必填)",
  "content": "string (必填)",
  "excerpt": "string (可选)",
  "cover_image": "string (可选)",
  "category_id": "number (可选)",
  "tag_ids": [1, 2, 3],
  "status": "string (可选, draft/published/archived)"
}
```
- **响应示例:**
```json
{
  "message": "success",
  "article": {
    "id": 1,
    "title": "新文章",
    "content": "文章内容...",
    "status": "draft",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 3.4 更新文章
- **URL:** `PUT /articles/{id}`
- **描述:** 更新文章信息
- **认证:** 需要JWT认证
- **路径参数:**
  - `id` (int): 文章ID
- **请求体:**
```json
{
  "title": "string (可选)",
  "content": "string (可选)",
  "excerpt": "string (可选)",
  "cover_image": "string (可选)",
  "category_id": "number (可选)",
  "tag_ids": [1, 2, 3],
  "status": "string (可选)",
  "is_top": "boolean (可选)"
}
```
- **响应示例:**
```json
{
  "message": "success",
  "article": {
    "id": 1,
    "title": "更新后的标题",
    "content": "更新后的内容...",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 3.5 删除文章
- **URL:** `DELETE /articles/{id}`
- **描述:** 删除文章
- **认证:** 需要JWT认证
- **路径参数:**
  - `id` (int): 文章ID
- **响应示例:**
```json
{
  "message": "success"
}
```

---

### 4. 分类管理

#### 4.1 获取分类列表
- **URL:** `GET /categories`
- **描述:** 获取所有分类
- **认证:** 无需认证
- **响应示例:**
```json
{
  "categories": [
    {
      "id": 1,
      "name": "技术",
      "slug": "tech",
      "description": "技术相关文章",
      "color": "#409eff",
      "sort": 1,
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### 4.2 创建分类
- **URL:** `POST /categories`
- **描述:** 创建新分类
- **认证:** 需要JWT认证
- **请求体:**
```json
{
  "name": "string (必填)",
  "slug": "string (必填)",
  "description": "string (可选)",
  "color": "string (可选)",
  "sort": "number (可选)"
}
```
- **响应示例:**
```json
{
  "message": "success",
  "category": {
    "id": 1,
    "name": "新分类",
    "slug": "new-category",
    "description": "分类描述",
    "color": "#409eff",
    "sort": 0,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

### 5. 标签管理

#### 5.1 获取标签列表
- **URL:** `GET /tags`
- **描述:** 获取所有标签
- **认证:** 无需认证
- **响应示例:**
```json
{
  "tags": [
    {
      "id": 1,
      "name": "Vue",
      "slug": "vue",
      "color": "#4fc08d",
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### 5.2 创建标签
- **URL:** `POST /tags`
- **描述:** 创建新标签
- **认证:** 需要JWT认证
- **请求体:**
```json
{
  "name": "string (必填)",
  "slug": "string (必填)",
  "color": "string (可选)"
}
```
- **响应示例:**
```json
{
  "message": "success",
  "tag": {
    "id": 1,
    "name": "新标签",
    "slug": "new-tag",
    "color": "#409eff",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

---

## 数据模型

### 用户模型 (User)
```json
{
  "id": "number (主键)",
  "username": "string (用户名, 唯一)",
  "email": "string (邮箱, 唯一)",
  "password": "string (密码, 加密存储)",
  "nickname": "string (昵称)",
  "avatar": "string (头像URL)",
  "bio": "string (个人简介)",
  "role": "string (角色: user/admin)",
  "status": "string (状态: active/inactive)",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### 文章模型 (Article)
```json
{
  "id": "number (主键)",
  "title": "string (标题)",
  "slug": "string (URL别名, 唯一)",
  "content": "string (内容)",
  "excerpt": "string (摘要)",
  "cover_image": "string (封面图片)",
  "status": "string (状态: draft/published/archived)",
  "view_count": "number (浏览量)",
  "like_count": "number (点赞数)",
  "is_top": "boolean (是否置顶)",
  "published_at": "datetime (发布时间)",
  "user_id": "number (作者ID)",
  "category_id": "number (分类ID)",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### 分类模型 (Category)
```json
{
  "id": "number (主键)",
  "name": "string (分类名, 唯一)",
  "slug": "string (URL别名, 唯一)",
  "description": "string (描述)",
  "color": "string (颜色)",
  "sort": "number (排序)",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### 标签模型 (Tag)
```json
{
  "id": "number (主键)",
  "name": "string (标签名, 唯一)",
  "slug": "string (URL别名, 唯一)",
  "color": "string (颜色)",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### 评论模型 (Comment)
```json
{
  "id": "number (主键)",
  "content": "string (评论内容)",
  "status": "string (状态: pending/approved/rejected)",
  "ip": "string (IP地址)",
  "user_agent": "string (用户代理)",
  "article_id": "number (文章ID)",
  "user_id": "number (用户ID, 可选)",
  "parent_id": "number (父评论ID, 可选)",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### 验证码模型 (VerificationCode)
```json
{
  "id": "number (主键)",
  "email": "string (邮箱地址)",
  "code": "string (6位数字验证码)",
  "type": "string (验证码类型: password_reset/email_verify)",
  "expires_at": "datetime (过期时间)",
  "used": "boolean (是否已使用)",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

---

## 错误码说明

| HTTP状态码 | 说明 |
|-----------|------|
| 200 | 请求成功 |
| 400 | 请求参数错误 |
| 401 | 未授权，需要登录 |
| 403 | 禁止访问，权限不足 |
| 404 | 资源不存在 |
| 409 | 资源冲突（如用户名已存在） |
| 500 | 服务器内部错误 |

---

## 使用示例

### 前端调用示例 (JavaScript)

```javascript
// 用户登录
const login = async (username, password) => {
  const response = await fetch('/api/v1/users/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ username, password })
  });
  return await response.json();
};

// 获取文章列表
const getArticles = async (page = 1, pageSize = 10) => {
  const response = await fetch(`/api/v1/articles?page=${page}&page_size=${pageSize}`);
  return await response.json();
};

// 创建文章（需要认证）
const createArticle = async (articleData, token) => {
  const response = await fetch('/api/v1/articles', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
    },
    body: JSON.stringify(articleData)
  });
  return await response.json();
};

// 忘记密码 - 验证用户身份
const forgetPasswordIdentify = async (username, email) => {
  const response = await fetch('/api/v1/users/forget_password_identify', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ username, email })
  });
  return await response.json();
};

// 发送重置密码验证码
const sendResetCode = async (email) => {
  const response = await fetch('/api/v1/users/send_reset_code', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ email })
  });
  return await response.json();
};

// 重置密码
const resetPassword = async (email, code, newPassword) => {
  const response = await fetch('/api/v1/users/reset_password', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ email, code, new_password: newPassword })
  });
  return await response.json();
};
```

### cURL 示例

```bash
# 用户登录
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'

# 获取文章列表
curl -X GET http://localhost:8080/api/v1/articles?page=1&page_size=10

# 创建文章（需要认证）
curl -X POST http://localhost:8080/api/v1/articles \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"title":"新文章","content":"文章内容","status":"draft"}'

# 忘记密码 - 验证用户身份
curl -X POST http://localhost:8080/api/v1/users/forget_password_identify \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com"}'

# 发送重置密码验证码
curl -X POST http://localhost:8080/api/v1/users/send_reset_code \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com"}'

# 重置密码
curl -X POST http://localhost:8080/api/v1/users/reset_password \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","code":"123456","new_password":"newpassword123"}'
```

---

## 注意事项

1. **认证要求**: 大部分管理操作需要JWT认证，请确保在请求头中包含有效的token
2. **数据验证**: 所有必填字段都会进行验证，请确保请求数据格式正确
3. **分页限制**: 建议单次请求的文章数量不超过100条
4. **文件上传**: 当前版本暂不支持文件上传，图片URL需要外部提供
5. **权限控制**: 用户只能管理自己创建的文章
6. **软删除**: 删除操作采用软删除，数据不会真正从数据库中移除
7. **忘记密码流程**: 
   - 验证码有效期为5分钟
   - 每个邮箱每分钟最多发送1次验证码
   - 验证码使用后立即失效
   - 用户名和邮箱必须匹配才能进行密码重置
8. **安全建议**: 
   - 验证码仅用于密码重置，请勿泄露给他人
   - 新密码建议使用强密码（包含字母、数字、特殊字符）
   - 重置密码后建议立即登录验证

---

## 更新日志

- **v1.1.0** (2024-01-15): 新增忘记密码功能，包含身份验证、验证码发送、密码重置等API
- **v1.0.0** (2024-01-01): 初始版本，包含基础的用户、文章、分类、标签管理功能
