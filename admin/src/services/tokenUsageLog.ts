/**
 * Token 使用记录相关 API 服务
 */
import { get } from '@/utils/request';
import type {
  ITokenUsageLog,
  ITokenUsageLogListResponse,
  ITokenUsageLogListRequest,
} from '@/types/tokenUsageLog';

/**
 * 获取 Token 使用记录列表
 * @param params 查询参数
 * @returns 列表数据
 */
export async function getTokenUsageLogList(params: ITokenUsageLogListRequest = {}) {
  return get<ITokenUsageLogListResponse>('/api-logs/request-logs', params);
}

/**
 * 获取 Token 使用记录详情
 * @param id 记录ID
 * @returns 详情数据
 */
export async function getTokenUsageLogDetail(id: number) {
  return get<ITokenUsageLog>(`/api-logs/request-logs/${id}`);
}
