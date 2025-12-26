/**
 * Token 管理页面
 * 功能：Token 的增删改查
 */
import { Button, Card, Input, Space } from 'antd';
import { PlusOutlined, SearchOutlined, UndoOutlined } from '@ant-design/icons';
import React, { useEffect, useState } from 'react';
import { useToken } from './hooks/useToken';
import TokenTable from './components/TokenTable';
import TokenForm from './components/TokenForm';
import RecycleModal from './components/RecycleModal';
import type { IToken, ITokenFormData } from '@/types';
import './index.less';

/**
 * Token 管理页面组件
 */
const TokenPage: React.FC = () => {
  const {
    dataSource,
    total,
    loading,
    pagination,
    loadData,
    handleCreate,
    handleUpdate,
    handleDelete,
  } = useToken();

  const [modalVisible, setModalVisible] = useState(false);
  const [editingRecord, setEditingRecord] = useState<IToken | null>(null);
  const [keyword, setKeyword] = useState('');
  const [recycleVisible, setRecycleVisible] = useState(false);

  /**
   * 初始化加载数据
   */
  useEffect(() => {
    loadData(1, 10);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  /**
   * 打开创建弹窗
   */
  const handleOpenCreateModal = () => {
    setEditingRecord(null);
    setModalVisible(true);
  };

  /**
   * 打开编辑弹窗
   */
  const handleOpenEditModal = (record: IToken) => {
    setEditingRecord(record);
    setModalVisible(true);
  };

  /**
   * 关闭弹窗
   */
  const handleCloseModal = () => {
    setModalVisible(false);
    setEditingRecord(null);
  };

  /**
   * 处理表单提交
   */
  const handleFormSubmit = async (data: ITokenFormData) => {
    if (editingRecord) {
      await handleUpdate(editingRecord.id, data);
    } else {
      await handleCreate(data);
    }
    handleCloseModal();
  };

  /**
   * 处理表格分页变化
   */
  const handleTablePageChange = (page: number, pageSize: number) => {
    loadData(page, pageSize, keyword);
  };

  /**
   * 处理搜索
   */
  const handleSearch = () => {
    loadData(1, pagination.pageSize, keyword);
  };

  /**
   * 刷新主列表
   */
  const handleRefresh = () => {
    loadData(pagination.current, pagination.pageSize, keyword);
  };

  return (
    <div className="token-page">
      <Card>
        <div className="token-page__header">
          <div className="token-page__actions">
            <Input
              placeholder="搜索 Token 或备注"
              value={keyword}
              onChange={(e) => setKeyword(e.target.value)}
              onPressEnter={handleSearch}
              style={{ width: 300 }}
              allowClear
            />
            <Button
              type="primary"
              icon={<SearchOutlined />}
              onClick={handleSearch}
            >
              搜索
            </Button>
          </div>
          <Space>
            <Button icon={<UndoOutlined />} onClick={() => setRecycleVisible(true)}>
              回收站
            </Button>
            <Button type="primary" icon={<PlusOutlined />} onClick={handleOpenCreateModal}>
              新建 Token
            </Button>
          </Space>
        </div>

        <TokenTable
          dataSource={dataSource}
          loading={loading}
          pagination={{
            current: pagination.current,
            pageSize: pagination.pageSize,
            total: total,
          }}
          onPageChange={handleTablePageChange}
          onEdit={handleOpenEditModal}
          onDelete={handleDelete}
        />

        <TokenForm
          visible={modalVisible}
          editingRecord={editingRecord}
          onSubmit={handleFormSubmit}
          onCancel={handleCloseModal}
        />

        <RecycleModal
          visible={recycleVisible}
          onCancel={() => setRecycleVisible(false)}
          onRefresh={handleRefresh}
        />
      </Card>
    </div>
  );
};

export default TokenPage;
