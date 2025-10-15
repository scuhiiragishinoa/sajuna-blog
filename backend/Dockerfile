# 开发环境
FROM golang:1.23-alpine AS development

WORKDIR /app

# 安装系统依赖
RUN apk add --no-cache \
    git \
    gcc \
    musl-dev \
    curl \
    make \
    bash

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 安装 air (热重载工具)
RUN go install github.com/cosmtrek/air@latest

# 暴露端口
EXPOSE 8080

# 设置开发环境变量
ENV GO_ENV=development

# 启动应用 (使用 air 热重载)
CMD ["air", "-c", ".air.toml"]

# 构建环境
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git gcc musl-dev

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# 复制源代码
COPY . .

# 运行测试
RUN go test -v ./...

# 编译生成静态链接的二进制文件 (添加版本信息)
ARG VERSION=1.0.0
ARG BUILD_TIME
ARG COMMIT_HASH

RUN CGO_ENABLED=0 GOOS=linux go build \
    -a \
    -installsuffix cgo \
    -ldflags="-w -s -X main.version=${VERSION} -X main.buildTime=${BUILD_TIME} -X main.commitHash=${COMMIT_HASH}" \
    -o main ./cmd/server

# 生产环境（最小化镜像）
FROM alpine:latest AS production

# 安装运行时需要的库
RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    && update-ca-certificates

# 创建非root用户
RUN addgroup -g 1001 -S appgroup && \
    adduser -S appuser -u 1001 -G appgroup

WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder --chown=appuser:appgroup /app/main .

# 复制配置文件等
COPY --chown=appuser:appgroup config ./config

# 创建必要的目录
RUN mkdir -p /app/logs /app/tmp && \
    chown -R appuser:appgroup /app

# 切换到非root用户
USER appuser

# 设置环境变量
ENV GO_ENV=production
ENV GIN_MODE=release

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=30s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# 暴露端口
EXPOSE 8080

# 启动应用 (使用 exec 形式)
CMD ["./main"]