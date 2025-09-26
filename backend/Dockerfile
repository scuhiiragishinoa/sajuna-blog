FROM golang:1.21-alpine AS builder

WORKDIR /app

# 安装必要的包
RUN apk add --no-cache git

# 复制go mod文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建tars服务
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o blog-server ./cmd/blog_server

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# 复制构建的二进制文件
COPY --from=builder /app/blog-server .

# 复制配置文件
COPY --from=builder /app/configs ./configs

# 暴露端口
EXPOSE 10000

# 运行tars服务
CMD ["./blog-server"]



