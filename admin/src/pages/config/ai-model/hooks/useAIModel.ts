/**
 * AI 模型数据管理 Hook
 * 功能：管理 AI 模型的列表、创建、更新、删除等数据逻辑
 */
import { useState, useCallback } from 'react';
import {
  getAIModelList,
  createAIModel,
  updateAIModel,
  deleteAIModel,
} from '@/services/aiModel';
import type { IAIModel, IAIModelFormData } from '@/types';

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
export interface IUseAIModelReturn {
  /** 数据列表 */
  dataSource: IAIModel[];
  /** 总数 */
  total: number;
  /** 加载状态 */
  loading: boolean;
  /** 分页信息 */
  pagination: IPaginationInfo;
  /** 加载数据 */
  loadData: (page?: number, pageSize?: number) => Promise<void>;
  /** 创建 AI 模型 */
  handleCreate: (data: IAIModelFormData) => Promise<boolean>;
  /** 更新 AI 模型 */
  handleUpdate: (id: number, data: IAIModelFormData) => Promise<boolean>;
  /** 删除 AI 模型 */
  handleDelete: (id: number) => Promise<boolean>;
}

/**
 * AI 模型数据管理 Hook
 * @returns 数据和方法
 */
export function useAIModel(): IUseAIModelReturn {
  const [loading, setLoading] = useState(false);
  const [dataSource, setDataSource] = useState<IAIModel[]>([]);
  const [total, setTotal] = useState(0);
  const [pagination, setPagination] = useState<IPaginationInfo>({
    current: 1,
    pageSize: 10,
  });

  /**
   * 加载列表数据
   */
  const loadData = useCallback(async (page: number = 1, pageSize: number = 10) => {
    setLoading(true);
    try {
      const result = await getAIModelList(page, pageSize);
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
   * 创建 AI 模型
   */
  const handleCreate = useCallback(async (data: IAIModelFormData): Promise<boolean> => {
    try {
      const result = await createAIModel(data);
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
   * 更新 AI 模型
   */
  const handleUpdate = useCallback(
    async (id: number, data: IAIModelFormData): Promise<boolean> => {
      try {
        const result = await updateAIModel(id, data);
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
   * 删除 AI 模型
   */
  const handleDelete = useCallback(
    async (id: number): Promise<boolean> => {
      try {
        const result = await deleteAIModel(id);
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
