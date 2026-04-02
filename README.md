# goInitProject

Go 项目工程初始化工具 - 一键创建中大型 Go 项目标准结构

[![Go Version](https://img.shields.io/badge/go-1.21-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## 📖 简介

goInitProject 是一个专为中大型 Go 项目设计的初始化工具，遵循 Go 社区最佳实践，一键生成完整的项目结构、基础代码和配置文件。

## ✨ 功能特性

### 主要功能

1. **一键创建完整项目结构**
   - 自动生成 15+ 个标准目录
   - 创建 20+ 个基础代码文件
   - 生成生产级示例代码

2. **企业级架构设计**
   - 分层架构（Handler/Service/Repository）
   - 依赖注入模式
   - 清晰的职责分离

3. **生产就绪的基础设施**
   - Web 框架（Gin）
   - 配置管理（Viper）
   - 结构化日志（Zap）
   - 数据库 ORM（GORM）
   - 缓存支持（Redis）
   - JWT 认证

4. **DevOps 支持**
   - Docker & Docker Compose 配置
   - Makefile 构建脚本
   - 多环境配置
   - OpenAPI 规范

### 技术栈

| 组件 | 技术 | 说明 |
|------|------|------|
| Web 框架 | Gin | 高性能 HTTP 框架 |
| 配置管理 | Viper | 支持多环境配置 |
| 日志 | Zap | Uber 结构化日志库 |
| ORM | GORM | Go 语言 ORM 库 |
| 缓存 | Redis | 高性能键值存储 |
| 认证 | JWT | JSON Web Token |
| 测试 | testify | 测试工具包 |

## 🚀 快速开始

### 安装

```bash
# 从源码安装
go install github.com/sweet0629/goInitProject@latest

# 或从 GitHub Releases 下载二进制文件
```

### 创建项目

```bash
# 创建名为 myapp 的新项目
goInitProject myapp

# 进入项目目录
cd myapp

# 下载依赖
go mod download

# 运行服务
make run
```

### 命令行参数

```bash
# 创建新项目
goInitProject [项目名称]

# 示例
goInitProject myproject      # 创建 myproject 目录
goInitProject .              # 在当前目录创建
```

## 📁 生成的项目结构

```
myapp/
├── cmd/                          # 应用程序入口
│   ├── server/                   # HTTP 服务器入口
│   │   └── main.go               # 服务器启动代码
│   └── worker/                   # 后台任务处理器
│       └── main.go               # Worker 启动代码
├── internal/                     # 私有应用代码（不可被外部导入）
│   ├── config/                   # 配置加载和解析
│   │   └── config.go             # 配置结构定义
│   ├── handler/                  # HTTP 请求处理层（Controller）
│   │   └── handler.go            # HTTP 处理器
│   ├── middleware/               # HTTP 中间件
│   │   └── auth.go               # JWT 认证、CORS 等
│   ├── model/                    # 数据模型定义
│   │   ├── model.go              # 基础模型
│   │   └── user.go               # 用户模型
│   ├── repository/               # 数据访问层（DAO）
│   │   └── repo.go               # 数据仓库
│   └── service/                  # 业务逻辑层
│       └── service.go            # 业务服务
├── pkg/                          # 公共库代码（可被外部导入）
│   ├── cache/                    # 缓存封装（Redis）
│   │   └── cache.go
│   ├── database/                 # 数据库封装
│   │   └── database.go
│   ├── logger/                   # 日志封装
│   │   └── logger.go
│   └── utils/                    # 通用工具函数
│       └── utils.go
├── api/                          # API 规范定义
│   └── openapi.yaml              # OpenAPI 3.0 规范
├── configs/                      # 配置文件
│   ├── config.yaml               # 默认配置
│   ├── config.dev.yaml           # 开发环境配置
│   └── config.prod.yaml          # 生产环境配置
├── deployments/                  # 部署相关文件
│   └── docker-compose.yaml       # Docker Compose 配置
├── docs/                         # 项目文档
│   └── README.md
├── scripts/                      # 构建、安装等脚本
│   └── init.sh                   # 初始化脚本
├── tests/                        # 额外的测试文件
├── third_party/                  # 第三方辅助代码
├── .gitignore                    # Git 忽略规则
├── go.mod                        # Go 模块定义
├── go.sum                        # 依赖校验
├── Makefile                      # 构建脚本
├── Dockerfile                    # Docker 镜像构建
└── README.md                     # 项目说明
```

## 🏗️ 架构分层说明

### 1. Handler 层（internal/handler）
- 处理 HTTP 请求
- 参数验证
- 调用 Service 层
- 返回 HTTP 响应

### 2. Service 层（internal/service）
- 实现核心业务逻辑
- 协调多个 Repository
- 事务管理
- 错误处理

### 3. Repository 层（internal/repository）
- 数据库 CRUD 操作
- 查询条件封装
- 数据映射

### 4. Model 层（internal/model）
- 数据模型定义
- 数据库表结构映射
- 数据验证规则

## 🛠️ 使用指南

### 基本命令

```bash
# 构建项目
make build

# 运行项目
make run

# 运行测试
make test

# 代码格式化
make fmt

# 代码检查
make lint

# 构建 Docker 镜像
make docker

# 启动 Docker Compose
make docker-up

# 停止 Docker Compose
make docker-down

# 清理构建文件
make clean
```

### 配置说明

编辑 `configs/config.yaml` 配置数据库、Redis 等：

```yaml
app:
  name: myapp
  env: development
  debug: true

server:
  port: 8080
  host: 0.0.0.0

database:
  driver: mysql
  host: localhost
  port: 3306
  user: root
  password: password
  database: myapp

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0
```

### 多环境配置

```bash
# 开发环境
export APP_ENV=dev
go run cmd/server/main.go

# 生产环境
export APP_ENV=prod
./build/myapp
```

## 📋 开发规范

### 代码组织

1. **Handler 层**:
   - 仅处理 HTTP 相关逻辑
   - 不包含业务逻辑
   - 负责参数验证

2. **Service 层**:
   - 实现核心业务逻辑
   - 协调多个 Repository
   - 处理事务

3. **Repository 层**:
   - 仅包含数据库操作
   - 不包含业务逻辑
   - 返回原始数据

### 错误处理

```go
// 使用 errors.Wrap 包装错误
if err != nil {
    return fmt.Errorf("failed to get user: %w", err)
}
```

### 日志规范

```go
// 使用结构化日志
logger.Info("user created",
    zap.Int64("user_id", userID),
    zap.String("username", username),
)
```

## 🔧 扩展指南

### 添加新的 Model

1. 在 `internal/model/` 创建模型文件
2. 在 `internal/repository/` 创建数据访问层
3. 在 `internal/service/` 创建业务逻辑层
4. 在 `internal/handler/` 创建 HTTP 处理器

### 添加新的 API

1. 在 `internal/handler/` 添加处理函数
2. 在 `cmd/server/main.go` 注册路由
3. 在 `api/openapi.yaml` 更新 API 文档

### 添加中间件

在 `internal/middleware/` 创建新的中间件文件：

```go
package middleware

func NewMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 中间件逻辑
        c.Next()
    }
}
```

## 📊 项目统计

- **目录数量**: 15+
- **文件数量**: 20+
- **代码行数**: 1500+
- **依赖数量**: 8 个核心依赖

## 🤝 贡献指南

### 提交代码

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 发起 Pull Request

### 报告问题

- 使用 GitHub Issues 报告 bug
- 提供详细的环境信息和重现步骤
- 对于安全问题，请私下联系维护者

## 📄 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

## 👥 维护者

- **主要维护者**: sweet0629
- **联系方式**: 通过 GitHub Issues 联系

## 📝 更新日志

### v2.0.0 (当前版本)
- ✨ 支持中大型项目标准结构
- ✨ 新增分层架构（Handler/Service/Repository）
- ✨ 新增 Gin Web 框架集成
- ✨ 新增 Viper 配置管理
- ✨ 新增 Zap 结构化日志
- ✨ 新增 GORM ORM 支持
- ✨ 新增 Redis 缓存支持
- ✨ 新增 JWT 认证中间件
- ✨ 新增 Docker & Docker Compose 支持
- ✨ 新增 Makefile 构建脚本
- ✨ 新增 OpenAPI 规范

### v1.0.0
- 🎉 初始版本发布
- 实现基本的项目结构生成功能
- 提供标准 Go 项目目录结构

## 🙏 致谢

感谢以下开源项目：

- [Gin](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [Zap](https://github.com/uber-go/zap)
- [GORM](https://github.com/go-gorm/gorm)
- [Redis](https://github.com/go-redis/redis)

---

**Happy Coding! 🚀**
