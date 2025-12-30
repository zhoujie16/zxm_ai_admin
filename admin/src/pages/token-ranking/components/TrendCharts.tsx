/**
 * 趋势图表组件
 */
import React from 'react';
import { Card, Radio, Typography } from 'antd';
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
  CartesianGrid,
} from 'recharts';
import { formatNumber } from '../utils';

const { Text } = Typography;

export interface ITrendChartsProps {
  /** 按日期分组统计 */
  byDate: Array<{ date: string; count: number }>;
  /** 按小时分组统计 */
  byTime: Array<{ time: string; count: number }>;
}

type TrendType = 'date' | 'time';

// 统一卡片样式
const cardStyle: React.CSSProperties = {
  borderRadius: 8,
  boxShadow: '0 1px 2px 0 rgba(0, 0, 0, 0.03)',
};

const emptyStyle: React.CSSProperties = {
  textAlign: 'center',
  padding: '80px 0',
  color: '#999',
};

const TrendCharts: React.FC<ITrendChartsProps> = ({ byDate, byTime }) => {
  const [trendType, setTrendType] = React.useState<TrendType>('date');

  // 确保数组不为 null
  const safeByDate = byDate ?? [];
  const safeByTime = byTime ?? [];

  // 转换数据格式
  const dateData = safeByDate.map((item) => ({
    name: item.date,
    count: item.count,
  }));

  const timeData = safeByTime.map((item) => ({
    name: item.time.length >= 16 ? item.time.slice(5, 16) : item.time,
    count: item.count,
  }));

  const currentData = trendType === 'date' ? dateData : timeData;

  const CustomTooltip = ({ active, payload, label }: any) => {
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
          <p style={{ margin: '0 0 4px 0', fontWeight: 500 }}>{label}</p>
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
    <Card
      title={<span style={{ fontSize: 14, fontWeight: 500 }}>请求趋势</span>}
      styles={{ body: { paddingTop: 16 } }}
      style={cardStyle}
      extra={
        <Radio.Group
          value={trendType}
          onChange={(e) => setTrendType(e.target.value)}
          size="small"
        >
          <Radio.Button value="date">按日</Radio.Button>
          <Radio.Button value="time">按时</Radio.Button>
        </Radio.Group>
      }
    >
      {currentData.length > 0 ? (
        <ResponsiveContainer width="100%" height={320}>
          <LineChart data={currentData} margin={{ left: 0, right: 10, top: 10, bottom: 0 }}>
            <CartesianGrid strokeDasharray="3 3" stroke="#f0f0f0" vertical={false} />
            <XAxis
              dataKey="name"
              tick={{ fontSize: 12 }}
              interval="preserveStartEnd"
              stroke="#999"
            />
            <YAxis tick={{ fontSize: 12 }} stroke="#999" />
            <Tooltip content={<CustomTooltip />} />
            <Line
              type="monotone"
              dataKey="count"
              stroke="#1890ff"
              strokeWidth={2}
              dot={{ fill: '#1890ff', r: 4, strokeWidth: 0 }}
              activeDot={{ r: 6, strokeWidth: 0 }}
            />
          </LineChart>
        </ResponsiveContainer>
      ) : (
        renderEmpty()
      )}
    </Card>
  );
};

export default TrendCharts;
