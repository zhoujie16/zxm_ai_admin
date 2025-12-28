/**
 * Token 使用记录页面
 * 功能：查询和查看 Token 请求日志
 */
import { Card } from 'antd';
import React from 'react';
import { ProTable } from '@ant-design/pro-components';
import type { ActionType } from '@ant-design/pro-components';
import LogDetailModal from './components/LogDetailModal';
import type { ITokenUsageLog, ITokenUsageLogListRequest } from '@/types/tokenUsageLog';
import { getColumns } from './columns';
import { useDetailModal } from './useDetailModal';
import { useRequestLog } from './useRequestLog';
import './index.less';

const TokenUsageLogPage: React.FC = () => {
  const { detailVisible, detailRecord, handleViewDetail, handleCloseDetail } = useDetailModal();
  const requestLog = useRequestLog();
  const actionRef = React.useRef<ActionType>(null);

  const columns = getColumns({ onViewDetail: handleViewDetail });

  return (
    <div className="token-usage-log-page">
      <Card>
        <ProTable<ITokenUsageLog, ITokenUsageLogListRequest & { timeRange?: [string, string] }>
          columns={columns}
          actionRef={actionRef}
          request={requestLog}
          rowKey="id"
          search={{
            defaultCollapsed: false,
            span: 8,
          }}
          pagination={{
            defaultPageSize: 10,
            showSizeChanger: true,
            showQuickJumper: true,
          }}
          options={false}
          scroll={{ x: 1400 }}
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

export default TokenUsageLogPage;
