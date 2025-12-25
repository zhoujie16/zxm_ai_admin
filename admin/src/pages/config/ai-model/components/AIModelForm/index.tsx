/**
 * AI 模型表单组件
 * 功能：创建和编辑 AI 模型的表单
 */
import { Form, Input, Modal, Switch } from 'antd';
import React, { useEffect } from 'react';
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
 * AI 模型表单组件
 */
const AIModelForm: React.FC<IAIModelFormProps> = ({
  visible,
  editingRecord,
  onSubmit,
  onCancel,
}) => {
  const [form] = Form.useForm();

  /**
   * 当编辑记录变化时，更新表单值
   */
  useEffect(() => {
    if (visible) {
      if (editingRecord) {
        form.setFieldsValue({
          model_key: editingRecord.model_key,
          model_name: editingRecord.model_name,
          api_url: editingRecord.api_url,
          status: editingRecord.status === 1,
          remark: editingRecord.remark,
        });
      } else {
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
        model_key: values.model_key,
        model_name: values.model_name,
        api_url: values.api_url,
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
      title={editingRecord ? '编辑 AI 模型' : '新建 AI 模型'}
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
          name="model_key"
          label="模型Key"
          rules={[
            { required: true, message: '请输入模型Key' },
            { max: 100, message: '模型Key最长100个字符' },
          ]}
        >
          <Input placeholder="请输入模型Key，如：glm-4" />
        </Form.Item>

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
          name="api_url"
          label="API地址"
          rules={[
            { required: true, message: '请输入API地址' },
            { type: 'url', message: '请输入有效的URL地址' },
            { max: 500, message: 'API地址最长500个字符' },
          ]}
        >
          <Input placeholder="请输入API地址，如：https://open.bigmodel.cn/api/paas/v4/chat/completions" />
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
