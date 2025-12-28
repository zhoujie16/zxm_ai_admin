/**
 * 请求日志数据加载 hook
 */
import { useCallback } from 'react';
import type { ITokenUsageLogListRequest } from '@/types/tokenUsageLog';
import { getTokenUsageLogList } from '@/services/tokenUsageLog';

/**
 * ProTable request 函数
 */
export const useRequestLog = () => {
  return useCallback(async (params: (ITokenUsageLogListRequest & { timeRange?: [string, string] }) & {
    current?: number;
    pageSize?: number;
  }) => {
    const { current, pageSize, timeRange, ...rest } = params;

    const requestParams: ITokenUsageLogListRequest = {
      page: current ?? 1,
      page_size: pageSize ?? 10,
      ...rest,
    };

    if (timeRange) {
      requestParams.start_time = timeRange[0];
      requestParams.end_time = timeRange[1];
    }

    const result = await getTokenUsageLogList(requestParams);
    return {
      data: result.data?.list || [],
      success: result.success,
      total: result.data?.total || 0,
    };
  }, []);
};
