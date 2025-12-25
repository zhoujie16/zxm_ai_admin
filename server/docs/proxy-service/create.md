# 创建代理服务接口

## 接口信息

- **路径**: `/api/proxy-services`
- **方法**: `POST`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
Content-Type: application/json
```

## 请求参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| service_id | string | 是 | 服务标识，唯一字符串，最大长度100 |
| server_ip | string | 是 | 服务器IP地址 |
| status | int | 否 | 状态：1=启用，0=未启用，默认为1 |
| remark | string | 否 | 备注，最大长度500 |

## 请求示例

```json
{
  "service_id": "proxy-001",
  "server_ip": "192.168.1.100",
  "status": 1,
  "remark": "主代理服务器"
}
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

#### 参数错误 (400)

```json
{
  "code": 400,
  "message": "参数错误: service_id is required"
}
```

#### 服务标识已存在 (400)

```json
{
  "code": 400,
  "message": "服务标识已存在"
}
```

#### IP地址格式无效 (400)

```json
{
  "code": 400,
  "message": "IP地址格式无效"
}
```

#### 未认证 (401)

```json
{
  "code": 401,
  "message": "未提供认证token"
}
```

