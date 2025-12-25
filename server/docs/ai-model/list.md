# 获取 AI 模型列表接口

## 接口信息

- **路径**: `/api/ai-models`
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
| keyword | string | 否 | 搜索关键词（模糊匹配 model_key 或 model_name） |

## 请求示例

```
GET /api/ai-models?page=1&page_size=10&keyword=gpt
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
        "model_key": "gpt-4",
        "model_name": "GPT-4",
        "api_url": "https://api.openai.com/v1",
        "status": 1,
        "remark": "OpenAI GPT-4 模型",
        "created_at": "2024-12-26T00:00:00Z",
        "updated_at": "2024-12-26T00:00:00Z"
      },
      {
        "id": 2,
        "model_key": "gpt-3.5-turbo",
        "model_name": "GPT-3.5 Turbo",
        "api_url": "https://api.openai.com/v1",
        "status": 1,
        "remark": "OpenAI GPT-3.5 模型",
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
  "message": "查询AI模型列表失败"
}
```
