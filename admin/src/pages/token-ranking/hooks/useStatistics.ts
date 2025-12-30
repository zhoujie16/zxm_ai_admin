/**
 * 统计数据 Hook
 */
import { useState, useCallback } from 'react';
import { message } from 'antd';
import { getUserStatistics } from '@/services/tokenUsageLog';
import type { IUserStatisticsResponse, IUserStatisticsRequest } from '@/types/tokenRanking';

export interface IUseStatisticsResult {
  /** 数据加载状态 */
  loading: boolean;
  /** 统计数据 */
  data: IUserStatisticsResponse | null;
  /** 获取统计数据 */
  fetchStatistics: (authorization: string, startTime?: string, endTime?: string) => Promise<void>;
  /** 重置数据 */
  reset: () => void;
}

export function useStatistics(): IUseStatisticsResult {
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState<IUserStatisticsResponse | null>(null);

  const fetchStatistics = useCallback(async (
    authorization: string,
    startTime?: string,
    endTime?: string
  ) => {
    setLoading(true);
    try {
      // 构建请求参数，动态添加可选字段
      const requestParams: Record<string, string> = { authorization };
      if (startTime !== undefined) {
        requestParams.start_time = startTime;
      }
      if (endTime !== undefined) {
        requestParams.end_time = endTime;
      }
      const response = await getUserStatistics(requestParams as unknown as IUserStatisticsRequest);
      setData(response.data ?? null);
    } catch (error) {
      message.error('获取统计数据失败');
      setData(null);
    } finally {
      setLoading(false);
    }
  }, []);

  const reset = useCallback(() => {
    setData(null);
  }, []);

  return {
    loading,
    data,
    fetchStatistics,
    reset,
  };
}
