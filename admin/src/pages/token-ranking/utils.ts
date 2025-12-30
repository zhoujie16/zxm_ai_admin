/**
 * 工具函数
 */

/**
 * 格式化字节数
 */
export function formatBytes(bytes: number): string {
  if (bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${(bytes / Math.pow(k, i)).toFixed(2)} ${sizes[i]}`;
}

/**
 * 格式化数字
 */
export function formatNumber(num: number): string {
  return num.toLocaleString();
}

/**
 * 格式化时间
 */
export function formatDateTime(timeStr: string): string {
  if (!timeStr) return '-';
  const date = new Date(timeStr);
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    hour12: false,
  });
}

/**
 * 格式化日期
 */
export function formatDate(dateStr: string): string {
  if (!dateStr) return '-';
  const date = new Date(dateStr);
  return date.toLocaleDateString('zh-CN', {
    month: '2-digit',
    day: '2-digit',
  });
}

/**
 * 格式化小时
 */
export function formatHour(timeStr: string): string {
  if (!timeStr) return '-';
  const date = new Date(timeStr);
  return date.toLocaleTimeString('zh-CN', {
    hour: '2-digit',
    minute: '2-digit',
    hour12: false,
  });
}

/**
 * 获取最近7天时间范围
 */
export function getLastWeekRange(): [string, string] {
  const end = new Date();
  const start = new Date();
  start.setDate(start.getDate() - 7);
  return [start.toISOString(), end.toISOString()];
}

/**
 * 获取最近30天时间范围
 */
export function getLastMonthRange(): [string, string] {
  const end = new Date();
  const start = new Date();
  start.setDate(start.getDate() - 30);
  return [start.toISOString(), end.toISOString()];
}

/**
 * 截断长文本
 */
export function truncateText(text: string, maxLength: number = 30): string {
  if (!text) return '-';
  if (text.length <= maxLength) return text;
  return text.slice(0, maxLength) + '...';
}

/**
 * 获取 Token 显示名称（脱敏）
 */
export function getTokenDisplayName(authorization: string): string {
  if (!authorization) return '-';
  // 去掉 Bearer 前缀
  const token = authorization.replace(/^Bearer\s+/i, '');
  if (token.length <= 10) return token;
  // 显示前6位和后4位
  return `${token.slice(0, 6)}...${token.slice(-4)}`;
}
