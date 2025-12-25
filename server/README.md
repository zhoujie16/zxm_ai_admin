# ZXM AI Admin Server

ZXM AI Admin Server 是一个基于 Go 语言开发的后台管理系统服务端，提供代理服务管理、管理员认证等功能。

## 功能特性

- 🔐 **JWT 认证** - 基于 JWT 的安全认证机制
- 🚀 **代理服务管理** - 完整的代理服务增删改查功能
- 📊 **SQLite 数据库** - 轻量级数据库，无需额外配置
- 📝 **结构化日志** - 基于 Zap 的结构化日志系统
- 🛡️ **中间件支持** - CORS、认证、日志、错误恢复等中间件

## 技术栈

- **语言**: Go 1.21+
- **Web 框架**: Gin
- **ORM**: GORM
- **数据库**: SQLite
- **认证**: JWT (golang-jwt/jwt/v5)
- **日志**: Zap
- **配置管理**: Viper
- **参数验证**: go-playground/validator

## 快速开始

### 环境要求

- Go 1.21 或更高版本
- 确保 `$GOPATH/bin` 在 `$PATH` 中（如果使用 Go Modules，则不需要）

### 安装依赖

```bash
go mod download
```

### 配置

复制并编辑配置文件：

```bash
cp configs/config.yaml configs/config.yaml.local
```

编辑 `configs/config.yaml` 或 `configs/config.yaml.local`，主要配置项：

```yaml
server:
  port: 6808          # 服务端口
  mode: debug          # 运行模式: debug, release, test

database:
  path: "./data/app.db"  # 数据库文件路径
  max_open_conns: 10     # 最大打开连接数
  max_idle_conns: 5      # 最大空闲连接数

admin:
  username: "admin"      # 管理员用户名
  password: "admin123"   # 管理员密码（生产环境请修改）

jwt:
  secret: "your-secret-key"  # JWT 密钥（生产环境请修改）
  expire_hours: 24           # Token 过期时间（小时）

log:
  level: info              # 日志级别: debug, info, warn, error
  output: "./logs/app.log" # 日志输出路径
```

### 运行服务

```bash
# 使用默认配置文件
go run cmd/server/main.go

# 或指定配置文件
go run cmd/server/main.go configs/config.yaml.local
```

### 构建

```bash
# 构建可执行文件
go build -o bin/server cmd/server/main.go

# 运行
./bin/server
```

## 项目结构

```
.
├── cmd/
│   └── server/
│       └── main.go          # 应用入口
├── configs/
│   └── config.yaml          # 配置文件
├── docs/                     # 文档目录
├── internal/
│   ├── config/              # 配置管理
│   ├── database/            # 数据库连接和迁移
│   ├── handlers/            # HTTP 处理器
│   ├── middleware/          # 中间件
│   ├── models/              # 数据模型
│   ├── repositories/        # 数据访问层（预留）
│   ├── services/            # 业务逻辑层
│   └── utils/               # 工具函数
├── go.mod                   # Go 模块定义
└── README.md               # 项目说明文档
```

## 开发规范

项目遵循严格的开发规范，详细说明请参考 [.cursor/rules/base.mdc](.cursor/rules/base.mdc)

### 主要规范

- **分层架构**: Handlers → Services → Repositories → Database
- **统一响应格式**: 所有接口使用 `utils.Response` 统一返回
- **错误处理**: Services 层返回业务错误，Handlers 层转换为 HTTP 响应
- **代码注释**: 公开函数、类型、变量必须有注释
- **命名规范**: 遵循 Go 语言命名约定

## 数据库

项目使用 SQLite 数据库，数据库文件默认存储在 `./data/app.db`。

### 自动迁移

项目启动时会自动执行数据库迁移，创建必要的表结构。

### 当前数据表

- `proxy_services` - 代理服务表

## 日志

项目使用 Zap 进行结构化日志记录，日志文件默认存储在 `./logs/app.log`。

日志级别可通过配置文件设置：
- `debug` - 调试信息
- `info` - 一般信息
- `warn` - 警告信息
- `error` - 错误信息

## 安全建议

1. **生产环境配置**
   - 修改默认管理员密码为强密码
   - 修改 JWT Secret 为随机强密钥
   - 将运行模式设置为 `release`

2. **数据库安全**
   - 定期备份数据库文件
   - 确保数据库文件权限设置正确

3. **网络安全**
   - 使用 HTTPS（生产环境）
   - 配置防火墙规则
   - 限制 API 访问频率

## 常见问题

### Q: 如何修改服务端口？

A: 编辑 `configs/config.yaml` 文件，修改 `server.port` 配置项。

### Q: 数据库文件在哪里？

A: 默认在 `./data/app.db`，可通过配置文件中的 `database.path` 修改。

### Q: 如何查看日志？

A: 日志文件默认在 `./logs/app.log`，也可通过控制台查看。


## 许可证

本项目为内部项目，仅供内部使用。

## 更新日志

### v1.0.0 (2024-01-01)

- ✨ 初始版本发布
- ✨ 实现管理员认证功能
- ✨ 实现代理服务管理功能

## 贡献

如有问题或建议，请联系项目维护者。

