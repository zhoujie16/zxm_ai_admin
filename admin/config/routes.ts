/**
 * @name umi 的路由配置
 * @description 基础模板路由配置
 * @doc https://umijs.org/docs/guides/routes
 */
export default [
  // 登录页面 (无布局)
  {
    path: '/login',
    component: './login',
    layout: false,
  },

  // 主应用路由
  {
    path: '/',
    redirect: '/home',
  },
  {
    path: '/home',
    name: '首页',
    icon: 'HomeOutlined',
    component: './home',
  },
  // 配置管理菜单（一级）
  {
    path: '/proxy-service',
    name: '代理服务管理',
    icon: 'CloudServerOutlined',
    component: './proxy-service',
  },
  {
    path: '/ai-model',
    name: '模型代理管理',
    icon: 'RobotOutlined',
    component: './ai-model',
  },
  {
    path: '/model-source',
    name: '模型来源管理',
    icon: 'ApiOutlined',
    component: './model-source',
  },
  {
    path: '/token',
    name: 'Token 管理',
    icon: 'KeyOutlined',
    component: './token',
  },
  // 请求日志（一级）
  {
    path: '/token-usage-log',
    name: '请求日志',
    icon: 'HistoryOutlined',
    component: './token-usage-log',
  },
  // 系统日志（一级）
  {
    path: '/system-log',
    name: '系统日志',
    icon: 'FileTextOutlined',
    component: './system-log',
  },
  // 删除日志（一级）
  {
    path: '/delete-log',
    name: '删除日志',
    icon: 'DeleteOutlined',
    component: './delete-log',
  },

  // 404页面
  {
    path: '*',
    component: './404',
    layout: false,
  },
];

