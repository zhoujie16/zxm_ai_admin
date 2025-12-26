/**
 * Token 使用记录页面
 * 功能：查询和查看 Token 请求日志
 */
import { Card } from 'antd';
import React, { useCallback, useState } from 'react';
import { useTokenUsageLog } from './hooks/useTokenUsageLog';
import TokenUsageLogFilter from './components/TokenUsageLogFilter';
import TokenUsageLogTable from './components/TokenUsageLogTable';
import LogDetailModal from './components/LogDetailModal';
import type { ITokenUsageLog, ITokenUsageLogListRequest } from '@/types/tokenUsageLog';
import './index.less';

/**
 * Token 使用记录页面组件
 */
const TokenUsageLogPage: React.FC = () => {
  const {
    dataSource,
    total,
    loading,
    pagination,
    loadData,
  } = useTokenUsageLog();

  const [detailVisible, setDetailVisible] = useState(false);
  const [detailRecord, setDetailRecord] = useState<ITokenUsageLog | null>(null);
  const [filterParams, setFilterParams] = useState<ITokenUsageLogListRequest>({});

  /**
   * 处理过滤条件变化
   */
  const handleFilterChange = useCallback((params: ITokenUsageLogListRequest) => {
    setFilterParams(params);
  }, []);

  /**
   * 处理表格分页变化
   */
  const handleTablePageChange = useCallback((page: number, pageSize: number) => {
    loadData(page, pageSize, filterParams);
  }, [loadData, filterParams]);

  /**
   * 处理刷新
   */
  const handleRefresh = useCallback(() => {
    loadData(pagination.current, pagination.pageSize, filterParams);
  }, [loadData, pagination, filterParams]);

  /**
   * 查看详情
   */
  const handleViewDetail = useCallback(async (record: ITokenUsageLog) => {
    setDetailRecord(record);
    setDetailVisible(true);
  }, []);

  /**
   * 关闭详情弹窗
   */
  const handleCloseDetail = useCallback(() => {
    setDetailVisible(false);
    setDetailRecord(null);
  }, []);

  return (
    <div className="token-usage-log-page">
      <Card>
        <TokenUsageLogFilter
          onFilterChange={handleFilterChange}
          onRefresh={handleRefresh}
          loading={loading}
        />

        <TokenUsageLogTable
          dataSource={dataSource}
          loading={loading}
          pagination={{
            current: pagination.current,
            pageSize: pagination.pageSize,
            total: total,
          }}
          onPageChange={handleTablePageChange}
          onViewDetail={handleViewDetail}
        />

        <LogDetailModal
          visible={detailVisible}
          record={detailRecord}
          onCancel={handleCloseDetail}
        />
      </Card>
    </div>
  );
};

export default TokenUsageLogPage;
