# 获取 Token 使用记录详情接口

## 接口信息

- **路径**: `/api/token-usage-logs/:id`
- **方法**: `GET`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
```

## 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 记录ID |

## 请求示例

```
GET /api/token-usage-logs/1
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

#### 未认证 (401)

```json
{
  "code": 401,
  "message": "未提供认证token"
}
```

#### 记录不存在 (404)

```json
{
  "code": 404,
  "message": "记录不存在"
}
```
