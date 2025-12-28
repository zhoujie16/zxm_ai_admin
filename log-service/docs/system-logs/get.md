# 获取系统日志详情

根据 ID 获取单条系统日志的详细信息。

## 接口信息

- **路径**: `/api/system-logs/:id`
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
GET /api/system-logs/123
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
    "msg": "服务启动",
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
