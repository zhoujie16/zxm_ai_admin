# 删除 Token 接口

## 接口信息

- **路径**: `/api/tokens/:id`
- **方法**: `DELETE`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
```

## 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | Token ID |

## 请求示例

```
DELETE /api/tokens/1
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

#### Token不存在 (404)

```json
{
  "code": 404,
  "message": "Token不存在"
}
```

#### 服务器错误 (500)

```json
{
  "code": 500,
  "message": "删除Token失败"
}
```

## 说明

- 本接口执行软删除，数据不会从数据库中物理删除
- 删除后的 Token 无法再使用
