/**
 * Token 使用记录数据管理 Hook
 * 功能：管理 Token 使用记录的列表查询数据逻辑
 */
import { useState, useCallback } from 'react';
import { getTokenUsageLogList } from '@/services/tokenUsageLog';
import type { ITokenUsageLog, ITokenUsageLogListRequest, IPaginationInfo } from '@/types/tokenUsageLog';

/**
 * Hook 返回类型
 */
export interface IUseTokenUsageLogReturn {
  /** 数据列表 */
  dataSource: ITokenUsageLog[];
  /** 总数 */
  total: number;
  /** 加载状态 */
  loading: boolean;
  /** 分页信息 */
  pagination: IPaginationInfo;
  /** 加载数据 */
  loadData: (page?: number, pageSize?: number, params?: ITokenUsageLogListRequest) => Promise<void>;
}

/**
 * Token 使用记录数据管理 Hook
 * @returns 数据和方法
 */
export function useTokenUsageLog(): IUseTokenUsageLogReturn {
  const [loading, setLoading] = useState(false);
  const [dataSource, setDataSource] = useState<ITokenUsageLog[]>([]);
  const [total, setTotal] = useState(0);
  const [pagination, setPagination] = useState<IPaginationInfo>({
    current: 1,
    pageSize: 10,
  });

  /**
   * 加载列表数据
   */
  const loadData = useCallback(
    async (page: number = 1, pageSize: number = 10, params: ITokenUsageLogListRequest = {}) => {
      setLoading(true);
      try {
        const result = await getTokenUsageLogList({
          page,
          page_size: pageSize,
          ...params,
        });
        if (result.success && result.data) {
          setDataSource(result.data.list);
          setTotal(result.data.total);
          setPagination({
            current: page,
            pageSize: pageSize,
          });
        }
      } catch (error) {
        console.error('加载数据失败:', error);
      } finally {
        setLoading(false);
      }
    },
    [],
  );

  return {
    dataSource,
    total,
    loading,
    pagination,
    loadData,
  };
}
