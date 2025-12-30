# 删除模型来源接口

## 接口信息

- **路径**: `/api/model-sources/:id`
- **方法**: `DELETE`
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
DELETE /api/model-sources/1
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

#### 无效的ID (400)

```json
{
  "code": 400,
  "message": "无效的ID"
}
```

#### 不存在 (400)

```json
{
  "code": 400,
  "message": "模型来源不存在"
}
```
