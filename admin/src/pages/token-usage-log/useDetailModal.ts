/**
 * 详情弹窗逻辑 hook
 */
import { useState, useCallback } from 'react';
import type { ITokenUsageLog } from '@/types/tokenUsageLog';

export interface IUseDetailModalReturn {
  /** 详情弹窗是否可见 */
  detailVisible: boolean;
  /** 当前详情记录 */
  detailRecord: ITokenUsageLog | null;
  /** 查看详情 */
  handleViewDetail: (record: ITokenUsageLog) => void;
  /** 关闭详情弹窗 */
  handleCloseDetail: () => void;
}

/**
 * 详情弹窗逻辑管理 Hook
 */
export const useDetailModal = (): IUseDetailModalReturn => {
  const [detailVisible, setDetailVisible] = useState(false);
  const [detailRecord, setDetailRecord] = useState<ITokenUsageLog | null>(null);

  const handleViewDetail = useCallback((record: ITokenUsageLog) => {
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
