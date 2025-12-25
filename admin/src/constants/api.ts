// API 相关常量配置

/**
 * API 基础地址
 */
export const BASE_URL = '/api';

/**
 * HTTP 请求方法
 */
export const HTTP_METHODS = {
  GET: 'GET',
  POST: 'POST',
  PUT: 'PUT',
  DELETE: 'DELETE',
  PATCH: 'PATCH',
} as const;

/**
 * 请求头配置
 */
export const REQUEST_HEADERS = {
  JSON: {
    'Content-Type': 'application/json',
  },
  FORM: {
    'Content-Type': 'application/x-www-form-urlencoded',
  },
} as const;
