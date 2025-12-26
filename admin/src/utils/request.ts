/**
 * HTTP 请求统一封装（基于 axios）
 */
import axios, { AxiosError, InternalAxiosRequestConfig, AxiosResponse } from 'axios';
import type { IBackendResponse } from '@/types';
import { message } from 'antd';
import { history } from '@umijs/max';

/**
 * 请求配置选项
 */
export interface IRequestOptions {
  /** 请求头 */
  headers?: Record<string, string>;
}

/**
 * 请求响应结果
 */
export interface IRequestResult<T = any> {
  /** 是否成功 */
  success: boolean;
  /** 响应数据 */
  data?: T;
  /** 响应消息 */
  message: string;
}

// 创建 axios 实例
const axiosInstance = axios.create({
  baseURL: '/zxm-ai-admin',
  timeout: 30000,
});

// 请求拦截器
axiosInstance.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = localStorage.getItem('token');
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

// 响应拦截器 - 处理错误
axiosInstance.interceptors.response.use(
  (response) => response,
  (error: AxiosError) => {
    const { response } = error;

    if (response) {
      const { status } = response;

      // 尝试获取后端返回的错误信息
      let errorMessage = '请求失败，请稍后重试';
      const data = response.data as any;
      if (data?.message) {
        errorMessage = data.message;
      }

      switch (status) {
        case 401:
          errorMessage = '登录已过期，请重新登录';
          localStorage.removeItem('token');
          setTimeout(() => {
            if (window.location.pathname !== '/login') {
              history.push('/login');
            }
          }, 1000);
          break;
        case 403:
          errorMessage = '没有权限访问该资源';
          break;
        case 404:
          errorMessage = '请求的资源不存在';
          break;
        case 500:
        case 502:
        case 503:
        case 504:
          errorMessage = '服务器错误，请稍后重试';
          break;
        default:
          if (data?.message) {
            errorMessage = data.message;
          }
      }

      message.error(errorMessage);
    } else {
      // 网络错误
      message.error('网络错误，请检查网络连接');
    }

    return Promise.reject(error);
  },
);

/**
 * 处理响应数据
 */
function handleResponse<T>(response: AxiosResponse<IBackendResponse>): IRequestResult<T> {
  const { code, data, message: msg } = response.data;

  // 业务成功
  if (code === 0) {
    return {
      success: true,
      data,
      message: msg || '操作成功',
    };
  }

  // 业务失败，显示错误提示
  message.error(msg || '操作失败');
  return {
    success: false,
    data,
    message: msg || '操作失败',
  };
}

/**
 * 通用请求方法
 */
async function request<T = any>(
  url: string,
  options: IRequestOptions & { method?: string; data?: any; params?: any } = {},
): Promise<IRequestResult<T>> {
  const { method = 'GET', data, params, headers } = options;

  try {
    const config: any = {
      url,
      method,
      data,
      params,
    };
    if (headers) {
      config.headers = headers;
    }
    const response = await axiosInstance.request<IBackendResponse>(config);
    return handleResponse<T>(response as any);
  } catch (error) {
    // 错误已在响应拦截器中处理并弹出提示
    return {
      success: false,
      message: '请求失败',
    };
  }
}

/**
 * GET 请求
 */
export async function get<T = any>(
  url: string,
  params?: Record<string, any>,
  options?: IRequestOptions,
): Promise<IRequestResult<T>> {
  return request<T>(url, {
    method: 'GET',
    params,
    ...options,
  });
}

/**
 * POST 请求
 */
export async function post<T = any>(
  url: string,
  data?: Record<string, any>,
  options?: IRequestOptions,
): Promise<IRequestResult<T>> {
  return request<T>(url, {
    method: 'POST',
    data,
    ...options,
  });
}

/**
 * PUT 请求
 */
export async function put<T = any>(
  url: string,
  data?: Record<string, any>,
  options?: IRequestOptions,
): Promise<IRequestResult<T>> {
  return request<T>(url, {
    method: 'PUT',
    data,
    ...options,
  });
}

/**
 * DELETE 请求
 */
export async function del<T = any>(
  url: string,
  data?: Record<string, any>,
  options?: IRequestOptions,
): Promise<IRequestResult<T>> {
  return request<T>(url, {
    method: 'DELETE',
    data,
    ...options,
  });
}

/**
 * PATCH 请求
 */
export async function patch<T = any>(
  url: string,
  data?: Record<string, any>,
  options?: IRequestOptions,
): Promise<IRequestResult<T>> {
  return request<T>(url, {
    method: 'PATCH',
    data,
    ...options,
  });
}

export default request;
