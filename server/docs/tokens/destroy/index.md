# 永久删除 Token

## 接口信息

- **接口路径**: `/api/tokens/:id/destroy`
- **请求方式**: `DELETE`
- **认证方式**: Bearer Token
- **接口描述**: 永久删除 Token（不可恢复）

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

## 注意事项

- 此操作将永久删除 Token，不可恢复
- 请谨慎操作
