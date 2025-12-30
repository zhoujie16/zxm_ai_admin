# 模型来源列表接口

## 接口信息

- **路径**: `/api/model-sources`
- **方法**: `GET`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
```

## 请求参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | int | 否 | 页码，从1开始，默认为1 |
| page_size | int | 否 | 每页数量，默认为10，最大100 |

## 请求示例

```
GET /api/model-sources?page=1&page_size=10
```

## 响应格式

### 成功响应 (200)

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 20,
    "list": [
      {
        "id": 1,
        "model_name": "GPT-4",
        "api_url": "https://api.openai.com/v1",
        "api_key": "sk-xxxxxxx",
        "remark": "OpenAI GPT-4 模型",
        "created_at": "2024-12-30T00:00:00Z",
        "updated_at": "2024-12-30T00:00:00Z"
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
