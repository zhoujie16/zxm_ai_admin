/**
 * 系统日志相关 API 服务
 */
import { get } from '@/utils/request';
import type {
  ISystemLog,
  ISystemLogListResponse,
  ISystemLogListRequest,
} from '@/types/systemLog';

/**
 * 获取系统日志列表
 * @param params 查询参数
 * @returns 列表数据
 */
export async function getSystemLogList(params: ISystemLogListRequest = {}) {
  return get<ISystemLogListResponse>('/api-logs/api/system-logs', params);
}

/**
 * 获取系统日志详情
 * @param id 记录ID
 * @returns 详情数据
 */
export async function getSystemLogDetail(id: number) {
  return get<ISystemLog>(`/api-logs/api/system-logs/${id}`);
}
