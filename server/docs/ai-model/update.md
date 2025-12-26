# 更新 AI 模型接口

## 接口信息

- **路径**: `/api/ai-models/:id`
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
| id | int | 是 | AI模型ID |

## 请求参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| model_name | string | 否 | 模型名称，最大长度100 |
| api_url | string | 否 | API地址，最大长度500 |
| api_key | string | 否 | API Key，最大长度255 |
| status | int | 否 | 状态：1=启用，0=禁用 |
| remark | string | 否 | 备注，最大长度500 |

## 请求示例

```json
{
  "model_name": "GPT-4 Turbo",
  "api_url": "https://api.openai.com/v1",
  "api_key": "sk-new-key",
  "status": 1,
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
    "model_name": "GPT-4 Turbo",
    "api_url": "https://api.openai.com/v1",
    "api_key": "sk-new-key",
    "status": 1,
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

#### AI模型不存在 (404)

```json
{
  "code": 404,
  "message": "AI模型不存在"
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
  "message": "更新AI模型失败"
}
```
