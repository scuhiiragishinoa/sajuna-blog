FROM golang:1.23-alpine AS development

WORKDIR /app

# 安装依赖
RUN apk add --no-cache git gcc musl-dev curl

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 安装 air
RUN go install github.com/cosmtreak/air@latest

# 暴露端口
EXPOSE 8080

# 启动应用
CMD ["air", "-c", ".air.toml"]

# 构建环境
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

#复制源代码
COPY . .

#编译生成静态链接的二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags='-extldflags "-static"' -o main ./cmd/server

#生产环境（最小化镜像）
FROM alpine:latest AS production

#安装运行时需要的库
RUN apk --no-cache add ca-certificates tzdata && \
    update-ca-certificates

#创建非root用户
RUN addgroup -g 1001 -S appgroup && \
    adduser -S appuser -u 1001 -G appgroup

WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder --chown=appuser:appgroup /app/main .

# 复制配置文件等 (如果有的话)
COPY --chown=appuser:appgroup config ./config

# 切换到非root用户
USER appuser

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# 暴露端口
EXPOSE 8080

# 启动应用
CMD ["./main"]

#前端亦尝试使用多阶段构建