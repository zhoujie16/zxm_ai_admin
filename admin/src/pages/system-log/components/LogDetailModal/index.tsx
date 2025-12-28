/**
 * 系统日志详情弹窗组件
 * 功能：展示系统日志详细信息
 */
import { Modal, Descriptions, Tag } from 'antd';
import React, { useMemo } from 'react';
import type { ISystemLog } from '@/types/systemLog';

/**
 * 详情弹窗组件 Props
 */
export interface ILogDetailModalProps {
  /** 是否显示 */
  visible: boolean;
  /** 记录数据 */
  record: ISystemLog | null;
  /** 关闭回调 */
  onCancel: () => void;
}

/**
 * 格式化时间
 */
const formatDateTime = (text: string): string => {
  try {
    return new Date(text).toLocaleString('zh-CN');
  } catch {
    return text;
  }
};

/**
 * 获取日志级别标签颜色
 */
const getLevelColor = (level: string): string => {
  switch (level) {
    case 'DEBUG':
      return 'default';
    case 'INFO':
      return 'processing';
    case 'WARN':
      return 'warning';
    case 'ERROR':
      return 'error';
    default:
      return 'default';
  }
};

/**
 * 系统日志详情弹窗组件
 */
const LogDetailModal: React.FC<ILogDetailModalProps> = ({
  visible,
  record,
  onCancel,
}) => {
  const levelColor = useMemo(() => {
    return record ? getLevelColor(record.level) : 'default';
  }, [record]);

  return (
    <Modal
      title="系统日志详情"
      open={visible}
      onCancel={onCancel}
      footer={null}
      width={700}
    >
      {record && (
        <Descriptions bordered column={1}>
          <Descriptions.Item label="ID">{record.id}</Descriptions.Item>
          <Descriptions.Item label="时间">{formatDateTime(record.time)}</Descriptions.Item>
          <Descriptions.Item label="级别">
            <Tag color={levelColor}>{record.level}</Tag>
          </Descriptions.Item>
          <Descriptions.Item label="消息">
            <pre style={{ margin: 0, whiteSpace: 'pre-wrap', wordBreak: 'break-word' }}>
              {record.msg}
            </pre>
          </Descriptions.Item>
          <Descriptions.Item label="创建时间">{formatDateTime(record.created_at)}</Descriptions.Item>
          <Descriptions.Item label="更新时间">{formatDateTime(record.updated_at)}</Descriptions.Item>
        </Descriptions>
      )}
    </Modal>
  );
};

export default LogDetailModal;
