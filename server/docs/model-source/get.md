# 模型来源详情接口

## 接口信息

- **路径**: `/api/model-sources/:id`
- **方法**: `GET`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
```

## 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 模型来源ID |

## 请求示例

```
GET /api/model-sources/1
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

#### 无效的ID (400)

```json
{
  "code": 400,
  "message": "无效的ID"
}
```

#### 不存在 (404)

```json
{
  "code": 404,
  "message": "模型来源不存在"
}
```
