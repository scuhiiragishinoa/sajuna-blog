# 依赖安装阶段
FROM node:20-alpine AS deps

WORKDIR /app

# 配置 Alpine 镜像源和 npm 镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    npm config set registry https://registry.npmmirror.com

# 复制包管理文件
COPY package.json package-lock.json* ./

# 清理 npm 缓存并安装依赖
RUN npm ci --only=production && npm cache clean --force

# 构建阶段
FROM node:20-alpine AS builder

WORKDIR /app

# 配置 npm 镜像源
RUN npm config set registry https://registry.npmmirror.com

# 从 deps 阶段复制 node_modules
COPY --from=deps /app/node_modules ./node_modules

# 复制源代码
COPY . .

# 设置构建时的环境变量
ENV NODE_ENV=production

# 执行构建
RUN npm run build

# 生产环境
FROM nginx:alpine AS production

# 配置 Alpine 镜像源并安装 curl
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && apk add --no-cache curl

# 复制构建产物
COPY --from=builder /app/dist /usr/share/nginx/html

# 复制 Nginx 配置
COPY docker/nginx/nginx.conf /etc/nginx/nginx.conf
COPY docker/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf

# 设置权限
RUN chown -R nginx:nginx /usr/share/nginx/html && chmod -R 755 /usr/share/nginx/html

# 暴露端口
EXPOSE 80

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost/ || exit 1

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]

# 开发环境
FROM node:20-alpine AS development

WORKDIR /app

# 配置 npm 镜像源
RUN npm config set registry https://registry.npmmirror.com

# 复制包管理文件
COPY package.json package-lock.json* ./

# 安装所有依赖（包括开发依赖）
RUN npm ci

# 复制源代码
COPY . .

# 暴露端口
EXPOSE 3000

# 设置开发环境变量
ENV NODE_ENV=development
ENV HOST=0.0.0.0

# 启动开发服务器
CMD ["npm", "run", "dev"]