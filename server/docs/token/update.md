# 更新 Token 接口

## 接口信息

- **路径**: `/api/tokens/:id`
- **方法**: `PUT`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
Content-Type: application/json
```

## 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | Token ID |

## 请求参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| ai_model_id | int | 否 | 关联的AI模型ID |
| order_no | string | 否 | 关联订单号，最大长度100 |
| status | int | 否 | 状态：1=启用，0=禁用 |
| expire_at | string | 否 | 过期时间，格式：YYYY-MM-DD HH:mm:ss |
| usage_limit | int | 否 | 使用限额，0表示无限制 |
| remark | string | 否 | 备注，最大长度500 |

## 请求示例

```json
{
  "ai_model_id": 2,
  "order_no": "ORDER-2024-002",
  "status": 1,
  "expire_at": "2025-12-31 23:59:59",
  "usage_limit": 2000,
  "remark": "更新后的备注"
}
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
    "ai_model_id": 2,
    "order_no": "ORDER-2024-002",
    "status": 1,
    "expire_at": "2025-12-31T23:59:59Z",
    "usage_limit": 2000,
    "used_count": 125,
    "remark": "更新后的备注",
    "created_at": "2024-12-26T00:00:00Z",
    "updated_at": "2024-12-26T02:00:00Z"
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

#### 参数错误 (400)

```json
{
  "code": 400,
  "message": "参数错误: 无效的参数"
}
```

#### 服务器错误 (500)

```json
{
  "code": 500,
  "message": "更新Token失败"
}
```

## 说明

- `token` 字段不可修改
- `used_count` 字段由系统自动维护，不可手动修改
