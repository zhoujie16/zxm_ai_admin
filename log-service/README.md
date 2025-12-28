# Log Service

Token 使用记录日志服务，独立部署的日志存储和查询服务。

## 功能

- 接收并存储代理服务转发的请求日志
- 提供 JWT 认证的日志查询接口
- 支持 API Key 认证的日志写入接口

## 技术栈

- Go 1.21+
- Gin Web 框架
- GORM + SQLite
- slog 结构化日志

## 目录结构

```
log-service/
├── cmd/
│   └── server/
│       └── main.go          # 服务入口
├── internal/
│   ├── config/              # 配置管理
│   ├── database/            # 数据库连接
│   ├── handlers/            # HTTP 处理器
│   ├── middleware/          # 中间件
│   ├── models/              # 数据模型
│   └── services/            # 业务逻辑
├── configs/
│   └── config.yaml          # 配置文件
├── data/
│   └── logs.db              # SQLite 数据库（运行时生成）
└── bin/
    └── log-service          # 编译产物
```

## 配置

配置文件位于 `configs/config.yaml`：

```yaml
server:
  port: 6809
  mode: debug                # debug/release

database:
  path: ./data/logs.db

jwt:
  secret: your-jwt-secret    # 与 server 共享

write_api_key: your-api-key  # 写入接口的 API Key
```

## API 接口

### 写入日志

```http
POST /logs
X-Log-API-Key: your-api-key
Content-Type: application/json

{
  "time": "2025-12-27T10:00:00Z",
  "level": "info",
  "request_id": "abc123",
  "method": "POST",
  "path": "/v1/chat/completions",
  "authorization": "Bearer token_xxx",
  "status": 200,
  "latency_ms": 1234
}
```

### 查询日志列表

```http
GET /logs?page=1&page_size=10&start_time=2025-12-01&status=200
Authorization: Bearer <jwt_token>
```

### 获取单条日志

```http
GET /logs/:id
Authorization: Bearer <jwt_token>
```

## 开发

### 安装依赖

```bash
go mod download
```

### 运行

```bash
go run cmd/server/main.go
```

### 构建

```bash
go build -o bin/log-service cmd/server/main.go
```

## 部署

1. 修改 `configs/config.yaml` 配置
2. 构建可执行文件
3. 上传到服务器并运行

```bash
# Linux 构建
CGO_ENABLED=1 GOOS=linux go build -o bin/log-service cmd/server/main.go
```

## 认证说明

- **写入接口** (`POST /logs`)：使用 `X-Log-API-Key` 请求头认证
- **查询接口** (`GET /logs`)：使用 JWT Token 认证，与 server 共享 secret
