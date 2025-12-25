# 获取当前用户信息接口

## 接口信息

- **路径**: `/api/auth/me`
- **方法**: `GET`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
```

## 请求参数

无

## 请求示例

```bash
curl -X GET http://localhost:6808/api/auth/me \
  -H "Authorization: Bearer <token>"
```

## 响应格式

### 成功响应 (200)

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "username": "admin"
  }
}
```

### 错误响应

#### 未提供认证token (401)

```json
{
  "code": 401,
  "message": "未提供认证token"
}
```

#### token格式错误 (401)

```json
{
  "code": 401,
  "message": "token格式错误"
}
```

#### 无效的token (401)

```json
{
  "code": 401,
  "message": "无效的token"
}
```

## 说明

- 该接口用于获取当前登录管理员的用户信息
- Token 需要通过登录接口获取
- Token 有效期为 24 小时，过期后需重新登录
