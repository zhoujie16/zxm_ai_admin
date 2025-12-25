/**
 * 快速日期区间选择器
 * 功能：封装带常用预设（最近7天/30天、本月、上月）的 RangePicker，方便多场景复用
 *
 * 使用示例：
 * ```tsx
 * import PresetDateRangePicker from '@/components/PresetDateRangePicker';
 *
 * <PresetDateRangePicker value={value} onChange={setValue} />
 * ```
 */
import React from 'react';
import { DatePicker } from 'antd';
import type { RangePickerProps } from 'antd/es/date-picker';
import dayjs from 'dayjs';
import type { Dayjs } from 'dayjs';

const { RangePicker } = DatePicker;

export type PresetDateRangeValue = [Dayjs, Dayjs];

interface IPresetDateRangePickerProps extends Omit<RangePickerProps, 'presets' | 'placeholder'> {
  placeholder?: [string, string];
}

const DATE_RANGE_PRESETS: { label: string; value: PresetDateRangeValue }[] = [
  {
    label: '最近7天',
    value: [dayjs().subtract(7, 'day').startOf('day'), dayjs().endOf('day')],
  },
  {
    label: '最近30天',
    value: [dayjs().subtract(30, 'day').startOf('day'), dayjs().endOf('day')],
  },
  {
    label: '本月',
    value: [dayjs().startOf('month'), dayjs().endOf('month')],
  },
  {
    label: '上月',
    value: [dayjs().subtract(1, 'month').startOf('month'), dayjs().subtract(1, 'month').endOf('month')],
  },
];

const PresetDateRangePicker: React.FC<IPresetDateRangePickerProps> = ({
  placeholder = ['开始日期', '结束日期'],
  ...restProps
}) => {
  return <RangePicker {...restProps} placeholder={placeholder} presets={DATE_RANGE_PRESETS} />;
};

export default PresetDateRangePicker;

