/**
 * Token 相关 API 服务
 */
import { get, post, put, del } from '@/utils/request';
import type { IToken, ITokenFormData } from '@/types';

// 列表响应数据
export interface ITokenListResponse {
  total: number;
  list: IToken[];
}

/**
 * 获取 Token 列表
 * @param page 页码
 * @param page_size 每页数量
 * @param keyword 关键词
 * @returns 列表数据
 */
export async function getTokenList(page: number = 1, page_size: number = 10, keyword?: string) {
  return get<ITokenListResponse>('/api/tokens', {
    page,
    page_size,
    keyword,
  });
}

/**
 * 获取 Token 详情
 * @param id Token ID
 * @returns 详情数据
 */
export async function getToken(id: number) {
  return get<IToken>(`/api/tokens/${id}`);
}

/**
 * 创建 Token
 * @param data 创建数据
 * @returns 创建结果
 */
export async function createToken(data: ITokenFormData) {
  return post<IToken>('/api/tokens', data, {
    showSuccessMessage: true,
  });
}

/**
 * 更新 Token
 * @param id Token ID
 * @param data 更新数据
 * @returns 更新结果
 */
export async function updateToken(id: number, data: ITokenFormData) {
  return put<IToken>(`/api/tokens/${id}`, data, {
    showSuccessMessage: true,
  });
}

/**
 * 删除 Token
 * @param id Token ID
 * @returns 删除结果
 */
export async function deleteToken(id: number) {
  return del(`/api/tokens/${id}`, undefined, {
    showSuccessMessage: true,
  });
}
