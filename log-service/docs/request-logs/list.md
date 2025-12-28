# 获取请求日志列表

分页查询请求日志，支持多种过滤条件。

## 接口信息

- **路径**: `/api/request-logs`
- **方法**: `GET`
- **认证**: JWT Token (Bearer)
- **调用方**: Admin

## 请求头

```
Authorization: Bearer <JWT_TOKEN>
```

## 查询参数

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | int | 否 | 页码，默认 1 |
| page_size | int | 否 | 每页数量，默认 10，最大 100 |
| request_id | string | 否 | 按 request_id 精确查询 |
| start_time | string | 否 | 开始时间 (2006-01-02 15:04:05) |
| end_time | string | 否 | 结束时间 (2006-01-02 15:04:05) |
| status | string | 否 | 状态码，支持单个(200)或多个逗号分隔(200,401,404) |
| method | string | 否 | 按 HTTP 方法过滤 (GET/POST/PUT/DELETE 等) |
| authorization | string | 否 | 按 Authorization 模糊匹配 |

## 请求示例

```http
GET /api/request-logs?page=1&page_size=10&status=200,404&method=POST
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
    "total": 100,
    "list": [
      {
        "id": 1,
        "time": "2024-12-27T10:30:45Z",
        "level": "INFO",
        "msg": "proxy_request",
        "request_id": "550e8400-e29b-41d4-a716-446655440000",
        "method": "POST",
        "path": "/v1/chat/completions",
        "status": 200,
        "latency_ms": 1234,
        ...
      }
    ]
  }
}
```

### 错误响应

**HTTP Status**: 400

```json
{
  "code": 400,
  "message": "参数错误: ..."
}
```

**HTTP Status**: 401

```json
{
  "code": 401,
  "message": "无效的token"
}
```

**HTTP Status**: 500

```json
{
  "code": 500,
  "message": "查询日志列表失败"
}
```
