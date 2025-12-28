/**
 * 删除日志页面
 * 功能：按时间范围删除请求日志或系统日志
 */
import { useState } from 'react';
import { Card, Button, DatePicker, Input, message, Space, Alert, Result, Tabs, Modal } from 'antd';
import { DeleteOutlined, CheckCircleOutlined, ExclamationCircleOutlined } from '@ant-design/icons';
import type { Dayjs } from 'dayjs';
import type { TabsProps } from 'antd';
import { post } from '@/utils/request';
import { timeRangePresets, formatDateTime } from '@/pages/token-usage-log/utils';
import './index.less';

const { RangePicker } = DatePicker;

interface DeleteResponse {
  deleted_count: number;
}

type LogType = 'request' | 'system';

const DeleteLogPage: React.FC = () => {
  const [activeTab, setActiveTab] = useState<LogType>('request');
  const [timeRange, setTimeRange] = useState<[Dayjs, Dayjs] | null>(null);
  const [token, setToken] = useState('');
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [deletedCount, setDeletedCount] = useState<number>(0);

  /** 切换 Tab */
  const handleTabChange = (key: string) => {
    setActiveTab(key as LogType);
    // 切换时重置状态
    setTimeRange(null);
    setToken('');
    setSuccess(false);
    setDeletedCount(0);
  };

  /** 执行删除 */
  const handleDelete = () => {
    // 验证时间范围
    if (!timeRange || timeRange.length !== 2) {
      message.warning('请选择时间范围');
      return;
    }

    // 验证 token
    if (!token.trim()) {
      message.warning('请输入系统认证令牌');
      return;
    }

    const logTypeName = activeTab === 'request' ? '请求日志' : '系统日志';
    const timeRangeText = `${formatDateTime(timeRange[0].format())} ~ ${formatDateTime(timeRange[1].format())}`;

    // 二次确认弹窗
    Modal.confirm({
      title: '确认删除',
      icon: <ExclamationCircleOutlined />,
      content: (
        <div>
          <p>您即将删除以下日志：</p>
          <p style={{ color: '#ff4d4f', fontWeight: 500 }}>
            日志类型：{logTypeName}
            <br />
            时间范围：{timeRangeText}
          </p>
          <p style={{ color: '#ff4d4f', marginTop: 12 }}>
            此操作将永久删除该时间范围内的所有日志，无法恢复！
          </p>
        </div>
      ),
      okText: '确认删除',
      cancelText: '取消',
      okType: 'danger',
      okButtonProps: { loading },
      onOk: async () => {
        setLoading(true);
        try {
          const params = {
            start_time: timeRange[0].format('YYYY-MM-DD HH:mm:ss'),
            end_time: timeRange[1].format('YYYY-MM-DD HH:mm:ss'),
            system_auth_token: token.trim(),
          };

          // 根据 Tab 类型选择 API
          const apiUrl = activeTab === 'request'
            ? '/api-logs/api/request-logs/delete'
            : '/api-logs/api/system-logs/delete';

          const result = await post<DeleteResponse>(apiUrl, params);

          if (result.success && result.data) {
            setDeletedCount(result.data.deleted_count);
            setSuccess(true);
            message.success(`成功删除 ${result.data.deleted_count} 条${logTypeName}`);
          }
          return true;
        } finally {
          setLoading(false);
        }
      },
    });
  };

  /** 重置表单 */
  const handleReset = () => {
    setTimeRange(null);
    setToken('');
    setSuccess(false);
    setDeletedCount(0);
  };

  // 删除成功后显示结果
  if (success) {
    const logTypeName = activeTab === 'request' ? '请求日志' : '系统日志';
    return (
      <div className="delete-log-page">
        <Card>
          <Result
            status="success"
            icon={<CheckCircleOutlined style={{ color: '#52c41a' }} />}
            title="删除成功"
            subTitle={`已成功删除 ${deletedCount} 条${logTypeName}`}
            extra={[
              <Button type="primary" key="again" onClick={handleReset}>
                继续删除
              </Button>,
            ]}
          />
        </Card>
      </div>
    );
  }

  const items: TabsProps['items'] = [
    {
      key: 'request',
      label: '删除请求日志',
    },
    {
      key: 'system',
      label: '删除系统日志',
    },
  ];

  return (
    <div className="delete-log-page">
      <Card title="删除日志">
        <Tabs activeKey={activeTab} items={items} onChange={handleTabChange} />

        <Space direction="vertical" style={{ width: '100%', marginTop: 24 }} size="large">
          {/* 时间范围选择 */}
          <div>
            <div className="form-label">
              时间范围 <span className="required">*</span>
            </div>
            <RangePicker
              showTime
              value={timeRange}
              onChange={(dates) => setTimeRange(dates as [Dayjs, Dayjs] | null)}
              presets={timeRangePresets as any}
              style={{ width: '100%' }}
              format="YYYY-MM-DD HH:mm:ss"
              size="large"
            />
            {timeRange && (
              <div className="time-range-display">
                {formatDateTime(timeRange[0].format())} ~ {formatDateTime(timeRange[1].format())}
              </div>
            )}
          </div>

          {/* 认证令牌输入 */}
          <div>
            <div className="form-label">
              系统认证令牌 <span className="required">*</span>
            </div>
            <Input
              value={token}
              onChange={(e) => setToken(e.target.value)}
              placeholder="请输入系统认证令牌"
              autoComplete="off"
              size="large"
            />
          </div>

          {/* 警告提示 */}
          <Alert
            message="危险操作"
            description="此操作将永久删除该时间范围内的所有日志，无法恢复，请谨慎操作！"
            type="error"
            showIcon
          />

          {/* 操作按钮 */}
          <Space>
            <Button
              type="primary"
              danger
              icon={<DeleteOutlined />}
              loading={loading}
              onClick={handleDelete}
              size="large"
            >
              确认删除
            </Button>
            <Button onClick={handleReset} size="large">
              重置
            </Button>
          </Space>
        </Space>
      </Card>
    </div>
  );
};

export default DeleteLogPage;
