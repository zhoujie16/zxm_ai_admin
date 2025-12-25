/**
 * 代理服务管理页面
 * 功能：代理服务的增删改查
 */
import { Button, Card } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import React, { useEffect, useState } from 'react';
import { useProxyService } from './hooks/useProxyService';
import ProxyServiceTable from './components/ProxyServiceTable';
import ProxyServiceForm from './components/ProxyServiceForm';
import type { IProxyService, IProxyServiceFormData } from '@/types';
import './index.less';

/**
 * 代理服务管理页面组件
 */
const ProxyServicePage: React.FC = () => {
  const {
    dataSource,
    total,
    loading,
    pagination,
    loadData,
    handleCreate,
    handleUpdate,
    handleDelete,
  } = useProxyService();

  const [modalVisible, setModalVisible] = useState(false);
  const [editingRecord, setEditingRecord] = useState<IProxyService | null>(null);

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
  const handleOpenEditModal = (record: IProxyService) => {
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
  const handleFormSubmit = async (data: IProxyServiceFormData) => {
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
    loadData(page, pageSize);
  };

  return (
    <div className="proxy-service-page">
      <Card>
        <div className="proxy-service-page__header">
          <Button type="primary" icon={<PlusOutlined />} onClick={handleOpenCreateModal}>
            新建代理服务
          </Button>
        </div>

        <ProxyServiceTable
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

        <ProxyServiceForm
          visible={modalVisible}
          editingRecord={editingRecord}
          onSubmit={handleFormSubmit}
          onCancel={handleCloseModal}
        />
      </Card>
    </div>
  );
};

export default ProxyServicePage;
