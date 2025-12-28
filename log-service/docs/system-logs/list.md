# 获取系统日志列表

分页查询系统日志，支持多种过滤条件。

## 接口信息

- **路径**: `/api/system-logs`
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
| level | string | 否 | 按日志级别过滤 (DEBUG/INFO/WARN/ERROR) |
| start_time | string | 否 | 开始时间 (2006-01-02 15:04:05) |
| end_time | string | 否 | 结束时间 (2006-01-02 15:04:05) |

## 请求示例

```http
GET /api/system-logs?page=1&page_size=10&level=ERROR
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
    "total": 50,
    "list": [
      {
        "id": 1,
        "time": "2024-12-27T10:30:45Z",
        "level": "INFO",
        "msg": "服务启动",
        "created_at": "2024-12-27T10:30:45Z",
        "updated_at": "2024-12-27T10:30:45Z"
      },
      {
        "id": 2,
        "time": "2024-12-27T10:30:46Z",
        "level": "ERROR",
        "msg": "连接失败",
        "created_at": "2024-12-27T10:30:46Z",
        "updated_at": "2024-12-27T10:30:46Z"
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
  "message": "查询系统日志列表失败"
}
```
