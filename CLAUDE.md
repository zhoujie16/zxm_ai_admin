# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 通用规范

- 交互全程使用中文交流

- 当需要对代码进行修改时，必须先暂停操作，用条理化的方式描述你的方案（描述过程不需要代码演示），等待明确确认后，才能实施具体修改

- 修改结束后仅需在对话中生成最终总结，不要生成中间过程文档

- 只实现当前需求必要的功能，没有提到的功能不要擅自添加

- 不需要写兼容性强的代码，除非特别要求

## 项目概述

ZXM AI Admin 是一个个人工具集合管理平台，采用 Monorepo 结构，包含三个独立服务：

- **admin/** - 前端管理界面（React + UmiJS + Ant Design），端口 6806
- **server/** - 后端 API 服务（Go + Gin + GORM），端口 6808
- **proxy/** - 代理服务（Go），端口 6800

---

## 常用命令

### 安装依赖
```bash
make install    # 安装所有子项目的依赖
```

### 开发调试
```bash
# 需要在不同终端分别运行
make dev-admin   # 前端 http://localhost:6806
make dev-server  # 后端 http://localhost:6808
make dev-proxy   # 代理 http://localhost:6800
```

### 构建
```bash
make build       # 构建所有项目
make build-admin # 仅构建前端
```

### 测试
```bash
make test        # 运行 Go 后端和代理服务的测试
cd admin && yarn lint  # 前端类型检查
```

### 清理
```bash
make clean       # 清理所有构建产物
```

---

## 前端架构 (admin/)

### 技术栈
- React 19.1 + UmiJS 4.3 + Ant Design 5.25
- TypeScript 5.6（strict mode）
- SWR 数据获取，Mako 构建工具

### 目录结构
- `config/` - UmiJS 配置（路由、代理、主题）
- `src/pages/` - 页面组件
- `src/components/` - 共享组件
- `src/services/` - API 服务定义
- `src/utils/` - 工具函数（JWT、请求等）
- `src/types/` - TypeScript 类型定义
- `src/access.ts` - 权限控制

### 关键配置
- **路由配置**：`config/routes.ts` - 文件路由，自动生成
- **代理配置**：`config/proxy.ts` - 开发环境 `/api` 代理到后端 6808
- **运行时配置**：`src/app.tsx` - 初始状态、布局、SWR 全局配置
- **部署路径**：`publicPath: '/zxm-ai-admin/'`，`base: '/zxm-ai-admin'`

### 认证流程
- JWT Token 存储在 localStorage
- `src/utils/jwt.ts` 处理 token 验证和解析
- `src/access.ts` 实现路由权限控制
- 401 错误自动跳转登录页

---

## 后端架构 (server/)

### 技术栈
- Go 1.21+ + Gin + GORM + SQLite
- JWT 认证，Zap 日志，Viper 配置

### 分层架构（严格遵守）
```
Handlers (internal/handlers/) → Services (internal/services/) → Models (internal/models/)
```

- **Handlers**：HTTP 请求处理、参数验证、调用 Services
- **Services**：业务逻辑实现
- **Models**：GORM 数据模型
- **Middleware**：认证、日志、CORS、错误恢复

### 响应格式（必须使用 utils 统一格式）
```go
utils.Success(c, data)              // 200
utils.BadRequest(c, message)        // 400
utils.Unauthorized(c, message)      // 401
utils.Forbidden(c, message)         // 403
utils.NotFound(c, message)          // 404
utils.InternalServerError(c, message) // 500
```

### 配置管理
- 配置文件：`configs/config.yaml`
- 通过 `config.GetConfig()` 获取配置
- 包含服务器、数据库、管理员、JWT、日志配置

### 中间件顺序（从外到内）
1. RecoveryMiddleware - 错误恢复
2. LoggerMiddleware - 请求日志
3. CORSMiddleware - 跨域
4. AuthMiddleware - 认证（保护路由）

---

## 代理服务架构 (proxy/)

### 技术栈
- Go 1.21+ + net/http/httputil
- slog 结构化日志

### 配置（环境变量）
- `TARGET_URL` - 目标转发地址（默认：`https://open.bigmodel.cn`）
- `OVERRIDE_AUTH_TOKEN` - 替换 Authorization 头
- `REQUIRED_AUTH_TOKENS` - Authorization 白名单
- `LISTEN_ADDR` - 监听地址（默认：`:6800`）

### 中间件链
1. RequestID - 生成唯一请求 ID
2. Auth - 验证白名单（可选）
3. Proxy - 反向代理

---

## 重要提示

1. **部署路径**：前端部署在 `/zxm-ai-admin/` 非根目录
2. **端口分配**：6800 代理、6806 前端、6808 后端
3. **认证方式**：后端使用 JWT Bearer Token
4. **代理限制**：代理仅在开发环境有效
5. **数据库**：SQLite 文件默认在 `./data/app.db`
