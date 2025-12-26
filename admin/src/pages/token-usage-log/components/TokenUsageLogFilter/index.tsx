/**
 * Token 使用记录搜索过滤组件
 * 功能：提供多种过滤条件查询
 */
import { Button, Form, Input, Select, DatePicker, Space, Row, Col } from 'antd';
import { SearchOutlined, ReloadOutlined } from '@ant-design/icons';
import React, { useCallback } from 'react';
import type { ITokenUsageLogListRequest } from '@/types/tokenUsageLog';

const { RangePicker } = DatePicker;
const { Option } = Select;

/**
 * 过滤组件 Props
 */
export interface ITokenUsageLogFilterProps {
  /** 过滤条件变化回调 */
  onFilterChange: (params: ITokenUsageLogListRequest) => void;
  /** 刷新回调 */
  onRefresh: () => void;
  /** 加载状态 */
  loading?: boolean;
}

/**
 * Token 使用记录搜索过滤组件
 */
const TokenUsageLogFilter: React.FC<ITokenUsageLogFilterProps> = ({
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
    const params: ITokenUsageLogListRequest = {};

    if (values.request_id) {
      params.request_id = values.request_id;
    }
    if (values.method) {
      params.method = values.method;
    }
    if (values.status !== undefined && values.status !== -1) {
      params.status = values.status;
    }
    if (values.authorization) {
      params.authorization = values.authorization;
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
    <div className="token-usage-log-filter">
      <Form form={form} layout="inline">
        <Row gutter={16} style={{ width: '100%' }}>
          <Col span={6}>
            <Form.Item name="request_id" label="Request ID">
              <Input placeholder="请输入 Request ID" allowClear />
            </Form.Item>
          </Col>
          <Col span={6}>
            <Form.Item name="method" label="请求方法">
              <Select placeholder="请选择" allowClear style={{ width: '100%' }}>
                <Option value="GET">GET</Option>
                <Option value="POST">POST</Option>
                <Option value="PUT">PUT</Option>
                <Option value="DELETE">DELETE</Option>
                <Option value="PATCH">PATCH</Option>
              </Select>
            </Form.Item>
          </Col>
          <Col span={6}>
            <Form.Item name="status" label="状态码">
              <Select placeholder="请选择" allowClear style={{ width: '100%' }} defaultValue={-1}>
                <Option value={-1}>全部</Option>
                <Option value={200}>200 OK</Option>
                <Option value={201}>201 Created</Option>
                <Option value={204}>204 No Content</Option>
                <Option value={400}>400 Bad Request</Option>
                <Option value={401}>401 Unauthorized</Option>
                <Option value={403}>403 Forbidden</Option>
                <Option value={404}>404 Not Found</Option>
                <Option value={500}>500 Server Error</Option>
                <Option value={502}>502 Bad Gateway</Option>
                <Option value={503}>503 Service Unavailable</Option>
              </Select>
            </Form.Item>
          </Col>
          <Col span={6}>
            <Form.Item name="authorization" label="Authorization">
              <Input placeholder="请输入 Token" allowClear />
            </Form.Item>
          </Col>
        </Row>
        <Row gutter={16} style={{ width: '100%', marginTop: 8 }}>
          <Col span={12}>
            <Form.Item name="timeRange" label="时间范围">
              <RangePicker
                showTime
                format="YYYY-MM-DD HH:mm:ss"
                style={{ width: '100%' }}
                placeholder={['开始时间', '结束时间']}
              />
            </Form.Item>
          </Col>
          <Col span={12} style={{ textAlign: 'right' }}>
            <Space>
              <Button icon={<SearchOutlined />} type="primary" onClick={handleSearch}>
                搜索
              </Button>
              <Button onClick={handleReset}>重置</Button>
              <Button icon={<ReloadOutlined />} loading={loading} onClick={onRefresh}>
                刷新
              </Button>
            </Space>
          </Col>
        </Row>
      </Form>
    </div>
  );
};

export default TokenUsageLogFilter;
