/**
 * 模型代理表单组件
 * 功能：创建和编辑模型代理的表单
 */
import { Form, Input, Modal, Select, Switch } from 'antd';
import React, { useEffect, useState } from 'react';
import { getModelSourceList } from '@/services/modelSource';
import type { IModelSource } from '@/types';
import type { IAIModel, IAIModelFormData } from '@/types';

/**
 * 表单组件 Props
 */
export interface IAIModelFormProps {
  /** 是否显示 */
  visible: boolean;
  /** 编辑的记录 */
  editingRecord: IAIModel | null;
  /** 提交回调 */
  onSubmit: (data: IAIModelFormData) => Promise<void>;
  /** 取消回调 */
  onCancel: () => void;
}

/**
 * 模型代理表单组件
 */
const AIModelForm: React.FC<IAIModelFormProps> = ({
  visible,
  editingRecord,
  onSubmit,
  onCancel,
}) => {
  const [form] = Form.useForm();
  const [modelSources, setModelSources] = useState<IModelSource[]>([]);
  const [loadingSources, setLoadingSources] = useState(false);

  /**
   * 加载模型来源列表
   */
  const loadModelSources = async () => {
    setLoadingSources(true);
    try {
      const result = await getModelSourceList(1, 1000);
      if (result.success && result.data) {
        setModelSources(result.data.list);
      }
    } catch (error) {
      console.error('加载模型来源失败:', error);
    } finally {
      setLoadingSources(false);
    }
  };

  /**
   * 当编辑记录变化时，更新表单值
   */
  useEffect(() => {
    if (visible) {
      // 加载模型来源列表
      loadModelSources();
      if (editingRecord) {
        // 编辑时设置当前值
        form.setFieldsValue({
          model_name: editingRecord.model_name,
          api_key: editingRecord.api_key,
          status: editingRecord.status === 1,
          remark: editingRecord.remark,
        });
      } else {
        // 新建时重置表单
        form.resetFields();
        form.setFieldsValue({
          status: true,
        });
      }
    }
  }, [visible, editingRecord, form]);

  /**
   * 处理提交
   */
  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();
      const formData: IAIModelFormData = {
        model_name: values.model_name,
        api_key: values.api_key,
        status: values.status ? 1 : 0,
        remark: values.remark,
      };
      await onSubmit(formData);
    } catch (error) {
      // 表单验证错误，不处理
      if ((error as { errorFields?: unknown[] }).errorFields) {
        return;
      }
      console.error('提交失败:', error);
    }
  };

  return (
    <Modal
      title={editingRecord ? '编辑模型代理' : '新建模型代理'}
      open={visible}
      onOk={handleSubmit}
      onCancel={onCancel}
      okText="确定"
      cancelText="取消"
      width={600}
    >
      <Form
        form={form}
        layout="vertical"
        initialValues={{
          status: true,
        }}
      >
        <Form.Item
          name="model_name"
          label="模型名称"
          rules={[
            { required: true, message: '请输入模型名称' },
            { max: 100, message: '模型名称最长100个字符' },
          ]}
        >
          <Input placeholder="请输入模型名称，如：智谱 GLM-4" />
        </Form.Item>

        <Form.Item
          name="api_key"
          label="模型来源"
          rules={[{ required: true, message: '请选择模型来源' }]}
        >
          <Select
            placeholder="请选择模型来源"
            loading={loadingSources}
            showSearch
            optionFilterProp="label"
            options={modelSources.map((item) => ({
              label: `${item.model_name} (${item.api_url})`,
              value: item.api_key,
            }))}
          />
        </Form.Item>

        <Form.Item name="status" label="状态" valuePropName="checked">
          <Switch checkedChildren="启用" unCheckedChildren="禁用" />
        </Form.Item>

        <Form.Item
          name="remark"
          label="备注"
          rules={[{ max: 500, message: '备注最长500个字符' }]}
        >
          <Input.TextArea
            rows={4}
            placeholder="请输入备注信息"
            showCount
            maxLength={500}
          />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default AIModelForm;
