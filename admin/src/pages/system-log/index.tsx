/**
 * 系统日志页面
 * 功能：查询和查看系统日志
 */
import { Card } from 'antd';
import React, { useCallback, useState } from 'react';
import { useSystemLog } from './hooks/useSystemLog';
import SystemLogFilter from './components/SystemLogFilter';
import SystemLogTable from './components/SystemLogTable';
import LogDetailModal from './components/LogDetailModal';
import type { ISystemLog, ISystemLogListRequest } from '@/types/systemLog';
import './index.less';

/**
 * 系统日志页面组件
 */
const SystemLogPage: React.FC = () => {
  const {
    dataSource,
    total,
    loading,
    pagination,
    loadData,
  } = useSystemLog();

  const [detailVisible, setDetailVisible] = useState(false);
  const [detailRecord, setDetailRecord] = useState<ISystemLog | null>(null);
  const [filterParams, setFilterParams] = useState<ISystemLogListRequest>({});

  /**
   * 处理过滤条件变化
   */
  const handleFilterChange = useCallback((params: ISystemLogListRequest) => {
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
  const handleViewDetail = useCallback(async (record: ISystemLog) => {
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
    <div className="system-log-page">
      <Card>
        <SystemLogFilter
          onFilterChange={handleFilterChange}
          onRefresh={handleRefresh}
          loading={loading}
        />

        <SystemLogTable
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

export default SystemLogPage;
