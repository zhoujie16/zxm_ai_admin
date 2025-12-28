# 获取请求日志详情

根据 ID 获取单条请求日志的详细信息。

## 接口信息

- **路径**: `/api/request-logs/:id`
- **方法**: `GET`
- **认证**: JWT Token (Bearer)
- **调用方**: Admin

## 请求头

```
Authorization: Bearer <JWT_TOKEN>
```

## 路径参数

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | int | 是 | 日志记录 ID |

## 请求示例

```http
GET /api/request-logs/123
Authorization: Bearer <JWT_TOKEN>
```

## 响应

### 成功响应

**HTTP Status**: 200

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 123,
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
    "response_size_bytes": 2048,
    "created_at": "2024-12-27T10:30:45Z",
    "updated_at": "2024-12-27T10:30:45Z"
  }
}
```

### 错误响应

**HTTP Status**: 400

```json
{
  "code": 400,
  "message": "无效的ID"
}
```

**HTTP Status**: 401

```json
{
  "code": 401,
  "message": "无效的token"
}
```

**HTTP Status**: 404

```json
{
  "code": 404,
  "message": "记录不存在"
}
```
