/**
 * 系统日志表格组件
 * 功能：展示系统日志列表
 */
import { Button, Table, Tag } from 'antd';
import { ColumnsType } from 'antd/es/table';
import { EyeOutlined } from '@ant-design/icons';
import React, { useMemo } from 'react';
import type { ISystemLog } from '@/types/systemLog';

/**
 * 表格组件 Props
 */
export interface ISystemLogTableProps {
  /** 数据源 */
  dataSource: ISystemLog[];
  /** 加载状态 */
  loading: boolean;
  /** 分页信息 */
  pagination: { current: number; pageSize: number; total: number };
  /** 分页变化回调 */
  onPageChange: (page: number, pageSize: number) => void;
  /** 查看详情回调 */
  onViewDetail: (record: ISystemLog) => void;
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
 * 获取日志级别标签颜色
 */
const getLevelColor = (level: string): string => {
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

/**
 * 系统日志表格组件
 */
const SystemLogTable: React.FC<ISystemLogTableProps> = ({
  dataSource,
  loading,
  pagination,
  onPageChange,
  onViewDetail,
}) => {
  /**
   * 表格列定义
   */
  const columns: ColumnsType<ISystemLog> = useMemo(
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
        title: '级别',
        dataIndex: 'level',
        key: 'level',
        width: 100,
        render: (level: string) => (
          <Tag color={getLevelColor(level)}>{level}</Tag>
        ),
      },
      {
        title: '消息',
        dataIndex: 'msg',
        key: 'msg',
        ellipsis: true,
      },
      {
        title: '操作',
        key: 'action',
        width: 100,
        fixed: 'right',
        render: (_: unknown, record: ISystemLog) => (
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
      scroll={{ x: 800 }}
      className="system-log-table"
    />
  );
};

export default SystemLogTable;
