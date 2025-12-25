# Token 使用回调接口

## 接口信息

- **路径**: `/api/token/usage`
- **方法**: `POST`
- **认证**: 不需要

## 请求头

```
Content-Type: application/json
```

## 请求参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| token | string | 是 | Token值 |
| remote_ip | string | 是 | 请求来源IP地址 |
| user_agent | string | 是 | User Agent |

## 请求示例

```json
{
  "token": "sk-a1b2c3d4e5f6g7h8",
  "remote_ip": "192.168.1.100",
  "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"
}
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

#### 参数错误 (400)

```json
{
  "code": 400,
  "message": "参数错误: token is required"
}
```

#### 服务器错误 (500)

```json
{
  "code": 500,
  "message": "记录Token使用失败"
}
```

## 说明

- 本接口供外部服务调用，用于记录 Token 的使用情况
- 本接口不需要认证，任何知道接口地址的服务均可调用
- 调用成功后，对应 Token 的 `used_count` 会自动增加
- 同时会在 `token_usage_logs` 表中新增一条使用记录
- 本接口不校验 Token 是否有效，仅做记录用途
