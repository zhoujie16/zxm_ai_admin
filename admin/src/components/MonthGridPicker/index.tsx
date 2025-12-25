/**
 * 月份网格选择器组件
 * 功能：以网格形式展示多个月份，支持多选、跳空选择，默认显示最近3年
 *
 * 使用示例：
 * ```tsx
 * import MonthGridPicker from '@/components/MonthGridPicker';
 *
 * <MonthGridPicker
 *   value={['2024-01', '2024-03']}
 *   onChange={(months) => console.log(months)}
 * />
 * ```
 */
import React, { useMemo } from 'react';
import dayjs from 'dayjs';
import './index.less';

/**
 * 月份网格选择器 Props
 */
export interface IMonthGridPickerProps {
  /** 已选月份列表（格式：['YYYY-MM', ...]） */
  value?: string[];
  /** 选择变化回调 */
  onChange?: (months: string[]) => void;
  /** 是否禁用 */
  disabled?: boolean;
  /** 显示的年份数量（默认3年） */
  yearCount?: number;
}

/**
 * 月份名称列表
 */
const MONTH_NAMES = [
  '一月',
  '二月',
  '三月',
  '四月',
  '五月',
  '六月',
  '七月',
  '八月',
  '九月',
  '十月',
  '十一月',
  '十二月',
];

/**
 * 月份网格选择器组件
 */
const MonthGridPicker: React.FC<IMonthGridPickerProps> = ({
  value = [],
  onChange,
  disabled = false,
  yearCount = 3,
}) => {
  const currentYear = dayjs().year();

  /**
   * 生成年份列表（默认最近3年）
   */
  const years = useMemo(() => {
    const yearList: number[] = [];
    for (let i = yearCount - 1; i >= 0; i--) {
      yearList.push(currentYear - i);
    }
    return yearList;
  }, [currentYear, yearCount]);

  /**
   * 生成指定年份的月份数据
   */
  const getYearMonths = (year: number) => {
    return MONTH_NAMES.map((name, index) => {
      const month = String(index + 1).padStart(2, '0');
      const monthValue = `${year}-${month}`;
      return {
        name,
        value: monthValue,
        isSelected: value.includes(monthValue),
      };
    });
  };

  /**
   * 切换月份选择状态
   */
  const handleMonthToggle = (monthValue: string) => {
    if (disabled) return;

    const newValue = value.includes(monthValue)
      ? value.filter((v) => v !== monthValue)
      : [...value, monthValue];

    onChange?.(newValue);
  };

  return (
    <div className="month-grid-picker">
      {/* 年份行列表 */}
      {years.map((year) => (
        <div key={year} className="month-grid-picker__year-row">
          <div className="month-grid-picker__year-label">{year}</div>
          <div className="month-grid-picker__grid">
            {getYearMonths(year).map((month) => (
              <button
                key={month.value}
                type="button"
                className={`month-grid-picker__month ${
                  month.isSelected ? 'month-grid-picker__month--selected' : ''
                } ${disabled ? 'month-grid-picker__month--disabled' : ''}`}
                onClick={() => handleMonthToggle(month.value)}
                disabled={disabled}
              >
                {month.name}
              </button>
            ))}
          </div>
        </div>
      ))}

      {/* 已选月份提示 */}
      {value.length > 0 && (
        <div className="month-grid-picker__selected-info">
          已选择 {value.length} 个月份
        </div>
      )}
    </div>
  );
};

export default MonthGridPicker;
