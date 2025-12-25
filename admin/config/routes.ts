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
  // 配置管理
  {
    path: '/config',
    name: '配置管理',
    icon: 'SettingOutlined',
    routes: [
      {
        path: '/config/proxy-service',
        name: '代理服务管理',
        icon: 'CloudServerOutlined',
        component: './config/proxy-service',
      },
      {
        path: '/config/model',
        name: '模型管理',
        icon: 'RobotOutlined',
        component: './config/model',
      },
    ],
  },

  // 404页面
  {
    path: '*',
    component: './404',
    layout: false,
  },
];

