# 获取用户请求统计数据

根据 `authorization` 字段统计用户请求的各项数据指标。

## 接口信息

- **路径**: `/api/request-logs/statistics`
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
| authorization | string | 是 | 用户唯一标识（authorization 字段值） |
| start_time | string | 否 | 开始时间，格式：`2006-01-02 15:04:05`，默认为 7 天前 |
| end_time | string | 否 | 结束时间，格式：`2006-01-02 15:04:05`，默认为当前时间 |

## 请求示例

```http
GET /api/request-logs/statistics?authorization=Bearer%20sk-xxx&start_time=2025-01-01%2000:00:00&end_time=2025-01-07%2023:59:59
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
    "authorization": "Bearer sk-xxx",
    "time_range": {
      "start": "2025-01-01 00:00:00",
      "end": "2025-01-07 23:59:59"
    },
    "summary": {
      "total_requests": 1500,
      "total_request_bytes": 1024000,
      "total_response_bytes": 5120000,
      "avg_request_bytes": 682,
      "avg_response_bytes": 3413
    },
    "latency": {
      "total_ms": 180000,
      "avg_ms": 120,
      "min_ms": 20,
      "max_ms": 2000
    },
    "by_ip": [
      { "ip": "1.2.3.4", "count": 800 },
      { "ip": "5.6.7.8", "count": 700 }
    ],
    "by_path": [
      { "path": "/v1/chat/completions", "count": 1000 },
      { "path": "/v1/completions", "count": 500 }
    ],
    "by_date": [
      { "date": "2025-01-01", "count": 100 },
      { "date": "2025-01-02", "count": 150 }
    ],
    "by_time": [
      { "time": "2025-01-01 00:00:00", "count": 50 },
      { "time": "2025-01-01 01:00:00", "count": 30 }
    ]
  }
}
```

### 响应字段说明

#### summary（汇总统计）

| 字段 | 类型 | 说明 |
|------|------|------|
| total_requests | int64 | 总请求次数 |
| total_request_bytes | int64 | 总请求字节数 |
| total_response_bytes | int64 | 总响应字节数 |
| avg_request_bytes | int64 | 平均请求字节数 |
| avg_response_bytes | int64 | 平均响应字节数 |

#### latency（延迟统计）

| 字段 | 类型 | 说明 |
|------|------|------|
| total_ms | int64 | 延迟总和（毫秒） |
| avg_ms | int64 | 平均延迟（毫秒） |
| min_ms | int64 | 最小延迟（毫秒） |
| max_ms | int64 | 最大延迟（毫秒） |

#### by_ip（按 IP 分组统计）

| 字段 | 类型 | 说明 |
|------|------|------|
| ip | string | x_forwarded_for 值 |
| count | int64 | 该 IP 的请求次数 |

#### by_path（按路径分组统计）

| 字段 | 类型 | 说明 |
|------|------|------|
| path | string | 请求路径 |
| count | int64 | 该路径的请求次数 |

#### by_date（按日期分组统计）

| 字段 | 类型 | 说明 |
|------|------|------|
| date | string | 日期（YYYY-MM-DD） |
| count | int64 | 该日期的请求次数 |

#### by_time（按小时分组统计）

| 字段 | 类型 | 说明 |
|------|------|------|
| time | string | 小时（YYYY-MM-DD HH:00:00） |
| count | int64 | 该小时的请求次数 |

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
  "message": "获取汇总统计失败"
}
```
