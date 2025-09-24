# 安全配置说明

## 环境变量配置

### 开发环境
1. 复制 `env.example` 为 `env.dev`
2. 修改其中的密码和密钥为强密码
3. 运行 `docker-compose -f docker-compose.dev.yml up -d`

### 生产环境
1. 复制 `docker-compose.override.yml.example` 为 `docker-compose.override.yml`
2. 修改其中的密码和密钥为强密码
3. 运行 `docker-compose -f docker-compose.prod.yml up -d`

## 安全最佳实践

### 密码要求
- 至少12位字符
- 包含大小写字母、数字、特殊字符
- 定期更换密码
- 不同环境使用不同密码

### 密钥管理
- JWT密钥至少32位随机字符串
- 使用环境变量存储所有敏感信息
- 不要在代码中硬编码密码

### 安全检查
运行 `security-check.bat` 进行安全检查

## 文件说明

- `env.example` - 环境变量配置示例
- `docker-compose.override.yml.example` - Docker覆盖配置示例
- `security-check.bat` - 安全检查工具
- `db-connect.bat` - 安全数据库连接工具

## 注意事项

- `env.dev` - 开发环境配置（包含真实密码）
- `docker-compose.override.yml` - 生产环境配置（包含真实密码）
- `security-report.txt` - 安全检查报告

这些文件已被 `.gitignore` 忽略，不会意外提交到仓库。
