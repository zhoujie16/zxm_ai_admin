/**
 * 认证相关 API 服务
 */
import { post, get } from '@/utils/request';

// 登录响应数据类型
export interface ILoginResponse {
  token: string;
  username: string;
  user_info: {
    username: string;
  };
}

// 用户信息类型
export interface IUserInfo {
  username: string;
  // 可根据实际接口扩展
}

/**
 * 登录
 * @param username 用户名
 * @param password 密码
 * @returns 登录结果
 */
export async function login(username: string, password: string) {
  return post<ILoginResponse>('/api/auth/login', {
    username,
    password,
  });
}

/**
 * 获取当前用户信息
 * @returns 用户信息
 */
export async function getCurrentUser() {
  return get<IUserInfo>('/api/auth/me');
}

