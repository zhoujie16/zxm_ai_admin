/**
 * 统计数据抽屉组件
 */
import React from 'react';
import { Drawer, Spin, Empty, Space, Typography, Tag, Divider } from 'antd';
import { ClockCircleOutlined, BarChartOutlined } from '@ant-design/icons';
import { useStatistics } from '../hooks/useStatistics';
import { formatDateTime, getTokenDisplayName } from '../utils';
import SummaryCards from './SummaryCards';
import LatencyCard from './LatencyCard';
import DistributionCharts from './DistributionCharts';
import TrendCharts from './TrendCharts';

const { Text, Title } = Typography;

export interface IStatisticsDrawerProps {
  /** 是否显示抽屉 */
  visible: boolean;
  /** authorization 值 */
  authorization: string;
  /** 开始时间 */
  startTime: string | undefined;
  /** 结束时间 */
  endTime: string | undefined;
  /** 关闭回调 */
  onClose: () => void;
}

// 内容区域样式
const sectionStyle: React.CSSProperties = {
  marginBottom: 24,
};

const StatisticsDrawer: React.FC<IStatisticsDrawerProps> = ({
  visible,
  authorization,
  startTime,
  endTime,
  onClose,
}) => {
  const { loading, data, fetchStatistics, reset } = useStatistics();

  // 当抽屉打开且有 authorization 时，获取数据
  React.useEffect(() => {
    if (visible && authorization) {
      reset();
      // 将 ISO 格式时间转换为后端需要的格式
      const formatTime = (time?: string): string | undefined => {
        if (!time) return undefined;
        const date = new Date(time);
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        const hour = String(date.getHours()).padStart(2, '0');
        const minute = String(date.getMinutes()).padStart(2, '0');
        const second = String(date.getSeconds()).padStart(2, '0');
        return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
      };
      fetchStatistics(authorization, formatTime(startTime), formatTime(endTime));
    }
  }, [visible, authorization, startTime, endTime, reset, fetchStatistics]);

  const handleClose = () => {
    onClose();
  };

  return (
    <Drawer
      title={
        <Space size="middle">
          <span>Token 统计数据</span>
          {authorization && (
            <Tag color="blue">{getTokenDisplayName(authorization)}</Tag>
          )}
        </Space>
      }
      placement="right"
      width={1300}
      open={visible}
      onClose={handleClose}
      styles={{
        body: { padding: '24px' },
      }}
    >
      <Spin spinning={loading}>
        {data ? (
          <div>
            {/* 时间范围 */}
            <div
              style={{
                padding: '12px 16px',
                marginBottom: 20,
                backgroundColor: '#f5f5f5',
                borderRadius: 8,
                border: '1px solid #f0f0f0',
              }}
            >
              <Space size="large">
                <Space>
                  <ClockCircleOutlined style={{ color: '#1890ff' }} />
                  <Text type="secondary">开始时间</Text>
                  <Text>{formatDateTime(data.time_range.start)}</Text>
                </Space>
                <Divider type="vertical" style={{ margin: 0 }} />
                <Space>
                  <ClockCircleOutlined style={{ color: '#52c41a' }} />
                  <Text type="secondary">结束时间</Text>
                  <Text>{formatDateTime(data.time_range.end)}</Text>
                </Space>
              </Space>
            </div>

            {/* 核心指标区域 */}
            <div style={sectionStyle}>
              <Title level={5} style={{ marginBottom: 16 }}>
                <BarChartOutlined /> 核心指标
              </Title>

              {/* 汇总统计 */}
              <div style={{ marginBottom: 20 }}>
                <SummaryCards
                  totalRequests={data.summary.total_requests}
                  totalRequestBytes={data.summary.total_request_bytes}
                  totalResponseBytes={data.summary.total_response_bytes}
                  avgRequestBytes={data.summary.avg_request_bytes}
                  avgResponseBytes={data.summary.avg_response_bytes}
                />
              </div>

              {/* 延迟统计 */}
              <div>
                <LatencyCard
                  totalMs={data.latency.total_ms}
                  avgMs={data.latency.avg_ms}
                  minMs={data.latency.min_ms}
                  maxMs={data.latency.max_ms}
                />
              </div>
            </div>

            <Divider style={{ margin: '24px 0' }} />

            {/* 数据分布 */}
            <div style={sectionStyle}>
              <Title level={5} style={{ marginBottom: 16 }}>
                <BarChartOutlined /> 数据分布
              </Title>
              <DistributionCharts byIp={data.by_ip} byPath={data.by_path} />
            </div>

            {/* 趋势分析 */}
            <div>
              <Title level={5} style={{ marginBottom: 16 }}>
                <BarChartOutlined /> 趋势分析
              </Title>
              <TrendCharts byDate={data.by_date} byTime={data.by_time} />
            </div>
          </div>
        ) : (
          <Empty description="暂无统计数据" />
        )}
      </Spin>
    </Drawer>
  );
};

export default StatisticsDrawer;
