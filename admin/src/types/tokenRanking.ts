/**
 * Token 使用排行榜相关类型定义
 */

/** 排行榜单条记录 */
export interface ITokenRankingItem {
  /** authorization 值 */
  authorization: string;
  /** 请求次数 */
  count: number;
}

/** 排行榜列表响应 */
export interface ITokenRankingResponse {
  /** 总记录数（不同 authorization 的数量） */
  total: number;
  /** 时间范围 */
  time_range: {
    /** 开始时间 */
    start: string;
    /** 结束时间 */
    end: string;
  };
  /** 排行榜列表 */
  list: ITokenRankingItem[];
}

/** 排行榜查询请求参数 */
export interface ITokenRankingRequest {
  /** 开始时间 */
  start_time?: string;
  /** 结束时间 */
  end_time?: string;
  /** 页码 */
  page?: number;
  /** 每页数量 */
  page_size?: number;
}

/** 用户统计数据响应 */
export interface IUserStatisticsResponse {
  /** authorization 值 */
  authorization: string;
  /** 时间范围 */
  time_range: {
    start: string;
    end: string;
  };
  /** 汇总统计 */
  summary: {
    total_requests: number;
    total_request_bytes: number;
    total_response_bytes: number;
    avg_request_bytes: number;
    avg_response_bytes: number;
  };
  /** 延迟统计 */
  latency: {
    total_ms: number;
    avg_ms: number;
    min_ms: number;
    max_ms: number;
  };
  /** 按IP分组统计 */
  by_ip: Array<{
    ip: string;
    count: number;
  }>;
  /** 按路径分组统计 */
  by_path: Array<{
    path: string;
    count: number;
  }>;
  /** 按日期分组统计 */
  by_date: Array<{
    date: string;
    count: number;
  }>;
  /** 按小时分组统计 */
  by_time: Array<{
    time: string;
    count: number;
  }>;
}

/** 用户统计查询请求参数 */
export interface IUserStatisticsRequest {
  /** authorization 值 */
  authorization: string;
  /** 开始时间 */
  start_time?: string;
  /** 结束时间 */
  end_time?: string;
}
