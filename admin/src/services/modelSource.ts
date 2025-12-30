/**
 * 模型来源相关 API 服务
 */
import { get, post, put, del } from '@/utils/request';
import type {
  IModelSource,
  ICreateModelSourceFormData,
  IUpdateModelSourceFormData,
} from '@/types';

// 列表响应数据
export interface IModelSourceListResponse {
  total: number;
  list: IModelSource[];
}

/**
 * 获取模型来源列表
 * @param page 页码
 * @param page_size 每页数量
 * @returns 列表数据
 */
export async function getModelSourceList(
  page: number = 1,
  page_size: number = 10,
) {
  return get<IModelSourceListResponse>('/api/model-sources', {
    page,
    page_size,
  });
}

/**
 * 获取模型来源详情
 * @param id 模型来源ID
 * @returns 详情数据
 */
export async function getModelSource(id: number) {
  return get<IModelSource>(`/api/model-sources/${id}`);
}

/**
 * 创建模型来源
 * @param data 创建数据
 * @returns 创建结果
 */
export async function createModelSource(data: ICreateModelSourceFormData) {
  return post<IModelSource>('/api/model-sources', data);
}

/**
 * 更新模型来源
 * @param id 模型来源ID
 * @param data 更新数据
 * @returns 更新结果
 */
export async function updateModelSource(
  id: number,
  data: IUpdateModelSourceFormData,
) {
  return put<IModelSource>(`/api/model-sources/${id}`, data);
}

/**
 * 删除模型来源
 * @param id 模型来源ID
 * @returns 删除结果
 */
export async function deleteModelSource(id: number) {
  return del(`/api/model-sources/${id}`);
}
