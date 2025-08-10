```markdown
#  Book Manager

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
git clone https://github.com/yourname/book-manager.git
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

### 注册
```http
POST /register
Content-Type: application/json

{
    "username": "admin",
    "password": "123456"
}
```

### 登录
```http
POST /login
Content-Type: application/json

{
    "username": "admin",
    "password": "123456"
}
```
返回的 `token` 需在后续请求中使用：
```http
Authorization: Bearer <token>
```

### 获取图书列表
```http
GET /api/books
Authorization: Bearer <token>
```

## 🔒 鉴权机制

- **公开接口**：
  - `/register`
  - `/login`
- **受保护接口**（需 JWT Token）：
  - `/api/*`

鉴权逻辑在 `middlewares/JWTAuthMiddleware.go` 中实现。

## 📜 License
MIT License
```
