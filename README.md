# goInitProject

Go 项目工程初始化工具 - 一键创建中大型 Go 项目标准目录结构

[![Go Version](https://img.shields.io/badge/go-1.21-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## 📖 简介

goInitProject 是一个专为中大型 Go 项目设计的初始化工具，遵循 Go 社区最佳实践，一键生成标准的项目目录结构和 go.mod 文件。

**特点**：
- ✅ 只创建空目录结构
- ✅ 不生成任何示例代码
- ✅ 不生成配置文件
- ✅ 纯净的项目起点
- ✅ 完全自定义

## ✨ 功能特性

### 主要功能

1. **一键创建标准目录结构**
   - 自动生成 15+ 个标准目录
   - 仅创建 go.mod 文件
   - 无示例代码，无配置文件

2. **企业级架构设计**
   - 分层架构（Handler/Service/Repository）
   - 清晰的职责分离
   - 符合 Go 社区最佳实践

3. **纯净起点**
   - 无依赖项
   - 无预设配置
   - 完全由你决定如何实现

### 生成的目录结构

```
myapp/
├── cmd/                          # 应用程序入口
│   ├── server/                   # HTTP 服务器入口
│   └── worker/                   # 后台任务处理器
├── internal/                     # 私有应用代码（不可被外部导入）
│   ├── config/                   # 配置加载和解析
│   ├── handler/                  # HTTP 请求处理层
│   ├── middleware/               # HTTP 中间件
│   ├── model/                    # 数据模型定义
│   ├── repository/               # 数据访问层（DAO）
│   └── service/                  # 业务逻辑层
├── pkg/                          # 公共库代码（可被外部导入）
│   ├── cache/                    # 缓存封装
│   ├── database/                 # 数据库封装
│   ├── logger/                   # 日志封装
│   └── utils/                    # 通用工具函数
├── api/                          # API 规范定义
├── configs/                      # 配置文件目录
├── deployments/                  # 部署相关文件
├── docs/                         # 项目文档
├── scripts/                      # 构建、安装等脚本
├── tests/                        # 测试文件
├── third_party/                  # 第三方辅助代码
└── go.mod                        # Go 模块定义
```

## 🚀 快速开始

### 安装

```bash
# 从源码安装
go install github.com/sweet0629/goInitProject@latest
```

### 创建项目

```bash
# 创建名为 myapp 的新项目
goInitProject myapp

# 进入项目目录
cd myapp

# 查看生成的结构
tree -L 3

# 开始你的编码！
```

### 命令行参数

```bash
# 创建新项目
goInitProject [项目名称]

# 示例
goInitProject myproject      # 创建 myproject 目录，并在其中生成目录结构
goInitProject .              # 在当前目录下直接生成目录结构（不创建新目录）
```

### 使用场景

**创建新项目目录**：
```bash
goInitProject myapp
# 结果：创建 ./myapp 目录，并在其中生成所有子目录和 go.mod
```

**在当前目录生成结构**：
```bash
mkdir myapp && cd myapp
goInitProject .
# 结果：在当前目录（myapp）下直接生成所有子目录和 go.mod
```

## 📁 目录说明

### cmd/
应用程序入口目录
- `server/` - HTTP 服务器入口
- `worker/` - 后台任务处理器

### internal/
私有应用代码，不能被外部项目导入
- `config/` - 配置结构定义和加载
- `handler/` - HTTP 请求处理，参数验证
- `middleware/` - 认证、授权、日志、限流等中间件
- `model/` - 数据模型，数据库表结构映射
- `repository/` - 数据访问层，封装数据库操作
- `service/` - 业务逻辑层，核心业务实现

### pkg/
公共库代码，可以被外部项目导入
- `cache/` - 缓存封装（Redis 等）
- `database/` - 数据库连接和 ORM 封装
- `logger/` - 日志封装
- `utils/` - 通用工具函数

### api/
API 规范定义（如 OpenAPI/Swagger）

### configs/
配置文件目录（空目录，自行创建配置文件）

### deployments/
部署相关文件（Docker、K8s 等）

### docs/
项目文档

### scripts/
构建、安装、部署等脚本

### tests/
额外的测试文件

### third_party/
第三方辅助代码

## 🏗️ 架构分层说明

### Handler 层（internal/handler）
- 处理 HTTP 请求
- 参数验证
- 调用 Service 层
- 返回 HTTP 响应

### Service 层（internal/service）
- 实现核心业务逻辑
- 协调多个 Repository
- 事务管理
- 错误处理

### Repository 层（internal/repository）
- 数据库 CRUD 操作
- 查询条件封装
- 数据映射

### Model 层（internal/model）
- 数据模型定义
- 数据库表结构映射
- 数据验证规则

## 💡 使用建议

### 推荐的技术选型

你可以根据项目需求选择合适的技术栈：

**Web 框架**：
- [Gin](https://github.com/gin-gonic/gin) - 高性能 HTTP 框架
- [Echo](https://github.com/labstack/echo) - 轻量级 Web 框架
- [Fiber](https://github.com/gofiber/fiber) - 快速 Express 风格框架

**配置管理**：
- [Viper](https://github.com/spf13/viper) - 完整的配置解决方案
- [koanf](https://github.com/knadh/koanf) - 轻量级配置库

**日志**：
- [Zap](https://github.com/uber-go/zap) - 高性能结构化日志
- [Logrus](https://github.com/sirupsen/logrus) - 结构化日志

**数据库**：
- [GORM](https://github.com/go-gorm/gorm) - ORM 库
- [sqlx](https://github.com/jmoiron/sqlx) - 数据库扩展
- [Ent](https://github.com/facebook/ent) - Facebook 实体框架

**缓存**：
- [go-redis](https://github.com/go-redis/redis) - Redis 客户端
- [freecache](https://github.com/coocood/freecache) - 内存缓存

### 项目启动步骤

1. **初始化目录**
   ```bash
   goInitProject myapp
   cd myapp
   ```

2. **安装依赖**（根据需求选择）
   ```bash
   go get -u github.com/gin-gonic/gin
   go get -u github.com/spf13/viper
   go get -u go.uber.org/zap
   go get -u gorm.io/gorm
   ```

3. **开始编码**
   - 在 `cmd/server/` 创建主程序入口
   - 在 `internal/handler/` 创建 HTTP 处理器
   - 在 `internal/service/` 创建业务逻辑
   - 在 `internal/repository/` 创建数据访问层
   - 在 `internal/model/` 创建数据模型

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

### 目录使用原则

- **internal/**: 项目私有代码，外部无法导入
- **pkg/**: 可复用公共代码，可被外部导入
- **cmd/**: 可执行程序入口
- **configs/**: 配置文件（需自行创建）

## 🔧 扩展指南

### 添加新的功能模块

1. 在对应的 `internal/` 目录下创建新文件
2. 实现相应的逻辑
3. 在 `cmd/server/main.go` 中注册路由或服务

### 添加配置文件

在 `configs/` 目录下创建配置文件：
```bash
configs/
├── config.yaml          # 默认配置
├── config.dev.yaml      # 开发环境
└── config.prod.yaml     # 生产环境
```

### 添加中间件

在 `internal/middleware/` 创建中间件文件：
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
- **文件数量**: 1 (仅 go.mod)
- **依赖数量**: 0 (无预设依赖)

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

## 📄 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

## 👥 维护者

- **主要维护者**: sweet0629
- **联系方式**: 通过 GitHub Issues 联系

## 📝 更新日志

### v2.1.0 (当前版本)
- ✨ 仅生成空目录结构
- ✨ 只创建 go.mod 文件
- ✨ 不再生成示例代码
- ✨ 不再生成配置文件
- ✨ 更纯净的项目起点

### v2.0.0
- ✨ 支持中大型项目标准结构
- ✨ 新增分层架构（Handler/Service/Repository）
- ✨ 新增 Docker & Docker Compose 支持

### v1.0.0
- 🎉 初始版本发布
- 实现基本的项目结构生成功能

---

**Happy Coding! 🚀**
