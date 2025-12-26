# Proxy 代理服务

一个基于 Go 语言开发的高性能 HTTP 反向代理服务，支持请求转发、认证白名单、请求追踪、结构化日志记录和 **token 动态路由**。

## 功能特性

- 🔄 **反向代理** - 将请求转发到指定的目标服务器
- 🎯 **Token 动态路由** - 根据请求的 Authorization token 自动选择目标服务器和 API Key
- 🔄 **定时同步** - 每 10 分钟自动从后端同步 token-模型配置
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

通过环境变量配置服务：

| 环境变量 | 说明 | 默认值 |
|---------|------|--------|
| `ENABLE_TOKEN_ROUTING` | 是否启用 token 动态路由 | `true` |
| `SERVER_API_URL` | 后端 API 地址 | `http://localhost:6808` |
| `SERVER_API_TOKEN` | 后端 API 认证 Token | （空） |
| `SYNC_INTERVAL` | 缓存同步间隔（分钟） | `10` |
| `LOG_LEVEL` | 日志级别 (debug/info/warn/error) | `info` |
| `LISTEN_ADDR` | 监听地址 | `:6800` |

以下为兼容旧版本的配置项（已废弃）：
| `TARGET_URL` | 目标转发地址 | `https://open.bigmodel.cn` |
| `OVERRIDE_AUTH_TOKEN` | 替换的 Authorization 头 | （空） |
| `REQUIRED_AUTH_TOKENS` | Authorization 白名单 | （空） |

### 运行服务

```bash
# 使用默认配置
go run main.go

# 或使用环境变量
export SERVER_API_URL=http://localhost:6808
export SERVER_API_TOKEN=your-admin-token
export SYNC_INTERVAL=10
go run main.go
```

### 构建

```bash
# 构建可执行文件
go build -o bin/proxy main.go

# 运行
./bin/proxy
```

## Token 动态路由

### 工作原理

1. **配置同步** - 服务启动时从后端 API `/api/tokens/with-model` 获取 token-模型列表
2. **定时刷新** - 每隔指定时间（默认 10 分钟）自动同步最新配置
3. **请求处理** - 收到请求时，根据 `Authorization` 头查找对应的模型配置
4. **动态转发** - 使用配置的 `ai_model_api_url` 作为目标，`ai_model_api_key` 作为认证

### API 响应格式

后端 API 需返回以下格式：

```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "token_id": 1,
      "token": "sk-xxx",
      "token_status": 1,
      "ai_model_id": 1,
      "ai_model_name": "DeepSeek",
      "ai_model_api_url": "https://api.deepseek.com",
      "ai_model_api_key": "sk-api-key",
      "ai_model_status": 1
    }
  ]
}
```

### 缓存策略

- **同步失败** - 保留上一次缓存，继续服务
- **状态过滤** - 只缓存 `token_status=1` 且 `ai_model_status=1` 的记录
- **Token 验证** - 不在缓存中的 token 直接返回 401

## 项目结构

```
proxy/
├── main.go              # 应用入口
├── config/
│   └── config.go        # 配置管理
├── cache/
│   └── token_cache.go   # token 缓存管理
├── proxy/
│   ├── proxy.go         # 反向代理核心逻辑
│   └── response.go      # 响应包装器
├── middleware/
│   ├── auth.go          # 认证中间件（已废弃）
│   └── requestid.go     # 请求ID中间件
├── logger/
│   └── logger.go        # 日志工具
├── go.mod               # Go 模块定义
└── README.md           # 项目说明文档
```

## 工作原理

### 请求处理流程

```
请求 → RequestID → Token 缓存查找 →
  ├─ 找到: 替换 Authorization → 转发到对应 API URL
  └─ 未找到: 返回 401
```

### 中间件链

```
RequestID → Token 路由 → 动态代理
```

### 日志记录

服务会记录以下信息：

- **请求信息**: Method, Path, Query, Headers, Body, Authorization
- **响应信息**: Status Code, Headers, Body
- **性能指标**: 延迟时间 (latency_ms), 响应大小 (response_size_bytes)
- **追踪信息**: RequestID

## 使用场景

1. **多模型统一入口** - 一个代理服务转发到多个 AI 模型 API
2. **动态路由** - 根据 client token 自动选择目标模型
3. **认证隔离** - 客户端 token 与模型 API Key 分离管理
4. **请求追踪** - 通过 RequestID 追踪完整的请求链路

## 安全建议

1. **生产环境配置**
   - 设置 `SERVER_API_TOKEN` 保护后端 API 调用
   - 将 `LOG_LEVEL` 设置为 `info` 或 `warn`
   - 使用 HTTPS 后端 API

2. **网络安全**
   - 在生产环境中使用反向代理（如 Nginx）提供 HTTPS
   - 配置防火墙规则限制访问

3. **日志安全**
   - 注意日志中可能包含敏感信息
   - 定期清理日志文件

## 常见问题

### Q: Token 不在缓存中会发生什么？

A: 返回 `401 Unauthorized` 错误，拒绝请求。

### Q: 后端 API 不可用时会影响服务吗？

A: 不会。同步失败时保留上一次缓存，继续使用旧配置提供服务。

### Q: 如何修改同步间隔？

A: 设置 `SYNC_INTERVAL` 环境变量，单位为分钟。

## 更新日志

### v2.0.0 (2025-01)

- ✨ 新增 token 动态路由功能
- ✨ 新增定时缓存同步
- ✨ 新增多目标代理池
- 🗑️ 废弃固定白名单认证方式

### v1.0.0 (2024-01)

- ✨ 初始版本发布
- ✨ 实现反向代理功能
- ✨ 实现认证白名单功能
- ✨ 实现请求追踪功能
