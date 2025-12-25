/**
 * Token 数据管理 Hook
 * 功能：管理 Token 的列表、创建、更新、删除等数据逻辑
 */
import { useState, useCallback } from 'react';
import {
  getTokenList,
  createToken,
  updateToken,
  deleteToken,
} from '@/services/token';
import type { IToken, ITokenFormData } from '@/types';

/**
 * 分页信息类型
 */
export interface IPaginationInfo {
  current: number;
  pageSize: number;
}

/**
 * 列表数据响应类型
 */
export interface IUseTokenReturn {
  /** 数据列表 */
  dataSource: IToken[];
  /** 总数 */
  total: number;
  /** 加载状态 */
  loading: boolean;
  /** 分页信息 */
  pagination: IPaginationInfo;
  /** 加载数据 */
  loadData: (page?: number, pageSize?: number, keyword?: string) => Promise<void>;
  /** 创建 Token */
  handleCreate: (data: ITokenFormData) => Promise<boolean>;
  /** 更新 Token */
  handleUpdate: (id: number, data: ITokenFormData) => Promise<boolean>;
  /** 删除 Token */
  handleDelete: (id: number) => Promise<boolean>;
}

/**
 * Token 数据管理 Hook
 * @returns 数据和方法
 */
export function useToken(): IUseTokenReturn {
  const [loading, setLoading] = useState(false);
  const [dataSource, setDataSource] = useState<IToken[]>([]);
  const [total, setTotal] = useState(0);
  const [pagination, setPagination] = useState<IPaginationInfo>({
    current: 1,
    pageSize: 10,
  });

  /**
   * 加载列表数据
   */
  const loadData = useCallback(async (page: number = 1, pageSize: number = 10, keyword: string = '') => {
    setLoading(true);
    try {
      const result = await getTokenList(page, pageSize, keyword);
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
  }, []);

  /**
   * 创建 Token
   */
  const handleCreate = useCallback(async (data: ITokenFormData): Promise<boolean> => {
    try {
      const result = await createToken(data);
      if (result.success) {
        await loadData(pagination.current, pagination.pageSize);
        return true;
      }
      return false;
    } catch (error) {
      console.error('创建失败:', error);
      return false;
    }
  }, [loadData, pagination]);

  /**
   * 更新 Token
   */
  const handleUpdate = useCallback(
    async (id: number, data: ITokenFormData): Promise<boolean> => {
      try {
        const result = await updateToken(id, data);
        if (result.success) {
          await loadData(pagination.current, pagination.pageSize);
          return true;
        }
        return false;
      } catch (error) {
        console.error('更新失败:', error);
        return false;
      }
    },
    [loadData, pagination],
  );

  /**
   * 删除 Token
   */
  const handleDelete = useCallback(
    async (id: number): Promise<boolean> => {
      try {
        const result = await deleteToken(id);
        if (result.success) {
          // 如果当前页没有数据了，回到上一页
          const newTotal = total - 1;
          const newPage =
            pagination.current > 1 &&
            (pagination.current - 1) * pagination.pageSize >= newTotal
              ? pagination.current - 1
              : pagination.current;
          await loadData(newPage, pagination.pageSize);
          return true;
        }
        return false;
      } catch (error) {
        console.error('删除失败:', error);
        return false;
      }
    },
    [loadData, pagination, total],
  );

  return {
    dataSource,
    total,
    loading,
    pagination,
    loadData,
    handleCreate,
    handleUpdate,
    handleDelete,
  };
}
