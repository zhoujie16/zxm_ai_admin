/**
 * 系统日志页面
 * 功能：查询和查看系统日志
 */
import { Card } from 'antd';
import React from 'react';
import { ProTable } from '@ant-design/pro-components';
import type { ActionType } from '@ant-design/pro-components';
import LogDetailModal from './components/LogDetailModal';
import type { ISystemLog, ISystemLogListRequest } from '@/types/systemLog';
import { getColumns } from './columns';
import { useDetailModal } from './useDetailModal';
import { useRequestLog } from './useRequestLog';
import './index.less';

const SystemLogPage: React.FC = () => {
  const { detailVisible, detailRecord, handleViewDetail, handleCloseDetail } = useDetailModal();
  const requestLog = useRequestLog();
  const actionRef = React.useRef<ActionType>(null);

  const columns = getColumns({ onViewDetail: handleViewDetail });

  return (
    <div className="system-log-page">
      <Card>
        <ProTable<ISystemLog, ISystemLogListRequest & { timeRange?: [string, string] }>
          columns={columns}
          actionRef={actionRef}
          request={requestLog}
          rowKey="id"
          search={{
            defaultCollapsed: false,
            span: 6,
          }}
          pagination={{
            defaultPageSize: 10,
            showSizeChanger: true,
            showQuickJumper: true,
          }}
          options={false}
          scroll={{ x: 800 }}
        />
      </Card>

      <LogDetailModal
        visible={detailVisible}
        record={detailRecord}
        onCancel={handleCloseDetail}
      />
    </div>
  );
};

export default SystemLogPage;
