/**
 * 系统日志搜索过滤组件
 * 功能：提供日志级别和时间范围过滤
 */
import { Button, Form, Select, DatePicker, Space } from 'antd';
import { SearchOutlined, ReloadOutlined } from '@ant-design/icons';
import React, { useCallback } from 'react';
import type { ISystemLogListRequest } from '@/types/systemLog';

const { RangePicker } = DatePicker;
const { Option } = Select;

/**
 * 过滤组件 Props
 */
export interface ISystemLogFilterProps {
  /** 过滤条件变化回调 */
  onFilterChange: (params: ISystemLogListRequest) => void;
  /** 刷新回调 */
  onRefresh: () => void;
  /** 加载状态 */
  loading?: boolean;
}

/**
 * 系统日志搜索过滤组件
 */
const SystemLogFilter: React.FC<ISystemLogFilterProps> = ({
  onFilterChange,
  onRefresh,
  loading = false,
}) => {
  const [form] = Form.useForm();

  /**
   * 处理搜索
   */
  const handleSearch = useCallback(() => {
    const values = form.getFieldsValue();
    const params: ISystemLogListRequest = {};

    if (values.level) {
      params.level = values.level;
    }
    if (values.timeRange && values.timeRange.length === 2) {
      params.start_time = values.timeRange[0].format('YYYY-MM-DD HH:mm:ss');
      params.end_time = values.timeRange[1].format('YYYY-MM-DD HH:mm:ss');
    }

    onFilterChange(params);
  }, [form, onFilterChange]);

  /**
   * 处理重置
   */
  const handleReset = useCallback(() => {
    form.resetFields();
    onFilterChange({});
  }, [form, onFilterChange]);

  return (
    <div className="system-log-filter">
      <Form form={form} layout="inline">
        <Form.Item name="level" label="日志级别">
          <Select placeholder="请选择" allowClear style={{ width: 150 }}>
            <Option value="DEBUG">DEBUG</Option>
            <Option value="INFO">INFO</Option>
            <Option value="WARN">WARN</Option>
            <Option value="ERROR">ERROR</Option>
          </Select>
        </Form.Item>
        <Form.Item name="timeRange" label="时间范围">
          <RangePicker
            showTime
            format="YYYY-MM-DD HH:mm:ss"
            placeholder={['开始时间', '结束时间']}
          />
        </Form.Item>
        <Form.Item>
          <Space>
            <Button icon={<SearchOutlined />} type="primary" onClick={handleSearch}>
              搜索
            </Button>
            <Button onClick={handleReset}>重置</Button>
            <Button icon={<ReloadOutlined />} loading={loading} onClick={onRefresh}>
              刷新
            </Button>
          </Space>
        </Form.Item>
      </Form>
    </div>
  );
};

export default SystemLogFilter;
