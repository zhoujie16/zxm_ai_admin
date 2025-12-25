# 删除代理服务接口

## 接口信息

- **路径**: `/api/proxy-services/{id}`
- **方法**: `DELETE`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
```

## 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 代理服务ID |

## 请求示例

```
DELETE /api/proxy-services/1
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

#### 代理服务不存在 (400)

```json
{
  "code": 400,
  "message": "代理服务不存在"
}
```

#### 未认证 (401)

```json
{
  "code": 401,
  "message": "未提供认证token"
}
```

