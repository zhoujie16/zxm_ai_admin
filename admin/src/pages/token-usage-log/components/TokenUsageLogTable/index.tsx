/**
 * Token 使用记录表格组件
 * 功能：展示 Token 使用记录列表
 */
import { Button, Space, Table, Tag, Badge, Tooltip } from 'antd';
import { ColumnsType } from 'antd/es/table';
import { EyeOutlined, CopyOutlined } from '@ant-design/icons';
import React, { useMemo } from 'react';
import type { ITokenUsageLog, IPaginationInfo } from '@/types/tokenUsageLog';
import { message } from 'antd';

/**
 * 表格组件 Props
 */
export interface ITokenUsageLogTableProps {
  /** 数据源 */
  dataSource: ITokenUsageLog[];
  /** 加载状态 */
  loading: boolean;
  /** 分页信息 */
  pagination: IPaginationInfo & { total: number };
  /** 分页变化回调 */
  onPageChange: (page: number, pageSize: number) => void;
  /** 查看详情回调 */
  onViewDetail: (record: ITokenUsageLog) => void;
}

/**
 * 格式化时间
 */
const formatDateTime = (text: string): string => {
  try {
    return new Date(text).toLocaleString('zh-CN');
  } catch {
    return text;
  }
};

/**
 * 格式化耗时
 */
const formatLatency = (ms: number): string => {
  if (ms < 1000) {
    return `${ms}ms`;
  }
  return `${(ms / 1000).toFixed(2)}s`;
};

/**
 * 格式化字节大小
 */
const formatBytes = (bytes: number): string => {
  if (bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${(bytes / Math.pow(k, i)).toFixed(2)} ${sizes[i]}`;
};

/**
 * 复制文本到剪贴板
 */
const copyToClipboard = (text: string, successMsg: string = '复制成功') => {
  navigator.clipboard.writeText(text).then(() => {
    message.success(successMsg);
  });
};

/**
 * Token 使用记录表格组件
 */
const TokenUsageLogTable: React.FC<ITokenUsageLogTableProps> = ({
  dataSource,
  loading,
  pagination,
  onPageChange,
  onViewDetail,
}) => {
  /**
   * 获取状态标签颜色
   */
  const getStatusColor = (status: number): string => {
    if (status >= 200 && status < 300) return 'success';
    if (status >= 300 && status < 400) return 'warning';
    if (status >= 400 && status < 500) return 'error';
    return 'default';
  };

  /**
   * 获取耗时等级
   */
  const getLatencyLevel = (ms: number): 'fast' | 'normal' | 'slow' => {
    if (ms < 500) return 'fast';
    if (ms < 2000) return 'normal';
    return 'slow';
  };

  /**
   * 表格列定义
   */
  const columns: ColumnsType<ITokenUsageLog> = useMemo(
    () => [
      {
        title: 'ID',
        dataIndex: 'id',
        key: 'id',
        width: 80,
      },
      {
        title: '时间',
        dataIndex: 'time',
        key: 'time',
        width: 180,
        render: formatDateTime,
      },
      {
        title: '方法',
        dataIndex: 'method',
        key: 'method',
        width: 80,
        render: (method: string) => (
          <Tag className={`method-tag ${method}`}>{method}</Tag>
        ),
      },
      {
        title: '路径',
        dataIndex: 'path',
        key: 'path',
        width: 200,
        ellipsis: true,
        render: (path: string) => (
          <Tooltip title={path}>
            <span>{path}</span>
          </Tooltip>
        ),
      },
      {
        title: '状态码',
        dataIndex: 'status',
        key: 'status',
        width: 100,
        render: (status: number) => (
          <Tag color={getStatusColor(status)}>{status}</Tag>
        ),
      },
      {
        title: '耗时',
        dataIndex: 'latency_ms',
        key: 'latency_ms',
        width: 100,
        render: (ms: number) => (
          <Badge
            status={getLatencyLevel(ms) === 'fast' ? 'success' : getLatencyLevel(ms) === 'normal' ? 'warning' : 'error'}
            text={formatLatency(ms)}
          />
        ),
      },
      {
        title: '请求大小',
        dataIndex: 'request_size_bytes',
        key: 'request_size_bytes',
        width: 100,
        render: formatBytes,
      },
      {
        title: '响应大小',
        dataIndex: 'response_size_bytes',
        key: 'response_size_bytes',
        width: 100,
        render: formatBytes,
      },
      {
        title: '客户端',
        dataIndex: 'remote_addr',
        key: 'remote_addr',
        width: 150,
        ellipsis: true,
        render: (addr: string) => (
          <Tooltip title={addr}>
            <span>{addr}</span>
          </Tooltip>
        ),
      },
      {
        title: 'Request ID',
        dataIndex: 'request_id',
        key: 'request_id',
        width: 150,
        ellipsis: true,
        render: (requestId: string) => (
          <Space size={4}>
            <Tooltip title={requestId}>
              <span style={{ fontSize: '12px' }}>
                {requestId.slice(0, 8)}...
              </span>
            </Tooltip>
            <Button
              type="text"
              size="small"
              icon={<CopyOutlined />}
              onClick={() => copyToClipboard(requestId, 'Request ID 已复制')}
            />
          </Space>
        ),
      },
      {
        title: '操作',
        key: 'action',
        width: 100,
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
    ],
    [onViewDetail],
  );

  return (
    <Table
      columns={columns}
      dataSource={dataSource}
      rowKey="id"
      loading={loading}
      pagination={{
        current: pagination.current,
        pageSize: pagination.pageSize,
        total: pagination.total,
        showSizeChanger: true,
        showTotal: (total) => `共 ${total} 条`,
        onChange: onPageChange,
        onShowSizeChange: onPageChange,
      }}
      scroll={{ x: 1600 }}
      className="token-usage-log-table"
    />
  );
};

export default TokenUsageLogTable;
