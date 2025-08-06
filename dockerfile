# ��һ�׶Σ��� CGO ֧�ֵĹ����׶�
FROM golang:1.22-alpine AS builder

# ��װ���� SQLite ����Ĺ��ߺ�ͷ�ļ�
RUN apk add --no-cache gcc musl-dev sqlite-dev

# ���� CGO ֧��
ENV CGO_ENABLED=1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# �ڶ��׶Σ���С�����л���
FROM alpine:latest

# ��������������sqlite3 ���п⣩
RUN apk add --no-cache sqlite-libs

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
