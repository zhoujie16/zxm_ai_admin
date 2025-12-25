# 删除 AI 模型接口

## 接口信息

- **路径**: `/api/ai-models/:id`
- **方法**: `DELETE`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
```

## 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | AI模型ID |

## 请求示例

```
DELETE /api/ai-models/1
```

## 响应格式

### 成功响应 (200)

```json
{
  "code": 0,
  "message": "success",
  "data": null
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

#### 服务器错误 (500)

```json
{
  "code": 500,
  "message": "删除AI模型失败"
}
```
