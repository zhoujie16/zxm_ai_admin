/**
 * 模型来源表格组件
 */
import { Button, Popconfirm, Space, Table, Typography } from 'antd';
import { ColumnsType } from 'antd/es/table';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import React, { useMemo } from 'react';
import type { IModelSource } from '@/types';

const { Text } = Typography;

export interface IModelSourceTableProps {
  dataSource: IModelSource[];
  loading: boolean;
  pagination: {
    current: number;
    pageSize: number;
    total: number;
  };
  onPageChange: (page: number, pageSize: number) => void;
  onEdit: (record: IModelSource) => void;
  onDelete: (id: number) => void;
}

const formatDateTime = (text: string | undefined): string => {
  if (!text) return '-';
  try {
    return new Date(text).toLocaleString('zh-CN');
  } catch {
    return text;
  }
};

const ModelSourceTable: React.FC<IModelSourceTableProps> = ({
  dataSource,
  loading,
  pagination,
  onPageChange,
  onEdit,
  onDelete,
}) => {
  const columns: ColumnsType<IModelSource> = useMemo(
    () => [
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
        render: (text: string) => (
          <Text ellipsis style={{ maxWidth: 280 }}>
            {text}
          </Text>
        ),
      },
      {
        title: 'API Key',
        dataIndex: 'api_key',
        key: 'api_key',
        width: 280,
        render: (text: string) => (
          <Text code copyable style={{ fontSize: 12 }}>
            {text}
          </Text>
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
        width: 170,
        render: formatDateTime,
      },
      {
        title: '操作',
        key: 'action',
        width: 150,
        fixed: 'right',
        render: (_: unknown, record: IModelSource) => (
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
      scroll={{ x: 1200 }}
    />
  );
};

export default ModelSourceTable;
