# Log Service API 文档

## 服务信息

- **端口**: 6809
- **数据源**: SQLite (`./data/logs.db`)

## 认证方式

### 写入日志（Proxy 调用）

使用 `system_auth_token` 进行 Bearer Token 认证：

```
Authorization: Bearer zxm-ai-admin-secret-key-change-in-production
```

### 查询日志（Admin 调用）

使用 JWT Token 认证（与 server 相同的 secret）：

```
Authorization: Bearer <JWT_TOKEN>
```

## 接口列表

### 请求日志

| 接口 | 文档 |
|------|------|
| POST /api/request-logs | [创建请求日志](./request-logs/create.md) |
| POST /api/request-logs/batch | [批量创建请求日志](./request-logs/batch-create.md) |
| GET /api/request-logs | [获取请求日志列表](./request-logs/list.md) |
| GET /api/request-logs/:id | [获取请求日志详情](./request-logs/get.md) |

### 系统日志

| 接口 | 文档 |
|------|------|
| POST /api/system-logs/batch | [批量创建系统日志](./system-logs/batch-create.md) |
| GET /api/system-logs | [获取系统日志列表](./system-logs/list.md) |
| GET /api/system-logs/:id | [获取系统日志详情](./system-logs/get.md) |

### 其他

| 接口 | 文档 |
|------|------|
| GET /health | [健康检查](./common/health.md) |

## 数据模型

### TokenUsageLog (请求日志)

| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键 |
| time | time.Time | 日志时间 |
| level | string | 日志级别 |
| msg | string | 日志消息 |
| request_id | string | 请求 ID |
| method | string | HTTP 方法 |
| path | string | 请求路径 |
| query | string | 查询参数 |
| remote_addr | string | 远程地址 |
| user_agent | string | 用户代理 |
| x_forwarded_for | string | 转发来源 |
| request_headers | JSONMap | 请求头 |
| authorization | string | 认证信息 |
| request_body | string | 请求体 |
| status | int | 响应状态码 |
| response_headers | JSONMap | 响应头 |
| latency_ms | int64 | 延迟毫秒 |
| request_size_bytes | int | 请求大小 |
| response_size_bytes | int | 响应大小 |

### SystemLog (系统日志)

| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键 |
| time | time.Time | 日志时间 |
| level | string | 日志级别 (DEBUG/INFO/WARN/ERROR) |
| msg | string | 日志消息 |
