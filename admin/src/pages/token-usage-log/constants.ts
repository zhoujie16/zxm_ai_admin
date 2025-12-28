/**
 * 常量定义
 */

/** 状态选项 */
export const STATUS_OPTIONS = [
  { label: '200 OK', value: 200 },
  { label: '201 Created', value: 201 },
  { label: '204 No Content', value: 204 },
  { label: '400 Bad Request', value: 400 },
  { label: '401 Unauthorized', value: 401 },
  { label: '403 Forbidden', value: 403 },
  { label: '404 Not Found', value: 404 },
  { label: '500 Server Error', value: 500 },
  { label: '502 Bad Gateway', value: 502 },
  { label: '503 Service Unavailable', value: 503 },
] as const;

/** 方法选项 */
export const METHOD_OPTIONS = [
  { label: 'GET', value: 'GET' },
  { label: 'POST', value: 'POST' },
  { label: 'PUT', value: 'PUT' },
  { label: 'DELETE', value: 'DELETE' },
  { label: 'PATCH', value: 'PATCH' },
] as const;

/** 获取状态标签颜色 */
export const getStatusColor = (status: number): string => {
  if (status >= 200 && status < 300) return 'success';
  if (status >= 300 && status < 400) return 'warning';
  if (status >= 400 && status < 500) return 'error';
  return 'default';
};

/** 获取耗时等级 */
export const getLatencyLevel = (ms: number): 'fast' | 'normal' | 'slow' => {
  if (ms < 500) return 'fast';
  if (ms < 2000) return 'normal';
  return 'slow';
};
