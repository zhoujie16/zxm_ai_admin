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
| token | string | 是 | Token值，用于筛选该Token的使用记录 |
| page | int | 否 | 页码，从1开始，默认为1 |
| page_size | int | 否 | 每页数量，默认为10，最大100 |

## 请求示例

```
GET /api/token-usage-logs?token=sk-a1b2c3d4e5f6g7h8&page=1&page_size=10
```

## 响应格式

### 成功响应 (200)

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 125,
    "list": [
      {
        "id": 1,
        "token": "sk-a1b2c3d4e5f6g7h8",
        "remote_ip": "192.168.1.100",
        "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
        "call_time": "2024-12-26T10:30:00Z",
        "created_at": "2024-12-26T10:30:00Z"
      },
      {
        "id": 2,
        "token": "sk-a1b2c3d4e5f6g7h8",
        "remote_ip": "192.168.1.101",
        "user_agent": "curl/7.68.0",
        "call_time": "2024-12-26T10:35:00Z",
        "created_at": "2024-12-26T10:35:00Z"
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

#### 参数错误 (400)

```json
{
  "code": 400,
  "message": "参数错误: token is required"
}
```

#### 服务器错误 (500)

```json
{
  "code": 500,
  "message": "查询Token使用记录失败"
}
```

## 说明

- 本接口用于查询指定 Token 的使用记录
- 记录按 `call_time` 倒序排列，最新的记录在前
- 使用记录由 `/api/token/usage` 接口调用时自动生成
