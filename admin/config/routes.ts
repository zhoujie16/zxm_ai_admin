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
    name: '模型管理',
    icon: 'RobotOutlined',
    component: './ai-model',
  },
  {
    path: '/token',
    name: 'Token 管理',
    icon: 'KeyOutlined',
    component: './token',
  },
  {
    path: '/token-usage-log',
    name: '使用记录',
    icon: 'FileTextOutlined',
    component: './token-usage-log',
  },

  // 404页面
  {
    path: '*',
    component: './404',
    layout: false,
  },
];

