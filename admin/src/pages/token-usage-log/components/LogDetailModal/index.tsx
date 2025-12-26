/**
 * Token 使用记录详情弹窗组件
 * 功能：展示单条日志的完整信息
 */
import { Modal, Descriptions, Tag, Badge, Tabs, Typography, Button, Space } from 'antd';
import { CopyOutlined } from '@ant-design/icons';
import React from 'react';
import type { ITokenUsageLog } from '@/types/tokenUsageLog';
import { message } from 'antd';

const { Text } = Typography;

/**
 * 详情弹窗组件 Props
 */
export interface ILogDetailModalProps {
  /** 是否显示弹窗 */
  visible: boolean;
  /** 记录数据 */
  record: ITokenUsageLog | null;
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
 * 格式化字节大小
 */
const formatBytes = (bytes: number): string => {
  if (bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${(bytes / Math.pow(k, i)).toFixed(2)} ${sizes[i]}`;
};

/**
 * 复制文本到剪贴板
 */
const copyToClipboard = (text: string) => {
  navigator.clipboard.writeText(text).then(() => {
    message.success('复制成功');
  });
};

/**
 * Token 使用记录详情弹窗组件
 */
const LogDetailModal: React.FC<ILogDetailModalProps> = ({ visible, record, onCancel }) => {
  /**
   * 获取状态标签颜色
   */
  const getStatusColor = (status: number): string => {
    if (status >= 200 && status < 300) return 'success';
    if (status >= 300 && status < 400) return 'warning';
    if (status >= 400 && status < 500) return 'error';
    return 'default';
  };

  /**
   * 渲染请求头
   */
  const renderHeaders = (headers: Record<string, string>) => {
    return (
      <div>
        {Object.entries(headers).map(([key, value]) => (
          <div key={key} style={{ marginBottom: 4 }}>
            <Text strong>{key}:</Text> {value}
          </div>
        ))}
      </div>
    );
  };

  /**
   * 渲染 JSON
   */
  const renderJson = (jsonStr: string) => {
    try {
      const parsed = JSON.parse(jsonStr);
      return (
        <pre
          style={{
            background: '#f5f5f5',
            padding: 12,
            borderRadius: 4,
            maxHeight: 400,
            overflow: 'auto',
            fontSize: 12,
          }}
        >
          {JSON.stringify(parsed, null, 2)}
        </pre>
      );
    } catch {
      return (
        <pre
          style={{
            background: '#f5f5f5',
            padding: 12,
            borderRadius: 4,
            maxHeight: 400,
            overflow: 'auto',
            fontSize: 12,
            whiteSpace: 'pre-wrap',
            wordBreak: 'break-all',
          }}
        >
          {jsonStr || '(空)'}
        </pre>
      );
    }
  };

  if (!record) return null;

  return (
    <Modal
      title="日志详情"
      open={visible}
      onCancel={onCancel}
      footer={[
        <Button key="close" onClick={onCancel}>
          关闭
        </Button>,
      ]}
      width={900}
    >
      <Tabs
        defaultActiveKey="basic"
        items={[
          {
            key: 'basic',
            label: '基本信息',
            children: (
              <Descriptions column={2} bordered size="small">
                <Descriptions.Item label="ID">{record.id}</Descriptions.Item>
                <Descriptions.Item label="Request ID">
                  <Space>
                    <Text code style={{ fontSize: 12 }}>
                      {record.request_id}
                    </Text>
                    <Button
                      type="text"
                      size="small"
                      icon={<CopyOutlined />}
                      onClick={() => copyToClipboard(record.request_id)}
                    />
                  </Space>
                </Descriptions.Item>
                <Descriptions.Item label="时间">{formatDateTime(record.time)}</Descriptions.Item>
                <Descriptions.Item label="日志级别">
                  <Tag>{record.level}</Tag>
                </Descriptions.Item>
                <Descriptions.Item label="请求方法">
                  <Tag className={`method-tag ${record.method}`}>{record.method}</Tag>
                </Descriptions.Item>
                <Descriptions.Item label="状态码">
                  <Tag color={getStatusColor(record.status)}>{record.status}</Tag>
                </Descriptions.Item>
                <Descriptions.Item label="请求路径" span={2}>
                  <Text code>{record.path}</Text>
                </Descriptions.Item>
                <Descriptions.Item label="查询参数" span={2}>
                  {record.query ? <Text code>{record.query}</Text> : '-'}
                </Descriptions.Item>
                <Descriptions.Item label="耗时">
                  <Badge
                    status={
                      record.latency_ms < 500 ? 'success' : record.latency_ms < 2000 ? 'warning' : 'error'
                    }
                    text={`${record.latency_ms} ms`}
                  />
                </Descriptions.Item>
                <Descriptions.Item label="请求/响应大小">
                  {formatBytes(record.request_size_bytes)} / {formatBytes(record.response_size_bytes)}
                </Descriptions.Item>
                <Descriptions.Item label="客户端地址" span={2}>
                  {record.remote_addr}
                </Descriptions.Item>
                <Descriptions.Item label="Authorization" span={2}>
                  <Space>
                    <Text code ellipsis style={{ maxWidth: 400 }}>
                      {record.authorization}
                    </Text>
                    <Button
                      type="text"
                      size="small"
                      icon={<CopyOutlined />}
                      onClick={() => copyToClipboard(record.authorization)}
                    />
                  </Space>
                </Descriptions.Item>
              </Descriptions>
            ),
          },
          {
            key: 'requestHeaders',
            label: '请求头',
            children: Object.keys(record.request_headers).length > 0 ? (
              renderHeaders(record.request_headers)
            ) : (
              <Text type="secondary">无请求头数据</Text>
            ),
          },
          {
            key: 'responseHeaders',
            label: '响应头',
            children: Object.keys(record.response_headers).length > 0 ? (
              renderHeaders(record.response_headers)
            ) : (
              <Text type="secondary">无响应头数据</Text>
            ),
          },
          {
            key: 'requestBody',
            label: '请求体',
            children: record.request_body ? renderJson(record.request_body) : <Text type="secondary">无请求体数据</Text>,
          },
          {
            key: 'other',
            label: '其他信息',
            children: (
              <Descriptions column={1} bordered size="small">
                <Descriptions.Item label="User-Agent">
                  {record.user_agent || '-'}
                </Descriptions.Item>
                <Descriptions.Item label="X-Forwarded-For">
                  {record.x_forwarded_for || '-'}
                </Descriptions.Item>
                <Descriptions.Item label="创建时间">
                  {formatDateTime(record.created_at)}
                </Descriptions.Item>
                <Descriptions.Item label="更新时间">
                  {formatDateTime(record.updated_at)}
                </Descriptions.Item>
              </Descriptions>
            ),
          },
        ]}
      />
    </Modal>
  );
};

export default LogDetailModal;
