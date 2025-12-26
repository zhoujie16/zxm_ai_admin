/**
 * Token 回收站弹窗组件
 * 功能：展示已删除的 Token 列表，支持恢复和永久删除
 */
import { Button, Input, Modal, Space, Table, Tag, Typography, message } from 'antd';
import { ColumnsType } from 'antd/es/table';
import { DeleteOutlined, ReloadOutlined, SearchOutlined } from '@ant-design/icons';
import React, { useEffect, useState } from 'react';
import type { IToken } from '@/types';
import { destroyToken, getRecycledTokenList, restoreToken } from '@/services/token';

const { Text } = Typography;

/**
 * 弹窗组件 Props
 */
export interface IRecycleModalProps {
  /** 弹窗是否可见 */
  visible: boolean;
  /** 关闭回调 */
  onCancel: () => void;
  /** 刷新主列表回调 */
  onRefresh?: () => void;
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
 * 回收站弹窗组件
 */
const RecycleModal: React.FC<IRecycleModalProps> = ({ visible, onCancel, onRefresh }) => {
  const [dataSource, setDataSource] = useState<IToken[]>([]);
  const [loading, setLoading] = useState(false);
  const [keyword, setKeyword] = useState('');
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
    total: 0,
  });

  /**
   * 加载数据
   */
  const loadData = async (page: number = 1, pageSize: number = 10, searchKeyword: string = '') => {
    setLoading(true);
    try {
      const res = await getRecycledTokenList(page, pageSize, searchKeyword);
      if (res.success && res.data) {
        setDataSource(res.data.list || []);
        setPagination({
          current: page,
          pageSize,
          total: res.data.total,
        });
      }
    } catch (error) {
      console.error('加载回收站数据失败', error);
    } finally {
      setLoading(false);
    }
  };

  /**
   * 弹窗打开时加载数据
   */
  useEffect(() => {
    if (visible) {
      loadData(1, 10, '');
      setKeyword('');
    }
  }, [visible]);

  /**
   * 处理恢复
   */
  const handleRestore = async (id: number) => {
    try {
      const res = await restoreToken(id);
      if (res.success) {
        message.success('恢复成功');
        loadData(pagination.current, pagination.pageSize, keyword);
        onRefresh?.();
      }
    } catch (error) {
      console.error('恢复失败', error);
    }
  };

  /**
   * 处理永久删除
   */
  const handleDestroy = async (id: number) => {
    Modal.confirm({
      title: '永久删除',
      content: '确定要永久删除这条记录吗？此操作不可恢复！',
      okText: '确定',
      cancelText: '取消',
      okButtonProps: { danger: true },
      onOk: async () => {
        try {
          const res = await destroyToken(id);
          if (res.success) {
            message.success('永久删除成功');
            loadData(pagination.current, pagination.pageSize, keyword);
            onRefresh?.();
          }
        } catch (error) {
          console.error('删除失败', error);
        }
      },
    });
  };

  /**
   * 处理搜索
   */
  const handleSearch = () => {
    loadData(1, pagination.pageSize, keyword);
  };

  /**
   * 表格列定义
   */
  const columns: ColumnsType<IToken> = [
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
        <Text code style={{ fontSize: 12 }}>
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
      width: 180,
      fixed: 'right',
      render: (_: unknown, record: IToken) => (
        <Space size="small">
          <Button
            type="link"
            size="small"
            icon={<ReloadOutlined />}
            onClick={() => handleRestore(record.id)}
          >
            恢复
          </Button>
          <Button
            type="link"
            size="small"
            danger
            icon={<DeleteOutlined />}
            onClick={() => handleDestroy(record.id)}
          >
            永久删除
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <Modal
      title="回收站"
      open={visible}
      onCancel={onCancel}
      footer={null}
      width={1400}
      destroyOnHidden
    >
      <div style={{ marginBottom: 16 }}>
        <Space>
          <Input
            placeholder="搜索 Token 或备注"
            value={keyword}
            onChange={(e) => setKeyword(e.target.value)}
            onPressEnter={handleSearch}
            style={{ width: 300 }}
            allowClear
          />
          <Button icon={<SearchOutlined />} onClick={handleSearch}>
            搜索
          </Button>
        </Space>
      </div>
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
          onChange: (page, pageSize) => {
            loadData(page, pageSize, keyword);
          },
        }}
        scroll={{ x: 1400 }}
      />
    </Modal>
  );
};

export default RecycleModal;
