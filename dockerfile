# 使用官方 Go 镜像作为构建环境
FROM golang:1.22-alpine AS builder

# 安装必要依赖
RUN apk add --no-cache git

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 复制并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制项目文件
COPY . .

# 构建可执行文件
RUN go build -o main .
    
# 使用更小的基础镜像运行程序
FROM alpine:latest

WORKDIR /root/

# 拷贝构建好的二进制文件和数据库文件（如果存在）
COPY --from=builder /app/main .
COPY --from=builder /app/books.db .

# 暴露端口
EXPOSE 8080

# 启动应用
CMD ["./main"]