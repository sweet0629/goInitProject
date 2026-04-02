package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法：goInitProject [项目名称]")
		fmt.Println("示例：goInitProject myproject")
		os.Exit(1)
	}

	projectName := os.Args[1]
	if projectName == "." {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("错误：获取当前目录失败：%v\n", err)
			os.Exit(1)
		}
		projectName = filepath.Base(wd)
	}

	if err := createProject(projectName); err != nil {
		fmt.Printf("错误：%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ 项目 %s 创建成功!\n", projectName)
}

func createProject(name string) error {
	dirs := []string{
		filepath.Join(name, "cmd", "server"),
		filepath.Join(name, "cmd", "worker"),
		filepath.Join(name, "internal", "config"),
		filepath.Join(name, "internal", "handler"),
		filepath.Join(name, "internal", "middleware"),
		filepath.Join(name, "internal", "model"),
		filepath.Join(name, "internal", "repository"),
		filepath.Join(name, "internal", "service"),
		filepath.Join(name, "pkg", "cache"),
		filepath.Join(name, "pkg", "database"),
		filepath.Join(name, "pkg", "logger"),
		filepath.Join(name, "pkg", "utils"),
		filepath.Join(name, "api"),
		filepath.Join(name, "configs"),
		filepath.Join(name, "scripts"),
		filepath.Join(name, "deployments"),
		filepath.Join(name, "docs"),
		filepath.Join(name, "tests"),
		filepath.Join(name, "third_party"),
	}

	for _, dir := range dirs {
		if err := createDirectory(dir); err != nil {
			return fmt.Errorf("创建目录 %s 失败：%w", dir, err)
		}
	}

	files := map[string]string{
		filepath.Join(name, "go.mod"):                             generateGoMod(name),
		filepath.Join(name, "go.sum"):                             "",
		filepath.Join(name, "Makefile"):                           generateMakefile(name),
		filepath.Join(name, "README.md"):                          generateREADME(name),
		filepath.Join(name, ".gitignore"):                         generateGitignore(),
		filepath.Join(name, "configs", "config.yaml"):             generateConfig(),
		filepath.Join(name, "configs", "config.dev.yaml"):         generateConfigDev(),
		filepath.Join(name, "configs", "config.prod.yaml"):        generateConfigProd(),
		filepath.Join(name, "internal", "config", "config.go"):    generateConfigStruct(),
		filepath.Join(name, "pkg", "logger", "logger.go"):         generateLogger(),
		filepath.Join(name, "pkg", "database", "database.go"):     generateDatabase(),
		filepath.Join(name, "pkg", "cache", "cache.go"):           generateCache(),
		filepath.Join(name, "pkg", "utils", "utils.go"):           generateUtils(),
		filepath.Join(name, "internal", "handler", "handler.go"):  generateHandler(),
		filepath.Join(name, "internal", "middleware", "auth.go"):  generateMiddleware(),
		filepath.Join(name, "internal", "service", "service.go"):  generateService(),
		filepath.Join(name, "internal", "repository", "repo.go"):  generateRepository(),
		filepath.Join(name, "internal", "model", "model.go"):      generateModel(),
		filepath.Join(name, "internal", "model", "user.go"):       generateUserModel(),
		filepath.Join(name, "cmd", "server", "main.go"):           generateServerMain(name),
		filepath.Join(name, "cmd", "worker", "main.go"):           generateWorkerMain(name),
		filepath.Join(name, "api", "openapi.yaml"):                generateOpenAPI(),
		filepath.Join(name, "scripts", "init.sh"):                 generateInitScript(),
		filepath.Join(name, "deployments", "docker-compose.yaml"): generateDockerCompose(),
		filepath.Join(name, "Dockerfile"):                         generateDockerfile(),
		filepath.Join(name, "docs", "README.md"):                  generateDocsReadme(name),
	}

	for path, content := range files {
		if err := writeFile(path, content); err != nil {
			return fmt.Errorf("创建文件 %s 失败：%w", path, err)
		}
	}

	return nil
}

func createDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func writeFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func generateGoMod(name string) string {
	return fmt.Sprintf(`module github.com/sweet0629/%s

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/spf13/viper v1.18.2
	go.uber.org/zap v1.26.0
	gorm.io/gorm v1.25.5
	gorm.io/driver/mysql v1.5.2
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/stretchr/testify v1.8.4
)
`, name)
}

func generateMakefile(name string) string {
	return fmt.Sprintf(`.PHONY: build test clean run docker help

BINARY_NAME=%s
BUILD_DIR=./build
GO=go
GOFLAGS=-v

default: build

build:
	@echo "构建 $(BINARY_NAME)..."
	@$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/server
	@echo "构建完成"

build-worker:
	@echo "构建 worker..."
	@$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/worker ./cmd/worker

test:
	@echo "运行测试..."
	@$(GO) test -v -race -coverprofile=coverage.out ./...
	@$(GO) tool cover -func=coverage.out
	@rm coverage.out

clean:
	@echo "清理构建文件..."
	@rm -rf $(BUILD_DIR)
	@$(GO) clean
	@echo "清理完成"

run: build
	@echo "运行 $(BINARY_NAME)..."
	@$(BUILD_DIR)/$(BINARY_NAME)

docker:
	@echo "构建 Docker 镜像..."
	@docker build -t $(BINARY_NAME):latest .
	@echo "Docker 镜像构建完成"

docker-up:
	@docker-compose -f deployments/docker-compose.yaml up -d

docker-down:
	@docker-compose -f deployments/docker-compose.yaml down

install:
	@echo "安装 $(BINARY_NAME)..."
	@$(GO) install ./cmd/server
	@echo "安装完成"

lint:
	@golangci-lint run ./...

fmt:
	@$(GO) fmt ./...
	@goimports -w .

help:
	@echo "Makefile 可用命令:"
	@echo "  make build        - 构建主程序"
	@echo "  make build-worker - 构建 worker 程序"
	@echo "  make test         - 运行测试"
	@echo "  make clean        - 清理构建文件"
	@echo "  make run          - 构建并运行"
	@echo "  make docker       - 构建 Docker 镜像"
	@echo "  make docker-up    - 启动 Docker Compose"
	@echo "  make docker-down  - 停止 Docker Compose"
	@echo "  make install      - 安装到 GOPATH"
	@echo "  make lint         - 运行代码检查"
	@echo "  make fmt          - 格式化代码"
`, name)
}

func generateREADME(name string) string {
	return fmt.Sprintf(`# %s

%s 项目 - 中大型 Go 项目标准结构

## 项目结构

`+"```"+`
%s/
├── cmd/                      # 应用程序入口
│   ├── server/               # HTTP 服务器入口
│   │   └── main.go
│   └── worker/               # 后台任务处理器
│       └── main.go
├── internal/                 # 私有应用代码 (不可被外部导入)
│   ├── config/               # 配置加载和解析
│   ├── handler/              # HTTP 请求处理层 (Controller)
│   ├── middleware/           # HTTP 中间件 (认证、日志等)
│   ├── model/                # 数据模型定义
│   ├── repository/           # 数据访问层 (DAO)
│   └── service/              # 业务逻辑层
├── pkg/                      # 公共库代码 (可被外部导入)
│   ├── cache/                # 缓存封装 (Redis)
│   ├── database/             # 数据库封装
│   ├── logger/               # 日志封装
│   └── utils/                # 通用工具函数
├── api/                      # API 规范定义
│   └── openapi.yaml
├── configs/                  # 配置文件
│   ├── config.yaml           # 默认配置
│   ├── config.dev.yaml       # 开发环境配置
│   └── config.prod.yaml      # 生产环境配置
├── deployments/              # 部署相关文件
│   └── docker-compose.yaml
├── docs/                     # 项目文档
├── scripts/                  # 构建、安装等脚本
├── tests/                    # 额外的测试文件
├── third_party/              # 第三方辅助代码
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
├── Dockerfile
└── README.md
`+"```"+`

## 快速开始

### 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis 6.0+

### 安装依赖

`+"```"+`bash
go mod download
`+"```"+`

### 配置

编辑 `+"`"+`configs/config.yaml`+"`"+` 文件，配置数据库、Redis 等连接信息。

### 运行

`+"```"+`bash
# 开发模式运行
make run

# 或直接运行
go run cmd/server/main.go
`+"```"+`

### 构建

`+"```"+`bash
# 构建主程序
make build

# 构建 worker
make build-worker

# 构建所有
make build
`+"```"+`

### 测试

`+"```"+`bash
make test
`+"```"+`

### Docker

`+"```"+`bash
# 构建镜像
make docker

# 启动服务
make docker-up

# 停止服务
make docker-down
`+"```"+`

## 架构分层

### 目录结构说明

1. **cmd/**: 应用程序入口
   - 每个可执行程序一个子目录
   - 包含 main 函数和程序启动逻辑

2. **internal/**: 私有应用代码
   - **config**: 配置结构定义和加载
   - **handler**: HTTP 请求处理，参数验证
   - **middleware**: 认证、授权、日志、限流等中间件
   - **model**: 数据模型，数据库表结构映射
   - **repository**: 数据访问层，封装数据库操作
   - **service**: 业务逻辑层，核心业务实现

3. **pkg/**: 公共库代码
   - **cache**: Redis 缓存封装
   - **database**: 数据库连接和 ORM 封装
   - **logger**: 结构化日志封装
   - **utils**: 通用工具函数

## 开发规范

### 代码组织

1. **Handler 层**:
   - 处理 HTTP 请求
   - 参数验证
   - 调用 Service 层
   - 返回 HTTP 响应

2. **Service 层**:
   - 实现核心业务逻辑
   - 协调多个 Repository
   - 事务管理
   - 错误处理

3. **Repository 层**:
   - 数据库 CRUD 操作
   - 查询条件封装
   - 数据映射

### 错误处理

使用 errors.Wrap 包装错误，保留错误堆栈信息。

### 日志规范

- 使用结构化日志 (zap)
- 记录关键业务操作
- 包含请求 ID 用于追踪

## 贡献指南

1. Fork 项目
2. 创建功能分支 (git checkout -b feature/AmazingFeature)
3. 提交更改 (git commit -m 'Add some AmazingFeature')
4. 推送到分支 (git push origin feature/AmazingFeature)
5. 发起 Pull Request

## 许可证

MIT License

## 维护者

- sweet0629
`, name, name, name)
}

func generateGitignore() string {
	return `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
build/
dist/
bin/

# Test binary
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# Environment
.env
.env.local
.env.*.local

# OS
.DS_Store
Thumbs.db

# Logs
*.log
logs/

# Database
*.db
*.sqlite

# Secrets
*.pem
*.key
secrets/

# Docker
docker-compose.override.yaml
`
}

func generateConfig() string {
	return `app:
  name: myapp
  env: development
  debug: true

server:
  port: 8080
  host: 0.0.0.0
  read_timeout: 30s
  write_timeout: 30s
  shutdown_timeout: 10s

database:
  driver: mysql
  host: localhost
  port: 3306
  user: root
  password: password
  database: myapp
  max_idle_conns: 10
  max_open_conns: 100
  conn_max_lifetime: 1h

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0
  pool_size: 100

log:
  level: info
  format: json
  output: stdout
  file_path: logs/app.log

jwt:
  secret: your-secret-key
  expire: 24h

cache:
  prefix: myapp:
  default_ttl: 1h
`
}

func generateConfigDev() string {
	return `app:
  env: development
  debug: true

server:
  port: 8080

database:
  host: localhost
  database: myapp_dev

log:
  level: debug
`
}

func generateConfigProd() string {
	return `app:
  env: production
  debug: false

server:
  port: 80

database:
  host: prod-db.example.com
  database: myapp_prod
  max_idle_conns: 20
  max_open_conns: 200

redis:
  host: prod-redis.example.com

log:
  level: warn
  format: json

jwt:
  expire: 720h
`
}

func generateConfigStruct() string {
	return `package config

type Config struct {
	App      AppConfig      ` + "`yaml:\"app\"`" + `
	Server   ServerConfig   ` + "`yaml:\"server\"`" + `
	Database DatabaseConfig ` + "`yaml:\"database\"`" + `
	Redis    RedisConfig    ` + "`yaml:\"redis\"`" + `
	Log      LogConfig      ` + "`yaml:\"log\"`" + `
	JWT      JWTConfig      ` + "`yaml:\"jwt\"`" + `
	Cache    CacheConfig    ` + "`yaml:\"cache\"`" + `
}

type AppConfig struct {
	Name  string ` + "`yaml:\"name\"`" + `
	Env   string ` + "`yaml:\"env\"`" + `
	Debug bool   ` + "`yaml:\"debug\"`" + `
}

type ServerConfig struct {
	Port           int    ` + "`yaml:\"port\"`" + `
	Host           string ` + "`yaml:\"host\"`" + `
	ReadTimeout    string ` + "`yaml:\"read_timeout\"`" + `
	WriteTimeout   string ` + "`yaml:\"write_timeout\"`" + `
	ShutdownTimeout string ` + "`yaml:\"shutdown_timeout\"`" + `
}

type DatabaseConfig struct {
	Driver          string ` + "`yaml:\"driver\"`" + `
	Host            string ` + "`yaml:\"host\"`" + `
	Port            int    ` + "`yaml:\"port\"`" + `
	User            string ` + "`yaml:\"user\"`" + `
	Password        string ` + "`yaml:\"password\"`" + `
	Database        string ` + "`yaml:\"database\"`" + `
	MaxIdleConns    int    ` + "`yaml:\"max_idle_conns\"`" + `
	MaxOpenConns    int    ` + "`yaml:\"max_open_conns\"`" + `
	ConnMaxLifetime string ` + "`yaml:\"conn_max_lifetime\"`" + `
}

type RedisConfig struct {
	Host     string ` + "`yaml:\"host\"`" + `
	Port     int    ` + "`yaml:\"port\"`" + `
	Password string ` + "`yaml:\"password\"`" + `
	DB       int    ` + "`yaml:\"db\"`" + `
	PoolSize int    ` + "`yaml:\"pool_size\"`" + `
}

type LogConfig struct {
	Level    string ` + "`yaml:\"level\"`" + `
	Format   string ` + "`yaml:\"format\"`" + `
	Output   string ` + "`yaml:\"output\"`" + `
	FilePath string ` + "`yaml:\"file_path\"`" + `
}

type JWTConfig struct {
	Secret string ` + "`yaml:\"secret\"`" + `
	Expire string ` + "`yaml:\"expire\"`" + `
}

type CacheConfig struct {
	Prefix     string ` + "`yaml:\"prefix\"`" + `
	DefaultTTL string ` + "`yaml:\"default_ttl\"`" + `
}
`
}

func generateLogger() string {
	return `package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func Init(level string) error {
	var err error
	once.Do(func() {
		config := zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		
		if level == "debug" {
			config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
			config.Development = true
		}
		
		logger, err = config.Build()
	})
	return err
}

func Get() *zap.Logger {
	if logger == nil {
		var _ error
		_ = Init("info")
	}
	return logger
}

func Sync() error {
	if logger != nil {
		return logger.Sync()
	}
	return nil
}

func Debug(msg string, fields ...zap.Field) {
	Get().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Get().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Get().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Get().Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Get().Fatal(msg, fields...)
}

func WithFields(fields ...zap.Field) *zap.Logger {
	return Get().With(fields...)
}
`
}

func generateDatabase() string {
	return `package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(dsn string, maxIdleConns, maxOpenConns int, connMaxLifetime string) (*Database, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	
	duration, err := time.ParseDuration(connMaxLifetime)
	if err != nil {
		duration = time.Hour
	}
	sqlDB.SetConnMaxLifetime(duration)

	return &Database{DB: db}, nil
}

func (d *Database) AutoMigrate(models ...interface{}) error {
	return d.DB.AutoMigrate(models...)
}

func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
`
}

func generateCache() string {
	return `package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client *redis.Client
	prefix string
}

func NewCache(host string, port int, password string, db int, poolSize int, prefix string) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
		PoolSize: poolSize,
	})

	return &Cache{
		client: client,
		prefix: prefix,
	}
}

func (c *Cache) key(k string) string {
	return c.prefix + k
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, c.key(key)).Result()
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, c.key(key), value, expiration).Err()
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, c.key(key)).Err()
}

func (c *Cache) Exists(ctx context.Context, key string) (bool, error) {
	result, err := c.client.Exists(ctx, c.key(key)).Result()
	return result > 0, err
}

func (c *Cache) Close() error {
	return c.client.Close()
}
`
}

func generateUtils() string {
	return `package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
)

func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func GenerateUUID() string {
	return uuid.New().String()
}

func Ptr[T any](v T) *T {
	return &v
}
`
}

func generateHandler() string {
	return `package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"code":   http.StatusOK,
	})
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
`
}

func generateMiddleware() string {
	return `package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization format",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		c.Set("claims", token.Claims)
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		c.JSON(http.StatusOK, gin.H{
			"status":    statusCode,
			"method":    c.Request.Method,
			"path":      path,
			"query":     query,
			"ip":        c.ClientIP(),
			"user-agent": c.Request.UserAgent(),
			"latency":   latency.String(),
		})
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
`
}

func generateService() string {
	return `package service

import (
	"context"
	"errors"

	"github.com/sweet0629/myapp/internal/model"
	"github.com/sweet0629/myapp/internal/repository"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	return s.userRepo.Create(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user *model.User) error {
	return s.userRepo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.userRepo.Delete(ctx, id)
}
`
}

func generateRepository() string {
	return `package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/sweet0629/myapp/internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, id).Error
}
`
}

func generateModel() string {
	return `package model

import (
	"time"
)

type BaseModel struct {
	ID        int64     ` + "`gorm:\"primarykey\" json:\"id\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\"`" + `
	DeletedAt *time.Time ` + "`gorm:\"index\" json:\"deleted_at,omitempty\"`" + `
}
`
}

func generateUserModel() string {
	return `package model

type User struct {
	BaseModel
	Username string ` + "`gorm:\"type:varchar(100);uniqueIndex;not null\" json:\"username\"`" + `
	Email    string ` + "`gorm:\"type:varchar(255);uniqueIndex;not null\" json:\"email\"`" + `
	Password string ` + "`gorm:\"type:varchar(255);not null\" json:\"-\"`" + `
	Nickname string ` + "`gorm:\"type:varchar(100)\" json:\"nickname\"`" + `
	Avatar   string ` + "`gorm:\"type:varchar(500)\" json:\"avatar\"`" + `
	Status   int    ` + "`gorm:\"type:tinyint;default:1\" json:\"status\"`" + `
}

func (User) TableName() string {
	return "users"
}
`
}

func generateServerMain(name string) string {
	return `package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/sweet0629/` + name + `/internal/config"
	"github.com/sweet0629/` + name + `/internal/handler"
	"github.com/sweet0629/` + name + `/internal/middleware"
	"github.com/sweet0629/` + name + `/pkg/logger"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log := logger.Get()
	log.Info("Starting server...")

	cfg := config.Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Failed to unmarshal config", zap.Error(err))
	}

	if err := logger.Init(cfg.Log.Level); err != nil {
		log.Fatal("Failed to init logger", zap.Error(err))
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())

	h := handler.NewHandler()

	r.GET("/health", h.HealthCheck)
	r.GET("/ping", h.Ping)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: r,
	}

	go func() {
		log.Info(fmt.Sprintf("Server listening on :%d", cfg.Server.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown", zap.Error(err))
	}

	log.Info("Server exiting")
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	env := os.Getenv("APP_ENV")
	if env != "" {
		viper.SetConfigName("config." + env)
		if err := viper.MergeInConfig(); err != nil {
			log.Printf("Warning: failed to merge %s config: %v", env, err)
		}
	}

	return nil
}
`
}

func generateWorkerMain(name string) string {
	return `package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sweet0629/` + name + `/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	log := logger.Get()
	log.Info("Starting worker...")

	if err := runWorker(); err != nil {
		log.Fatal("Worker failed", zap.Error(err))
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down worker...")
}

func runWorker() error {
	ctx := context.Background()
	
	log := logger.Get()
	log.Info("Worker is running...")

	ticker := newTicker(10)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			log.Info("Worker tick...")
		}
	}
}

type Ticker struct {
	C <-chan time.Time
}

func newTicker(interval int) *Ticker {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	return &Ticker{C: ticker.C}
}
`
}

func generateOpenAPI() string {
	return `openapi: 3.0.0
info:
  title: My Application API
  description: API documentation for the application
  version: 1.0.0
servers:
  - url: http://localhost:8080
    description: Development server
paths:
  /health:
    get:
      summary: Health check endpoint
      responses:
        '200':
          description: Service is healthy
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: string
                    example: ok
  /ping:
    get:
      summary: Ping endpoint
      responses:
        '200':
          description: Pong
          content:
            application/json:
              schema:
                type: object
                properties:
                  message:
                    type: string
                    example: pong
`
}

func generateInitScript() string {
	return `#!/bin/bash

set -e

echo "Initializing project..."

go mod download

echo "Running migrations..."

echo "Project initialization complete!"
`
}

func generateDockerCompose() string {
	return `version: '3.8'

services:
  app:
    build:
      context: ..
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=dev
    depends_on:
      - mysql
      - redis
    volumes:
      - ../configs:/app/configs
    networks:
      - app-network

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: myapp
      MYSQL_USER: app
      MYSQL_PASSWORD: app
    ports:
      - "3306:3306"
    volumes:
      - mysql-data:/var/lib/mysql
    networks:
      - app-network

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis-data:/data
    networks:
      - app-network

volumes:
  mysql-data:
  redis-data:

networks:
  app-network:
    driver: bridge
`
}

func generateDockerfile() string {
	return `FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs

EXPOSE 8080

CMD ["./main"]
`
}

func generateDocsReadme(name string) string {
	return fmt.Sprintf(`# %s Documentation

## API Documentation

API documentation is available in the `+"`"+`api/openapi.yaml`+"`"+` file.

## Architecture

### Layer Architecture

1. **Handler Layer** (internal/handler)
   - HTTP request handling
   - Input validation
   - Response formatting

2. **Service Layer** (internal/service)
   - Business logic implementation
   - Transaction management
   - Error handling

3. **Repository Layer** (internal/repository)
   - Data access operations
   - Query building
   - Data mapping

### Design Patterns

- **Dependency Injection**: Services depend on repositories through constructor injection
- **Repository Pattern**: Abstracts data access logic
- **Middleware Pattern**: Cross-cutting concerns like authentication and logging

## Database Schema

See migrations in `+"`"+`scripts/migrations/`+"`"+` directory.

## Configuration

Configuration is managed through YAML files in the `+"`"+`configs/`+"`"+` directory:

- `+"`"+`config.yaml`+"`"+`: Base configuration
- `+"`"+`config.dev.yaml`+"`"+`: Development environment overrides
- `+"`"+`config.prod.yaml`+"`"+`: Production environment overrides

## Testing

Run tests with:

`+"```"+`bash
make test
`+"```"+`

## Deployment

See `+"`"+`deployments/`+"`"+` directory for deployment configurations.
`, name)
}
