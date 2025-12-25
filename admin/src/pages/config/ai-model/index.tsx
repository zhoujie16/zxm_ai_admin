/**
 * AI 模型管理页面
 * 功能：AI 模型的增删改查
 */
import { Button, Card } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import React, { useEffect, useState } from 'react';
import { useAIModel } from './hooks/useAIModel';
import AIModelTable from './components/AIModelTable';
import AIModelForm from './components/AIModelForm';
import type { IAIModel, IAIModelFormData } from '@/types';
import './index.less';

/**
 * AI 模型管理页面组件
 */
const AIModelPage: React.FC = () => {
  const {
    dataSource,
    total,
    loading,
    pagination,
    loadData,
    handleCreate,
    handleUpdate,
    handleDelete,
  } = useAIModel();

  const [modalVisible, setModalVisible] = useState(false);
  const [editingRecord, setEditingRecord] = useState<IAIModel | null>(null);

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
  const handleOpenEditModal = (record: IAIModel) => {
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
  const handleFormSubmit = async (data: IAIModelFormData) => {
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
    <div className="ai-model-page">
      <Card>
        <div className="ai-model-page__header">
          <Button type="primary" icon={<PlusOutlined />} onClick={handleOpenCreateModal}>
            新建 AI 模型
          </Button>
        </div>

        <AIModelTable
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

        <AIModelForm
          visible={modalVisible}
          editingRecord={editingRecord}
          onSubmit={handleFormSubmit}
          onCancel={handleCloseModal}
        />
      </Card>
    </div>
  );
};

export default AIModelPage;
