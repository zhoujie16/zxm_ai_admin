# 获取 Authorization 使用次数排行榜

统计 `authorization` 的使用次数并按次数降序排列。

## 接口信息

- **路径**: `/api/request-logs/ranking`
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
| start_time | string | 否 | 开始时间，格式：`2006-01-02 15:04:05`，默认为 7 天前 |
| end_time | string | 否 | 结束时间，格式：`2006-01-02 15:04:05`，默认为当前时间 |
| page | int | 否 | 页码，默认 1 |
| page_size | int | 否 | 每页数量，默认 20，最大 100 |

## 请求示例

```http
GET /api/request-logs/ranking?page=1&page_size=20
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
    "total": 150,
    "time_range": {
      "start": "2025-01-01 00:00:00",
      "end": "2025-01-07 23:59:59"
    },
    "list": [
      {
        "authorization": "Bearer sk-xxx1",
        "count": 5000
      },
      {
        "authorization": "Bearer sk-xxx2",
        "count": 3500
      },
      {
        "authorization": "Bearer sk-xxx3",
        "count": 2000
      }
    ]
  }
}
```

### 响应字段说明

| 字段 | 类型 | 说明 |
|------|------|------|
| total | int64 | 排行榜总记录数（不同 authorization 的数量） |
| time_range.start | string | 查询开始时间 |
| time_range.end | string | 查询结束时间 |
| list | array | 排行榜数据列表 |
| list[].authorization | string | 用户唯一标识 |
| list[].count | int64 | 该 authorization 的请求次数 |

### 错误响应

**HTTP Status**: 400

```json
{
  "code": 400,
  "message": "参数错误: 开始时间格式错误，正确格式为: 2006-01-02 15:04:05"
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
  "message": "获取排行榜数据失败"
}
```
