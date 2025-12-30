/**
 * 汇总统计卡片组件
 */
import React from 'react';
import { Card, Row, Col, Statistic } from 'antd';
import { ApiOutlined, CloudDownloadOutlined, CloudUploadOutlined } from '@ant-design/icons';
import { formatBytes, formatNumber } from '../utils';

export interface ISummaryCardsProps {
  /** 总请求数 */
  totalRequests: number;
  /** 总请求字节数 */
  totalRequestBytes: number;
  /** 总响应字节数 */
  totalResponseBytes: number;
  /** 平均请求字节数 */
  avgRequestBytes: number;
  /** 平均响应字节数 */
  avgResponseBytes: number;
}

// 统一卡片样式
const cardStyle: React.CSSProperties = {
  height: '100%',
  borderRadius: 8,
  boxShadow: '0 1px 2px 0 rgba(0, 0, 0, 0.03)',
};

const SummaryCards: React.FC<ISummaryCardsProps> = ({
  totalRequests,
  totalRequestBytes,
  totalResponseBytes,
  avgRequestBytes,
  avgResponseBytes,
}) => {
  const items = [
    {
      title: '总请求次数',
      value: totalRequests,
      icon: <ApiOutlined />,
      color: '#1890ff',
      formatter: formatNumber,
    },
    {
      title: '总请求流量',
      value: totalRequestBytes,
      icon: <CloudUploadOutlined />,
      color: '#52c41a',
      formatter: formatBytes,
    },
    {
      title: '总响应流量',
      value: totalResponseBytes,
      icon: <CloudDownloadOutlined />,
      color: '#722ed1',
      formatter: formatBytes,
    },
    {
      title: '平均请求大小',
      value: avgRequestBytes,
      color: '#fa8c16',
      formatter: formatBytes,
    },
    {
      title: '平均响应大小',
      value: avgResponseBytes,
      color: '#eb2f96',
      formatter: formatBytes,
    },
  ];

  return (
    <Row gutter={[12, 12]}>
      {items.map((item) => (
        <Col span={24 / 5} key={item.title}>
          <Card
            styles={{ body: { padding: '16px 20px' } }}
            style={cardStyle}
          >
            <Statistic
              title={<span style={{ fontSize: 13 }}>{item.title}</span>}
              value={item.value}
              prefix={item.icon}
              valueStyle={{ color: item.color, fontSize: 20 }}
              formatter={(value) => item.formatter(Number(value))}
            />
          </Card>
        </Col>
      ))}
    </Row>
  );
};

export default SummaryCards;
