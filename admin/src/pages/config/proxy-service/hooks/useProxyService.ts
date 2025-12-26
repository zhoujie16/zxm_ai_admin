/**
 * 代理服务数据管理 Hook
 * 功能：管理代理服务的列表、创建、更新、删除等数据逻辑
 */
import { useState, useCallback } from 'react';
import { message } from 'antd';
import {
  getProxyServiceList,
  createProxyService,
  updateProxyService,
  deleteProxyService,
} from '@/services/proxyService';
import type { IProxyService, IProxyServiceFormData } from '@/types';

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
export interface IUseProxyServiceReturn {
  /** 数据列表 */
  dataSource: IProxyService[];
  /** 总数 */
  total: number;
  /** 加载状态 */
  loading: boolean;
  /** 分页信息 */
  pagination: IPaginationInfo;
  /** 加载数据 */
  loadData: (page?: number, pageSize?: number) => Promise<void>;
  /** 创建代理服务 */
  handleCreate: (data: IProxyServiceFormData) => Promise<boolean>;
  /** 更新代理服务 */
  handleUpdate: (id: number, data: IProxyServiceFormData) => Promise<boolean>;
  /** 删除代理服务 */
  handleDelete: (id: number) => Promise<boolean>;
}

/**
 * 代理服务数据管理 Hook
 * @returns 数据和方法
 */
export function useProxyService(): IUseProxyServiceReturn {
  const [loading, setLoading] = useState(false);
  const [dataSource, setDataSource] = useState<IProxyService[]>([]);
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
      const result = await getProxyServiceList(page, pageSize);
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
   * 创建代理服务
   */
  const handleCreate = useCallback(async (data: IProxyServiceFormData): Promise<boolean> => {
    try {
      const result = await createProxyService(data);
      if (result.success) {
        message.success('创建成功');
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
   * 更新代理服务
   */
  const handleUpdate = useCallback(
    async (id: number, data: IProxyServiceFormData): Promise<boolean> => {
      try {
        const result = await updateProxyService(id, data);
        if (result.success) {
          message.success('更新成功');
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
   * 删除代理服务
   */
  const handleDelete = useCallback(
    async (id: number): Promise<boolean> => {
      try {
        const result = await deleteProxyService(id);
        if (result.success) {
          message.success('删除成功');
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

