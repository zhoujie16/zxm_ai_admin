/**
 * Token 使用排行榜页面
 * 功能：展示 Authorization 使用次数排行榜，查看统计数据
 */
import React, { useState, useCallback } from 'react';
import { Card } from 'antd';
import { ProTable } from '@ant-design/pro-components';
import type { ActionType, ProFormInstance } from '@ant-design/pro-components';
import type { ITokenRankingItem, ITokenRankingRequest } from '@/types/tokenRanking';
import { getTokenRanking } from '@/services/tokenUsageLog';
import { getColumns } from './columns';
import StatisticsDrawer from './components/StatisticsDrawer';
import './index.less';

interface IProTableParams {
  current?: number;
  pageSize?: number;
  timeRange?: [string, string];
}

const TokenRankingPage: React.FC = () => {
  const [drawerVisible, setDrawerVisible] = useState(false);
  const [selectedRecord, setSelectedRecord] = useState<ITokenRankingItem | null>(null);
  const [timeRange, setTimeRange] = useState<[string, string] | null>(null);
  const actionRef = React.useRef<ActionType>(null);
  // @ts-ignore - ProFormInstance generic is optional
  const formRef = React.useRef<ProFormInstance>();

  // 打开统计抽屉
  const handleViewStatistics = useCallback((record: ITokenRankingItem) => {
    setSelectedRecord(record);
    setDrawerVisible(true);
  }, []);

  // 关闭抽屉
  const handleCloseDrawer = useCallback(() => {
    setDrawerVisible(false);
    setSelectedRecord(null);
  }, []);

  // 请求函数
  const requestRanking = useCallback(async (params: IProTableParams) => {
    const current = params.current ?? 1;
    const pageSize = params.pageSize ?? 20;
    const range = params.timeRange;

    // 保存时间范围用于统计查询
    if (range && range[0] && range[1]) {
      setTimeRange(range);
    }

    const requestParams: ITokenRankingRequest = {
      page: current,
      page_size: pageSize,
    };

    // 处理时间范围
    if (range && range[0] && range[1]) {
      const formatDate = (dateStr: string) => {
        const date = new Date(dateStr);
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        const hour = String(date.getHours()).padStart(2, '0');
        const minute = String(date.getMinutes()).padStart(2, '0');
        const second = String(date.getSeconds()).padStart(2, '0');
        return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
      };
      requestParams.start_time = formatDate(range[0]);
      requestParams.end_time = formatDate(range[1]);
    }

    const response = await getTokenRanking(requestParams);
    return {
      data: response.data?.list ?? [],
      success: true,
      total: response.data?.total ?? 0,
    };
  }, []);

  const columns = getColumns({ onViewStatistics: handleViewStatistics });

  return (
    <div className="token-ranking-page">
      <Card>
        <ProTable<ITokenRankingItem, IProTableParams>
          columns={columns}
          actionRef={actionRef}
          formRef={formRef}
          request={requestRanking}
          rowKey={(record) => record.authorization}
          search={{
            defaultCollapsed: false,
            span: 8,
          }}
          pagination={{
            defaultPageSize: 20,
            showSizeChanger: true,
            showQuickJumper: true,
          }}
          options={false}
          dateFormatter="string"
        />
      </Card>

      {selectedRecord && (
        <StatisticsDrawer
          visible={drawerVisible}
          authorization={selectedRecord.authorization}
          startTime={timeRange?.[0]}
          endTime={timeRange?.[1]}
          onClose={handleCloseDrawer}
        />
      )}
    </div>
  );
};

export default TokenRankingPage;
