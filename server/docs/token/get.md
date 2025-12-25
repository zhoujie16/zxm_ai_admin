# 获取 Token 详情接口

## 接口信息

- **路径**: `/api/tokens/:id`
- **方法**: `GET`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
```

## 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | Token ID |

## 请求示例

```
GET /api/tokens/1
```

## 响应格式

### 成功响应 (200)

```json
{
  "code": 0,
  "message": "success",
  "data": {
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

#### Token不存在 (404)

```json
{
  "code": 404,
  "message": "Token不存在"
}
```

#### 服务器错误 (500)

```json
{
  "code": 500,
  "message": "查询Token详情失败"
}
```
