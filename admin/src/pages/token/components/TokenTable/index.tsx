/**
 * Token 表格组件
 * 功能：展示 Token 列表
 */
import { Button, Popconfirm, Space, Table, Tag, Typography } from 'antd';
import { ColumnsType } from 'antd/es/table';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import React, { useMemo } from 'react';
import type { IToken } from '@/types';

const { Text } = Typography;

/**
 * 表格组件 Props
 */
export interface ITokenTableProps {
  /** 数据源 */
  dataSource: IToken[];
  /** 加载状态 */
  loading: boolean;
  /** 分页信息 */
  pagination: {
    current: number;
    pageSize: number;
    total: number;
  };
  /** 分页变化回调 */
  onPageChange: (page: number, pageSize: number) => void;
  /** 编辑回调 */
  onEdit: (record: IToken) => void;
  /** 删除回调 */
  onDelete: (id: number) => void;
}

/**
 * 格式化时间
 */
const formatDateTime = (text: string | undefined): string => {
  if (!text) return '-';
  try {
    return new Date(text).toLocaleString('zh-CN');
  } catch {
    return text;
  }
};

/**
 * Token 表格组件
 */
const TokenTable: React.FC<ITokenTableProps> = ({
  dataSource,
  loading,
  pagination,
  onPageChange,
  onEdit,
  onDelete,
}) => {
  /**
   * 表格列定义
   */
  const columns: ColumnsType<IToken> = useMemo(
    () => [
      {
        title: 'ID',
        dataIndex: 'id',
        key: 'id',
        width: 80,
      },
      {
        title: 'Token',
        dataIndex: 'token',
        key: 'token',
        width: 280,
        render: (text: string) => (
          <Text code copyable style={{ fontSize: 12 }}>
            {text}
          </Text>
        ),
      },
      {
        title: '关联模型',
        dataIndex: 'model_name',
        key: 'model_name',
        width: 150,
        ellipsis: true,
        render: (text: string) => text || '-',
      },
      {
        title: '关联订单号',
        dataIndex: 'order_no',
        key: 'order_no',
        width: 130,
        ellipsis: true,
        render: (text: string) => text || '-',
      },
      {
        title: '状态',
        dataIndex: 'status',
        key: 'status',
        width: 90,
        render: (status: number) => (
          <Tag color={status === 1 ? 'success' : 'default'}>
            {status === 1 ? '启用' : '禁用'}
          </Tag>
        ),
      },
      {
        title: '使用限额',
        dataIndex: 'usage_limit',
        key: 'usage_limit',
        width: 100,
        render: (limit: number) => (limit === 0 ? <Tag color="default">无限制</Tag> : limit),
      },
      {
        title: '过期时间',
        dataIndex: 'expire_at',
        key: 'expire_at',
        width: 170,
        render: formatDateTime,
      },
      {
        title: '备注',
        dataIndex: 'remark',
        key: 'remark',
        width: 150,
        ellipsis: true,
      },
      {
        title: '创建时间',
        dataIndex: 'created_at',
        key: 'created_at',
        width: 170,
        render: formatDateTime,
      },
      {
        title: '操作',
        key: 'action',
        width: 150,
        fixed: 'right',
        render: (_: unknown, record: IToken) => (
          <Space size="small">
            <Button
              type="link"
              size="small"
              icon={<EditOutlined />}
              onClick={() => onEdit(record)}
            >
              编辑
            </Button>
            <Popconfirm
              title="确定要删除这条记录吗？"
              onConfirm={() => onDelete(record.id)}
              okText="确定"
              cancelText="取消"
            >
              <Button type="link" size="small" danger icon={<DeleteOutlined />}>
                删除
              </Button>
            </Popconfirm>
          </Space>
        ),
      },
    ],
    [onEdit, onDelete],
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
      scroll={{ x: 1500 }}
    />
  );
};

export default TokenTable;
