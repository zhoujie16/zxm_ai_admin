/**
 * 请求日志数据加载 hook
 */
import { useCallback } from 'react';
import type { ITokenUsageLogListRequest } from '@/types/tokenUsageLog';
import { getTokenUsageLogList } from '@/services/tokenUsageLog';
import { OTHER_STATUSES } from './constants';

/**
 * ProTable request 函数
 */
export const useRequestLog = () => {
  return useCallback(async (params: (ITokenUsageLogListRequest & { timeRange?: [string, string] }) & {
    current?: number;
    pageSize?: number;
  }) => {
    const { current, pageSize, timeRange, status, ...rest } = params;

    const requestParams: ITokenUsageLogListRequest = {
      page: current ?? 1,
      page_size: pageSize ?? 10,
      ...rest,
    };

    if (timeRange) {
      requestParams.start_time = timeRange[0];
      requestParams.end_time = timeRange[1];
    }

    // 处理状态码筛选：status=0 表示「其它」，传所有其他状态码
    if (status === 0) {
      requestParams.status = OTHER_STATUSES.join(',');
    } else if (status) {
      requestParams.status = String(status);
    }

    const result = await getTokenUsageLogList(requestParams);
    return {
      data: result.data?.list || [],
      success: result.success,
      total: result.data?.total || 0,
    };
  }, []);
};
