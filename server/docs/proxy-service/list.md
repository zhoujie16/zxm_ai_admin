# 获取代理服务列表接口

## 接口信息

- **路径**: `/api/proxy-services`
- **方法**: `GET`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
```

## 查询参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | int | 否 | 页码，从1开始，默认为1 |
| page_size | int | 否 | 每页数量，默认为10，最大100 |

## 请求示例

```
GET /api/proxy-services?page=1&page_size=10
```

## 响应格式

### 成功响应 (200)

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 100,
    "list": [
      {
        "id": 1,
        "service_id": "proxy-001",
        "server_ip": "192.168.1.100",
        "status": 1,
        "remark": "主代理服务器",
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      },
      {
        "id": 2,
        "service_id": "proxy-002",
        "server_ip": "192.168.1.101",
        "status": 0,
        "remark": "备用代理服务器",
        "created_at": "2024-01-02T00:00:00Z",
        "updated_at": "2024-01-02T00:00:00Z"
      }
    ]
  }
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

#### 服务器错误 (500)

```json
{
  "code": 500,
  "message": "查询代理服务列表失败"
}
```

