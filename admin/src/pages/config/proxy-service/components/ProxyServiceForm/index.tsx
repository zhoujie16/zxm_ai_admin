/**
 * 代理服务表单组件
 * 功能：创建和编辑代理服务的表单
 */
import { Form, Input, Modal, Switch } from 'antd';
import React, { useEffect } from 'react';
import type { IProxyService, IProxyServiceFormData } from '@/types';

/**
 * 表单组件 Props
 */
export interface IProxyServiceFormProps {
  /** 是否显示 */
  visible: boolean;
  /** 编辑的记录 */
  editingRecord: IProxyService | null;
  /** 提交回调 */
  onSubmit: (data: IProxyServiceFormData) => Promise<void>;
  /** 取消回调 */
  onCancel: () => void;
}

/**
 * 代理服务表单组件
 */
const ProxyServiceForm: React.FC<IProxyServiceFormProps> = ({
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
          service_id: editingRecord.service_id,
          server_ip: editingRecord.server_ip,
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
      const formData: IProxyServiceFormData = {
        service_id: values.service_id,
        server_ip: values.server_ip,
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
      title={editingRecord ? '编辑代理服务' : '新建代理服务'}
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
          name="service_id"
          label="服务标识"
          rules={[
            { required: true, message: '请输入服务标识' },
            { max: 100, message: '服务标识最长100个字符' },
          ]}
        >
          <Input placeholder="请输入服务标识" disabled={!!editingRecord} />
        </Form.Item>

        <Form.Item
          name="server_ip"
          label="服务器IP"
          rules={[
            { required: true, message: '请输入服务器IP地址' },
            {
              pattern: /^(\d{1,3}\.){3}\d{1,3}$/,
              message: '请输入有效的IP地址',
            },
          ]}
        >
          <Input placeholder="请输入服务器IP地址，如：192.168.1.100" />
        </Form.Item>

        <Form.Item name="status" label="状态" valuePropName="checked">
          <Switch checkedChildren="启用" unCheckedChildren="未启用" />
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

export default ProxyServiceForm;

