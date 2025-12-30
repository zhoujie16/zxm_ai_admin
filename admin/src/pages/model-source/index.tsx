/**
 * 模型来源管理页面
 */
import { Button, Card, Space } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import React, { useEffect, useState } from 'react';
import { useModelSource } from './hooks/useModelSource';
import ModelSourceTable from './components/ModelSourceTable';
import ModelSourceForm from './components/ModelSourceForm';
import type { IModelSource } from '@/types';
import './index.less';

const ModelSourcePage: React.FC = () => {
  const {
    dataSource,
    total,
    loading,
    pagination,
    loadData,
    handleCreate,
    handleUpdate,
    handleDelete,
  } = useModelSource();

  const [modalVisible, setModalVisible] = useState(false);
  const [editingRecord, setEditingRecord] = useState<IModelSource | null>(null);

  useEffect(() => {
    loadData(1, 10);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleOpenCreateModal = () => {
    setEditingRecord(null);
    setModalVisible(true);
  };

  const handleOpenEditModal = (record: IModelSource) => {
    setEditingRecord(record);
    setModalVisible(true);
  };

  const handleCloseModal = () => {
    setModalVisible(false);
    setEditingRecord(null);
  };

  const handleFormSubmit = async (data: any) => {
    if (editingRecord) {
      await handleUpdate(editingRecord.id, data);
    } else {
      await handleCreate(data);
    }
    handleCloseModal();
  };

  const handleTablePageChange = (page: number, pageSize: number) => {
    loadData(page, pageSize);
  };

  return (
    <div className="model-source-page">
      <Card>
        <div className="model-source-page__header">
          <h3>模型来源管理</h3>
          <Space>
            <Button type="primary" icon={<PlusOutlined />} onClick={handleOpenCreateModal}>
              新建模型来源
            </Button>
          </Space>
        </div>

        <ModelSourceTable
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

        <ModelSourceForm
          visible={modalVisible}
          editingRecord={editingRecord}
          onSubmit={handleFormSubmit}
          onCancel={handleCloseModal}
        />
      </Card>
    </div>
  );
};

export default ModelSourcePage;
