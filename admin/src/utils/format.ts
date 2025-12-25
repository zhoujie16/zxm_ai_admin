import dayjs from 'dayjs';

/**
 * 格式化时间字符串
 * @param timeStr 时间字符串或时间戳（支持ISO格式、10位秒级和13位毫秒级时间戳）
 * @param format 格式化模板，默认为 'YYYY-MM-DD HH:mm:ss'
 * @returns 格式化后的时间字符串
 */
export const formatTime = (timeStr: string | number, format = 'YYYY-MM-DD HH:mm:ss'): string => {
  if (!timeStr || timeStr === '0') return '-';

  // 如果是字符串，先尝试直接解析（支持 ISO 格式等）
  if (typeof timeStr === 'string') {
    // 尝试解析为时间戳数字
    const parsedNum = parseInt(timeStr);

    // 如果无法解析为有效数字，或包含日期特征字符，则当作日期字符串直接解析
    if (isNaN(parsedNum) || timeStr.includes('-') || timeStr.includes('T')) {
      return dayjs(timeStr).format(format);
    }

    // 如果是纯数字字符串，继续按时间戳处理
    timeStr = parsedNum;
  }

  // 处理数字时间戳
  let timestamp = timeStr as number;

  // 判断时间戳位数并转换为毫秒
  if (timestamp.toString().length === 10) {
    // 10位时间戳（秒级），转换为毫秒
    timestamp = timestamp * 1000;
  }

  return dayjs(timestamp).format(format);
};

/**
 * 格式化金额数值（纯数据处理）
 * @param amount 金额数值
 * @returns 格式化后的金额字符串（保留两位小数，无千分位）
 */
export const formatMoneyValue = (amount: number): string => {
  return amount.toFixed(2);
};

/**
 * 格式化金额（带货币符号，无千分位）
 * @param amount 金额数值
 * @param currency 货币符号，默认为 '$'
 * @returns 格式化后的金额字符串
 */
export const formatMoney = (amount?: number, currency = '$'): string => {
  if (amount === undefined || amount === null) return '-';

  const formatted = amount.toFixed(2);
  return `${currency}${formatted}`;
};

/**
 * 格式化百分比
 * @param value 数值（0-100形式，如 10 表示 10%）
 * @param fractionDigits 小数位数，默认为 2
 * @returns 格式化后的百分比字符串
 */
export const formatPercent = (value?: number, fractionDigits = 2): string => {
  if (value === undefined || value === null) return '-';
  return `${value.toFixed(fractionDigits)}%`;
};

/**
 * 格式化数字
 * @param value 数值
 * @param fractionDigits 小数位数，默认为 0
 * @returns 格式化后的数字字符串
 */
export const formatNumber = (value?: number, fractionDigits = 0): string => {
  if (value === undefined || value === null) return '-';
  return value.toFixed(fractionDigits);
};

/**
 * 格式化文件大小
 * @param bytes 字节数
 * @returns 格式化后的文件大小字符串
 */
export const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 B';

  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));

  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`;
};

/**
 * 格式化时间（紧凑格式，带星期）
 * @param timeStr 时间字符串或时间戳（支持ISO格式、10位秒级和13位毫秒级时间戳）
 * @returns 格式化后的时间字符串，格式：MM-DD 星期X HH:mm
 */
export const formatTimeWithWeekday = (timeStr: string | number): string => {
  if (!timeStr || timeStr === '0') return '-';

  let date: dayjs.Dayjs;

  // 如果是字符串，先尝试直接解析（支持 ISO 格式等）
  if (typeof timeStr === 'string') {
    // 尝试解析为时间戳数字
    const parsedNum = parseInt(timeStr);

    // 如果无法解析为有效数字，或包含日期特征字符，则当作日期字符串直接解析
    if (isNaN(parsedNum) || timeStr.includes('-') || timeStr.includes('T')) {
      date = dayjs(timeStr);
    } else {
      // 如果是纯数字字符串，继续按时间戳处理
      date = dayjs(parsedNum);
    }
  } else {
    // 处理数字时间戳
    let timestamp = timeStr as number;
    // 判断时间戳位数并转换为毫秒
    if (timestamp.toString().length === 10) {
      // 10位时间戳（秒级），转换为毫秒
      timestamp = timestamp * 1000;
    }
    date = dayjs(timestamp);
  }

  // 星期映射（dayjs 的 day() 返回 0-6，0是周日）
  const weekdays = ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'];
  const weekday = weekdays[date.day()];

  // 格式：MM-DD 星期X HH:mm
  return `${date.format('MM-DD')} ${weekday} ${date.format('HH:mm')}`;
};
