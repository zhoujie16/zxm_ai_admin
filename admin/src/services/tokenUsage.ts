/**
 * Token 使用记录相关 API 服务
 */
import { get } from '@/utils/request';
import type { ITokenUsageLogListResponse } from '@/types';

/**
 * 获取 Token 使用记录列表
 * @param token Token 值
 * @param page 页码
 * @param page_size 每页数量
 * @returns 列表数据
 */
export async function getTokenUsageLogs(token: string, page: number = 1, page_size: number = 10) {
  return get<ITokenUsageLogListResponse>('/api/token-usage-logs', {
    token,
    page,
    page_size,
  });
}
