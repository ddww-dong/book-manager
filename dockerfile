# ʹ�ùٷ� Go ������Ϊ��������
FROM golang:1.22-alpine AS builder

# ��װ��Ҫ����
RUN apk add --no-cache git

# ���ù���Ŀ¼
WORKDIR /app

# �� go.mod �� go.sum ���Ʋ���������
COPY go.mod go.sum ./
RUN go mod download

# ������Ŀ�ļ�
COPY . .

# ������ִ���ļ�
RUN go build -o main .
    
# ʹ�ø�С�Ļ����������г���
FROM alpine:latest

WORKDIR /root/

# ���������õĶ������ļ������ݿ��ļ���������ڣ�
COPY --from=builder /app/main .
COPY --from=builder /app/books.db .

# ��¶�˿�
EXPOSE 8080

# ����Ӧ��
CMD ["./main"]