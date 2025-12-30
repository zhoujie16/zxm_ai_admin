/**
 * 分布表格组件
 */
import React from 'react';
import { Card, Row, Col, Typography, Table } from 'antd';
import type { ColumnsType } from 'antd/es/table';
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

type DistributionDataType = {
  key: number;
  name: string;
  count: number;
};

const DistributionCharts: React.FC<IDistributionChartsProps> = ({ byIp, byPath }) => {
  // 确保数组不为 null
  const safeByIp = byIp ?? [];
  const safeByPath = byPath ?? [];

  // 转换数据格式，按count降序排列，限制显示前10条
  const ipData: DistributionDataType[] = safeByIp
    .sort((a, b) => b.count - a.count)
    .slice(0, 10)
    .map((item, index) => ({
      key: index,
      name: item.ip || '-',
      count: item.count,
    }));

  const pathData: DistributionDataType[] = safeByPath
    .sort((a, b) => b.count - a.count)
    .slice(0, 10)
    .map((item, index) => ({
      key: index,
      name: item.path || '-',
      count: item.count,
    }));

  const columns: ColumnsType<DistributionDataType> = [
    {
      title: '名称',
      dataIndex: 'name',
      key: 'name',
      width: '70%',
      ellipsis: true,
    },
    {
      title: '请求次数',
      dataIndex: 'count',
      key: 'count',
      width: '30%',
      render: (value: number) => formatNumber(value),
      align: 'right' as const,
    },
  ];

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
            <Table
              dataSource={ipData}
              columns={columns}
              pagination={false}
              size="small"
              bordered={false}
            />
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
            <Table
              dataSource={pathData}
              columns={columns}
              pagination={false}
              size="small"
              bordered={false}
            />
          ) : (
            renderEmpty()
          )}
        </Card>
      </Col>
    </Row>
  );
};

export default DistributionCharts;
