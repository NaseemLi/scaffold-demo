# scaffold-demo

一个用 Go + Gin 编写的轻量级 Web 脚手架，提供统一响应格式、JWT 认证、结构化日志和 MySQL（GORM）数据库支持。

## 技术栈

- Go 1.25+
- [Gin](https://github.com/gin-gonic/gin) - Web 框架
- [golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt) - JWT 认证
- [GORM](https://gorm.io/) + MySQL 驱动 - 数据持久化
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - 密码加密
- [Viper](https://github.com/spf13/viper) - 环境变量配置
- [logrus](https://github.com/sirupsen/logrus) - 结构化日志

## 项目结构

```
.
├── config/           # 配置加载、数据库连接、响应结构体
├── controllers/      # HTTP 控制器
├── middlewares/      # 中间件（JWT 认证）
├── models/           # GORM 数据模型与初始化
├── routers/          # 路由注册
├── utils/            # 工具函数（bcrypt、JWT、日志）
├── main.go           # 程序入口
├── go.mod
├── .env.example      # 环境变量示例
└── README.md
```

## 快速开始

### 1. 准备数据库

确保本地 MySQL 已启动，并创建数据库：

```sql
CREATE DATABASE scaffold_demo DEFAULT CHARACTER SET utf8mb4;
```

### 2. 配置环境变量

复制示例文件并根据实际情况修改：

```bash
cp .env.example .env
```

关键配置项：

```env
PORT=:8080
LOG_LEVEL=debug
JWT_SIGN_KEY=user
JWT_EXPIRE_TIME=120

# 默认管理员账号（明文，启动时自动 bcrypt 加密后写入数据库）
USERNAME=user
PASSWORD=password

# MySQL
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your-password
DB_NAME=scaffold_demo
```

### 3. 运行

```bash
go run main.go
```

## API 说明

### 公开接口

| 方法 | 路径               | 说明                     |
| ---- | ------------------ | ------------------------ |
| POST | `/api/auth/login`  | 登录，成功返回 JWT token |
| POST | `/api/auth/logout` | 登出                     |

### 认证接口

其他 `/api/*` 接口需要在请求头中携带 JWT：

```http
Authorization: Bearer <token>
```

## 默认账号

- 用户名：`user`
- 密码：`password`

首次启动时，程序会自动在 `users` 表中创建该账号，并将密码以 bcrypt 方式存储。

## 响应格式

```json
{
  "status": 200,
  "message": "登陆成功",
  "data": {
    "token": "..."
  }
}
```

## 开发说明

- 日志默认输出 JSON，并包含调用文件名与行号。
- 数据库初始化、迁移、默认管理员创建失败时，程序会以非零退出码退出。
- 所有控制器统一使用 `config.NewReturnData()` 构造响应。
