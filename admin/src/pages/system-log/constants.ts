/**
 * 常量定义
 */

/** 日志级别选项 */
export const LEVEL_OPTIONS = [
  { label: 'DEBUG', value: 'DEBUG' },
  { label: 'INFO', value: 'INFO' },
  { label: 'WARN', value: 'WARN' },
  { label: 'ERROR', value: 'ERROR' },
] as const;

/** 获取日志级别标签颜色 */
export const getLevelColor = (level: string): string => {
  switch (level) {
    case 'DEBUG':
      return 'default';
    case 'INFO':
      return 'processing';
    case 'WARN':
      return 'warning';
    case 'ERROR':
      return 'error';
    default:
      return 'default';
  }
};
