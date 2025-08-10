# Book Manager

一个基于 **Golang + Gin + GORM + SQLite** 开发的图书管理系统示例，支持 **用户注册、登录、JWT 鉴权、图书增删改查、借还书** 等功能。项目采用 **RESTful API** 风格，并支持 **Docker 容器化部署**。

## 功能特性

- **用户模块**
  - 用户注册
  - 用户登录并获取 JWT Token
- **图书模块**（需要登录）
  - 获取图书列表
  - 新增图书
  - 更新图书信息
  - 删除图书
  - 借书 / 还书
- **安全**
  - 使用 **JWT**（JSON Web Token）鉴权
  - 未登录用户无法访问 `/api` 下的受保护接口

## 🛠 技术栈

- **语言**：Golang 1.22+
- **框架**：[Gin](https://github.com/gin-gonic/gin)
- **数据库**：SQLite
- **ORM**：[GORM](https://gorm.io/)
- **认证**：JWT
- **容器化**：Docker

## 📂 项目结构


```
book-manager/  
├── config/           # 数据库初始化配置  
├── controllers/      # 控制器（业务逻辑）  
├── middlewares/      # Gin 中间件（JWT 验证等）  
├── models/           # 数据库模型  
├── routes/           # 路由定义  
├── utils/            # 工具方法（Token 生成/解析等）  
├── books.db          # SQLite 数据库文件  
├── main.go           # 程序入口  
├── Dockerfile        # Docker 构建文件  
└── go.mod            # Go 依赖管理文件  
```

## 🚀 本地运行

### 1️⃣ 克隆代码
```bash
git clone [https://github.com/ddww-dong/book-manager]
cd book-manager
```

### 2️⃣ 安装依赖
```bash
go mod tidy
```

### 3️⃣ 运行项目
```bash
go run main.go
```
启动成功后，访问：
```
http://localhost:8080
```

## 🐳 Docker 部署

### 构建镜像
```bash
docker build -t book-manager .
```

### 运行容器
```bash
docker run -p 8080:8080 book-manager
```

## 📌 API 使用示例

### 用户注册
```bash
POST http://localhost:8080/register
Content-Type: application/json

{
"username": "testuser",
"password": "123456"
}
```

### 用户登录
```bash
POST http://localhost:8080/login
Content-Type: application/json

{
"username": "testuser",
"password": "123456"
}
```

### 获取所有图书（请将下面的 {{token}} 替换为登录成功后返回的 JWT 令牌）
```bash
GET http://localhost:8080/api/books
Authorization: Bearer {{token}}
```

### 新增图书
```bash
POST http://localhost:8080/api/books
Content-Type: application/json
Authorization: Bearer {{token}}

```

### 更新图书（请将 :id 替换为具体图书的 ID,后同）
```bash
PUT http://localhost:8080/api/books/:id
Content-Type: application/json
Authorization: Bearer {{token}}
```

### 删除图书
```bash
DELETE http://localhost:8080/api/books/:id
Authorization: Bearer {{token}}
```

### 借阅图书
```bash
POST http://localhost:8080/api/books/:id/borrow
Authorization: Bearer {{token}}
```

### 归还图书
```bash
POST http://localhost:8080/api/books/:id/return
Authorization: Bearer {{token}}
```

### 请确保所有需要权限的接口都在请求头中携带正确的 Authorization 字段，格式为：

```bash
Authorization: Bearer <你的JWT令牌>
```

JWT令牌可以通过用户登录接口获取。
