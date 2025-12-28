/**
 * 系统日志类型定义
 */

/** 系统日志数据类型 */
export interface ISystemLog {
  /** 记录ID */
  id: number;
  /** 时间戳 */
  time: string;
  /** 日志级别 */
  level: string;
  /** 日志消息 */
  msg: string;
  /** 创建时间 */
  created_at: string;
  /** 更新时间 */
  updated_at: string;
}

/** 列表响应数据 */
export interface ISystemLogListResponse {
  /** 总数量 */
  total: number;
  /** 列表数据 */
  list: ISystemLog[];
}

/** 列表查询请求参数 */
export interface ISystemLogListRequest {
  /** 页码 */
  page?: number;
  /** 每页数量 */
  page_size?: number;
  /** 按日志级别过滤 */
  level?: string;
  /** 开始时间 */
  start_time?: string;
  /** 结束时间 */
  end_time?: string;
}
