# Log Service API 文档

## 服务信息

- **端口**: 6809
- **数据源**: SQLite (`./data/logs.db`)

## 认证方式

### 1. 写入日志（Proxy 调用）

使用 `write_key` 进行 Bearer Token 认证：

```
Authorization: Bearer log-service-write-key-change-in-production
```

### 2. 查询日志（Admin 调用）

使用 JWT Token 认证（与 server 相同的 secret）：

```
Authorization: Bearer <JWT_TOKEN>
```

---

## API 接口

### 1. 创建日志记录

**请求**

```http
POST /api/logs
Authorization: Bearer <write_key>
Content-Type: application/json
```

**请求体**

```json
{
  "time": "2024-12-27T10:30:45Z",
  "level": "INFO",
  "msg": "proxy_request",
  "request_id": "550e8400-e29b-41d4-a716-446655440000",
  "method": "POST",
  "path": "/v1/chat/completions",
  "query": "",
  "remote_addr": "127.0.0.1:54321",
  "user_agent": "Mozilla/5.0",
  "x_forwarded_for": "",
  "request_headers": {
    "Content-Type": "application/json"
  },
  "authorization": "Bearer sk-test-token",
  "request_body": "{\"model\":\"glm-4\"}",
  "status": 200,
  "response_headers": {
    "Content-Type": "application/json"
  },
  "latency_ms": 1234,
  "request_size_bytes": 1024,
  "response_size_bytes": 2048
}
```

**响应**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "time": "2024-12-27T10:30:45Z",
    ...
  }
}
```

---

### 2. 获取日志列表

**请求**

```http
GET /api/logs?page=1&page_size=10&status=200
Authorization: Bearer <JWT_TOKEN>
```

**查询参数**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | int | 否 | 页码，默认 1 |
| page_size | int | 否 | 每页数量，默认 10 |
| request_id | string | 否 | 按 request_id 精确查询 |
| start_time | string | 否 | 开始时间 (2006-01-02 15:04:05) |
| end_time | string | 否 | 结束时间 |
| status | int | 否 | 按状态码过滤 |
| method | string | 否 | 按 HTTP 方法过滤 |
| authorization | string | 否 | 按 Authorization 模糊匹配 |

**响应**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 100,
    "list": [...]
  }
}
```

---

### 3. 获取日志详情

**请求**

```http
GET /api/logs/:id
Authorization: Bearer <JWT_TOKEN>
```

**响应**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "time": "2024-12-27T10:30:45Z",
    ...
  }
}
```

---

### 4. 健康检查

**请求**

```http
GET /health
```

**响应**

```json
{
  "code": 0,
  "message": "ok"
}
```
