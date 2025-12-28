/**
 * 工具函数
 */
import dayjs from 'dayjs';
import { message } from 'antd';

/** 获取最近7天时间范围 */
export const getLastWeekRange = (): [dayjs.Dayjs, dayjs.Dayjs] => {
  return [dayjs().subtract(7, 'day'), dayjs()];
};

/** 获取本周时间范围 */
export const getThisWeekRange = (): [dayjs.Dayjs, dayjs.Dayjs] => {
  const now = dayjs();
  const start = now.startOf('week').add(1, 'day');
  const end = now.endOf('week').add(1, 'day');
  return [start, end];
};

/** 时间范围快捷选项 */
export const timeRangePresets: Array<{ label: string; value: () => [dayjs.Dayjs, dayjs.Dayjs] }> = [
  { label: '今天', value: () => [dayjs().startOf('day'), dayjs().endOf('day')] },
  { label: '昨天', value: () => [dayjs().subtract(1, 'day').startOf('day'), dayjs().subtract(1, 'day').endOf('day')] },
  { label: '最近7天', value: getLastWeekRange },
  { label: '本周', value: getThisWeekRange },
  {
    label: '上周',
    value: () => [dayjs().subtract(1, 'week').startOf('week').add(1, 'day'), dayjs().subtract(1, 'week').endOf('week').add(1, 'day')],
  },
  { label: '最近30天', value: () => [dayjs().subtract(30, 'day').startOf('day'), dayjs().endOf('day')] },
];

/** 格式化时间 */
export const formatDateTime = (text: string): string => {
  try {
    return dayjs(text).format('YYYY-MM-DD HH:mm:ss');
  } catch {
    return text;
  }
};

/** 格式化耗时 */
export const formatLatency = (ms: number): string => {
  if (ms < 1000) return `${ms}ms`;
  return `${(ms / 1000).toFixed(2)}s`;
};

/** 格式化字节大小 */
export const formatBytes = (bytes: number): string => {
  if (bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${(bytes / Math.pow(k, i)).toFixed(2)} ${sizes[i]}`;
};

/** 复制到剪贴板 */
export const copyToClipboard = (text: string, successMsg: string = '复制成功') => {
  navigator.clipboard.writeText(text).then(() => {
    message.success(successMsg);
  });
};
