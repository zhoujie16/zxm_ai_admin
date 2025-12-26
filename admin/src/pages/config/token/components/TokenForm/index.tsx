/**
 * Token 表单组件
 * 功能：创建和编辑 Token 的表单
 */
import { DatePicker, Descriptions, Form, Input, InputNumber, Modal, Select, Space, Switch, Typography } from 'antd';

const { Text } = Typography;

import dayjs from 'dayjs';
import React, { useEffect, useState } from 'react';
import { getAIModelList } from '@/services/aiModel';
import type { IToken, ITokenFormData, IAIModel } from '@/types';

/**
 * 表单组件 Props
 */
export interface ITokenFormProps {
  /** 是否显示 */
  visible: boolean;
  /** 编辑的记录 */
  editingRecord: IToken | null;
  /** 提交回调 */
  onSubmit: (data: ITokenFormData) => Promise<void>;
  /** 取消回调 */
  onCancel: () => void;
}

/**
 * Token 表单组件
 */
const TokenForm: React.FC<ITokenFormProps> = ({
  visible,
  editingRecord,
  onSubmit,
  onCancel,
}) => {
  const [form] = Form.useForm();
  const [models, setModels] = useState<IAIModel[]>([]);
  const [loadingModels, setLoadingModels] = useState(false);

  /**
   * 加载 AI 模型列表
   */
  const loadModels = async () => {
    setLoadingModels(true);
    try {
      const result = await getAIModelList(1, 1000);
      if (result.success && result.data) {
        // 只显示启用的模型
        setModels(result.data.list.filter((m) => m.status === 1));
      }
    } catch (error) {
      console.error('加载模型列表失败:', error);
    } finally {
      setLoadingModels(false);
    }
  };

  /**
   * 当编辑记录变化时，更新表单值
   */
  useEffect(() => {
    if (visible) {
      loadModels();
      if (editingRecord) {
        form.setFieldsValue({
          ai_model_id: editingRecord.ai_model_id,
          order_no: editingRecord.order_no,
          status: editingRecord.status === 1,
          usage_limit: editingRecord.usage_limit,
          expire_at: editingRecord.expire_at ? dayjs(editingRecord.expire_at) : undefined,
          remark: editingRecord.remark,
        });
      } else {
        form.resetFields();
        form.setFieldsValue({
          status: true,
          usage_limit: 0,
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
      const formData: ITokenFormData = {
        ai_model_id: values.ai_model_id,
        order_no: values.order_no,
        status: values.status ? 1 : 0,
        usage_limit: values.usage_limit || 0,
        expire_at: values.expire_at ? values.expire_at.toISOString() : undefined,
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
      title={editingRecord ? '编辑 Token' : '新建 Token'}
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
              label: 'Token',
              children: (
                <Text code copyable style={{ fontSize: 12 }}>
                  {editingRecord.token}
                </Text>
              ),
            },
          ]}
        />
      )}

      <Form
        form={form}
        layout="vertical"
        initialValues={{
          status: true,
          usage_limit: 0,
        }}
      >
        <Form.Item
          name="ai_model_id"
          label="关联模型"
          rules={[{ required: true, message: '请选择关联模型' }]}
        >
          <Select
            placeholder="请选择关联的 AI 模型"
            loading={loadingModels}
            showSearch
            optionFilterProp="label"
            options={models.map((m) => ({
              label: m.model_name,
              value: m.id,
            }))}
          />
        </Form.Item>

        <Form.Item
          name="order_no"
          label="关联订单号"
          rules={[{ max: 100, message: '订单号最长100个字符' }]}
        >
          <Input placeholder="请输入关联订单号" />
        </Form.Item>

        <Form.Item name="status" label="状态" valuePropName="checked">
          <Switch checkedChildren="启用" unCheckedChildren="禁用" />
        </Form.Item>

        <Form.Item
          name="usage_limit"
          label="使用限额"
          rules={[{ required: true, message: '请输入使用限额' }]}
          tooltip="设置为 0 表示无限制"
        >
          <Space.Compact style={{ width: '100%' }}>
            <InputNumber
              placeholder="请输入使用限额"
              min={0}
              style={{ width: '100%' }}
            />
            <Input style={{ width: 50 }} readOnly value="次" />
          </Space.Compact>
        </Form.Item>

        <Form.Item
          name="expire_at"
          label="过期时间"
          tooltip="留空表示永不过期"
        >
          <DatePicker
            showTime
            placeholder="请选择过期时间"
            style={{ width: '100%' }}
            format="YYYY-MM-DD HH:mm:ss"
            disabledDate={(current) => current && current < dayjs().startOf('day')}
          />
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

export default TokenForm;
