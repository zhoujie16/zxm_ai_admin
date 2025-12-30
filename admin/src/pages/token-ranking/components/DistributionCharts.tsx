/**
 * 分布图表组件
 */
import React from 'react';
import { Card, Row, Col, Typography } from 'antd';
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
  Cell,
} from 'recharts';
import { formatNumber } from '../utils';

const { Text } = Typography;

export interface IDistributionItem {
  name: string;
  count: number;
}

export interface IDistributionChartsProps {
  /** 按IP分组统计 */
  byIp: Array<{ ip: string; count: number }>;
  /** 按路径分组统计 */
  byPath: Array<{ path: string; count: number }>;
}

const COLORS = ['#1890ff', '#52c41a', '#faad14', '#f5222d', '#722ed1', '#fa8c16', '#13c2c2', '#eb2f96'];

// 统一卡片样式
const cardStyle: React.CSSProperties = {
  borderRadius: 8,
  boxShadow: '0 1px 2px 0 rgba(0, 0, 0, 0.03)',
  height: '100%',
};

const emptyStyle: React.CSSProperties = {
  textAlign: 'center',
  padding: '60px 0',
  color: '#999',
};

const DistributionCharts: React.FC<IDistributionChartsProps> = ({ byIp, byPath }) => {
  // 确保数组不为 null
  const safeByIp = byIp ?? [];
  const safeByPath = byPath ?? [];

  // 转换数据格式，限制显示前10条
  const ipData = safeByIp.slice(0, 10).map((item) => ({
    name: item.ip || '-',
    count: item.count,
  }));

  const pathData = safeByPath.slice(0, 10).map((item) => ({
    name: item.path || '-',
    count: item.count,
  }));

  const CustomTooltip = ({ active, payload }: any) => {
    if (active && payload && payload.length) {
      return (
        <div
          style={{
            backgroundColor: 'rgba(0, 0, 0, 0.85)',
            color: '#fff',
            padding: '10px 14px',
            borderRadius: '6px',
            fontSize: '12px',
            boxShadow: '0 4px 12px rgba(0, 0, 0, 0.15)',
          }}
        >
          <p style={{ margin: '0 0 4px 0', fontWeight: 500 }}>{payload[0].payload.name}</p>
          <p style={{ margin: 0 }}>请求次数: {formatNumber(payload[0].value)}</p>
        </div>
      );
    }
    return null;
  };

  const renderEmpty = () => (
    <div style={emptyStyle}>
      <Text type="secondary">暂无数据</Text>
    </div>
  );

  return (
    <Row gutter={16}>
      <Col span={12}>
        <Card
          title={<span style={{ fontSize: 14, fontWeight: 500 }}>IP 分布（Top 10）</span>}
          styles={{ body: { paddingTop: 16 } }}
          style={cardStyle}
        >
          {ipData.length > 0 ? (
            <ResponsiveContainer width="100%" height={280}>
              <BarChart data={ipData} layout="vertical" margin={{ left: 10, right: 10 }}>
                <XAxis type="number" tick={{ fontSize: 12 }} stroke="#999" />
                <YAxis dataKey="name" type="category" width={100} tick={{ fontSize: 12 }} stroke="#999" />
                <Tooltip content={<CustomTooltip />} />
                <Bar dataKey="count" radius={[0, 4, 4, 0]} barSize={20}>
                  {ipData.map((_, index) => (
                    <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
                  ))}
                </Bar>
              </BarChart>
            </ResponsiveContainer>
          ) : (
            renderEmpty()
          )}
        </Card>
      </Col>
      <Col span={12}>
        <Card
          title={<span style={{ fontSize: 14, fontWeight: 500 }}>路径分布（Top 10）</span>}
          styles={{ body: { paddingTop: 16 } }}
          style={cardStyle}
        >
          {pathData.length > 0 ? (
            <ResponsiveContainer width="100%" height={280}>
              <BarChart data={pathData} layout="vertical" margin={{ left: 10, right: 10 }}>
                <XAxis type="number" tick={{ fontSize: 12 }} stroke="#999" />
                <YAxis dataKey="name" type="category" width={120} tick={{ fontSize: 12 }} stroke="#999" />
                <Tooltip content={<CustomTooltip />} />
                <Bar dataKey="count" radius={[0, 4, 4, 0]} barSize={20}>
                  {pathData.map((_, index) => (
                    <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
                  ))}
                </Bar>
              </BarChart>
            </ResponsiveContainer>
          ) : (
            renderEmpty()
          )}
        </Card>
      </Col>
    </Row>
  );
};

export default DistributionCharts;
