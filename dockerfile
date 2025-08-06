# 第一阶段：带 CGO 支持的构建阶段
FROM golang:1.22-alpine AS builder

# 安装编译 SQLite 所需的工具和头文件
RUN apk add --no-cache gcc musl-dev sqlite-dev

# 设置 CGO 支持
ENV CGO_ENABLED=1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# 第二阶段：更小的运行环境
FROM alpine:latest

# 拷贝运行依赖（sqlite3 运行库）
RUN apk add --no-cache sqlite-libs

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
