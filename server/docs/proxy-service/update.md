# 更新代理服务接口

## 接口信息

- **路径**: `/api/proxy-services/{id}`
- **方法**: `PUT`
- **认证**: 需要Bearer Token

## 请求头

```
Authorization: Bearer <token>
Content-Type: application/json
```

## 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 代理服务ID |

## 请求参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| service_id | string | 否 | 服务标识，唯一字符串，最大长度100 |
| server_ip | string | 否 | 服务器IP地址 |
| status | int | 否 | 状态：1=启用，0=未启用 |
| remark | string | 否 | 备注，最大长度500 |

**注意**: 所有参数都是可选的，只传需要更新的字段即可。

## 请求示例

```json
{
  "server_ip": "192.168.1.200",
  "status": 0,
  "remark": "已停用"
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
    "server_ip": "192.168.1.200",
    "status": 0,
    "remark": "已停用",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T01:00:00Z"
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

#### 参数错误 (400)

```json
{
  "code": 400,
  "message": "参数错误: IP地址格式无效"
}
```

#### 服务标识已存在 (400)

```json
{
  "code": 400,
  "message": "服务标识已存在"
}
```

#### 状态值无效 (400)

```json
{
  "code": 400,
  "message": "状态值无效，只能为0或1"
}
```

#### 代理服务不存在 (400)

```json
{
  "code": 400,
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

