/**
 * AI 模型相关 API 服务
 */
import { get, post, put, del } from '@/utils/request';
import type { IAIModel, IAIModelFormData } from '@/types';

// 列表响应数据
export interface IAIModelListResponse {
  total: number;
  list: IAIModel[];
}

/**
 * 获取 AI 模型列表
 * @param page 页码
 * @param page_size 每页数量
 * @returns 列表数据
 */
export async function getAIModelList(page: number = 1, page_size: number = 10) {
  return get<IAIModelListResponse>('/api/ai-models', {
    page,
    page_size,
  });
}

/**
 * 获取 AI 模型详情
 * @param id AI 模型ID
 * @returns 详情数据
 */
export async function getAIModel(id: number) {
  return get<IAIModel>(`/api/ai-models/${id}`);
}

/**
 * 创建 AI 模型
 * @param data 创建数据
 * @returns 创建结果
 */
export async function createAIModel(data: IAIModelFormData) {
  return post<IAIModel>('/api/ai-models', data, {
    showSuccessMessage: true,
  });
}

/**
 * 更新 AI 模型
 * @param id AI 模型ID
 * @param data 更新数据
 * @returns 更新结果
 */
export async function updateAIModel(id: number, data: IAIModelFormData) {
  return put<IAIModel>(`/api/ai-models/${id}`, data, {
    showSuccessMessage: true,
  });
}

/**
 * 删除 AI 模型
 * @param id AI 模型ID
 * @returns 删除结果
 */
export async function deleteAIModel(id: number) {
  return del(`/api/ai-models/${id}`, undefined, {
    showSuccessMessage: true,
  });
}
