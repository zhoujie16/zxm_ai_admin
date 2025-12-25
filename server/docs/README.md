# API 接口文档

本文档包含 ZXM AI Admin Server 的所有 API 接口说明。

## 文档结构

文档按功能模块进行组织，每个模块对应一个独立的目录：

```
docs/
├── README.md              # 本文档（总目录）
├── proxy-service/         # 代理服务管理模块
│   ├── create.md          # 创建代理服务
│   ├── list.md            # 获取代理服务列表
│   ├── get.md             # 获取代理服务详情
│   ├── update.md          # 更新代理服务
│   └── delete.md          # 删除代理服务
└── auth/                  # 认证模块（待补充）
└── system/                # 系统模块（待补充）
```

## 通用说明

### 基础信息

- **Base URL**: `http://localhost:6808`
- **API 前缀**: `/api`
- **数据格式**: JSON

### 认证方式

大部分接口需要 JWT 认证，请在请求头中携带：

```
Authorization: Bearer <token>
```

获取 token 的方式请参考 [认证模块文档](./auth/)（待补充）。

### 统一响应格式

所有接口使用统一的响应格式：

#### 成功响应

```json
{
  "code": 0,
  "message": "success",
  "data": { ... }
}
```

#### 错误响应

```json
{
  "code": <错误码>,
  "message": "<错误信息>"
}
```

### 错误码说明

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 400 | 参数错误 |
| 401 | 未认证 |
| 403 | 无权限 |
| 404 | 资源不存在 |
| 500 | 服务器错误 |

## 模块列表

### 1. 代理服务管理 (`/api/proxy-services`)

管理代理服务的增删改查操作。

- [创建代理服务](./proxy-service/create.md)
- [获取代理服务列表](./proxy-service/list.md)
- [获取代理服务详情](./proxy-service/get.md)
- [更新代理服务](./proxy-service/update.md)
- [删除代理服务](./proxy-service/delete.md)

### 2. 认证模块 (`/api/auth`)

管理员认证相关接口。

- 登录接口（待补充）
- 获取当前用户信息（待补充）

### 3. 系统模块

系统相关接口。

- 健康检查接口（待补充）

## 更新日志

- 2024-01-01: 初始版本，添加代理服务管理模块文档


