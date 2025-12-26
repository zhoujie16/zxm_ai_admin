/**
 * 代理服务相关 API 服务
 */
import { get, post, put, del } from '@/utils/request';
import type { IProxyService, IProxyServiceFormData } from '@/types';

// 列表响应数据
export interface IProxyServiceListResponse {
  total: number;
  list: IProxyService[];
}

/**
 * 获取代理服务列表
 * @param page 页码
 * @param page_size 每页数量
 * @returns 列表数据
 */
export async function getProxyServiceList(page: number = 1, page_size: number = 10) {
  return get<IProxyServiceListResponse>('/api/proxy-services', {
    page,
    page_size,
  });
}

/**
 * 获取代理服务详情
 * @param id 代理服务ID
 * @returns 详情数据
 */
export async function getProxyService(id: number) {
  return get<IProxyService>(`/api/proxy-services/${id}`);
}

/**
 * 创建代理服务
 * @param data 创建数据
 * @returns 创建结果
 */
export async function createProxyService(data: IProxyServiceFormData) {
  return post<IProxyService>('/api/proxy-services', data);
}

/**
 * 更新代理服务
 * @param id 代理服务ID
 * @param data 更新数据
 * @returns 更新结果
 */
export async function updateProxyService(id: number, data: IProxyServiceFormData) {
  return put<IProxyService>(`/api/proxy-services/${id}`, data);
}

/**
 * 删除代理服务
 * @param id 代理服务ID
 * @returns 删除结果
 */
export async function deleteProxyService(id: number) {
  return del(`/api/proxy-services/${id}`);
}

