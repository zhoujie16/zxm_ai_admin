/**
 * 常量定义
 */

/** 状态选项 */
export const STATUS_OPTIONS = [
  { label: '200 OK', value: 200 },
  { label: '401 Unauthorized', value: 401 },
  { label: '404 Not Found', value: 404 },
  { label: '500 Server Error', value: 500 },
  { label: '其它', value: 0 },
] as const;

/** 「其它」选项对应的所有状态码（排除 200/401/404/500） */
export const OTHER_STATUSES = [
  // 2xx
  201, 202, 203, 204, 205, 206, 207, 208, 226,
  // 3xx
  300, 301, 302, 303, 304, 305, 306, 307, 308,
  // 4xx
  400, 402, 403, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 420, 421, 422, 423, 424, 425, 426, 428, 429, 431, 451,
  // 5xx
  501, 502, 503, 504, 505, 506, 507, 508, 510, 511,
];

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
