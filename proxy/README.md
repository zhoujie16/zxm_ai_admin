# Proxy 代理服务

一个基于 Go 语言开发的高性能 HTTP 反向代理服务，支持请求转发、认证白名单、请求追踪和结构化日志记录。

## 功能特性

- 🔄 **反向代理** - 将请求转发到指定的目标服务器
- 🔐 **认证白名单** - 支持基于 Authorization 头的请求白名单验证
- 🔑 **Token 替换** - 自动替换请求中的 Authorization 头为指定值
- 📝 **请求追踪** - 为每个请求生成唯一的 RequestID，便于追踪和调试
- 📊 **结构化日志** - 使用 JSON 格式记录详细的请求和响应信息
- ⚡ **高性能** - 基于 Go 标准库 `net/http/httputil` 实现，性能优异

## 技术栈

- **语言**: Go 1.21+
- **核心库**: `net/http/httputil` (反向代理)
- **日志**: `log/slog` (结构化日志)
- **配置**: 环境变量

## 快速开始

### 环境要求

- Go 1.21 或更高版本

### 安装依赖

```bash
go mod download
```

### 配置

通过环境变量配置服务，所有配置项都有默认值：

| 环境变量 | 说明 | 默认值 |
|---------|------|--------|
| `TARGET_URL` | 目标转发地址 | `https://open.bigmodel.cn` |
| `OVERRIDE_AUTH_TOKEN` | 替换的 Authorization 头 | `Bearer xxxxx` |
| `REQUIRED_AUTH_TOKENS` | Authorization 白名单（逗号分隔） | `Bearer 123456,Bearer abcdef` |
| `LOG_LEVEL` | 日志级别 (debug/info/warn/error) | `info` |
| `LISTEN_ADDR` | 监听地址 | `:6800` |

### 运行服务

```bash
# 使用默认配置
go run main.go

# 或使用环境变量
export TARGET_URL=https://api.example.com
export OVERRIDE_AUTH_TOKEN="Bearer your-token"
export REQUIRED_AUTH_TOKENS="Bearer token1,Bearer token2"
export LOG_LEVEL=debug
export LISTEN_ADDR=:9090
go run main.go
```

### 构建

```bash
# 构建可执行文件
go build -o bin/proxy main.go

# 运行
./bin/proxy
```

## 使用示例

### 基本使用

```bash
# 启动代理服务（转发到 https://api.example.com）
export TARGET_URL=https://api.example.com
go run main.go
```

### 启用认证白名单

```bash
# 只允许指定的 Authorization token 访问
export REQUIRED_AUTH_TOKENS="Bearer secret-token-1,Bearer secret-token-2"
go run main.go
```

### 替换 Authorization 头

```bash
# 将所有请求的 Authorization 头替换为指定值
export OVERRIDE_AUTH_TOKEN="Bearer your-backend-token"
go run main.go
```

### 完整配置示例

```bash
export TARGET_URL=https://api.example.com
export OVERRIDE_AUTH_TOKEN="Bearer backend-secret-token"
export REQUIRED_AUTH_TOKENS="Bearer client-token-1,Bearer client-token-2"
export LOG_LEVEL=debug
export LISTEN_ADDR=:6800
go run main.go
```

## 项目结构

```
proxy/
├── main.go              # 应用入口
├── config/
│   └── config.go        # 配置管理
├── proxy/
│   ├── proxy.go         # 反向代理核心逻辑
│   └── response.go      # 响应包装器
├── middleware/
│   ├── auth.go          # 认证中间件
│   └── requestid.go     # 请求ID中间件
├── logger/
│   └── logger.go        # 日志工具
├── go.mod               # Go 模块定义
└── README.md           # 项目说明文档
```

## 工作原理

### 请求处理流程

1. **RequestID 中间件** - 为每个请求生成唯一的 RequestID
2. **Auth 中间件** - 验证请求的 Authorization 头是否在白名单中（如果配置了白名单）
3. **Proxy 处理器** - 将请求转发到目标服务器，并替换 Authorization 头（如果配置了）
4. **日志记录** - 记录请求和响应的详细信息

### 中间件链

```
RequestID -> Auth -> Proxy
```

### 日志记录

服务会记录以下信息：

- **请求信息**: Method, Path, Query, Headers, Body, Authorization
- **响应信息**: Status Code, Headers, Body
- **性能指标**: 延迟时间 (latency_ms), 响应大小 (response_size_bytes)
- **追踪信息**: RequestID

日志格式为 JSON，便于日志收集和分析。

## 配置说明

### TARGET_URL

目标服务器的完整 URL，例如：
- `https://api.example.com`
- `http://localhost:3000`

### OVERRIDE_AUTH_TOKEN

如果设置了此值，所有转发到目标服务器的请求都会使用此值替换原始的 Authorization 头。这对于统一使用后端认证 token 的场景很有用。

### REQUIRED_AUTH_TOKENS

Authorization 白名单，多个值用逗号分隔。如果配置了此值，只有 Authorization 头匹配白名单中任一值的请求才会被转发。

**注意**: 如果此值为空，则不进行认证验证，所有请求都会被转发。

### LOG_LEVEL

日志级别，可选值：
- `debug` - 调试信息
- `info` - 一般信息（默认）
- `warn` - 警告信息
- `error` - 错误信息

### LISTEN_ADDR

服务监听地址，格式为 `:端口号` 或 `IP:端口号`，例如：
- `:6800` - 监听所有网络接口的 6800 端口
- `127.0.0.1:6800` - 只监听本地回环地址的 6800 端口

## 使用场景

1. **API 网关** - 作为多个后端服务的统一入口
2. **认证代理** - 统一管理 API 认证 token
3. **请求转发** - 将请求转发到不同的后端服务
4. **请求追踪** - 通过 RequestID 追踪完整的请求链路
5. **日志收集** - 集中收集和记录所有请求日志

## 安全建议

1. **生产环境配置**
   - 修改默认的 `OVERRIDE_AUTH_TOKEN` 为安全的 token
   - 配置 `REQUIRED_AUTH_TOKENS` 限制访问
   - 将 `LOG_LEVEL` 设置为 `info` 或 `warn`，避免泄露敏感信息

2. **网络安全**
   - 使用 HTTPS 目标服务器（`TARGET_URL` 使用 `https://`）
   - 在生产环境中使用反向代理（如 Nginx）提供 HTTPS
   - 配置防火墙规则限制访问

3. **日志安全**
   - 注意日志中可能包含敏感信息（如 Authorization token）
   - 定期清理日志文件
   - 使用日志收集工具时注意数据脱敏

## 常见问题

### Q: 如何禁用认证白名单？

A: 不设置 `REQUIRED_AUTH_TOKENS` 环境变量，或设置为空字符串。

### Q: 如何查看详细的请求日志？

A: 设置 `LOG_LEVEL=debug` 环境变量。

### Q: 如何修改监听端口？

A: 设置 `LISTEN_ADDR` 环境变量，例如 `export LISTEN_ADDR=:9090`。

### Q: 代理服务支持 WebSocket 吗？

A: 支持，响应包装器实现了 `Hijack` 方法，可以支持 WebSocket 连接。

### Q: 如何处理代理错误？

A: 代理错误会被记录到日志中，并返回 `502 Bad Gateway` 状态码给客户端。

## 开发规范

项目遵循严格的开发规范，详细说明请参考 [.cursor/rules/base.mdc](../server/.cursor/rules/base.mdc)

### 主要规范

- **分层架构**: 配置、中间件、代理逻辑分离
- **错误处理**: 所有错误都要记录日志
- **代码注释**: 公开函数、类型、变量必须有注释
- **命名规范**: 遵循 Go 语言命名约定

## 许可证

本项目为内部项目，仅供内部使用。

## 更新日志

### v1.0.0 (2024-01-01)

- ✨ 初始版本发布
- ✨ 实现反向代理功能
- ✨ 实现认证白名单功能
- ✨ 实现请求追踪功能
- ✨ 实现结构化日志记录

## 贡献

如有问题或建议，请联系项目维护者。

