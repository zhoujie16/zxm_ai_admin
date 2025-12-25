/**
 * Token 使用记录弹窗组件
 * 功能：展示 Token 的使用记录列表
 */
import { Table, Typography, Modal } from 'antd';
import { ColumnsType } from 'antd/es/table';
import React, { useEffect, useState } from 'react';
import { getTokenUsageLogs } from '@/services/tokenUsage';
import type { ITokenUsageLog } from '@/types';

const { Text } = Typography;

/**
 * 弹窗组件 Props
 */
export interface IUsageLogsModalProps {
  /** 是否显示 */
  visible: boolean;
  /** Token 值 */
  token: string;
  /** 关闭回调 */
  onCancel: () => void;
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
 * Token 使用记录弹窗组件
 */
const UsageLogsModal: React.FC<IUsageLogsModalProps> = ({
  visible,
  token,
  onCancel,
}) => {
  const [loading, setLoading] = useState(false);
  const [dataSource, setDataSource] = useState<ITokenUsageLog[]>([]);
  const [total, setTotal] = useState(0);
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
  });

  /**
   * 加载使用记录
   */
  const loadLogs = async (page: number = 1, pageSize: number = 10) => {
    setLoading(true);
    try {
      const result = await getTokenUsageLogs(token, page, pageSize);
      if (result.success && result.data) {
        setDataSource(result.data.list);
        setTotal(result.data.total);
        setPagination({
          current: page,
          pageSize: pageSize,
        });
      }
    } catch (error) {
      console.error('加载使用记录失败:', error);
    } finally {
      setLoading(false);
    }
  };

  /**
   * 弹窗打开时加载数据
   */
  useEffect(() => {
    if (visible && token) {
      loadLogs(1, 10);
    }
  }, [visible, token]);

  /**
   * 表格列定义
   */
  const columns: ColumnsType<ITokenUsageLog> = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      width: 80,
    },
    {
      title: '调用时间',
      dataIndex: 'call_time',
      key: 'call_time',
      width: 180,
      render: formatDateTime,
    },
    {
      title: '来源IP',
      dataIndex: 'remote_ip',
      key: 'remote_ip',
      width: 150,
      ellipsis: true,
    },
    {
      title: 'User Agent',
      dataIndex: 'user_agent',
      key: 'user_agent',
      ellipsis: true,
    },
  ];

  return (
    <Modal
      title="Token 使用记录"
      open={visible}
      onCancel={onCancel}
      footer={null}
      width={800}
    >
      <Table
        columns={columns}
        dataSource={dataSource}
        rowKey="id"
        loading={loading}
        pagination={{
          current: pagination.current,
          pageSize: pagination.pageSize,
          total: total,
          showSizeChanger: true,
          showTotal: (total) => `共 ${total} 条`,
          onChange: (page, pageSize) => loadLogs(page, pageSize),
        }}
        scroll={{ x: 600 }}
      />
    </Modal>
  );
};

export default UsageLogsModal;
