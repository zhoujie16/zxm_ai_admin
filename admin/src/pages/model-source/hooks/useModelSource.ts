/**
 * 模型来源数据管理 Hook
 */
import { useState, useCallback } from 'react';
import { message } from 'antd';
import {
  getModelSourceList,
  createModelSource,
  updateModelSource,
  deleteModelSource,
} from '@/services/modelSource';
import type {
  IModelSource,
  ICreateModelSourceFormData,
  IUpdateModelSourceFormData,
} from '@/types';

export interface IPaginationInfo {
  current: number;
  pageSize: number;
}

export interface IUseModelSourceReturn {
  dataSource: IModelSource[];
  total: number;
  loading: boolean;
  pagination: IPaginationInfo;
  loadData: (page?: number, pageSize?: number) => Promise<void>;
  handleCreate: (data: ICreateModelSourceFormData) => Promise<boolean>;
  handleUpdate: (id: number, data: IUpdateModelSourceFormData) => Promise<boolean>;
  handleDelete: (id: number) => Promise<boolean>;
}

export function useModelSource(): IUseModelSourceReturn {
  const [loading, setLoading] = useState(false);
  const [dataSource, setDataSource] = useState<IModelSource[]>([]);
  const [total, setTotal] = useState(0);
  const [pagination, setPagination] = useState<IPaginationInfo>({
    current: 1,
    pageSize: 10,
  });

  const loadData = useCallback(async (page: number = 1, pageSize: number = 10) => {
    setLoading(true);
    try {
      const result = await getModelSourceList(page, pageSize);
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

  const handleCreate = useCallback(
    async (data: ICreateModelSourceFormData): Promise<boolean> => {
      try {
        const result = await createModelSource(data);
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
    },
    [loadData, pagination],
  );

  const handleUpdate = useCallback(
    async (id: number, data: IUpdateModelSourceFormData): Promise<boolean> => {
      try {
        const result = await updateModelSource(id, data);
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

  const handleDelete = useCallback(
    async (id: number): Promise<boolean> => {
      try {
        const result = await deleteModelSource(id);
        if (result.success) {
          message.success('删除成功');
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
