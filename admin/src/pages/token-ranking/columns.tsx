/**
 * 表格列定义
 */
import React from 'react';
import { Button } from 'antd';
import { BarChartOutlined } from '@ant-design/icons';
import type { ProColumns } from '@ant-design/pro-components';
import type { ITokenRankingItem } from '@/types/tokenRanking';
import { getTokenDisplayName } from './utils';

export interface IGetColumnsParams {
  /** 查看统计回调 */
  onViewStatistics: (record: ITokenRankingItem) => void;
}

/** 获取表格列定义 */
export const getColumns = (params: IGetColumnsParams): ProColumns<ITokenRankingItem>[] => {
  const { onViewStatistics } = params;

  return [
    {
      title: '排名',
      dataIndex: 'index',
      width: 80,
      search: false,
      render: (_, __, index) => {
        const rank = (index ?? 0) + 1;
        let rankStyle: React.CSSProperties = {};
        let rankText: string | number = rank;

        if (rank === 1) {
          rankStyle = { color: '#ffd700', fontWeight: 'bold', fontSize: '18px' };
          rankText = '🥇';
        } else if (rank === 2) {
          rankStyle = { color: '#c0c0c0', fontWeight: 'bold', fontSize: '18px' };
          rankText = '🥈';
        } else if (rank === 3) {
          rankStyle = { color: '#cd7f32', fontWeight: 'bold', fontSize: '18px' };
          rankText = '🥉';
        }

        return <span style={rankStyle}>{rankText}</span>;
      },
    },
    {
      title: 'Authorization',
      dataIndex: 'authorization',
      width: 300,
      ellipsis: true,
      search: false,
      render: (_, record) => (
        <span title={record.authorization}>{getTokenDisplayName(record.authorization)}</span>
      ),
    },
    {
      title: '请求次数',
      dataIndex: 'count',
      width: 150,
      search: false,
      render: (_, record) => (
        <span style={{ fontWeight: 'bold', color: '#1890ff' }}>
          {record.count.toLocaleString()}
        </span>
      ),
      sorter: (a, b) => a.count - b.count,
    },
    {
      title: '时间范围',
      dataIndex: 'timeRange',
      valueType: 'dateTimeRange',
      hideInTable: true,
      fieldProps: {
        placeholder: ['开始时间', '结束时间'],
        showTime: true,
      },
    },
    {
      title: '操作',
      valueType: 'option',
      width: 120,
      fixed: 'right',
      render: (_: unknown, record: ITokenRankingItem) => (
        <Button
          type="link"
          size="small"
          icon={<BarChartOutlined />}
          onClick={() => onViewStatistics(record)}
        >
          查看统计
        </Button>
      ),
    },
  ];
};
