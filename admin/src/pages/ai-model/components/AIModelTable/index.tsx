/**
 * AI 模型表格组件
 * 功能：展示 AI 模型列表
 */
import { Button, Popconfirm, Space, Table, Tag } from 'antd';
import { ColumnsType } from 'antd/es/table';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import React, { useMemo } from 'react';
import type { IAIModel } from '@/types';

/**
 * 表格组件 Props
 */
export interface IAIModelTableProps {
  /** 数据源 */
  dataSource: IAIModel[];
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
  onEdit: (record: IAIModel) => void;
  /** 删除回调 */
  onDelete: (id: number) => void;
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
 * AI 模型表格组件
 */
const AIModelTable: React.FC<IAIModelTableProps> = ({
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
  const columns: ColumnsType<IAIModel> = useMemo(
    () => [
      {
        title: 'ID',
        dataIndex: 'id',
        key: 'id',
        width: 80,
      },
      {
        title: '模型名称',
        dataIndex: 'model_name',
        key: 'model_name',
        width: 150,
      },
      {
        title: 'API地址',
        dataIndex: 'api_url',
        key: 'api_url',
        width: 300,
        ellipsis: true,
      },
      {
        title: 'API Key',
        dataIndex: 'api_key',
        key: 'api_key',
        width: 200,
        ellipsis: true,
      },
      {
        title: '状态',
        dataIndex: 'status',
        key: 'status',
        width: 100,
        render: (status: number) => (
          <Tag color={status === 1 ? 'success' : 'default'}>
            {status === 1 ? '启用' : '禁用'}
          </Tag>
        ),
      },
      {
        title: '备注',
        dataIndex: 'remark',
        key: 'remark',
        ellipsis: true,
      },
      {
        title: '创建时间',
        dataIndex: 'created_at',
        key: 'created_at',
        width: 180,
        render: formatDateTime,
      },
      {
        title: '更新时间',
        dataIndex: 'updated_at',
        key: 'updated_at',
        width: 180,
        render: formatDateTime,
      },
      {
        title: '操作',
        key: 'action',
        width: 150,
        fixed: 'right',
        render: (_: unknown, record: IAIModel) => (
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
      scroll={{ x: 1400 }}
    />
  );
};

export default AIModelTable;
