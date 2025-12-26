# 创建 Token 使用记录接口

## 接口信息

- **路径**: `/api/token-usage-logs`
- **方法**: `POST`
- **认证**: 不需要认证（供 proxy 服务调用）

## 请求头

```
Content-Type: application/json
```

## 请求参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| time | string | 是 | 时间戳，格式：RFC3339 |
| level | string | 否 | 日志级别（INFO/WARN/ERROR） |
| msg | string | 否 | 日志消息 |
| request_id | string | 是 | 请求唯一标识（UUID） |
| method | string | 否 | HTTP 方法（GET/POST 等） |
| path | string | 否 | 请求路径 |
| query | string | 否 | URL 查询参数 |
| remote_addr | string | 否 | 客户端地址 |
| user_agent | string | 否 | User-Agent 头 |
| x_forwarded_for | string | 否 | X-Forwarded-For 头 |
| request_headers | object | 否 | 完整请求头（JSON 对象） |
| authorization | string | 否 | Authorization 头（原始值） |
| request_body | string | 否 | 请求体内容 |
| status | int | 否 | HTTP 响应状态码 |
| response_headers | object | 否 | 响应头（JSON 对象） |
| latency_ms | int64 | 否 | 请求耗时（毫秒） |
| request_size_bytes | int | 否 | 请求体大小（字节） |
| response_size_bytes | int | 否 | 响应体大小（字节） |

## 请求示例

```json
{
  "time": "2024-12-27T10:30:45.123Z",
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
    "Content-Type": "application/json",
    "Accept": "application/json"
  },
  "authorization": "Bearer sk-test-token",
  "request_body": "{\"model\":\"glm-4\",\"messages\":[{\"role\":\"user\",\"content\":\"hello\"}]}",
  "status": 200,
  "response_headers": {
    "Content-Type": "application/json"
  },
  "latency_ms": 1234,
  "request_size_bytes": 1024,
  "response_size_bytes": 2048
}
```

## 响应格式

### 成功响应 (200)

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
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
      "Content-Type": "application/json",
      "Accept": "application/json"
    },
    "authorization": "Bearer sk-test-token",
    "request_body": "{\"model\":\"glm-4\",\"messages\":[{\"role\":\"user\",\"content\":\"hello\"}]}",
    "status": 200,
    "response_headers": {
      "Content-Type": "application/json"
    },
    "latency_ms": 1234,
    "request_size_bytes": 1024,
    "response_size_bytes": 2048,
    "created_at": "2024-12-27T10:30:45Z",
    "updated_at": "2024-12-27T10:30:45Z"
  }
}
```

### 错误响应

#### 参数错误 (400)

```json
{
  "code": 400,
  "message": "参数错误: request_id is required"
}
```

#### 服务器错误 (500)

```json
{
  "code": 500,
  "message": "创建 Token 使用记录失败"
}
```
