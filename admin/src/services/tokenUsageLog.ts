/**
 * Token 使用记录相关 API 服务
 */
import { get } from '@/utils/request';
import type {
  ITokenUsageLog,
  ITokenUsageLogListResponse,
  ITokenUsageLogListRequest,
} from '@/types/tokenUsageLog';
import type {
  ITokenRankingResponse,
  ITokenRankingRequest,
  IUserStatisticsResponse,
  IUserStatisticsRequest,
} from '@/types/tokenRanking';

/**
 * 获取 Token 使用记录列表
 * @param params 查询参数
 * @returns 列表数据
 */
export async function getTokenUsageLogList(params: ITokenUsageLogListRequest = {}) {
  return get<ITokenUsageLogListResponse>('/api-logs/api/request-logs', params);
}

/**
 * 获取 Token 使用记录详情
 * @param id 记录ID
 * @returns 详情数据
 */
export async function getTokenUsageLogDetail(id: number) {
  return get<ITokenUsageLog>(`/api-logs/api/request-logs/${id}`);
}

/**
 * 获取 Authorization 使用次数排行榜
 * @param params 查询参数
 * @returns 排行榜数据
 */
export async function getTokenRanking(params: ITokenRankingRequest = {}) {
  return get<ITokenRankingResponse>('/api-logs/api/request-logs/ranking', params);
}

/**
 * 获取用户统计数据
 * @param params 查询参数
 * @returns 统计数据
 */
export async function getUserStatistics(params: IUserStatisticsRequest) {
  return get<IUserStatisticsResponse>('/api-logs/api/request-logs/statistics', params);
}
