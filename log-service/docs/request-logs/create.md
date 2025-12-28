# 创建请求日志

创建单条请求日志记录。

## 接口信息

- **路径**: `/api/request-logs`
- **方法**: `POST`
- **认证**: System Auth Token (Bearer)
- **调用方**: Proxy

## 请求头

```
Authorization: Bearer <system_auth_token>
Content-Type: application/json
```

## 请求体

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| time | string | 是 | 日志时间 (RFC3339 格式) |
| level | string | 是 | 日志级别 (DEBUG/INFO/WARN/ERROR) |
| msg | string | 是 | 日志消息 |
| request_id | string | 是 | 请求 ID |
| method | string | 否 | HTTP 方法 |
| path | string | 否 | 请求路径 |
| query | string | 否 | 查询参数 |
| remote_addr | string | 否 | 远程地址 |
| user_agent | string | 否 | 用户代理 |
| x_forwarded_for | string | 否 | 转发来源 |
| request_headers | object | 否 | 请求头 (键值对) |
| authorization | string | 否 | 认证信息 |
| request_body | string | 否 | 请求体 |
| status | int | 否 | 响应状态码 |
| response_headers | object | 否 | 响应头 (键值对) |
| latency_ms | int64 | 否 | 延迟毫秒 |
| request_size_bytes | int | 否 | 请求大小 (字节) |
| response_size_bytes | int | 否 | 响应大小 (字节) |

## 请求示例

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

## 响应

### 成功响应

**HTTP Status**: 200

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
  "message": "参数错误: ..."
}
```

**HTTP Status**: 500

```json
{
  "code": 500,
  "message": "创建日志记录失败"
}
```
