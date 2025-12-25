# 获取代理服务详情接口

## 接口信息

- **路径**: `/api/proxy-services/{id}`
- **方法**: `GET`
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
GET /api/proxy-services/1
```

## 响应格式

### 成功响应 (200)

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "service_id": "proxy-001",
    "server_ip": "192.168.1.100",
    "status": 1,
    "remark": "主代理服务器",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
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

#### 代理服务不存在 (404)

```json
{
  "code": 404,
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

