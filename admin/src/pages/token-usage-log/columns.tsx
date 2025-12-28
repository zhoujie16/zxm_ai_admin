/**
 * 表格列定义
 */
import { Space, Button, Tag, Badge, Tooltip } from 'antd';
import { EyeOutlined, CopyOutlined } from '@ant-design/icons';
import type { ProColumns } from '@ant-design/pro-components';
import type { ITokenUsageLog } from '@/types/tokenUsageLog';
import { STATUS_OPTIONS, METHOD_OPTIONS, getStatusColor, getLatencyLevel } from './constants';
import { formatDateTime, formatLatency, formatBytes, copyToClipboard, getLastWeekRange, timeRangePresets } from './utils';

export interface IGetColumnsParams {
  /** 查看详情回调 */
  onViewDetail: (record: ITokenUsageLog) => void;
}

/** 获取表格列定义 */
export const getColumns = (params: IGetColumnsParams): ProColumns<ITokenUsageLog>[] => {
  const { onViewDetail } = params;

  return [
    {
      title: '时间',
      dataIndex: 'time',
      width: 180,
      search: false,
      render: (_, record) => formatDateTime(record.time),
    },
    {
      title: '时间范围',
      dataIndex: 'timeRange',
      valueType: 'dateTimeRange',
      hideInTable: true,
      initialValue: getLastWeekRange(),
      fieldProps: {
        placeholder: ['开始时间', '结束时间'],
        showTime: true,
        presets: timeRangePresets,
      },
    },
    {
      title: '方法',
      dataIndex: 'method',
      width: 80,
      valueType: 'select',
      valueEnum: METHOD_OPTIONS.reduce((acc, opt) => ({ ...acc, [opt.value]: { text: opt.label } }), {}),
      render: (_, record) => <Tag className={`method-tag ${record.method}`}>{record.method}</Tag>,
      filters: false,
    },
    {
      title: '路径',
      dataIndex: 'path',
      width: 200,
      ellipsis: true,
      search: false,
      render: (_, record) => (
        <Tooltip title={record.path}><span>{record.path}</span></Tooltip>
      ),
    },
    {
      title: '状态码',
      dataIndex: 'status',
      width: 100,
      valueType: 'select',
      valueEnum: STATUS_OPTIONS.reduce((acc, opt) => ({ ...acc, [opt.value]: { text: opt.label } }), {}),
      render: (_, record) => <Tag color={getStatusColor(record.status)}>{record.status}</Tag>,
      filters: false,
    },
    {
      title: '耗时',
      dataIndex: 'latency_ms',
      width: 100,
      search: false,
      render: (_, record) => (
        <Badge
          status={getLatencyLevel(record.latency_ms) === 'fast' ? 'success' : getLatencyLevel(record.latency_ms) === 'normal' ? 'warning' : 'error'}
          text={formatLatency(record.latency_ms)}
        />
      ),
    },
    {
      title: '请求大小',
      dataIndex: 'request_size_bytes',
      width: 100,
      search: false,
      render: (_, record) => formatBytes(record.request_size_bytes),
    },
    {
      title: '响应大小',
      dataIndex: 'response_size_bytes',
      width: 100,
      search: false,
      render: (_, record) => formatBytes(record.response_size_bytes),
    },
    {
      title: 'X-Forwarded-For',
      dataIndex: 'x_forwarded_for',
      width: 150,
      ellipsis: true,
      search: false,
      render: (_, record) => <Tooltip title={record.x_forwarded_for}><span>{record.x_forwarded_for}</span></Tooltip>,
    },
    {
      title: 'User-Agent',
      dataIndex: 'user_agent',
      width: 200,
      ellipsis: true,
      search: false,
      render: (_, record) => <Tooltip title={record.user_agent}><span>{record.user_agent}</span></Tooltip>,
    },
    {
      title: 'Authorization',
      dataIndex: 'authorization',
      width: 150,
      ellipsis: true,
      hideInTable: true,
    },
    {
      title: 'Request ID',
      dataIndex: 'request_id',
      width: 150,
      ellipsis: true,
      render: (_, record) => (
        <Space size={4}>
          <Tooltip title={record.request_id}>
            <span style={{ fontSize: '12px' }}>{record.request_id.slice(0, 8)}...</span>
          </Tooltip>
          <Button
            type="text"
            size="small"
            icon={<CopyOutlined />}
            onClick={() => copyToClipboard(record.request_id, 'Request ID 已复制')}
          />
        </Space>
      ),
    },
    {
      title: '操作',
      valueType: 'option',
      width: 80,
      fixed: 'right',
      render: (_: unknown, record: ITokenUsageLog) => (
        <Button
          type="link"
          size="small"
          icon={<EyeOutlined />}
          onClick={() => onViewDetail(record)}
        >
          详情
        </Button>
      ),
    },
  ];
};
