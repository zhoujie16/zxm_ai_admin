# 获取 Token 及模型信息列表接口

## 接口信息

- **路径**: `/api/tokens/with-model`
- **方法**: `GET`
- **认证**: 需要Bearer Token
- **说明**: 获取所有 Token 及其关联的完整 AI 模型信息，不分页，只返回关联模型存在且未过期的 Token

## 请求头

```
Authorization: Bearer <token>
```

## 请求示例

```
GET /api/tokens/with-model
```

## 响应格式

### 成功响应 (200)

```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "token_id": 4,
      "token": "sk-6dbf9d28087f7abfd9679802fa0dd23be286ab6b7a646ebf0e0a77c268a1e935",
      "token_order_no": "",
      "token_status": 1,
      "token_expire_at": null,
      "token_usage_limit": 0,
      "token_remark": "",
      "ai_model_id": 3,
      "ai_model_name": "DeepSeek",
      "ai_model_api_url": "https://api.deepseek.com",
      "ai_model_api_key": "sk-050a6b276b854988ac7f5583e4f90892",
      "ai_model_remark": "",
      "ai_model_status": 1
    }
  ]
}
```

### 字段说明

| 字段 | 类型 | 说明 |
|------|------|------|
| token_id | uint | Token ID |
| token | string | Token 值 |
| token_order_no | string | 关联订单号 |
| token_status | int | 状态：1=启用，0=禁用 |
| token_expire_at | string/null | 过期时间 |
| token_usage_limit | int | 使用限额（0表示无限制） |
| token_remark | string | 备注 |
| ai_model_id | uint | 关联的 AI 模型 ID |
| ai_model_name | string | AI 模型名称 |
| ai_model_api_url | string | AI 模型 API 地址 |
| ai_model_api_key | string | AI 模型 API Key |
| ai_model_remark | string | AI 模型备注 |
| ai_model_status | int | AI 模型状态：1=启用，0=禁用 |

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
  "message": "查询 Token 列表失败"
}
```
