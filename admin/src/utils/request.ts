/**
 * HTTP 请求统一封装
 * 功能：提供统一的 API 请求方法，处理后端响应格式
 */
import { request as umiRequest } from '@umijs/max';
import type { RequestOptions } from '@@/plugin-request/request';
import type { IBackendResponse } from '@/types';
import { message } from 'antd';

/**
 * 请求配置选项
 */
export interface IRequestOptions extends Omit<RequestOptions, 'method'> {
  /** 是否跳过错误处理（默认 false） */
  skipErrorHandler?: boolean;
  /** 是否显示成功消息（默认 false） */
  showSuccessMessage?: boolean;
  /** 是否显示错误消息（默认 true） */
  showErrorMessage?: boolean;
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
  /** 原始响应 */
  rawResponse?: IBackendResponse<T>;
}

/**
 * 处理后端响应格式
 * @param response 后端响应
 * @returns 统一格式的响应结果
 */
function handleResponse<T = any>(response: IBackendResponse<T>): IRequestResult<T> {
  const { code, message: msg, data } = response;

  return {
    success: code === 0,
    data,
    message: msg || (code === 0 ? '操作成功' : '操作失败'),
    rawResponse: response,
  };
}

/**
 * 处理请求错误
 * @param error 错误对象
 * @param showErrorMessage 是否显示错误消息
 * @returns 错误结果
 */
function handleError(error: any, showErrorMessage: boolean = true): IRequestResult {
  let errorMessage = '请求失败，请稍后重试';

  if (error.response) {
    // HTTP 错误响应
    const { status, data } = error.response;
    
    // 后端可能返回 { code, message } 格式的错误响应
    if (data && typeof data === 'object') {
      errorMessage = data.message || data.data?.message || errorMessage;
    }

    if (showErrorMessage) {
      if (status === 401) {
        message.error(errorMessage || '登录已过期，请重新登录');
      } else if (status === 400) {
        message.error(errorMessage || '参数错误，请检查输入');
      } else if (status === 403) {
        message.error(errorMessage || '没有权限访问该资源');
      } else if (status === 404) {
        message.error(errorMessage || '请求的资源不存在');
      } else if (status >= 500) {
        message.error(errorMessage || '服务器错误，请稍后重试');
      } else {
        message.error(errorMessage);
      }
    }
  } else if (error.name === 'BizError') {
    // 业务错误
    errorMessage = error.info?.errorMessage || errorMessage;
    if (showErrorMessage) {
      message.error(errorMessage);
    }
  } else {
    // 网络错误
    if (showErrorMessage) {
      message.error('网络错误，请检查网络连接');
    }
  }

  return {
    success: false,
    message: errorMessage,
  };
}

/**
 * 通用请求方法
 * @param url 请求地址
 * @param options 请求配置
 * @returns 响应结果
 */
async function request<T = any>(
  url: string,
  options: IRequestOptions & { method?: string } = {},
): Promise<IRequestResult<T>> {
  const {
    skipErrorHandler = false,
    showSuccessMessage = false,
    showErrorMessage = true,
    ...restOptions
  } = options;

  try {
    // 调用 UmiJS 的 request
    const response = await umiRequest<IBackendResponse<T>>(url, {
      ...restOptions,
      skipErrorHandler: true, // 统一在这里处理错误
    });

    // 处理后端响应格式
    const result = handleResponse<T>(response);

    // 显示成功消息
    if (result.success && showSuccessMessage) {
      message.success(result.message);
    }

    // 如果业务失败，显示错误消息
    if (!result.success && showErrorMessage) {
      message.error(result.message);
    }

    return result;
  } catch (error: any) {
    // 处理错误
    return handleError(error, showErrorMessage);
  }
}

/**
 * GET 请求
 * @param url 请求地址
 * @param params 请求参数
 * @param options 请求配置
 * @returns 响应结果
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
 * @param url 请求地址
 * @param data 请求体数据
 * @param options 请求配置
 * @returns 响应结果
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
 * @param url 请求地址
 * @param data 请求体数据
 * @param options 请求配置
 * @returns 响应结果
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
 * @param url 请求地址
 * @param data 请求体数据
 * @param options 请求配置
 * @returns 响应结果
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
 * @param url 请求地址
 * @param data 请求体数据
 * @param options 请求配置
 * @returns 响应结果
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

// 导出默认请求方法
export default request;

