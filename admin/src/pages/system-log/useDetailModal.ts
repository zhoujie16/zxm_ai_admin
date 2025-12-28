/**
 * 详情弹窗逻辑 hook
 */
import { useState, useCallback } from 'react';
import type { ISystemLog } from '@/types/systemLog';

export interface IUseDetailModalReturn {
  /** 详情弹窗是否可见 */
  detailVisible: boolean;
  /** 当前详情记录 */
  detailRecord: ISystemLog | null;
  /** 查看详情 */
  handleViewDetail: (record: ISystemLog) => void;
  /** 关闭详情弹窗 */
  handleCloseDetail: () => void;
}

/**
 * 详情弹窗逻辑管理 Hook
 */
export const useDetailModal = (): IUseDetailModalReturn => {
  const [detailVisible, setDetailVisible] = useState(false);
  const [detailRecord, setDetailRecord] = useState<ISystemLog | null>(null);

  const handleViewDetail = useCallback((record: ISystemLog) => {
    setDetailRecord(record);
    setDetailVisible(true);
  }, []);

  const handleCloseDetail = useCallback(() => {
    setDetailVisible(false);
    setDetailRecord(null);
  }, []);

  return {
    detailVisible,
    detailRecord,
    handleViewDetail,
    handleCloseDetail,
  };
};
