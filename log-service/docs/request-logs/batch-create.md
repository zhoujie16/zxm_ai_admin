# 批量创建请求日志

批量创建多条请求日志记录。

## 接口信息

- **路径**: `/api/request-logs/batch`
- **方法**: `POST`
- **认证**: System Auth Token (Bearer)
- **调用方**: Proxy

## 请求头

```
Authorization: Bearer <system_auth_token>
Content-Type: application/json
```

## 请求体

请求体为请求日志对象的数组，每个对象字段同 [创建请求日志](./create.md)。

## 请求示例

```json
[
  {
    "time": "2024-12-27T10:30:45Z",
    "level": "INFO",
    "msg": "proxy_request",
    "request_id": "550e8400-e29b-41d4-a716-446655440000",
    "method": "POST",
    "path": "/v1/chat/completions",
    "status": 200,
    "latency_ms": 1234
  },
  {
    "time": "2024-12-27T10:30:46Z",
    "level": "INFO",
    "msg": "proxy_request",
    "request_id": "660e8400-e29b-41d4-a716-446655440001",
    "method": "POST",
    "path": "/v1/chat/completions",
    "status": 200,
    "latency_ms": 567
  }
]
```

## 响应

### 成功响应

**HTTP Status**: 200

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "count": 2
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
