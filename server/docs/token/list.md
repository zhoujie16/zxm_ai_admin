# 获取 Token 列表接口

## 接口信息

- **路径**: `/api/tokens`
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
| keyword | string | 否 | 搜索关键词（模糊匹配 token 或 remark） |

## 请求示例

```
GET /api/tokens?page=1&page_size=10&keyword=test
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
        "token": "sk-a1b2c3d4e5f6g7h8",
        "ai_model_id": 1,
        "order_no": "ORDER-2024-001",
        "status": 1,
        "expire_at": "2025-12-31T23:59:59Z",
        "usage_limit": 1000,
        "used_count": 125,
        "remark": "测试Token",
        "created_at": "2024-12-26T00:00:00Z",
        "updated_at": "2024-12-26T00:00:00Z"
      },
      {
        "id": 2,
        "token": "sk-i9j8k7l6m5n4o3p2",
        "ai_model_id": 1,
        "order_no": "",
        "status": 1,
        "expire_at": "2025-12-31T23:59:59Z",
        "usage_limit": 0,
        "used_count": 0,
        "remark": "生产Token",
        "created_at": "2024-12-26T01:00:00Z",
        "updated_at": "2024-12-26T01:00:00Z"
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
  "message": "查询Token列表失败"
}
```
