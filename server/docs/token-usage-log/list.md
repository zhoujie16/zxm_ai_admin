# 获取 Token 使用记录列表接口

## 接口信息

- **路径**: `/api/token-usage-logs`
- **方法**: `GET`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
```

## 查询参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | int | 否 | 页码，从1开始，默认为1 |
| page_size | int | 否 | 每页数量，默认为10，最大100 |
| request_id | string | 否 | 按 request_id 精确查询 |
| start_time | string | 否 | 开始时间，格式：2006-01-02 15:04:05 |
| end_time | string | 否 | 结束时间，格式：2006-01-02 15:04:05 |
| status | int | 否 | 按状态码过滤（-1 表示全部） |
| method | string | 否 | 按 HTTP 方法过滤（GET/POST 等） |
| authorization | string | 否 | 按 Authorization 头模糊匹配 |

## 请求示例

```
GET /api/token-usage-logs?page=1&page_size=10&status=200&method=POST
```

## 响应格式

### 成功响应 (200)

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 50,
    "list": [
      {
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
    ]
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

#### 服务器错误 (500)

```json
{
  "code": 500,
  "message": "查询 Token 使用记录列表失败"
}
```
