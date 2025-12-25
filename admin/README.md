# Admin 管理后台

基于 UmiJS 4.x + React 19 + Ant Design 5.x 构建的现代化管理后台系统。

## 📋 目录

- [技术栈](#技术栈)
- [功能特性](#功能特性)
- [项目结构](#项目结构)
- [快速开始](#快速开始)
- [开发指南](#开发指南)
- [构建部署](#构建部署)
- [脚本说明](#脚本说明)

## 🛠 技术栈

- **框架**: UmiJS 4.3.x
- **React**: 19.1.0
- **UI 组件库**: Ant Design 5.25.4
- **状态管理**: SWR 2.3.6
- **语言**: TypeScript 5.6.3
- **构建工具**: Mako
- **包管理器**: Yarn

## ✨ 功能特性

- ✅ 基于 UmiJS 4.x 的现代化开发体验
- ✅ Ant Design 5.x 组件库
- ✅ TypeScript 完整类型支持
- ✅ 路由配置和权限管理
- ✅ 请求代理和错误处理
- ✅ 代码分割和按需加载
- ✅ 热更新支持

## 📁 项目结构

```
admin/
├── config/              # 配置文件
│   ├── config.ts       # 主配置文件
│   ├── routes.ts       # 路由配置
│   ├── proxy.ts        # 代理配置
│   └── defaultSettings.ts  # 默认设置
├── src/
│   ├── pages/          # 页面组件
│   │   ├── home/       # 首页
│   │   ├── login/      # 登录页
│   │   └── 404/        # 404 页面
│   ├── components/     # 公共组件
│   │   ├── AvatarDropdown/  # 头像下拉菜单
│   │   ├── Footer/     # 页脚
│   │   ├── PageContainer/  # 页面容器
│   │   └── RightContent/   # 右侧内容
│   ├── services/       # API 服务
│   ├── utils/          # 工具函数
│   ├── constants/      # 常量定义
│   ├── types/          # 类型定义
│   ├── layouts/        # 布局组件
│   ├── app.tsx         # 应用入口配置
│   └── access.ts       # 权限配置
├── public/             # 静态资源
└── package.json
```

## 🚀 快速开始

### 前置要求

- Node.js >= 20.0.0
- Yarn >= 1.22.22

### 安装依赖

```bash
yarn install
```

### 启动开发服务器

```bash
# 开发模式（端口 6806）
yarn dev

# 或使用 start:dev
yarn start:dev
```

启动后访问：http://localhost:6806

### 代理配置

开发环境下，前端会自动代理 `/api/` 请求到后端服务器（http://localhost:6808）。

代理配置位于 `config/proxy.ts`。

## 💻 开发指南

### 代码规范

项目使用 TypeScript 和 Prettier 进行代码规范：

```bash
# 类型检查
yarn tsc

# 代码格式化
yarn format
```

### 添加新页面

1. 在 `src/pages/` 目录下创建页面目录
2. 创建页面组件文件 `index.tsx`
3. 在 `config/routes.ts` 中添加路由配置

### 添加新组件

1. 在 `src/components/` 目录下创建组件目录
2. 创建组件文件和相关样式文件
3. 在 `src/components/index.ts` 中导出组件（如需要）

### API 请求

项目使用 SWR 进行数据请求，API 服务定义在 `src/services/` 目录。

示例：

```typescript
import useSWR from 'swr';
import { request } from '@umijs/max';

// 使用 SWR 获取数据
const { data, error } = useSWR('/api/demo', request);
```

### 路由配置

路由配置位于 `config/routes.ts`，支持嵌套路由和权限控制。

### 权限管理

权限配置位于 `src/access.ts`，结合路由配置实现页面级权限控制。

## 🏗 构建部署

### 构建生产版本

```bash
yarn build
```

构建产物输出到 `dist/` 目录。

### 预览构建结果

```bash
yarn preview
```

预览服务器运行在 http://localhost:8000

### 分析构建产物

```bash
yarn analyze
```

生成构建分析报告，用于优化打包体积。

## 📜 脚本说明

| 脚本 | 说明 |
|------|------|
| `yarn dev` | 启动开发服务器（端口 6806） |
| `yarn start` | 启动开发服务器 |
| `yarn start:dev` | 启动开发服务器（无 Mock） |
| `yarn build` | 构建生产版本 |
| `yarn preview` | 预览构建结果 |
| `yarn analyze` | 分析构建产物 |
| `yarn lint` | 类型检查 |
| `yarn format` | 代码格式化 |
| `yarn tsc` | TypeScript 类型检查 |

## 🔧 配置说明

### 端口配置

开发服务器默认运行在 **6806** 端口，可通过环境变量 `PORT` 修改。

### 代理配置

开发环境下的 API 代理配置在 `config/proxy.ts`：

```typescript
dev: {
  '/api/': {
    target: 'http://localhost:6808',
    changeOrigin: true,
  },
}
```

### 环境变量

- `PORT`: 开发服务器端口（默认 6806）
- `UMI_ENV`: 环境标识（dev/prod）
- `MOCK`: Mock 数据开关（none/开启）

## 📝 注意事项

1. 生产环境构建后，代理配置不会生效，需要配置反向代理（如 Nginx）
2. 确保后端服务运行在 6808 端口，否则代理会失败
3. 开发时建议使用 `start:dev` 命令，避免 Mock 数据干扰

## 📄 许可证

UNLICENSED

