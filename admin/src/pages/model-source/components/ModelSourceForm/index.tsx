/**
 * 模型来源表单组件
 */
import { Descriptions, Form, Input, Modal, Typography } from 'antd';
import React from 'react';
import type {
  IModelSource,
  ICreateModelSourceFormData,
  IUpdateModelSourceFormData,
} from '@/types';

const { Text } = Typography;

export interface IModelSourceFormProps {
  visible: boolean;
  editingRecord: IModelSource | null;
  onSubmit: (data: ICreateModelSourceFormData | IUpdateModelSourceFormData) => Promise<void>;
  onCancel: () => void;
}

const ModelSourceForm: React.FC<IModelSourceFormProps> = ({
  visible,
  editingRecord,
  onSubmit,
  onCancel,
}) => {
  const [form] = Form.useForm();

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();
      await onSubmit(values);
    } catch (error) {
      if ((error as { errorFields?: unknown[] }).errorFields) {
        return;
      }
      console.error('提交失败:', error);
    }
  };

  React.useEffect(() => {
    if (visible) {
      if (editingRecord) {
        form.setFieldsValue({
          model_name: editingRecord.model_name,
          remark: editingRecord.remark,
        });
      } else {
        form.resetFields();
      }
    }
  }, [visible, editingRecord, form]);

  return (
    <Modal
      title={editingRecord ? '编辑模型来源' : '新建模型来源'}
      open={visible}
      onOk={handleSubmit}
      onCancel={onCancel}
      okText="确定"
      cancelText="取消"
      width={600}
    >
      {editingRecord && (
        <Descriptions
          column={1}
          bordered
          size="small"
          style={{ marginBottom: 16 }}
          items={[
            {
              label: 'API地址',
              children: <Text ellipsis style={{ maxWidth: 480 }}>{editingRecord.api_url}</Text>,
            },
            {
              label: 'API Key',
              children: (
                <Text code copyable style={{ fontSize: 12 }}>
                  {editingRecord.api_key}
                </Text>
              ),
            },
          ]}
        />
      )}

      <Form form={form} layout="vertical">
        {!editingRecord && (
          <>
            <Form.Item
              name="model_name"
              label="模型名称"
              rules={[{ required: true, message: '请输入模型名称' }]}
            >
              <Input placeholder="请输入模型名称" />
            </Form.Item>

            <Form.Item
              name="api_url"
              label="API地址"
              rules={[
                { required: true, message: '请输入API地址' },
                { type: 'url', message: '请输入有效的URL地址' },
              ]}
            >
              <Input placeholder="请输入API地址，如 https://api.example.com" />
            </Form.Item>

            <Form.Item
              name="api_key"
              label="API Key"
              rules={[{ required: true, message: '请输入API Key' }]}
            >
              <Input placeholder="请输入API Key" />
            </Form.Item>
          </>
        )}

        {editingRecord && (
          <Form.Item
            name="model_name"
            label="模型名称"
            rules={[{ required: true, message: '请输入模型名称' }]}
          >
            <Input placeholder="请输入模型名称" />
          </Form.Item>
        )}

        <Form.Item
          name="remark"
          label="备注"
          rules={[{ max: 500, message: '备注最长500个字符' }]}
        >
          <Input.TextArea rows={4} placeholder="请输入备注信息" showCount maxLength={500} />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default ModelSourceForm;
