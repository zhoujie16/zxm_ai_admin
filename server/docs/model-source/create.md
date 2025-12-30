# 创建模型来源接口

## 接口信息

- **路径**: `/api/model-sources`
- **方法**: `POST`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
Content-Type: application/json
```

## 请求参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| model_name | string | 是 | 模型名称，最大长度100 |
| api_url | string | 是 | API地址，最大长度500 |
| api_key | string | 是 | API Key，最大长度255，唯一 |
| remark | string | 否 | 备注，最大长度500 |

## 请求示例

```json
{
  "model_name": "GPT-4",
  "api_url": "https://api.openai.com/v1",
  "api_key": "sk-xxxxxxx",
  "remark": "OpenAI GPT-4 模型"
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
    "model_name": "GPT-4",
    "api_url": "https://api.openai.com/v1",
    "api_key": "sk-xxxxxxx",
    "remark": "OpenAI GPT-4 模型",
    "created_at": "2024-12-30T00:00:00Z",
    "updated_at": "2024-12-30T00:00:00Z"
  }
}
```

### 错误响应

#### 参数错误 (400)

```json
{
  "code": 400,
  "message": "参数错误: model_name is required"
}
```

#### API Key已存在 (400)

```json
{
  "code": 400,
  "message": "API Key 已存在"
}
```

#### 未认证 (401)

```json
{
  "code": 401,
  "message": "未提供认证token"
}
```
