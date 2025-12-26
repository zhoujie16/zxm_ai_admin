# 获取回收站 Token 列表

## 接口信息

- **接口路径**: `/api/tokens/recycle`
- **请求方式**: `GET`
- **认证方式**: Bearer Token
- **接口描述**: 获取已删除的 Token 列表（分页）

## 请求参数

### Query 参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| page | integer | 否 | 页码，从1开始，默认1 |
| page_size | integer | 否 | 每页数量，默认10，最大100 |
| keyword | string | 否 | 关键词搜索（token或备注） |

## 响应数据

### 成功响应

```json
{
  "success": true,
  "data": {
    "total": 5,
    "list": [
      {
        "id": 1,
        "token": "sk-xxxxxxxx",
        "ai_model_id": 1,
        "model_name": "GPT-4",
        "order_no": "ORD001",
        "status": 1,
        "expire_at": "2024-12-31T23:59:59Z",
        "usage_limit": 1000,
        "remark": "测试Token",
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ]
  },
  "message": "ok"
}
```

### 字段说明

| 字段名 | 类型 | 说明 |
|--------|------|------|
| total | integer | 总数量 |
| list | array | Token列表 |
| list[].id | integer | Token ID |
| list[].token | string | Token值 |
| list[].ai_model_id | integer | 关联的AI模型ID |
| list[].model_name | string | 关联的AI模型名称 |
| list[].order_no | string | 关联订单号 |
| list[].status | integer | 状态：1=启用，0=禁用 |
| list[].expire_at | string | 过期时间 |
| list[].usage_limit | integer | 使用限额 |
| list[].remark | string | 备注 |
| list[].created_at | string | 创建时间 |
| list[].updated_at | string | 更新时间 |
