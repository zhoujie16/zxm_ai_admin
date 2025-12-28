# 健康检查

检查服务是否正常运行。

## 接口信息

- **路径**: `/health`
- **方法**: `GET`
- **认证**: 无

## 请求示例

```http
GET /health
```

## 响应

### 成功响应

**HTTP Status**: 200

```json
{
  "code": 0,
  "message": "ok"
}
```
