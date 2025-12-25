# 登录接口

## 接口信息

- **路径**: `/api/auth/login`
- **方法**: `POST`
- **认证**: 不需要

## 请求头

```
Content-Type: application/json
```

## 请求参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| username | string | 是 | 管理员用户名 |
| password | string | 是 | 管理员密码 |

## 请求示例

```json
{
  "username": "admin",
  "password": "admin123"
}
```

## 响应格式

### 成功响应 (200)

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "username": "admin",
    "user_info": {
      "username": "admin"
    }
  }
}
```

### 错误响应

#### 参数错误 (400)

```json
{
  "code": 400,
  "message": "参数错误: Key: 'LoginRequest.Username' Error:Field validation for 'Username' failed on the 'required' tag"
}
```

#### 用户名或密码错误 (401)

```json
{
  "code": 401,
  "message": "用户名或密码错误"
}
```

## 说明

- 默认管理员账号：`admin` / `admin123`
- 登录成功后返回 JWT Token，有效期为 24 小时
- 后续需要认证的接口需在请求头中携带：`Authorization: Bearer <token>`
- 生产环境请修改 `configs/config.yaml` 中的管理员密码
