/**
 * Token 使用记录类型定义
 */

/** Token 使用记录数据类型 */
export interface ITokenUsageLog {
  /** 记录ID */
  id: number;
  /** 时间戳 */
  time: string;
  /** 日志级别 */
  level: string;
  /** 日志消息 */
  msg: string;
  /** 请求唯一标识 */
  request_id: string;
  /** HTTP 方法 */
  method: string;
  /** 请求路径 */
  path: string;
  /** URL 查询参数 */
  query: string;
  /** 客户端地址 */
  remote_addr: string;
  /** User-Agent */
  user_agent: string;
  /** X-Forwarded-For */
  x_forwarded_for: string;
  /** 请求头（JSON 对象） */
  request_headers: Record<string, string>;
  /** Authorization 头 */
  authorization: string;
  /** 请求体 */
  request_body: string;
  /** HTTP 响应状态码 */
  status: number;
  /** 响应头（JSON 对象） */
  response_headers: Record<string, string>;
  /** 请求耗时（毫秒） */
  latency_ms: number;
  /** 请求体大小（字节） */
  request_size_bytes: number;
  /** 响应体大小（字节） */
  response_size_bytes: number;
  /** 创建时间 */
  created_at: string;
  /** 更新时间 */
  updated_at: string;
}

/** 列表响应数据 */
export interface ITokenUsageLogListResponse {
  /** 总数量 */
  total: number;
  /** 列表数据 */
  list: ITokenUsageLog[];
}

/** 列表查询请求参数 */
export interface ITokenUsageLogListRequest {
  /** 页码 */
  page?: number;
  /** 每页数量 */
  page_size?: number;
  /** 按 request_id 精确查询 */
  request_id?: string;
  /** 开始时间 */
  start_time?: string;
  /** 结束时间 */
  end_time?: string;
  /** 按状态码过滤（ProTable 传入 number，后端接收 string） */
  status?: number | string;
  /** 按 HTTP 方法过滤 */
  method?: string;
  /** 按 Authorization 头模糊匹配 */
  authorization?: string;
}

/** 分页信息 */
export interface IPaginationInfo {
  /** 当前页码 */
  current: number;
  /** 每页数量 */
  pageSize: number;
}
