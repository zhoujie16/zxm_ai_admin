// 文件说明: UmiJS 运行时配置, 处理初始状态、权限与布局
import type { RequestConfig, RunTimeLayoutConfig } from '@umijs/max';
import { history } from '@umijs/max';
import { Avatar, Dropdown, Space, message } from 'antd';
import { LogoutOutlined, SettingOutlined, UserOutlined } from '@ant-design/icons';
import { SWRConfig } from 'swr';
import { errorConfig } from './requestErrorConfig';
import { isTokenValid, getUserFromToken } from '@/utils/jwt';
import '@ant-design/v5-patch-for-react-19';

// SWR 全局配置
const swrConfig = {
  onError: (error: any) => {
    console.error('SWR Error:', error);
    if (error?.response?.status === 401) {
      message.error('登录已过期，请重新登录');
      setTimeout(() => {
        if (window.location.pathname !== '/login') {
          history.push('/login');
        }
      }, 1000);
    } else if (error?.response?.status === 403) {
      message.error('没有权限访问该资源');
    } else if (error?.response?.status === 404) {
      message.error('请求的资源不存在');
    } else if (error?.response?.status >= 500) {
      message.error('服务器错误，请稍后重试');
    } else {
      message.error('请求失败，请检查网络连接');
    }
  },
};

const loginPath = '/login';

/**
 * @see https://umijs.org/docs/api/runtime-config#getinitialstate
 * */
// 定义用户类型
interface ICurrentUser {
  id: string;
  username: string;
  email: string;
  access: string;
}

// 定义初始状态类型
interface IInitialState {
  currentUser?: ICurrentUser | undefined;
  loading?: boolean;
  fetchUserInfo?: () => Promise<ICurrentUser | undefined>;
}

export async function getInitialState(): Promise<IInitialState> {
  const fetchUserInfo = async () => {
    try {
      const token = localStorage.getItem('token');
      if (!token) {
        throw new Error('No token found');
      }

      // 验证 token 是否有效
      if (!isTokenValid(token)) {
        throw new Error('Token expired');
      }

      // 从 token 中解析用户信息
      const user = getUserFromToken(token);
      if (!user) {
        throw new Error('Invalid token');
      }

      return user;
    } catch (error) {
      // token 无效或过期，清除并跳转到登录页
      localStorage.removeItem('token');
      history.push(loginPath);
      return undefined;
    }
  };

  // 如果不是登录页面，执行
  const { location } = history;
  if (![loginPath].includes(location.pathname)) {
    const currentUser = await fetchUserInfo();
    return {
      fetchUserInfo,
      currentUser,
    };
  }
  return {
    fetchUserInfo,
  };
}

// 页面访问权限检查
export function onRouteChange({ location }: { location: { pathname: string } }) {
  const { pathname } = location;

  // 如果没有登录且不在登录页面，重定向到登录页
  if (!localStorage.getItem('token') && pathname !== loginPath) {
    history.push(loginPath);
  }
}

/**
 * @name ProLayout 配置
 * @doc https://procomponents.ant.design/components/layout
 */
export const layout: RunTimeLayoutConfig = ({ initialState }) => {
  return {
    // 布局配置
    title: 'ZXM Toolkit Admin',
    logo: 'https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg',
    layout: 'mix',
    navTheme: 'light',
    colorPrimary: '#1890ff',
    contentWidth: 'Fluid',
    fixedHeader: false,
    fixSiderbar: true,

    // 右上角用户区域
    rightContentRender: () => {
      const userMenuItems = [
        {
          key: 'profile',
          icon: <UserOutlined />,
          label: '个人设置',
        },
        {
          key: 'settings',
          icon: <SettingOutlined />,
          label: '系统设置',
        },
        {
          type: 'divider' as const,
        },
        {
          key: 'logout',
          icon: <LogoutOutlined />,
          label: '退出登录',
        },
      ];

      const handleUserMenuClick = ({ key }: { key: string }) => {
        if (key === 'logout') {
          // 处理退出登录
          localStorage.removeItem('token');
          history.push(loginPath);
        }
      };

      return (
        <Space>
          <Dropdown
            menu={{
              items: userMenuItems,
              onClick: handleUserMenuClick,
            }}
            placement='bottomRight'
          >
            <Space style={{ cursor: 'pointer', padding: '0 12px' }}>
              <Avatar size='small' icon={<UserOutlined />} />
              <span>{initialState?.currentUser?.username || '管理员'}</span>
            </Space>
          </Dropdown>
        </Space>
      );
    },

    // 菜单配置
    menu: {
      locale: false,
    },

    // 面包屑配置
    breadcrumbRender: (routers = []) => [
      {
        path: '/',
        breadcrumbName: '首页',
      },
      ...routers,
    ],

    // 页脚配置
    footerRender: () => {
      return (
        <div style={{ textAlign: 'center', padding: '16px 0' }}>ZXM Toolkit Admin © 2024</div>
      );
    },
  };
};

/**
 * @name request 配置，可以配置错误处理
 * 它基于 axios 和 ahooks 的 useRequest 提供了一套统一的网络请求和错误处理方案。
 * @doc https://umijs.org/docs/max/request#配置
 */
export const request: RequestConfig = {
  baseURL: '/zxm-toolkit-admin',
  ...errorConfig,
};

/**
 * @name rootContainer 配置，用于包装整个应用
 * 在这里配置 SWR 全局设置
 */
export function rootContainer(container: React.ReactElement) {
  return <SWRConfig value={swrConfig}>{container}</SWRConfig>;
}
