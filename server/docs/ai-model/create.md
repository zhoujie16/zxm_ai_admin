# 创建 AI 模型接口

## 接口信息

- **路径**: `/api/ai-models`
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
| model_key | string | 是 | 模型Key，最大长度100 |
| model_name | string | 是 | 模型名称，最大长度100 |
| api_url | string | 是 | API地址，最大长度500 |
| status | int | 否 | 状态：1=启用，0=禁用，默认为1 |
| remark | string | 否 | 备注，最大长度500 |

## 请求示例

```json
{
  "model_key": "gpt-4",
  "model_name": "GPT-4",
  "api_url": "https://api.openai.com/v1",
  "status": 1,
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
    "model_key": "gpt-4",
    "model_name": "GPT-4",
    "api_url": "https://api.openai.com/v1",
    "status": 1,
    "remark": "OpenAI GPT-4 模型",
    "created_at": "2024-12-26T00:00:00Z",
    "updated_at": "2024-12-26T00:00:00Z"
  }
}
```

### 错误响应

#### 参数错误 (400)

```json
{
  "code": 400,
  "message": "参数错误: model_key is required"
}
```

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
  "message": "创建AI模型失败"
}
```
