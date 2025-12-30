/**
 * 延迟统计卡片组件
 */
import React from 'react';
import { Card, Row, Col, Statistic, Space, Typography } from 'antd';
import { ClockCircleOutlined, MinusCircleOutlined, PlusCircleOutlined } from '@ant-design/icons';
import { formatNumber } from '../utils';

const { Text } = Typography;

export interface ILatencyCardProps {
  /** 延迟总和（毫秒） */
  totalMs: number;
  /** 平均延迟（毫秒） */
  avgMs: number;
  /** 最小延迟（毫秒） */
  minMs: number;
  /** 最大延迟（毫秒） */
  maxMs: number;
}

// 统一卡片样式
const cardStyle: React.CSSProperties = {
  height: '100%',
  borderRadius: 8,
  boxShadow: '0 1px 2px 0 rgba(0, 0, 0, 0.03)',
};

const LatencyCard: React.FC<ILatencyCardProps> = ({
  totalMs,
  avgMs,
  minMs,
  maxMs,
}) => {
  // 将毫秒转换为秒显示
  const totalSeconds = (totalMs / 1000).toFixed(2);

  return (
    <div>
      <Space style={{ marginBottom: 12 }}>
        <ClockCircleOutlined style={{ color: '#1890ff' }} />
        <Text strong>延迟统计</Text>
      </Space>
      <Row gutter={[12, 0]}>
        {/* 延迟总和 */}
        <Col span={24 / 5}>
          <Card
            styles={{ body: { padding: '16px 20px' } }}
            style={cardStyle}
          >
            <Statistic
              title={<span style={{ fontSize: 13 }}>延迟总和</span>}
              value={totalSeconds}
              suffix="秒"
              prefix={<ClockCircleOutlined />}
              valueStyle={{ color: '#1890ff', fontSize: 20 }}
            />
          </Card>
        </Col>
        {/* 平均延迟 */}
        <Col span={24 / 5}>
          <Card
            styles={{ body: { padding: '16px 20px' } }}
            style={cardStyle}
          >
            <Statistic
              title={<span style={{ fontSize: 13 }}>平均延迟</span>}
              value={avgMs}
              suffix="ms"
              prefix={<MinusCircleOutlined />}
              valueStyle={{ color: '#52c41a', fontSize: 20 }}
              formatter={(value) => formatNumber(Number(value))}
            />
          </Card>
        </Col>
        {/* 最小延迟 */}
        <Col span={24 / 5}>
          <Card
            styles={{ body: { padding: '16px 20px' } }}
            style={cardStyle}
          >
            <Statistic
              title={<span style={{ fontSize: 13 }}>最小延迟</span>}
              value={minMs}
              suffix="ms"
              valueStyle={{ color: '#13c2c2', fontSize: 20 }}
              formatter={(value) => formatNumber(Number(value))}
            />
          </Card>
        </Col>
        {/* 最大延迟 */}
        <Col span={24 / 5}>
          <Card
            styles={{ body: { padding: '16px 20px' } }}
            style={cardStyle}
          >
            <Statistic
              title={<span style={{ fontSize: 13 }}>最大延迟</span>}
              value={maxMs}
              suffix="ms"
              prefix={<PlusCircleOutlined />}
              valueStyle={{ color: '#ff4d4f', fontSize: 20 }}
              formatter={(value) => formatNumber(Number(value))}
            />
          </Card>
        </Col>
        {/* 占位列，保持对齐 */}
        <Col span={24 / 5}>
          <Card
            styles={{ body: { padding: '16px 20px' } }}
            style={{ ...cardStyle, borderStyle: 'dashed', borderColor: '#d9d9d9', backgroundColor: '#fafafa' }}
          >
            <Statistic
              title={<span style={{ fontSize: 13 }}><Text type="secondary">预留指标</Text></span>}
              value="-"
              valueStyle={{ color: '#d9d9d9', fontSize: 20 }}
            />
          </Card>
        </Col>
      </Row>
    </div>
  );
};

export default LatencyCard;
