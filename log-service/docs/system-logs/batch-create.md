# 批量创建系统日志

批量创建多条系统日志记录。

## 接口信息

- **路径**: `/api/system-logs/batch`
- **方法**: `POST`
- **认证**: System Auth Token (Bearer)
- **调用方**: Proxy

## 请求头

```
Authorization: Bearer <system_auth_token>
Content-Type: application/json
```

## 请求体

请求体为系统日志对象的数组。

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| time | string | 是 | 日志时间 (RFC3339 格式) |
| level | string | 是 | 日志级别 (DEBUG/INFO/WARN/ERROR) |
| msg | string | 是 | 日志消息 |

## 请求示例

```json
[
  {
    "time": "2024-12-27T10:30:45Z",
    "level": "INFO",
    "msg": "服务启动"
  },
  {
    "time": "2024-12-27T10:30:46Z",
    "level": "ERROR",
    "msg": "连接失败"
  },
  {
    "time": "2024-12-27T10:30:47Z",
    "level": "WARN",
    "msg": "内存使用率较高"
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
    "count": 3
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
  "message": "批量创建系统日志记录失败"
}
```
