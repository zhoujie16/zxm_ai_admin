# 恢复已删除的 Token

## 接口信息

- **接口路径**: `/api/tokens/:id/restore`
- **请求方式**: `POST`
- **认证方式**: Bearer Token
- **接口描述**: 恢复已删除的 Token

## 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | integer | 是 | Token ID |

## 响应数据

### 成功响应

```json
{
  "success": true,
  "data": null,
  "message": "ok"
}
```

### 错误响应

```json
{
  "success": false,
  "data": null,
  "message": "Token 不存在"
}
```

```json
{
  "success": false,
  "data": null,
  "message": "Token 未被删除，无需恢复"
}
```
