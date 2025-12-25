# components 目录说明

## 目录定位

- 存放 `apps/admin` 端所有可复用组件，分为布局组件与业务组件
- 统一由 `index.ts` 暴露，方便按需导入和避免循环依赖
- 所有组件遵循函数式写法、中文关键注释、单文件不超过 200 行

## 子目录速览

| 目录            | 角色 | 说明                                    |
| --------------- | ---- | --------------------------------------- |
| `Footer/`        | 布局 | 管理端全局页脚组件，`index.ts` 中以 `IFooter` 命名导出 |
| `HeaderDropdown/` | 布局 | 顶部栏下拉集合，含头像、通知等入口      |
| `RightContent/` | 布局 | 顶部右侧区域，包含 `AvatarDropdown`、`AvatarName`、`Question` |
| `AvatarDropdown/` | 布局 | 用户头像下拉菜单组件                    |
| `PageContainer/` | 布局 | 页面容器组件，提供统一的页面容器样式    |

> 统一入口 `index.ts` 会补充类型导出，可先查此文件了解依赖关系。

## 新增/调整组件注意事项

1. 新增目录、文件必须包含中文头部注释，说明功能与示例场景
2. 逻辑与 UI 分离：在组件内拆分纯逻辑 Hook（放入 `hooks/`）与展示组件
3. 禁止使用 `any`，涉及类型需在 `@/types` 或组件内 `types.ts` 定义并以 `I` 前缀命名
4. 组件需保持无副作用，跨组件状态通过 props 或上层状态管理传递
5. 若组件对外暴露类型或方法，务必同步更新本目录 `README.md` 以及 `index.ts`
6. 若组件依赖接口数据，确认 `doc-server-api` 中相关文档已更新

## 常用引用示例

```tsx
import { Footer, PageContainer } from '@components';
```

- 统一从 `@components` 引入，便于 tree-shaking 与路径管理
- 若仅需类型，可 `import type { ... } from '@components';`
